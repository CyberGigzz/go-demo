package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, req *http.Request) {
	var data struct {
		Email string
	}
	data.Email = req.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Email: ", req.FormValue("email"))
	fmt.Fprint(w, "Password: ", req.FormValue("password"))
}
