package controllers

import (
	"lenslocked.com/views"
	"net/http"
)

func NewGalleries() *Galleries {
	return &Galleries{
		NewView: views.NewView("bootstrap", "galleries/new"),
	}
}

type Galleries struct {
	NewView *views.View
}

func (g *Galleries) New(w http.ResponseWriter, r *http.Request) {
	g.NewView.Render(w, nil)
}
