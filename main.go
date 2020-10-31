package main

import (
	"github.com/irvandandung/goAPI/config"
	"github.com/irvandandung/goAPI/pkg/data"
	"log"
)

func main() {
	dataUser := map[string]string { "username":"dandung", "password":config.GetMD5Hash("admin") }
	response, err := data.InsertDataUser(dataUser)
	if err != nil {
		log.Fatal(err)
	}else{
		log.Println(response)
	}
}