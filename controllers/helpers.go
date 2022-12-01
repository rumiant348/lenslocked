package controllers

import (
	"github.com/gorilla/schema"
	"net/http"
)

func parseForm(r *http.Request, dst interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	return schema.NewDecoder().Decode(dst, r.PostForm)
}
