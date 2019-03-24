package controllers

import (
	"fmt"
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

	fmt.Fprintln(w, r.PostForm["email"])
	fmt.Fprintln(w, r.PostFormValue("email"))
}
