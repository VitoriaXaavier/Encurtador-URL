package main

import (
	"log"
	"net/http"
)

func main() {
    db := NewDatabase()
    urlManager := NewURLManager(db)

    http.HandleFunc("/urls", func(w http.ResponseWriter, r *http.Request) {
        ListURL(w, r, urlManager) })
    http.HandleFunc("/cria", func(w http.ResponseWriter, r *http.Request) {
        CreateURL(w, r, urlManager) })
    http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
        RedirectURL(w,r, urlManager) })
    http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
        DeletURL(w, r, urlManager) })

    log.Println("Servidor rodando em http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
