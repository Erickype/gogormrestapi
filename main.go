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

	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")

	http.ListenAndServe(":3000", router)
}
