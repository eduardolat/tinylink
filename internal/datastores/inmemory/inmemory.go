package inmemory

import (
	"errors"
	"sync"
	"time"

	"github.com/eduardolat/tinylink/internal/shortener"
)

type DataStore struct {
	data map[string]shortener.URLData
	mu   sync.RWMutex
}

func NewDataStore() *DataStore {
	return &DataStore{
		data: make(map[string]shortener.URLData),
	}
}

func (ds *DataStore) AutoMigrate() error {
	return nil
}

func (ds *DataStore) IsShortCodeAvailable(shortCode string) (bool, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	_, exists := ds.data[shortCode]

	return !exists, nil
}

func (ds *DataStore) IsURLAlreadyStored(originalURL string) (bool, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	for _, data := range ds.data {
		if data.OriginalURL == originalURL {
			return true, nil
		}
	}

	return false, nil
}

func (ds *DataStore) StoreURL(params shortener.StoreURLParams) (shortener.URLData, error) {
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

func (ds *DataStore) RetrieveURL(shortCode string) (shortener.URLData, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	data, exists := ds.data[shortCode]
	if !exists {
		return shortener.URLData{}, errors.New("URL not found")
	}

	return data, nil
}

func (ds *DataStore) UpdateURL(shortCode string, params shortener.UpdateURLParams) (shortener.URLData, error) {
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

func (ds *DataStore) DeleteURL(shortCode string) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	_, exists := ds.data[shortCode]
	if !exists {
		return errors.New("URL not found")
	}

	delete(ds.data, shortCode)

	return nil
}

func (ds *DataStore) IncrementClicks(shortCode string) error {
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

func (ds *DataStore) IncrementRedirects(shortCode string) error {
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

func (ds *DataStore) GetURLsByTag(tag string) ([]shortener.URLData, error) {
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

func (ds *DataStore) GetActiveURLs() ([]shortener.URLData, error) {
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

func (ds *DataStore) GetExpiredURLs() ([]shortener.URLData, error) {
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