package main

import (
	"github.com/joho/godotenv"
    "os"
    "log"
	// "github.com/irvandandung/goAPI/config"
	// "net/http"
	// "github.com/gorilla/mux"
	// "reflect"
)

func main() {
	if errLoadDotenv := godotenv.Load(); errLoadDotenv != nil {
		log.Println(errLoadDotenv.Error())
		os.Exit(1)
	}
	log.Println(os.Getenv("HOST_DATABASE"))
	// r := mux.NewRouter()

 //    log.Println(reflect.TypeOf(r))
 //    config.Routes(r)

 //    log.Fatal(http.ListenAndServe(":1234", r))
}