package views

import (
	"bytes"
	"errors"
	"github.com/gorilla/csrf"
	"html/template"
	"io"
	"lenslocked.com/context"
	"log"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir   = "views/layouts/"
	TemplateDir = "views/"
	TemplateExt = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
	files = append(files, getLayoutFiles()...)

	t, err := template.New("").Funcs(template.FuncMap{
		"csrfField": func() (template.HTML, error) {
			return "", errors.New("csrfField is not implemented")
		},
	}).ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{Template: t, Layout: layout}
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(w http.ResponseWriter, r *http.Request,
	data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	var vd Data
	switch d := data.(type) {
	case Data:
		// We need to do this so we can access the data in a var
		// with the type Data
		vd = d
	default:
		// If the data IS NOT of the type Data, we create one
		// and set the data to the Yield field
		data = Data{
			Yield: data,
		}
	}
	// Lookup and set the user
	vd.User = context.User(r.Context())
	var buf bytes.Buffer

	// implement the csrfField function
	csrfField := csrf.TemplateField(r)

	tpl := v.Template.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return csrfField
		},
	})

	err := tpl.ExecuteTemplate(&buf, v.Layout, vd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong. If the problem "+
			"persists, please email support@lenslocked.com",
			http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, r, nil)
}

// getLayoutFiles returns a slice of strings
// representing all layout files
// in LayoutDir with extension TemplateExt
func getLayoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

// addTemplatePath takes in a slice of strings
// representing file paths for templates, and it prepends
// the TemplateDir directory to each string in the slice
//
// Eg the input {"home"} would result in the output
// {"views/home"} if TemplateDir == "views/"
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

// addTemplateExt takes in a slice of strings
// representing the file paths for templates, and it appends
// the TemplateDir directory to each string in the slice
//
// Eg the input {"home"} would result in the output
// {"views/home"} if TemplateDir == "views/"
func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}
