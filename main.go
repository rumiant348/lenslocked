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
	r.HandleFunc("/cookietest", users.CookieTest).Methods("GET")
	r.HandleFunc("/galleries/new", galleries.New).Methods("GET")
	r.NotFoundHandler = http.Handler(static.NotFoundView)
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
