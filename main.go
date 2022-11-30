package main

import (
	"fmt"
	"lenslocked.com/controllers"
	"lenslocked.com/models/users"
	"net/http"

	"github.com/gorilla/mux"
)

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
		"dbname=%s sslmode=disable password=%s",
		host, port, userName, dbname, password)

	us, err := users.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.AutoMigrate()
	//us.DestructiveReset()

	// controllers init
	usersC := controllers.NewUsers(us)
	staticC := controllers.NewStatic()
	galleriesC := controllers.NewGalleries()

	// routes init
	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	r.HandleFunc("/cookietest", usersC.CookieTest).Methods("GET")
	r.HandleFunc("/galleriesC/new", galleriesC.New).Methods("GET")
	r.NotFoundHandler = http.Handler(staticC.NotFoundView)
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
