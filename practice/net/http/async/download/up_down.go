package main

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/xid"
)

func main() {
	var routes = mux.NewRouter()
	routes.HandleFunc("/photo/save", savePhoto).Methods("POST")
	routes.HandleFunc("/photo/get/{filename}", getPhoto).Methods("GET")

	http.Handle("/", routes)
	log.Fatal(http.ListenAndServe(":8555", nil))

}

const (
	serverKey    = "bleashup-file-server"
	photosFolder = "/Users/fpf/Downloads/"
)

type file struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}

func retrievePhotoInfo(w http.ResponseWriter, r *http.Request) (file, error) {
	fileInstance := file{}
	var err error
	if err = r.ParseMultipartForm(5 * MB); err != nil {
		fmt.Println(err.Error())
		return file{}, err
	}
	fileInstance.File, fileInstance.FileHeader, err = r.FormFile("file")
	fmt.Println(fileInstance.FileHeader.Filename, "********************")
	if err != nil {
		fmt.Println(err.Error())
		return file{}, err
	}
	return fileInstance, nil

}

const (
	//MB represents a megabyte
	MB = 1 << 20
)

func generateID() string {
	guid := xid.New()
	ID := guid.String()
	return ID
}

func savePhoto(w http.ResponseWriter, r *http.Request) {
	processSave(w, r, photosFolder)
}

func getPhoto(w http.ResponseWriter, r *http.Request) {
	processGet(w, r, photosFolder)
}

func processSave(w http.ResponseWriter, r *http.Request, parentDir string) {

	fileInstance := file{}
	fileInstance, err := retrievePhotoInfo(w, r)
	defer fileInstance.File.Close()
	if err != nil {
		fmt.Println(err)
	}

	ID := generateID() + "_" + generateID() + "_" + generateID()
	filenames := strings.Split(fileInstance.FileHeader.Filename, ".")
	fmt.Println(filenames[len(filenames)-1])

	var file *os.File
	images, err := filepath.Glob(parentDir + ID + ".*")
	if err != nil {
		fmt.Println(err.Error())
	}

	if images != nil {
		file, err = os.Create(parentDir + ID + "." + filenames[len(filenames)-1])
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	file, err = os.Create(parentDir + ID + "." + filenames[len(filenames)-1])
	if err != nil {
		fmt.Println(err.Error())
	}

	io.Copy(file, fileInstance.File)

	if err != nil {
		fmt.Println(err.Error())
	}

	w.Write([]byte(ID + "." + filenames[len(filenames)-1]))
}

func processGet(w http.ResponseWriter, r *http.Request, parentDir string) {
	var vars = mux.Vars(r)
	fileName := parentDir + vars["filename"]
	if fileName == "" {
		http.Error(w, "Get 'file' not specified in url.", 400)
		return
	}
	fmt.Println("Client requests: " + fileName)

	openFile, err := os.Open(fileName)
	defer openFile.Close()

	if err != nil {
		return
	}

	rr, ww := io.Pipe()

	fileStat, _ := openFile.Stat()                     //Get info from file
	fileSize := strconv.FormatInt(fileStat.Size(), 10) //Get file size as a string

	fileHeader := make([]byte, 32)
	openFile.Read(fileHeader)
	fileContentType := http.DetectContentType(fileHeader)

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Type", fileContentType)
	w.Header().Set("Content-Length", fileSize)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			ww.Close()
			wg.Done()
		}()
		if err != nil {
			return
		}

		if _, err = io.Copy(ww, openFile); err != nil {
			return
		}
	}()
	wg.Wait()

	openFile.Seek(0, 0)//从文件头重新读取数据
	io.Copy(w, rr)

	return
}
