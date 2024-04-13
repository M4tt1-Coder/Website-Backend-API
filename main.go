package main

import (
	"log"
	"net/http"

	"github.com/M4tt1-Coder/business/portfolio_website/API_GO/authenticator"
	"github.com/M4tt1-Coder/business/portfolio_website/API_GO/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//TODO - http request should only need one attemp to get the requested resource

func main() {
	r := mux.NewRouter()
	routes.AllRoutes(r)
	r.Use(authenticator.Authenticate)
	http.Handle("/", r)

	//a notification that the server is running
	println("Listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedOrigins([]string{"localhost:5173"}), //change that to "matthisgeissler.com"
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}),
	)(r)))
	//add a authentication
	//redirect to https
}
