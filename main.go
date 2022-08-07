package main

import (
	"net/http"

	"github.com/Erickype/GoGormRestApi/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", router)
}
