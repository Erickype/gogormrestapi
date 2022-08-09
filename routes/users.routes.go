package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Erickype/GoGormRestApi/db"
	"github.com/Erickype/GoGormRestApi/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {

	//conseguir datos de body
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	log.Printf("Body: %v", r.Body)
	//crear usuario
	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user"))
}
