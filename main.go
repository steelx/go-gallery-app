package main

import (
	"go-gallery-app/views"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	homeView *views.View
)

func main() {
	homeView = views.NewView("bootstrap", "views/home.tpl.html")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	http.ListenAndServe(":3000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		panic(err)
	}
}
