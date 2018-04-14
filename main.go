package main

import (
	_ "github.com/lib/pq"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	i := Impl{}
	i.InitDB()
	i.InitSchema()
	i.Seed()


	router := mux.NewRouter()
	router.HandleFunc("/usuario", i.GetAllUsuario).Methods("GET")
	router.HandleFunc("/cadastro", i.PostCadastro).Methods("POST")
	router.HandleFunc("/usuario/{id}", i.GetUsuario).Methods("GET")
	router.HandleFunc("/usuario/{id}", i.DeleteUsuario).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", router))
}





