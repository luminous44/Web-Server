package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHander(res http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(res,"404 not found",http.StatusNotFound)
	}

	if r.Method != "GET"{
		http.Error(res,"Method id not supported",http.StatusNotFound)
	}

	fmt.Fprintf(res,"Hello")
}


func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/home",formHander)
	http.HandleFunc("/hello",helloHander)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080",nil); err != nil{
      log.Fatal(err)
	}



}