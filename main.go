package main

import (
    "net/http"
    "html/template"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        t, err := template.ParseFiles("index.html")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        t.Execute(w, nil)
    })

    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":8080", nil)
}
