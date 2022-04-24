package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	// t, err := template.ParseFiles("app/views/templates/top.html")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// t.Execute(w, nil)

	// 共通化したlayoutを用いてviewを表示する（"layout"と"top"はこの時点でtemplateフォルダ内にあるhtmlファイルの差している）
	// generateHTML(w, nil, "layout", "top")

	// generateHTMLの第二引数にnilではなく任意のものを設置してviewに渡す場合
	// generateHTML(w, "Hello World!!!!!", "layout", "public_navbar", "top")

	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "private_navbar", "index")
	} else {
		http.Redirect(w, r, "/todos", 302)

	}
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
}
