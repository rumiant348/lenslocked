package controllers

import (
	"fmt"
	"lenslocked.com/models"
	"lenslocked.com/views"
	"net/http"
)

func NewGalleries(gs models.GalleryService) *Galleries {
	return &Galleries{
		NewView: views.NewView("bootstrap", "galleries/new"),
		gs:      gs,
	}
}

type Galleries struct {
	NewView *views.View
	gs      models.GalleryService
}

type GalleryForm struct {
	Title string `scheme:"title"`
}

func (g *Galleries) New(w http.ResponseWriter, r *http.Request) {
	g.NewView.Render(w, nil)
}

func (g *Galleries) Create(w http.ResponseWriter, r *http.Request) {
	var vd views.Data
	var form GalleryForm
	if err := parseForm(r, &form); err != nil {
		vd.SetAlert(err)
		g.NewView.Render(w, vd)
		return
	}
	gallery := models.Gallery{
		Title: form.Title,
	}
	if err := g.gs.Create(&gallery); err != nil {
		vd.SetAlert(err)
		g.NewView.Render(w, vd)
		return
	}
	fmt.Fprintln(w, gallery)
}
