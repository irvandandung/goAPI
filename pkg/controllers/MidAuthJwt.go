package controllers

import (
	"net/http"
	"log"
)

func MidAuthJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login" {
            next.ServeHTTP(w, r)
            return
        }
        
        log.Println(r.RequestURI)

        next.ServeHTTP(w, r)
    })
}