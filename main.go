package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"example.com/goserver/handlers"
)

const PORT = "0.0.0.0:5678"
const REQ_LOGGING = true

// Middleware: Ensures that file listings are not served
func noDirListing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Middleware: Enables request loging on console
func logReq(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if REQ_LOGGING {
			log.Printf("Request: %s\t%s %s", r.RemoteAddr, r.Method, r.RequestURI)
		}
		next.ServeHTTP(w, r)
	})
}

func main() {

	// u := models.User{Id: 123, Username: "John Doe", Email: "john.doe@example.com"}

	// page handlers
	http.Handle("GET /", logReq(http.HandlerFunc(handlers.HomeHandler)))

	http.Handle("GET /home", logReq(http.HandlerFunc(handlers.HomeHandler)))

	http.Handle("GET /contact", logReq(http.HandlerFunc(handlers.ContactHandler)))

	http.Handle("GET /about", logReq(http.HandlerFunc(handlers.AboutHandler)))

	http.Handle("GET /users", logReq(http.HandlerFunc(handlers.GetUsersHandler)))
	http.Handle("POST /users", logReq(http.HandlerFunc(handlers.PostUserHandler)))

	// static file handler
	fs := http.FileServer(http.Dir("./views/static"))
	http.Handle("GET /static/", logReq(noDirListing(http.StripPrefix("/static/", fs))))

	// favicon.ico handler
	http.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./views/static/img/favicon.ico")
	})

	fmt.Println("Starting server on " + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
