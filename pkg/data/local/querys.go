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

func QueryUpdate(db *sql.DB, table string, data map[string]string, wheredata map[string]string) (string, error) {
	sets := []string{}
	wheres := []string{}
	for key, value := range data {
		sets = append(sets, key+" = "+value)
	}
	for key, value := range wheredata{
		wheres = append(wheres, key+" = "+value)
	}
    log.Println("update "+table+" set "+strings.Join(sets, ", ")+" where "+strings.Join(wheres, ", "))
	update, err := db.Query("update "+table+" set "+strings.Join(sets, ", ")+" where "+strings.Join(wheres, ", "))
	if err != nil {
		log.Fatal(err)
		return "update gagal", err
	}

	defer update.Close()

	return "update success", nil
}

func QuerySelect(db *sql.DB, table string, data []string, wheredata map[string]string) (*sql.Rows) {
	query := "select "+strings.Join(data, ", ")+" from "+table
	wheres := []string{}
	if len(wheredata) != 0 {
		for key, value := range wheredata {
			wheres = append(wheres, key+value)
		}
		query = query+" where "+strings.Join(wheres, ", ")
	}
	rows, err := db.Query(query)
	if (err != nil){
		log.Fatal(err)
	}
	return rows
}