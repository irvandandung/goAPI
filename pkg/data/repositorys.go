package data

import (
	"github.com/irvandandung/goAPI/config"
	"github.com/irvandandung/goAPI/pkg/data/local"
	"log"
	"strconv"
)

func InsertDataUser(username string, password string) (string){
	data := map[string]string{
		"username" : username,
		"password" : config.GetMD5Hash(password),
		"role"     : "basic",
	}
	db := config.ConnectDB()
	defer db.Close()
	response, err := local.QueryInsert(db, "user", data)
	if(err != nil){
		log.Fatal(err.Error())
	}
	return response
}

func InsertGambarBuku(namegambar string) (string){
	data := map[string]string{"nama_gambar" : namegambar}
	db := config.ConnectDB()
	defer db.Close()
	response, err := local.QueryInsert(db, "list_gambar_buku", data)
	if(err != nil){
		log.Fatal(err.Error())
	}
	return response
} 

func InsertDataBuku(book Buku) (string){
	data := map[string]string{
		"judul" : book.Judul,
		"keterangan" : book.Keterangan,
		"pencipta" : book.Pencipta,
		"nama_gambar" : book.Nama_gambar,
		"tahun" : book.Tahun,
	}
	db := config.ConnectDB()
	defer db.Close()
	response, err := local.QueryInsert(db, "buku", data)
	if(err != nil){
		log.Fatal(err.Error())
	}
	return response
}

func UpdateDataUser(data map[string]string, wheredata map[string]string) (string, error){
	db := config.ConnectDB()
	defer db.Close()
	response, err := local.QueryUpdate(db, "user", data, wheredata)
	return response, err
}

func UpdateDataBuku(book Buku, id int) (string){
	data := map[string]string{
		"judul" : "'"+book.Judul+"'",
		"keterangan" : "'"+book.Keterangan+"'",
		"pencipta" : "'"+book.Pencipta+"'",
		"nama_gambar" : "'"+book.Nama_gambar+"'",
		"tahun" : "'"+book.Tahun+"'",
	}
	wheredata := map[string]string{
		"id" : strconv.Itoa(id),
	}
	db := config.ConnectDB()
	defer db.Close()
	response, err := local.QueryUpdate(db, "buku", data, wheredata)
	if(err != nil){
		log.Fatal(err.Error())
	}
	return response
}

func GetAllDataUsers() ([]Users){
	var user Users
	var list_user []Users
	wheredata := map[int]string{}
	db := config.ConnectDB()
	defer db.Close()
	var fields = []string{ "id", "username", "role" }
	rows := local.QuerySelect(db, "user", fields, wheredata)
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Role)
		if(err != nil){
			log.Fatal(err.Error())
		}
		list_user = append(list_user, user)
	}

	return list_user
}

func GetDataUser(username string, password string) (bool, Users){
	var user Users
	data := false
	wheredata := map[int]string{0:"username='"+username+"' AND", 1:"password='"+config.GetMD5Hash(password)+"'"}
	db := config.ConnectDB()
	defer db.Close()
	var fields = []string{"id", "username", "role"}
	rows := local.QuerySelect(db, "user", fields, wheredata)
	for rows.Next(){
		err := rows.Scan(&user.Id, &user.Username, &user.Role)
		if(err != nil){
			log.Fatal(err.Error())
		}
	}
	if(user.Username != ""){
		data = true
	}

	return data, user
}

func GetAllDataBuku() ([]Buku){
	var buku Buku
	var list_buku []Buku
	wheredata := map[int]string{}
	db := config.ConnectDB()
	defer db.Close()
	var fields = []string{ "id", "judul"}
	rows := local.QuerySelect(db, "buku", fields, wheredata)
	for rows.Next() {
		err := rows.Scan(&buku.Id, &buku.Judul)
		if(err != nil){
			log.Fatal(err.Error())
		}
		list_buku = append(list_buku, buku)
	}

	return list_buku
}

func GetDataBukuById(id int) (Buku) {
	var buku Buku
	idString := strconv.Itoa(id)
	wheredata := map[int]string{ 0:"id="+idString}
	db := config.ConnectDB()
	defer db.Close()
	var fields = []string{ "id", "judul", "keterangan", "pencipta", "nama_gambar", "tahun" }
	rows := local.QuerySelect(db, "buku", fields, wheredata)
	for rows.Next() {
		err := rows.Scan(&buku.Id, &buku.Judul, &buku.Keterangan, &buku.Pencipta, &buku.Nama_gambar, &buku.Tahun)
		if(err != nil){
			log.Fatal(err.Error())
		}
	}

	return buku
}