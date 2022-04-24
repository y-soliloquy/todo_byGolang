package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"todo_bygolang/app/models"
	"todo_bygolang/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))

		templates := template.Must(template.ParseFiles(files...))
		templates.ExecuteTemplate(w, "layout", data)
	}
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}

	return sess, err
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	// http.HandleFanc("URLのパス", route_auth.goなどのファイルで設定したパスに関する関数)
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
