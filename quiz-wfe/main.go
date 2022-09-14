package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
)

const (
	layoutsDir   = "templates/layouts"
	templatesDir = "templates"
	extension    = "/*.html"
)

type Quiz struct {
	QuizName string
	Question string
}

var (
	//go:embed templates/* templates/layouts/*
	files     embed.FS
	templates map[string]*template.Template
)

func LoadTemplates() error {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	tmplFiles, err := fs.ReadDir(files, templatesDir)
	if err != nil {
		return err
	}

	for _, tmpl := range tmplFiles {
		if tmpl.IsDir() {
			continue
		}

		pt, err := template.ParseFS(files, templatesDir+"/"+tmpl.Name(), layoutsDir+extension)
		if err != nil {
			return err
		}

		templates[tmpl.Name()] = pt
	}
	return nil
}

const index = "index.html"

func indexLoad(w http.ResponseWriter, r *http.Request) {
	t, ok := templates[index]
	if !ok {
		// TODO handle error
		return
	}

	data := Quiz{
		QuizName: "ACE Quiz",
	}

	if err := t.Execute(w, data); err != nil {
		// TODO handle error
	}
}
