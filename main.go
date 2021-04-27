package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	r.HandleFunc("/assets", get_list_assets).Methods("GET")
	r.HandleFunc("/assets", post_asset).Methods("POST")

	r.HandleFunc("/coins", get_list_coins).Methods("GET")
	r.HandleFunc("/coins", post_coin).Methods("POST")

	if err := http.ListenAndServe(":5555", handlers.CORS(headers, methods, origins)(r)); err != nil {
		log.Fatal(err)
	}

}
