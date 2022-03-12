package main

import (
	"go-crud-rest-api/controllers"
	"go-crud-rest-api/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "",
			DB:         "go_db",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/users", controllers.GetAllUser).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUserByID).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeletUserByID).Methods("DELETE")
}

func main() {
	initDB()

	log.Println("Starting the HTTP server on http://localhost:3000")
	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":3000", router))
}
