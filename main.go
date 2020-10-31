package main

import (
	"github.com/irvandandung/goAPI/config"
	"log"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	insert, err := db.Query("insert into user (username, password) values ('dandung', '"+config.GetMD5Hash("admin")+"')")
	if err != nil {
		log.Fatal(err)
	}else{
		log.Println("input success")
	}

	defer insert.Close()
}