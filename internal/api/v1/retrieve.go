package v1

import "net/http"

func (m *router) retrieve(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Query().Get("shortCode")
	if shortCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("shortCode query param is required"))
		return
	}

	url, err := m.shortener.RetrieveURL(shortCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(url))
}
