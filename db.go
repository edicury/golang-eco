package main

import (
	"github.com/jinzhu/gorm"
	"log"
	"github.com/satori/go.uuid"
)

type Impl struct {
	DB *gorm.DB
}

func (i *Impl) InitDB() {
	var err error
	connStr := "postgres://postgres:123quatro@localhost/goeco?sslmode=disable"
	i.DB, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	i.DB.LogMode(true)
}

func (i *Impl) InitSchema() {
	i.DB.AutoMigrate(&Usuario{})
}

func (i *Impl) Seed() {
	var usuario Usuario
	usuario.UsuarioId, _ = uuid.NewV4()
	usuario.Usuario = "bob"
	usuario.Idade = 20
	usuario.Nome = "bob bob"


	i.DB.Create(&usuario)
}