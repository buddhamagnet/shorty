package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/buddhamagnet/shorty/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Ping).Methods("GET")
	r.HandleFunc("/shorten", handlers.Shortener).Methods("POST")
	r.HandleFunc("/{id:[a-zA-Z0-9]{6}}", handlers.Redirector).Methods("GET")
	http.Handle("/", r)
	fmt.Printf("shorty running on port %s\n", os.Getenv("SHORTENER_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("SHORTENER_PORT"), nil))
}
