package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (m *ruoter) redirect(w http.ResponseWriter, r *http.Request) {
	shortCode := chi.URLParam(r, "shortCode")
	if shortCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("shortCode is required"))
		return
	}

	// Retrieve the URL from the shortener service
	url, err := m.shortener.RetrieveURL(shortCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Redirect the user to the URL
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
