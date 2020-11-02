package controllers

import(
	"net/http"
	"encoding/json"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    // w.Write([]byte("Gorilla!\n"))
    json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}