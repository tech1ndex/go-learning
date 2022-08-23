package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Quiz struct {
	QuizName string
	Question string
}

func main() {
	tmpl := template.Must(template.ParseFiles("template.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := Quiz{
			QuizName: "ACE Quiz",
		}
		tmpl.Execute(w, data)
	})
	fmt.Println("Serving on port 8080 :)")
	http.ListenAndServe(":8080", nil)
}
