package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"io"
	"path"
	"strconv"
)

var dir string
var port int
var staticHandler http.Handler

func init() {
	dir = path.Dir(os.Args[0])
	flag.IntVar(&port, "port", 8080, "服务器端口")

	flag.Parse()
	println(dir)
	dir = "/home/su"
	staticHandler = http.FileServer(http.Dir(dir))
}


func main() {
	http.HandleFunc("/", StaticServer)
	err := http.ListenAndServe(":"+ strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

func StaticServer(w http.ResponseWriter, req *http.Request) {
	println(req.URL.Path)
	if req.URL.Path == "/" {
		staticHandler.ServeHTTP(w, req)
		io.WriteString(w, "x\n")
		return
	}

	io.WriteString(w, "hello world!\n")
}
