package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/markbrown87/prototype-dev-station/helpers"
)

type PageData struct {
	Message string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var message string
		if r.Method == http.MethodPost {
			action := r.FormValue("action")
			if action == "Deploy Git" {
				helpers.DeployGit()
				message = "Deploy Git button pressed"
			} else if action == "Deploy Webtop" {
				helpers.DeployWebtop()
				message = "Deploy Webtop button pressed"
			}
		}
		data := PageData{
			Message: message,
		}
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, data)
	})

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
