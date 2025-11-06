package main

import (
	"fmt"
	"log"
	"net/http"
)

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