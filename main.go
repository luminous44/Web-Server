package main

import (
	"fmt"
	"log"
	"net/http"
)
func formHandler(res http.ResponseWriter, r *http.Request){

		if r.URL.Path != "/home" {
		http.Error(res,"404 not found",http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {

		fmt.Fprintf(res,"ParseForm() err: %v",err)
		return
		
	}
	fmt.Fprintf(res,"Post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(res,"Name : %s\n",name)
	fmt.Fprintf(res,"Address : %s\n",address)

}

func helloHandler(res http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(res,"404 not found",http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(res,"Method id not supported",http.StatusNotFound)
		return
 	}

	fmt.Fprintf(res,"Hello")
}


func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/home",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080",nil); err != nil{
      log.Fatal(err)
	}



}