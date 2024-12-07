package main

import (
    "fmt"
    "html/template"
    "net/http"
)

type PageData struct {
    Message string
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        message := r.FormValue("message")
        data := PageData{
            Message: message,
        }
        tmpl := template.Must(template.ParseFiles("index.html"))
        tmpl.Execute(w, data)
    })

    fmt.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}
