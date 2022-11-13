package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"todo_app/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	templateParseFiles, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatalln(err)
	}
	templates := template.Must(templateParseFiles, nil)
	fmt.Println(templateParseFiles)
	templates.ExecuteTemplate(w, "layout", data)
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
