package main

/// The same main but with different router as an example

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/julienschmidt/httprouter"
// )

// func notFound(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	w.WriteHeader(http.StatusNotFound)
// 	fmt.Fprint(w, "<h1>We could not find the page you "+
// 		"you were looking for :(</h1>"+
// 		"<p>Please email us if you keep being sent to an "+
// 		"invalid page.</p>")
// }

// func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
// }

// func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, "To get in touch, please send an email "+
// 		"to <a href=\"mailto:support@lenslocked.com\">"+
// 		"support@lenslocked.com</a>.")
// }

// func faq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, "<h1>FAQ</h1>"+
// 		"Welcome to the faq page!")
// }

// func main() {
// 	// r := mux.NewRouter()
// 	// r.NotFoundHandler = http.HandlerFunc(notFound)
// 	// r.HandleFunc("/", home)
// 	// r.HandleFunc("/contact", contact)
// 	// r.HandleFunc("/faq", faq)
// 	// http.ListenAndServe(":3000", r)
// 	router := httprouter.New()
// 	router.GET("/", home)
// 	router.GET("/contact", contact)
// 	router.GET("/faq", faq)
// 	router.NotFound = http.HandlerFunc(notFound)
// 	http.ListenAndServe(":3000", router)
// }
