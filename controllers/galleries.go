package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"lenslocked.com/context"
	"lenslocked.com/models"
	"lenslocked.com/views"
	"net/http"
	"strconv"
)

func NewGalleries(gs models.GalleryService) *Galleries {
	return &Galleries{
		NewView:  views.NewView("bootstrap", "galleries/new"),
		ShowView: views.NewView("bootstrap", "galleries/show"),
		gs:       gs,
	}
}

type Galleries struct {
	NewView  *views.View
	ShowView *views.View
	gs       models.GalleryService
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

	user := context.User(r.Context())
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	gallery := models.Gallery{
		Title:  form.Title,
		UserID: user.ID,
	}
	if err := g.gs.Create(&gallery); err != nil {
		vd.SetAlert(err)
		g.NewView.Render(w, vd)
		return
	}
	fmt.Fprintln(w, gallery)
}

func (g *Galleries) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid gallery ID", http.StatusNotFound)
		return
	}
	_ = id
	// tmp gallery until the search is not implemented
	gallery := models.Gallery{
		Title: "A temporary fake gallery with id: " + idStr,
	}
	var vd views.Data
	vd.Yield = gallery
	g.ShowView.Render(w, vd)
}
