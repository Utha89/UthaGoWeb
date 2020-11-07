package handler

import (
	"go_web_mengetahui_dia_method_post_atau_get/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

//funtion baru yg membutuhkan 2 paramater yaitu
//1. w http.ResponseWriter
//2. pointer nya yaitu r*http.Request
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	//body dr fnctionnya
	w.Write([]byte("Lagi Belajar Go Lang guys"))
}

func MarioHandler(w http.ResponseWriter, r *http.Request) {
	//body dr fnctionnya
	w.Write([]byte("Mario Balotelli"))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//body dr fnctionnya

	//ngehandle 4040
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//w.Write([]byte("Home Page"))
	//pakai package template.html untuk rendering ke web
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))

	//klo ada error kita handle
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm bro / sis", http.StatusInternalServerError)
		return
	}

	//disin kita akan parsing data dr function ke view caranya
	//data dummy
	//data := map[string]string{
	// data := map[string]interface{}{
	// 	"tittle":  "Belajar golang",
	// 	"content": "with utha ganteng",
	// }

	//Tanggal 7 November 2020
	//passing struct ke html
	//data := entity.Product{ID: 1, Name: "Sabun GIV", Price: 40000, Stock: 2}
	//next cara aksesnya di template

	//Slice datanya banyak yg nanti di looping
	data := [] /*namaStructnya*/ entity.Product{
		{ID: 1, Name: "Sabun GIV", Price: 40000, Stock: 2},
		{ID: 2, Name: "Sabun LUX", Price: 45000, Stock: 5},
		{ID: 3, Name: "Sabun DOVE", Price: 20000, Stock: 1},
	}

	// untuk rendering disin biar bisa jalanin yg "views", "index.html")
	// err = tmpl.Execute(w, nil)
	//kan sebelum nya nill yg kita parsing sekarang kita bisa ganti nil dengan data
	err = tmpl.Execute(w, data)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm bro / sis", http.StatusInternalServerError)
		return
	}

}

/*function product handler*/
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	//body dr fnctionnya
	//menangkap query berdasrkan id
	id := r.URL.Query().Get("id")

	//convery string ke integer pake nya Atoi
	//balikannya dia angka dan error
	idNumber, err := strconv.Atoi(id)

	if err != nil || idNumber < 1 {
		http.NotFound(w, r)
		return
	}
	//w.Write([]byte("Product Page"))

	//cetaknya pake Fprinf(w writer, "string", object bebas)
	//fmt.Fprintf(w, "Product Page : %d", idNumber)

	//web html layout
	//pakai package template.html untuk rendering ke web
	//untuk menggunakan layout dinamis tambahkan path.Join("views", "layout.html")
	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))

	//klo ada error kita handle
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm bro / sis", http.StatusInternalServerError)
		return
	}

	//disin kita akan parsing data dr function ke view caranya
	//data dummy
	//data := map[string]string{
	data := map[string]interface{}{

		"content": idNumber,
	}

	// untuk rendering disin biar bisa jalanin yg "views", "index.html")
	// err = tmpl.Execute(w, nil)
	//kan sebelum nya nill yg kita parsing sekarang kita bisa ganti nil dengan data
	err = tmpl.Execute(w, data)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm bro / sis", http.StatusInternalServerError)
		return
	}

}

//Tanggal 7 November 2020
//funct untuk mengetahi dia method post atau get

func PostGetHandler(w http.ResponseWriter, r *http.Request) {
	//set variabel dulu
	method := r.Method

	//pake sicth untuk ngecek
	switch method {
	case "GET":
		w.Write([]byte("Ini Method Get"))
	case "POST":
		w.Write([]byte("Ini Method Post"))
	default:
		http.Error(w, "Error is happening, keep calm bro / sis", http.StatusBadRequest)
	}
}
