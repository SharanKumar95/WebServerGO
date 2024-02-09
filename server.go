package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {

	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello!!")
	// })

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	 err:= http.ListenAndServe (":8090", nil)
	 if err!=nil {
		log.Fatal (err)
	}
}

func helloHandler (w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error (w, "404 not Found", http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error (w, "Only GET is supported", http.StatusNotFound)
	}

	fmt.Fprintf (w, "Hello there!")
}

func formHandler (w http.ResponseWriter, r *http.Request){
	err:=r.ParseForm()
	if err!=nil{
		// log.Fatal()
		fmt.Fprintf(w, "ParseForm err:%v", err)
	}

	fmt.Fprintf(w, "POST request SUccessful")
	name:=r.FormValue("name")
	address:=r.FormValue("address")

	fmt.Fprintf(w, "Name: %v\n", name)
	fmt.Fprintf(w, "Address: %v\n", address)
}
