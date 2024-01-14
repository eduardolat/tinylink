package shortener

type Shortener struct {
	dataStore DataStore
	shortGen  ShortGen
}

func NewShortener(
	dataStore DataStore,
	shortGen ShortGen,
) *Shortener {
	return &Shortener{
		dataStore: dataStore,
		shortGen:  shortGen,
	}
}

func (c *Shortener) ShortenURL(url string) (string, error) {
	shortCode, err := c.shortGen.Generate()
	if err != nil {
		return "", err
	}

	args := StoreURLParams{
		ShortCode:         shortCode,
		OriginalURL:       url,
		HTTPRedirectCode:  HTTPRedirectCodePermanent,
		IsActive:          true,
		DuplicateIfExists: true,
	}

	urlData, err := c.dataStore.StoreURL(args)
	if err != nil {
		return "", err
	}

	return urlData.ShortCode, nil
}

func (c *Shortener) RetrieveURL(shortCode string) (string, error) {
	data, err := c.dataStore.RetrieveURL(shortCode)
	if err != nil {
		return "", err
	}

	return data.OriginalURL, nil
}
