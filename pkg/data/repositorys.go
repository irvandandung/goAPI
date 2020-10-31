package data

import (
	"github.com/irvandandung/goAPI/config"
	"github.com/irvandandung/goAPI/pkg/data/local"
	"log"
)

func InsertDataUser(data map[string]string) (string, error){
	db := config.ConnectDB()
	defer db.Close()
	response, err := local.QueryInsert(db, "user", data)
	return response, err
}

func UpdateDataUser(data map[string]string, wheredata map[string]string) (string, error){
	db := config.ConnectDB()
	defer db.Close()
	response, err := local.QueryUpdate(db, "user", data, wheredata)
	return response, err
}

func GetAllDataUsers() ([]Users){
	var user Users
	var list_user []Users
	db := config.ConnectDB()
	defer db.Close()
	var fields = []string{ "username", "password" }
	rows := local.QuerySelect(db, "user", fields)
	for rows.Next() {
		err := rows.Scan(&user.Username, &user.Password)
		if(err != nil){
			log.Fatal(err.Error())
		}
		list_user = append(list_user, user)
	}

	return list_user
}