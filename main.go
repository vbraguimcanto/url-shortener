package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Shortener struct {
	storage Storage
}

type Redirector struct {
	storage Storage
}

type Storage interface {
	Save(shortLink, originalURL string) error
	Get(shortLink string) (string, error)
	IncrementRedirectCount(shortLink string)
}

type MemoryStorage struct {
	data           map[string]string
	redirectCounts map[string]int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:           make(map[string]string),
		redirectCounts: make(map[string]int),
	}
}

func (ms *MemoryStorage) Save(shortLink, originalURL string) error {
	ms.data[shortLink] = originalURL
	return nil
}

func (ms *MemoryStorage) Get(shortLink string) (string, error) {
	url, exists := ms.data[shortLink]
	if !exists {
		return "", fmt.Errorf("short link not found")
	}
	return url, nil
}

func (ms *MemoryStorage) IncrementRedirectCount(shortLink string) {
	ms.redirectCounts[shortLink]++
}

func generateShortLink() string {
	rand.Seed(time.Now().UnixNano())
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortLink := make([]byte, 5)
	for i := range shortLink {
		shortLink[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortLink)
}

func (s *Shortener) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	shortLink := generateShortLink()

	err = s.storage.Save(shortLink, payload.URL)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	resp := struct {
		ShortLink string `json:"short_link"`
	}{
		ShortLink: shortLink,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (r *Redirector) RedirectHandler(w http.ResponseWriter, req *http.Request) {
	shortLink := strings.TrimPrefix(req.URL.Path, "/")
	fmt.Println("Short link:", shortLink)

	url, err := r.storage.Get(shortLink)
	if err != nil {
		http.Error(w, "Short link not found", http.StatusNotFound)
		return
	}

	r.storage.IncrementRedirectCount(shortLink)

	http.Redirect(w, req, url, http.StatusSeeOther)
}

func main() {
	storage := NewMemoryStorage()
	shortener := &Shortener{storage: storage}
	redirector := &Redirector{storage: storage}

	http.HandleFunc("/shorten", shortener.ShortenHandler)
	http.HandleFunc("/", redirector.RedirectHandler)

	fmt.Println("Server listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
