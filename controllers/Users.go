package controllers

import (
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
	err := u.NewView.Template.ExecuteTemplate(w, u.newView.Layout, nil)
	if err != nil {
		panic(err)
	}
}
