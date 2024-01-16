package v1

import (
	"net/http"

	"github.com/eduardolat/tinylink/internal/shortener"
)

func (m *router) shorten(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("url query param is required"))
		return
	}

	shortCode, err := m.shortener.ShortenURL(shortener.StoreURLParams{
		OriginalURL: url,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(shortCode))
}
