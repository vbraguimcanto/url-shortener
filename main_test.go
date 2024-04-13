package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShortenHandler(t *testing.T) {
	storage := NewMemoryStorage()
	shortener := &Shortener{storage: storage}

	payload := []byte(`{"url": "https://www.google.com"}`)
	req, err := http.NewRequest("POST", "/shorten", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(shortener.ShortenHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var resp struct {
		ShortLink string `json:"short_link"`
	}
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.ShortLink) != 5 {
		t.Errorf("handler returned unexpected short link: got %s want length of 5", resp.ShortLink)
	}
}

func TestRedirectHandler(t *testing.T) {
	storage := NewMemoryStorage()
	redirector := &Redirector{storage: storage}

	shortLink := "abcde"
	originalURL := "https://www.google.com"
	storage.Save(shortLink, originalURL)

	req, err := http.NewRequest("GET", "/"+shortLink, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(redirector.RedirectHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusSeeOther)
	}

	location := rr.Header().Get("Location")
	if location != originalURL {
		t.Errorf("handler returned wrong location header: got %s want %s",
			location, originalURL)
	}
}

func TestMemoryStorage_Save_Get(t *testing.T) {
	storage := NewMemoryStorage()

	shortLink := "abcde"
	originalURL := "https://www.google.com"

	err := storage.Save(shortLink, originalURL)
	if err != nil {
		t.Fatal(err)
	}

	url, err := storage.Get(shortLink)
	if err != nil {
		t.Fatal(err)
	}

	if url != originalURL {
		t.Errorf("stored URL doesn't match original URL: got %s want %s",
			url, originalURL)
	}
}

func TestMemoryStorage_IncrementRedirectCount(t *testing.T) {
	storage := NewMemoryStorage()

	shortLink := "abcde"

	storage.IncrementRedirectCount(shortLink)

	count := storage.redirectCounts[shortLink]
	if count != 1 {
		t.Errorf("redirect count doesn't match: got %d want %d", count, 1)
	}
}