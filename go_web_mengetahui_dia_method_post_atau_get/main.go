package main

import (
	"go_web_mengetahui_dia_method_post_atau_get/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//mux.HandleFunc atau rooting halaman
	// func(pattern string, handler func(ResponseWriter, *Request))
	//pattern string isi /hello
	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/mario", handler.MarioHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/post-get", handler.PostGetHandler)
	//cara load file css di folder assets
	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting web on port 8080")

	//jalanin servernya
	//LListenAndServe func(addr string, handler Handler) error
	//parameter pertamanya addr string
	//keduanya si handle mux itu sendiri
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
