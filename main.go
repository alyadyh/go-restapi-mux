package main

import (
	"log"
	"net/http"

	"github.com/alyadyh/go-restapi-mux/controllers/studentcontroller"
	"github.com/alyadyh/go-restapi-mux/models"
	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/students", studentcontroller.Index).Methods("GET")
	r.HandleFunc("/students/{nim}", studentcontroller.Show).Methods("GET")
	r.HandleFunc("/students", studentcontroller.Create).Methods("POST")
	r.HandleFunc("/students/{nim}", studentcontroller.Update).Methods("PUT")
	r.HandleFunc("/students/{nim}", studentcontroller.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8082", r))
}
