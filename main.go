package main

import (
	"net/http"

	"github.com/Erickype/GoGormRestApi/db"
	"github.com/Erickype/GoGormRestApi/models"
	"github.com/Erickype/GoGormRestApi/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConecction()

	//crear modelos
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	//router
	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler)

	//user handler functions
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users", routes.DeleteUserHandler).Methods("DELETE")

	//Server start
	http.ListenAndServe(":3000", router)
}
