package inmemory

import (
	"errors"
	"sync"
	"time"

	"github.com/eduardolat/tinylink/internal/shortener"
)

type InMemoryDataStore struct {
	data map[string]shortener.URLData
	mu   sync.RWMutex
}

func NewInMemoryDataStore() *InMemoryDataStore {
	return &InMemoryDataStore{
		data: make(map[string]shortener.URLData),
	}
}

func (ds *InMemoryDataStore) StoreURL(params shortener.StoreURLParams) (shortener.URLData, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	data := shortener.URLData{
		IsActive:           params.IsActive,
		Description:        params.Description,
		Tags:               params.Tags,
		HTTPRedirectCode:   params.HTTPRedirectCode,
		OriginalURL:        params.OriginalURL,
		ShortCode:          params.ShortCode,
		Password:           params.Password,
		ExpiresAt:          params.ExpiresAt,
		CreatedByIP:        params.CreatedByIP,
		CreatedByUserAgent: params.CreatedByUserAgent,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	ds.data[params.ShortCode] = data

	return data, nil
}

func (ds *InMemoryDataStore) RetrieveURL(shortCode string) (shortener.URLData, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	data, exists := ds.data[shortCode]
	if !exists {
		return shortener.URLData{}, errors.New("URL not found")
	}

	return data, nil
}

func (ds *InMemoryDataStore) UpdateURL(shortCode string, params shortener.UpdateURLParams) (shortener.URLData, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	data, exists := ds.data[shortCode]
	if !exists {
		return shortener.URLData{}, errors.New("URL not found")
	}

	data.IsActive = params.IsActive
	data.Description = params.Description
	data.Tags = params.Tags
	data.HTTPRedirectCode = params.HTTPRedirectCode
	data.Password = params.Password
	data.ExpiresAt = params.ExpiresAt
	data.UpdatedAt = time.Now()

	ds.data[shortCode] = data

	return data, nil
}

func (ds *InMemoryDataStore) DeleteURL(shortCode string) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	_, exists := ds.data[shortCode]
	if !exists {
		return errors.New("URL not found")
	}

	delete(ds.data, shortCode)

	return nil
}

func (ds *InMemoryDataStore) IncrementClicks(shortCode string) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	data, exists := ds.data[shortCode]
	if !exists {
		return errors.New("URL not found")
	}

	data.Clicks.Int64++
	data.UpdatedAt = time.Now()

	ds.data[shortCode] = data

	return nil
}

func (ds *InMemoryDataStore) IncrementRedirects(shortCode string) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	data, exists := ds.data[shortCode]
	if !exists {
		return errors.New("URL not found")
	}

	data.Redirects.Int64++
	data.UpdatedAt = time.Now()

	ds.data[shortCode] = data

	return nil
}

func (ds *InMemoryDataStore) GetURLsByTag(tag string) ([]shortener.URLData, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	var urls []shortener.URLData

	for _, data := range ds.data {
		for _, t := range data.Tags {
			if t == tag {
				urls = append(urls, data)
				break
			}
		}
	}

	return urls, nil
}

func (ds *InMemoryDataStore) GetActiveURLs() ([]shortener.URLData, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	var urls []shortener.URLData

	for _, data := range ds.data {
		if data.IsActive {
			urls = append(urls, data)
		}
	}

	return urls, nil
}

func (ds *InMemoryDataStore) GetExpiredURLs() ([]shortener.URLData, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	var urls []shortener.URLData

	for _, data := range ds.data {
		if data.ExpiresAt.Valid && data.ExpiresAt.Time.Before(time.Now()) {
			urls = append(urls, data)
		}
	}

	return urls, nil
}
