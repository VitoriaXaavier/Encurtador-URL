package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var ErrURLNotFound = errors.New("URL not found")
var urlMap = make(map[string]string)

func ListURL(w http.ResponseWriter, r *http.Request, urlManager *URLManager) {
		// Configurar os detalhes da conexão
		connStr := "postgres://vitoria.xavier:root@localhost/url?sslmode=disable"

		// Abrir uma conexão com o banco de dados PostgreSQL
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		   // Executar a consulta SELECT na tabela URLCurta
		   rows, err := db.Query("SELECT * FROM url")
		   if err != nil {
			   http.Error(w, err.Error(), http.StatusInternalServerError)
			   return
		   }
		   defer rows.Close()
	   
		   // Criar uma estrutura para armazenar os resultados
		   var resultados []URLCurta
		   for rows.Next() {
			   var url URLCurta
			   err := rows.Scan(&url.ID, &url.OriginalURL, &url.Contador, &url.URLEncurtada)
			   if err != nil {
				   http.Error(w, err.Error(), http.StatusInternalServerError)
				   return
			   }
			   resultados = append(resultados, url)
		   }
	   
		   // Converter os resultados para JSON
		   jsonResponse, err := json.Marshal(resultados)
		   if err != nil {
			   http.Error(w, err.Error(), http.StatusInternalServerError)
			   return
		   }

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func CreateURL(w http.ResponseWriter, r *http.Request, urlManager *URLManager) {
	var url URLCurta
    err := json.NewDecoder(r.Body).Decode(&url)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
  	// Configurar os detalhes da conexão
	  connStr := "postgres://vitoria.xavier:root@localhost/url?sslmode=disable"

	  // Abrir uma conexão com o banco de dados PostgreSQL
	  db, err := sql.Open("postgres", connStr)
	  if err != nil {
		  log.Fatal(err)
	  }
	  defer db.Close()
  
  
	   // Executar a query de inserção
	  _, err = db.Exec("INSERT INTO url (id, originalurl, contador, urlEncurtada) VALUES ($1, $2, $3, $4)", url.ID, url.OriginalURL, url.Contador, url.URLEncurtada)
	  if err != nil {
			  log.Fatal(err)
		  }
  
		  rows, err := db.Query("SELECT * FROM url where id = '"+ url.ID +"'")
		  if err != nil {
			  log.Fatal(err)
		  }
		  defer rows.Close()
		     // Criar uma estrutura para armazenar os resultados
			 var resultados []URLCurta
			 for rows.Next() {
				 var url URLCurta
				 err := rows.Scan(&url.ID, &url.OriginalURL, &url.Contador, &url.URLEncurtada)
				 if err != nil {
					 http.Error(w, err.Error(), http.StatusInternalServerError)
					 return
				 }
				 resultados = append(resultados, url)
			 }
		 
			 // Converter os resultados para JSON
			 jsonResponse, err := json.Marshal(resultados)
			 if err != nil {
				 http.Error(w, err.Error(), http.StatusInternalServerError)
				 return
			 }
	
	
	w.Header().Set("Content-Type", "application-json")
	w.Write(jsonResponse)
}

func DeleteURL(w http.ResponseWriter, r *http.Request, urlManager *URLManager) {
	var url URLCurta
    err := json.NewDecoder(r.Body).Decode(&url)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	
    // Configurar os detalhes da conexão
    connStr := "postgres://vitoria.xavier:root@localhost/url?sslmode=disable"
	
    // Abrir uma conexão com o banco de dados PostgreSQL
    db, err := sql.Open("postgres", connStr)
    if err != nil {
		log.Fatal(err)
    }
    defer db.Close()
	
	rows, err := db.Query("DELETE FROM url WHERE id = '" + url.ID +"'")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

    // Responder com uma mensagem de sucesso
    w.Header().Set("Content-Type", "application/json")
    response := map[string]string{"message": "URL excluída com sucesso"}
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Write(jsonResponse)
}

func RedirectURL(w http.ResponseWriter, r *http.Request, urlManager *URLManager) {
	var url URLCurta
    err := json.NewDecoder(r.Body).Decode(&url)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	// Configurar os detalhes da conexão
	connStr := "postgres://vitoria.xavier:root@localhost/url?sslmode=disable"

	// Abrir uma conexão com o banco de dados PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM url where urlEncurtada = '"+ url.URLEncurtada +"'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	   // Criar uma estrutura para armazenar os resultados
	   var resultados []URLCurta
	   for rows.Next() {
		   var url URLCurta
		   err := rows.Scan(&url.ID, &url.OriginalURL, &url.Contador, &url.URLEncurtada)
		   if err != nil {
			   http.Error(w, err.Error(), http.StatusInternalServerError)
			   return
		   }
		   resultados = append(resultados, url)
	   }
   
	   // Converter os resultados para JSON
	   jsonResponse, err := json.Marshal(resultados)
	   if err != nil {
		   http.Error(w, err.Error(), http.StatusInternalServerError)
		   return
	   }


w.Header().Set("Content-Type", "application-json")
w.Write(jsonResponse)

}
