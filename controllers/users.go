package controllers

import (
	"fmt"
	"github.com/gorilla/schema"
	"go-gallery-app/views"
	"net/http"
)

func NewUsers(layout, template string) *Users {
	return &Users{
		NewView: views.NewView(layout, template),
	}
}

type Users struct {
	NewView *views.View
}

type SignUpForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	err := u.NewView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	decoder := schema.NewDecoder()
	var form SignUpForm
	err = decoder.Decode(&form, r.PostForm)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, form)
}
