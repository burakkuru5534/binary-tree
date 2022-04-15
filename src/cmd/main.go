package main

import (
	"example.com/m/src/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"


)

func main () {

	//router
	r := mux.NewRouter()
	//api endpoints
	r.Handle("/binary/tree",api.BinaryTreeHandler())

	//define options
	corsWrapper := cors.New(cors.Options{
	AllowedMethods: []string{"GET", "POST"},
	AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	//start server
	log.Fatal(http.ListenAndServe(":8081", corsWrapper.Handler(r)))
}

