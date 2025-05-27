package auth

import (
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Login endpoint"))
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/login", LoginHandler)
}
