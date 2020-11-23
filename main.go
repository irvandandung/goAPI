package main

import (
    "log"
	"github.com/irvandandung/goAPI/pkg/controllers"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

    //image static file
    fs := http.FileServer(http.Dir("./assets/images/"))
    r.PathPrefix("/assets/images/").Handler(http.StripPrefix("/assets/images/", fs))

	//User
    r.HandleFunc("/login", controllers.Login)
	r.HandleFunc("/user/myprofile", controllers.GetMyDataProfile)
	r.HandleFunc("/user/all-data-user", controllers.GetAllDataUser)
	r.HandleFunc("/user/add-data-user", controllers.AddUser)

	//Book
	r.HandleFunc("/book/addimage", controllers.SubmitPhoto)
	r.HandleFunc("/book/all-data-book", controllers.GetAllDataBook)
	r.HandleFunc("/book/get-detail-book", controllers.GetDataBook)
	r.HandleFunc("/book/add-data-book", controllers.AddDataBook)
	r.HandleFunc("/book/update-data-book", controllers.UpdateDataBook)
	
	//Auth Jwt
	r.Use(controllers.MidAuthJwt)

    log.Fatal(http.ListenAndServe(":1234", r))
}