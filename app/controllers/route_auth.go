package controllers

import (
	"log"
	"net/http"
	"todo_bygolang/app/models"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		generateHTML(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}

		if err := user.CreateUser(); err != nil {
			log.Fatalln(err)
		}

		http.Redirect(w, r, "/", 302)

	}
}
func login(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "login")
}
