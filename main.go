package main

import (
	"github.com/irvandandung/goAPI/config"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"reflect"
)

func main() {
	r := mux.NewRouter()

    log.Println(reflect.TypeOf(r))
    config.Routes(r)

    log.Fatal(http.ListenAndServe(":1234", r))
}