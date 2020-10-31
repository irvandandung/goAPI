package local

import (
	"database/sql"
	"strings"
	"log"
)

func QueryInsert(db *sql.DB, table string, data map[string]string) (string, error) {
	fields := []string{}
	values := []string{}
	for key, value := range data {
		fields = append(fields, key)
		values = append(values, value)
	}
	insert, err := db.Query("insert into "+table+" ("+strings.Join(fields, ", ")+") values ('"+strings.Join(values, "', '")+"')")
	if err != nil {
		log.Fatal(err)
		return "insert gagal", err
	}

	defer insert.Close()

	return "insert success", nil
}