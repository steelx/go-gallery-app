package controllers

import (
	"github.com/gorilla/schema"
	"net/http"
)

func ParseForm(r *http.Request, destination interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	decoder := schema.NewDecoder()
	return decoder.Decode(destination, r.PostForm)
}
