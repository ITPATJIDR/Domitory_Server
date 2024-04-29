package main

import (
	"Domitory_Server/database"
	"Domitory_Server/infrastructure"
)

func main() {
	infrastructure.InitENV()
	s, err := database.NewGormStore()

	if err != nil {
		panic("Can't Connect to Database")
	}

	infrastructure.Dispatch(*s)
}
