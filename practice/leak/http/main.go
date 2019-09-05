package main

import (
	"fmt"
	"net/http"

	_ "net/http/pprof"
)

func comdHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.Body
	defer r.Body.Close()
	//con, _ := ioutil.ReadAll(r.Body) //获取post的数据
	//m := make(map[string]interface{})
	//blocksize,_ :=strconv.Atoi(flag.Arg(0))
	//fmt.Print(body)
	//block := make([]byte,blocksize)
	//
	///*data := make([]byte, 1025)
	//r := bufio.NewReader(file)
	//r.Read(data)*/
	//for{
	//	n,err := body.Read(block)
	//	fmt.Print(n)
	//	if (err != nil && err != io.EOF){
	//		panic(err)
	//	}
	//	if (0 == n ){
	//		break
	//	}
	//}
	//if con != nil {
	//	re := json.Unmarshal([]byte(string(con)), &m)
	//	if re != nil {
	//		_ = re
	//	} else {
	//
	//	}
	//}
	fmt.Println("success")
}

func main() {
	http.HandleFunc("/", comdHandler)
	//http.HandleFunc("/ungetTasks",ungetTasks)
	//http.HandleFunc("/userlist",userList)
	http.ListenAndServe(":9731", nil)
}
