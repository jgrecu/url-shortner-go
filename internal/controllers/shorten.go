package controllers

import (
	"database/sql"
	"html/template"
	"net/http"
	"strings"

	"github.com/jgrecu/url-shortner-go/internal/db"
	"github.com/jgrecu/url-shortner-go/internal/url"
)

func Shorten(lite *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		originalURL := r.FormValue("url")
		if originalURL == "" {
			http.Error(w, "url not provided", http.StatusBadRequest)
			return
		}

		if !strings.HasPrefix(originalURL, "http://") && !strings.HasPrefix(originalURL, "https://") {
			originalURL = "https://" + originalURL
		}

		shortURL := url.Shorten(originalURL)

		if err := db.StoreURL(lite, shortURL, originalURL); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := map[string]string{
			"ShortURL": shortURL,
		}

		t, err := template.ParseFiles("internal/views/shorten.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func Proxy(lite *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := r.URL.Path[1:]
		if shortURL == "" {
			http.Error(w, "url not provided", http.StatusBadRequest)
			return
		}
		originalURL, err := db.GetOriginalURL(lite, shortURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, originalURL, http.StatusPermanentRedirect)
	}
}
