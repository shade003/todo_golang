package controllers

import (
	"fmt"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hoge")
	generateHTML(w, "Hello", "layout", "public_navbar", "top")
	// t, err := template.ParseFiles("app/views/templates/top.html")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// t.Execute(w, "hello") // データの引き渡し
}
