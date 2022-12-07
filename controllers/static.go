package controllers

import "lenslocked.com/views"

func NewStatic() *Static {
	return &Static{
		Home:         views.NewView("bootstrap", "static/home"),
		Contact:      views.NewView("bootstrap", "static/contact"),
		NotFoundView: views.NewView("bootstrap", "static/notFound"),
		Faq:          views.NewView("bootstrap", "static/faq"),
	}
}

type Static struct {
	Home         *views.View
	Contact      *views.View
	Faq          *views.View
	NotFoundView *views.View
}
