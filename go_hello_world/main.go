package main

import (
	"log"
	"net/http"
)

func main()  {
	mux:=http.NewServeMux()

	//mux.HandleFunc
	// func(pattern string, handler func(ResponseWriter, *Request))
	//pattern string isi /hello
	mux.HandleFunc("/hello",helloHandler)

	log.Println("Starting web on port 8080")

	//jalanin servernya 
	//LListenAndServe func(addr string, handler Handler) error
	//parameter pertamanya addr string
	//keduanya si handle mux itu sendiri
	err:=http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
//funtion baru yg membutuhkan 2 paramater yaitu 
//1. w http.ResponseWriter
//2. pointer nya yaitu r*http.Request
func helloHandler(w http.ResponseWriter, r*http.Request)  {
	//body dr fnctionnya 
	w.Write([]byte("Lagi Belajar Go Lang guys"))
}