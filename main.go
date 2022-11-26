package main

import (
	"fmt"
	"lenslocked.com/controllers"
	"lenslocked.com/models"
	"lenslocked.com/views"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	homeView     *views.View
	contactView  *views.View
	notFoundView *views.View
	faqView      *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(faqView.Render(w, nil))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	must(notFoundView.Render(w, nil))
}

// A helper function that panics on any error
func must(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	host     = "localhost"
	port     = 5432
	userName = "aru"
	password = ""
	dbname   = "lenslocked_dev"
)

func main() {
	// db init
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, userName, password, dbname)

	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.AutoMigrate()
	//us.DestructiveReset()

	// controllers init
	users := controllers.NewUsers(us)
	static := controllers.NewStatic()
	galleries := controllers.NewGalleries()

	// routes init
	r := mux.NewRouter()
	r.Handle("/", static.Home).Methods("GET")
	r.Handle("/contact", static.Contact).Methods("GET")
	r.Handle("/faq", static.Faq).Methods("GET")
	r.HandleFunc("/signup", users.New).Methods("GET")
	r.HandleFunc("/signup", users.Create).Methods("POST")
	r.Handle("/login", users.LoginView).Methods("GET")
	r.HandleFunc("/login", users.Login).Methods("POST")
	r.HandleFunc("/galleries/new", galleries.New).Methods("GET")
	r.NotFoundHandler = http.Handler(static.NotFoundView)
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
