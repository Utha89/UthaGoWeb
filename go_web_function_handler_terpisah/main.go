package main

import (

	"log"
	"net/http"
	"go_web_function_handler_terpisah/handler"
)

func main()  {
	mux:=http.NewServeMux()

	

	//mux.HandleFunc
	// func(pattern string, handler func(ResponseWriter, *Request))
	//pattern string isi /hello
	mux.HandleFunc("/",handler.HomeHandler)
	mux.HandleFunc("/hello",handler.HelloHandler)
	mux.HandleFunc("/mario",handler.MarioHandler)
	mux.HandleFunc("/product",handler.ProductHandler)
	
	
	log.Println("Starting web on port 8080")

	//jalanin servernya 
	//LListenAndServe func(addr string, handler Handler) error
	//parameter pertamanya addr string
	//keduanya si handle mux itu sendiri
	err:=http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
