package main

import (
	"log"
	"net/http"
	"strings"
   // "database/sql"
	_ "github.com/lib/pq"
)


func main() {
  
    
    db1 := NewDatabase()
    urlManager := NewURLManager(db1)

    http.HandleFunc("/urls", func(w http.ResponseWriter, r *http.Request) {
        ListURL(w, r, urlManager) })
    http.HandleFunc("/cria", func(w http.ResponseWriter, r *http.Request) {
        CreateURL(w, r, urlManager) })
    // http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
    //     RedirectURL(w,r, urlManager) })
    http.HandleFunc("/redirect/", handleRedirect)
    http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
        DeleteURL(w, r, urlManager) })

    log.Println("Servidor rodando em http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}


func handleRedirect(w http.ResponseWriter, r *http.Request) {
    db := NewDatabase()
    urlManager := NewURLManager(db)
    redirect := r.URL.Path[len("/redirect/"):]
    segments := strings.Split(redirect, "/")
    if len(segments) > 0 {
       // shortURL := segments[0]

        // Supondo que urlManager seja uma instância de URLManager que você já tenha criado anteriormente
        RedirectURL(w, r, urlManager)
    } else {
        http.Error(w, "Invalid URL", http.StatusBadRequest)
    }
}
