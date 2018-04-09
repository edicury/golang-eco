package main

import (
	"log"
	"fmt"
)

func main(){
	var db = getDB()
	rows, err := db.Query("SELECT nome FROM usuario")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var nome string
		if err := rows.Scan(&nome); err != nil {
			log.Fatal(err)
		}
		fmt.Print("Users:")
		fmt.Println("Name: ", nome)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

