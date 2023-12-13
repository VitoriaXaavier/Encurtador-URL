package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	//"io/ioutil"
	//"log"
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
		   rows, err := db.Query("SELECT ID, OriginalURL, Contador, URLEncurtada FROM URLCurta")
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
	urls, err := urlManager.GetURL()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// var formattedURLs []map[string]string
    // for _, url := range urls {
    //     formattedURL := map[string]string{
    //         "OriginalURL": url.OriginalURL,
    //         "ShortenedURL": urlMap[url.ID], 
    //     }
    //     formattedURLs = append(formattedURLs, formattedURL)
    // }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urls)
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
	  _, err = db.Exec("INSERT INTO url (ID, OriginalURL, Contador, URLEncurtada) VALUES ($1, $2, $3, $4)", url.ID, url.OriginalURL, url.Contador, url.URLEncurtada)
	  if err != nil {
			  log.Fatal(err)
		  }
  
		  rows, err := db.Query(`SELECT ID, OriginalURL, Contador, URLEncurtada FROM URLCurta`)
		  if err != nil {
			  log.Fatal(err)
		  }
		  defer rows.Close()
	
	var request struct {
		OriginalURL string `json:"original-url"`
		URLCurta string `json:"url-curta"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newURL := URLCurta{
		OriginalURL: request.OriginalURL,
		URLEncurtada: request.URLCurta,
		ID: generateID(),
		Contador: 0,
	}
	err = urlManager.db.SaveURL(newURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	urlMap[request.URLCurta] = request.OriginalURL

	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(newURL)
}

func DeletURL(w http.ResponseWriter, r *http.Request, urlManager *URLManager) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	err := urlManager.DeleteURL(id)
	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "URL deleted successfuly")
}


// func RedirectURL(w http.ResponseWriter, r *http.Request, urlManager *URLManager) {
// 	redirect := r.URL.Path[1:]
// 	if redirect == "" {
// 		http.Error(w, "Invalid URL", http.StatusBadRequest)
// 		return
// 	}
// 	log.Println("Passando aqui RedirectURL", redirect)

// 	originalURL, ok := urlMap[redirect]
// 	log.Println("Passando aqui RedirectURL", originalURL)
//     if !ok {
//         http.Error(w, "URL not found", http.StatusNotFound)
//         return
//     }
// 	log.Println("Passando aqui RedirectURL", originalURL)

// 	err := urlManager.RecordAccess(redirect)
// 	if err != nil {
// 		fmt.Println("Erro ao registrar acesso", err)
// 	}
// 	urls, err := urlManager.GetOriginalURL(originalURL)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	log.Println("Passando aqui RedirectURL", urls)
	
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(urls)
// 	//http.Redirect(w, r, originalURL, http.StatusFound)

// }

func RedirectURL(w http.ResponseWriter, r *http.Request, urlManager *URLManager) {
	// redirect := r.URL.Path[1:]
	// if redirect == "" {
	// 	http.Error(w, "Invalid URL", http.StatusBadRequest)
	// 	return
	// }
	// log.Println("Passando aqui RedirectURL", redirect)

	var request struct {
	 	originalURL URLCurta
	}
	url, err := urlManager.GetURL()
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		   return
	}
	if request.originalURL.OriginalURL == request.originalURL.URLEncurtada {
		log.Println("passando aqui", request.originalURL.OriginalURL, request.originalURL.URLEncurtada)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(url)
	}
	// originalURL, _ := urlMap[redirect]
	// log.Println("Passando aqui RedirectURL", originalURL)
    // // if !ok {
    // //     http.Error(w, "URL not found", http.StatusNotFound)
    // //     return
    // // }
	// log.Println("Passando aqui RedirectURL", originalURL)

	// // Realiza uma requisição GET para a URL original
	// resp, err := http.Get(originalURL)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// defer resp.Body.Close()

	// // Lê o conteúdo da resposta
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// // Escreve o conteúdo da URL original como resposta
	// w.Header().Set("Content-Type", "text/plain") // Altere o Content-Type conforme necessário
	// w.WriteHeader(resp.StatusCode)
	// w.Write(body)
}
