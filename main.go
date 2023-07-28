package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/santiagoHV/go-gorm-restapi/db"
	"github.com/santiagoHV/go-gorm-restapi/models"
	"github.com/santiagoHV/go-gorm-restapi/routes"
)

func main() {
	db.DBconnection()

	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Task{})

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
