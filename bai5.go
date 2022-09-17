package main

import (
	"log"
)

type User struct {
}


func CreateUserPartner() {
	err := engine.CreateTables(new(UserPartner))
	if err != nil {
		log.Fatal("Create table error: ", err)
	} else {
		log.Println("Create table successfully!")
	}
}

