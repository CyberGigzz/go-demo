package controllers

import (
	"fmt"
	"net/http"

	"github.com/CyberGigzz/go-demo/models"
)

type Users struct {
	Templates struct {
		New Template
		Signin Template
	}
	UserService  *models.UserService

}

func (u Users) New(w http.ResponseWriter, req *http.Request) {
	var data struct {
		Email string
	}
	data.Email = req.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")
	user, err := u.UserService.Create(data.Email, data.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User Created %+v", user)
}

func (u Users) Signin(w http.ResponseWriter, req *http.Request) {
	var data struct {
		Email string
	}
	data.Email = req.FormValue("email")
	u.Templates.Signin.Execute(w, data)
}
