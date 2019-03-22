package main

import (
	"galleryApp2019/views"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	homeView *views.View
)

func main() {
	homeView = views.NewView("views/home.tpl.html")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	http.ListenAndServe(":3000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
