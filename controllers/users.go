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

type SignUpForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// New is used to render a sign-up page
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	err := u.NewView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

// Create is used to process a sign-up form
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignUpForm
	err := ParseForm(r, &form)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, form)
}
