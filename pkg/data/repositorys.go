package data

import (
	"database/sql"
	"github.com/irvandandung/goAPI/config"
	"github.com/irvandandung/goAPI/pkg/data/local"
)

func GetDB() *sql.DB {
	db := config.ConnectDB()
	return db
}

func InsertDataUser(data map[string]string) (string, error){
	db := GetDB()
	defer db.Close()
	response, err := local.QueryInsert(db, "user", data)
	return response, err
}