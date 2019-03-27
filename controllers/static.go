package controllers

import "go-gallery-app/views"

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "views/static/home.tpl.html"),
		Contact: views.NewView("bootstrap", "views/static/contact.tpl.html"),
	}
}

type Static struct {
	Home    *views.View
	Contact *views.View
}
