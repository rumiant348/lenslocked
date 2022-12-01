package main

import (
	"fmt"
	"lenslocked.com/controllers"
	"lenslocked.com/models"
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

	//us, err := models.NewUserService(psqlInfo)
	services, err := models.NewServices(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer services.Close()
	services.AutoMigrate()
	//services.DestructiveReset()

	// controllers init
	usersC := controllers.NewUsers(services.User)
	staticC := controllers.NewStatic()
	galleriesC := controllers.NewGalleries(services.Gallery)

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
	r.Handle("/galleries/new", galleriesC.NewView).Methods("GET")
	r.HandleFunc("/galleries", galleriesC.Create).Methods("POST")
	r.NotFoundHandler = http.Handler(staticC.NotFoundView)
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
