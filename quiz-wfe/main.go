package main

import (
	"html/template"
	"net/http"
)

type Quiz struct {
	QuizName string
	Question string
}

func main() {
	tmpl := template.Must(template.ParseFiles("template.html"))

	//Add code to serve CSS
	http.Handle("/static/",
		http.StripPrefix("/static",
			http.FileServer(http.Dir("static"))))
	//Serve Site based on Data
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := Quiz{
			QuizName: "ACE Quiz",
		}
		tmpl.Execute(w, data)
	})
	//Actually Listen on Port 8080
	http.ListenAndServe(":8080", nil)
}
