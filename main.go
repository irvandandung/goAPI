package main

import (
	"github.com/irvandandung/goAPI/pkg/data"
	"log"
)

func main() {
	user := data.GetAllDataUsers()
	for key, val := range user {
		log.Println(val.Username)
		log.Println(val.Password)
		log.Println(key)
	}

	bukus := data.GetAllDataBuku()
	for key, val := range bukus {
		log.Println(val)
		log.Println(key)	
	}

	buku := data.GetDataBukuById(1)
	log.Println(buku)
}