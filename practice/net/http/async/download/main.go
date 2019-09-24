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
	if err != nil {
		fmt.Println(err)
	}
	ID := generateID() + "_" + generateID() + "_" + generateID()
	defer fileInstance.File.Close()
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
	var fileName = vars["filename"]
	//First of check if Get is set in the URL
	Filename := parentDir + fileName
	if Filename == "" {
		//Get not set, send a 400 bad request
		http.Error(w, "Get 'file' not specified in url.", 400)
		return
	}
	fmt.Println("Client requests: " + Filename)

	//Check if file exists and open
	Openfile, err := os.Open(Filename)
	defer Openfile.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		http.Error(w, "File not found.", 404)
		return
	}

	//File is found, create and send the correct headers
	rr, ww := io.Pipe()
	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+Filename)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)
	go func() {
		defer ww.Close()
		if err != nil {
			return
		}
		fmt.Println("io.Copy", Openfile)
		if _, err = io.Copy(ww, Openfile); err != nil {
			return
		}
	}()
	Openfile.Seek(0, 0)
	io.Copy(w, rr)
	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	//	Openfile.Seek(0, 0)
	//	io.Copy(w, Openfile) //'Copy' the file to the client
	return
}
