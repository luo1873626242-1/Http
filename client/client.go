package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main(){
	HttpGet()
}

func HttpGet(){
	tr := &http.Transport{}
	client := http.Client{
		Transport: tr,
	}
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8000/ping", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(bodyBytes))
}