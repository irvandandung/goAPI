package local

import (
	"database/sql"
	"strings"
	"log"
	"sort"
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

func QuerySelect(db *sql.DB, table string, data []string, wheredata map[int]string) (*sql.Rows) {
	query := "select "+strings.Join(data, ", ")+" from "+table
	if len(wheredata) != 0 {
		wheres := []string{}
		keyswhere := make([]int, 0, len(wheredata))
		for key := range wheredata {
			keyswhere = append(keyswhere, key)

		}
		sort.Ints(keyswhere)
		for _, k := range keyswhere{
			wheres = append(wheres, wheredata[k])
		}
		query = query+" where "+strings.Join(wheres, " ")
	}
	rows, err := db.Query(query)
	if (err != nil){
		log.Fatal(err)
	}
	return rows
}