package inmemory

import (
	"errors"
	"sort"
	"sync"
	"time"

	"github.com/eduardolat/tinylink/internal/datastores"
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

func (ds *DataStore) RetrieveURLByOriginalUrl(originalUrl string) (shortener.URLData, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	for _, data := range ds.data {
		if data.OriginalURL == originalUrl {
			return data, nil
		}
	}

	return shortener.URLData{}, errors.New("URL not found")
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

func (ds *DataStore) PaginateURLS(params shortener.PaginateURLSParams) (shortener.PaginateURLSResponse, error) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	allURLs := []shortener.URLData{}
	for _, data := range ds.data {
		if len(params.TagsFilter) == 0 {
			allURLs = append(allURLs, data)
			continue
		}

		if containsTags(data.Tags, params.TagsFilter) {
			allURLs = append(allURLs, data)
		}
	}

	sort.SliceStable(allURLs, func(i, j int) bool {
		return allURLs[i].CreatedAt.Before(allURLs[j].CreatedAt)
	})

	start := (params.Page - 1) * params.Size
	if start > len(allURLs) {
		return shortener.PaginateURLSResponse{}, errors.New("page number out of range")
	}

	end := start + params.Size
	if end > len(allURLs) {
		end = len(allURLs)
	}

	totalItems := len(allURLs)
	items := allURLs[start:end]
	pagination := datastores.CreatePaginationResponse(
		params.Page,
		params.Size,
		totalItems,
		items,
	)

	return pagination, nil
}

func containsTags(urlTags []string, filterTags []string) bool {
	for _, urlTag := range urlTags {
		for _, filterTag := range filterTags {
			if urlTag == filterTag {
				return true
			}
		}
	}

	return false
}
