package main

import (
	"log"
	"github.com/google/uuid"
)

type URLCurta struct {
	ID           string `json:"id"`
	OriginalURL  string `json:"originalURL"`
	Contador     int    `json:"contador"`
	URLEncurtada string `json:"urlEncurtada"`
}

type URLManager struct {
	db *Database
}

func NewURLManager(db *Database) *URLManager{
	return &URLManager{db: db}
}

func (u *URLManager) CreateURL(originalURL string) (*URLCurta, error) {
	id := generateID()

	newURL := URLCurta{
		ID: id,
		URLEncurtada: originalURL,
		Contador: 0,
	}
	log.Println("Passando aqui CreateURL ", newURL)

	err := u.db.SaveURL(newURL)
	if err != nil {
		return nil, err
	}
	log.Println("Passando aqui  CreateURL", newURL)
	return &newURL, nil
}
func (u *URLManager) GetURLCurta(curtaURL string) (URLCurta, error){
	url, err := u.db.GetURL(curtaURL)
	log.Println("Passando aqui  GetURLCurta", curtaURL)
	if err != nil {
		return url, err
	}
	log.Println("Passando aqui GetURLCurta", url)
	return url, nil
}

func (u *URLManager) GetURL() ([]URLCurta, error) {
	return u.db.GetAllURLs()
}

func (u *URLManager) DeleteURL(id string) error {
	return u.db.DeleteURL(id)
}

func (u *URLManager) GetOriginalURL(id string) (string, error) {
	url, err := u.db.GetURL(id)
	if err != nil {
		return " ", err
	}
	log.Println("Passando aqui GetOriginalURL", url)
	log.Println("Passando aqui GetOriginalURL original", url.OriginalURL)
	log.Println("Passando aqui GetOriginalURL encurtada", url.URLEncurtada)
	return url.OriginalURL, nil
}

func (u *URLManager) RecordAccess(id string) error {
	return u.db.IncrementAccessCount(id)
}

func generateID() string {
	return uuid.New().String()
}