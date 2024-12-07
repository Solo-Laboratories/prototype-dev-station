package main

import (
	"html/template"
	"net/http"
    "log"

	"github.com/markbrown87/prototype-dev-station/helpers"
)

type PageData struct {
	Message string
}

func googleRedirect(w http.ResponseWriter, r *http.Request) {
	log.Println("Redirecting to google.com...")
    http.Redirect(w, r, "https://www.google.com", http.StatusFound)
}

func main() {
	http.HandleFunc("/google", googleRedirect)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var message string
		if r.Method == http.MethodPost {
			action := r.FormValue("action")
			if action == "Deploy Git" {
				returnUrl := helpers.DeployGit()
				message = "Deploy Git button pressed. Visit: " + returnUrl
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

	log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}