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
	//Print Message to Indicate we are Serving
	fmt.Println("Serving on port 8080 :)")
	http.ListenAndServe(":8080", nil)
}
