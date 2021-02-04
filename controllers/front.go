package controllers

import (
	"net/http"
)

func RegisterControllers() {
	uc := newusercontroller()

	http.Handle("/users", *uc)
}
