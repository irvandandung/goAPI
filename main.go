package main

import (
	// "github.com/irvandandung/goAPI/config"
	"github.com/irvandandung/goAPI/pkg/data"
	"log"
)

func main() {
	data := data.GetAllDataUsers()
	for key, val := range data {
		log.Println(val.Username)
		log.Println(val.Password)
		log.Println(key)
	}
}