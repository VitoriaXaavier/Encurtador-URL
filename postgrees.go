package main

import (
    "sync"
)

type Database struct {
    mu   sync.Mutex
    data map[string]URLCurta
}

func NewDatabase() *Database {
    return &Database{
        data: make(map[string]URLCurta),
    }
}

func (db *Database) SaveURL(url URLCurta) error {
    db.mu.Lock()
    defer db.mu.Unlock()

    db.data[url.ID] = url
    return nil
}

func (db *Database) GetAllURLs() ([]URLCurta, error) {
    db.mu.Lock()
    defer db.mu.Unlock()

    var urls []URLCurta
    for _, v := range db.data {
        urls = append(urls, v)
    }
    return urls, nil
}

func (db *Database) GetURL(URLEncurtada string) (URLCurta, error) {
    db.mu.Lock()
    defer db.mu.Unlock()

    url, ok := db.data[URLEncurtada]
    if !ok {
        return URLCurta{}, ErrURLNotFound
    }
    return url, nil
}

func (db *Database) DeleteURL(id string) error {
    db.mu.Lock()
    defer db.mu.Unlock()

    delete(db.data, id)
    return nil
}

func (db *Database) IncrementAccessCount(id string) error {
    db.mu.Lock()
    defer db.mu.Unlock()

    url, ok := db.data[id]
    if !ok {
        return ErrURLNotFound
    }
    url.Contador++
    db.data[id] = url
    return nil
}


