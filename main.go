package main

import (
	"github.com/gorilla/mux"
	"go-gallery-app/controllers"
	"go-gallery-app/views"
	"net/http"
)

var (
	homeView *views.View
)

func main() {
	homeView = views.NewView("bootstrap", "views/home.tpl.html")
	usersCtrl := controllers.NewUsers("bootstrap", "views/users/signup.tpl.html")

	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")

	r.HandleFunc("/signup", usersCtrl.New).Methods("GET")
	r.HandleFunc("/signup", usersCtrl.Create).Methods("POST")

	http.ListenAndServe(":3000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

//45
