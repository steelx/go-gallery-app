package main

import (
	"github.com/gorilla/mux"
	"go-gallery-app/controllers"
	"net/http"
)

func main() {
	staticCtrl := controllers.NewStatic()
	usersCtrl := controllers.NewUsers("bootstrap", "views/users/signup.tpl.html")

	r := mux.NewRouter()
	r.Handle("/", staticCtrl.Home).Methods("GET")
	r.Handle("/contact", staticCtrl.Contact).Methods("GET")

	r.HandleFunc("/signup", usersCtrl.New).Methods("GET")
	r.HandleFunc("/signup", usersCtrl.Create).Methods("POST")

	http.ListenAndServe(":3000", r)
}

//49
