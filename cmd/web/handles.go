package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./ui/html/home.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	err = t.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./ui/html/404.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	err = t.ExecuteTemplate(w, "404.html", nil)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
}

func news(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./ui/html/news.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	err = t.ExecuteTemplate(w, "news.html", nil)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
}
