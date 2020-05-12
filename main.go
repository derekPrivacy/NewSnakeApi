package main

import (
	db "api/db"
	route "api/route"
	socketRoute "api/route/socket"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	db.DbInit()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})

	router := mux.NewRouter().StrictSlash(true)
	s := router.PathPrefix("/api").Subrouter()

	//test
	s.HandleFunc("/test", route.Test).Methods("GET")

	//user register and login
	s.HandleFunc("/createUser", route.CreateUser).Methods("POST")
	s.HandleFunc("/validateUser", route.ValidateUser).Methods("POST")

	s.HandleFunc("/socket", socketRoute.SocketHandler)
	s.HandleFunc("/socketAdd", socketRoute.SocketAdd)
	//make only add User to websocket connection
	//rest of the message type to http

	port, _ := os.LookupEnv("PORT")

	log.Fatal(http.ListenAndServe(":"+port, c.Handler(router)))
}
