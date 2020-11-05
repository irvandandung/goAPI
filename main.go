package main

import (
    "log"
	"github.com/irvandandung/goAPI/pkg/controllers"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
    r.HandleFunc("/login", controllers.Login)
	r.HandleFunc("/", controllers.YourHandler)
	r.HandleFunc("/user/myprofile", controllers.GetMyDataProfile)
	r.HandleFunc("/user/all-data-user", controllers.GetAllDataUser)
	r.Use(controllers.MidAuthJwt)

    log.Fatal(http.ListenAndServe(":1234", r))
}