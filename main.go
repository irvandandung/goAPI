package main

import (
    "log"
	"github.com/irvandandung/goAPI/pkg/controllers"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	//User
    r.HandleFunc("/login", controllers.Login)
	r.HandleFunc("/user/myprofile", controllers.GetMyDataProfile)
	r.HandleFunc("/user/all-data-user", controllers.GetAllDataUser)

	//Book
	r.HandleFunc("/book/addimage", controllers.SubmitPhoto)
	r.HandleFunc("/book/all-data-book", controllers.GetAllDataBook)
	r.HandleFunc("/book/get-detail-book", controllers.GetDataBook)

	//Auth Jwt
	r.Use(controllers.MidAuthJwt)

    log.Fatal(http.ListenAndServe(":1234", r))
}