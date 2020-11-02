package config

import(
	"github.com/gorilla/mux"
	"github.com/irvandandung/goAPI/pkg/controllers"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/", controllers.YourHandler)
}