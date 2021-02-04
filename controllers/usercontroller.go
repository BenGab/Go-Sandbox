package controllers

import "net/http"

type usercontroller struct{}

func (uc usercontroller) seveHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from user controller"))
}
