package main

import ("github.com/satori/go.uuid"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
)


type Usuario struct {
	UsuarioId uuid.UUID  `json:"usuario_id"`
	Nome   string      `json:"nome"`
	Idade int          `json:"idade"`
	Usuario string     `json:"usuario"`
}


func (i *Impl) GetAllUsuario(w http.ResponseWriter, r *http.Request) {
	usuario := []Usuario{}
	i.DB.Find(&usuario)
	json.NewEncoder(w).Encode(&usuario)
}

func (i *Impl) GetUsuario(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	var usuario Usuario

	i.DB.First(&usuario, "usuario = ?", params["id"])
	json.NewEncoder(w).Encode(&usuario)
}


func (i *Impl) PostCadastro(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)

	var params Usuario
	err := decoder.Decode(&params)
	if err != nil {
		panic(err)
	}


	var usuario Usuario
	_ = json.NewDecoder(r.Body).Decode(&usuario)
	u1, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}

	usuario.UsuarioId = u1
	usuario.Nome = params.Nome
	usuario.Idade = params.Idade
	usuario.Usuario = params.Usuario

	i.DB.Create(&usuario)
}

func (i *Impl) DeleteUsuario(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	var usuario Usuario

	i.DB.Delete(&usuario,"usuario = ?", params["id"])
}
