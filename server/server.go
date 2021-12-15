package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func PingHandle(w http.ResponseWriter, req *http.Request)(){
	fmt.Println(req.Header.Get("Content-Type"))
	w.Write([]byte("pong"))
}
func main(){
	http.HandleFunc("/ping", PingHandle)
	MyServer(":8000")
}

func MyServer(addr string){
	s := &http.Server{
		Addr:           addr,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}