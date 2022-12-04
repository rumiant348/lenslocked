package main

import (
	"fmt"
	"lenslocked.com/controllers"
	"lenslocked.com/middleware"
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

	// router init
	r := mux.NewRouter()

	// controllers init
	usersC := controllers.NewUsers(services.User)
	staticC := controllers.NewStatic()
	galleriesC := controllers.NewGalleries(services.Gallery, services.Image, r)

	// logging middleware?
	logging := middleware.LogRequest{}
	// cors middleware?

	// user middleware
	userMw := middleware.User{
		UserService: services.User,
	}
	// auth middleware
	requireUserMw := middleware.RequireUser{}

	newGallery := requireUserMw.Apply(galleriesC.NewView)
	createGallery := requireUserMw.ApplyFn(galleriesC.Create)
	editGallery := requireUserMw.ApplyFn(galleriesC.Edit)
	updateGallery := requireUserMw.ApplyFn(galleriesC.Update)
	deleteGallery := requireUserMw.ApplyFn(galleriesC.Delete)
	indexGallery := requireUserMw.ApplyFn(galleriesC.Index)
	imageUpload := requireUserMw.ApplyFn(galleriesC.ImageUpload)

	// Static routes
	r.Handle("/", logging.Apply(staticC.Home)).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")
	r.NotFoundHandler = http.Handler(staticC.NotFoundView)
	// User routes
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	r.HandleFunc("/cookietest", usersC.CookieTest).Methods("GET")
	// Gallery routes
	r.Handle("/galleries/new", newGallery).Methods("GET")
	r.HandleFunc("/galleries", createGallery).Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}", galleriesC.Show).Methods("GET").
		Name(controllers.ShowGallery)
	r.HandleFunc("/galleries/{id:[0-9]+}/edit", editGallery).Methods("GET").
		Name(controllers.EditGallery)
	r.HandleFunc("/galleries/{id:[0-9]+}/update", updateGallery).Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}/delete", deleteGallery).Methods("POST")
	r.Handle("/galleries", indexGallery).Methods("GET").
		Name(controllers.IndexGalleries)
	r.HandleFunc("/galleries/{id:[0-9]+}/images", imageUpload).Methods("POST")

	// Image routes
	imageHandler := http.FileServer(http.Dir("./images"))
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", imageHandler))

	err = http.ListenAndServe(":3000", userMw.Apply(r))
	if err != nil {
		panic(err)
	}
}
