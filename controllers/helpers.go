package controllers

import (
	"github.com/gorilla/schema"
	"net/http"
)

func ParseForm(r *http.Request, dst interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	return schema.NewDecoder().Decode(dst, r.PostForm)
}
