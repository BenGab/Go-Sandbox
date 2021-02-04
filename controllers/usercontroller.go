package controllers

import (
	"net/http"
	"regexp"
)

type usercontroller struct {
	IdPattern *regexp.Regexp
}

func (uc usercontroller) seveHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from user controller"))
}

func newusercontroller() *usercontroller {
	return &usercontroller{
		IdPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
