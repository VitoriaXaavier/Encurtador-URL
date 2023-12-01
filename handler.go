package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var ErrURLNotFound = errors.New("URL not found")
var urlMap = make(map[string]string)

func ListURL(w http.ResponseWriter, r *http.Request, urlManager *URLManager) {
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
}

func CreateURL(w http.ResponseWriter, r *http.Request, urlManager *URLManager) {
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
	err := urlManager.db.SaveURL(newURL)
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


func RedirectURL(w http.ResponseWriter, r *http.Request, urlManager *URLManager) {
	redirect := r.URL.Path[1:]
	if redirect == "" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	originalURL, ok := urlMap[redirect]
    if !ok {
        http.Error(w, "URL not found", http.StatusNotFound)
        return
    }

	err := urlManager.RecordAccess(redirect)
	if err != nil {
		fmt.Println("Erro ao registrar acesso", err)
	}
	urls, err := urlManager.RedirectURLCurta(originalURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urls)
	http.Redirect(w, r, originalURL, http.StatusFound)

}

