package shortener

type Client struct {
	dataStore DataStore
	shortGen  ShortGen
}

func NewClient(
	shortGen ShortGen,
	dataStore DataStore,
) *Client {
	return &Client{
		shortGen:  shortGen,
		dataStore: dataStore,
	}
}

func (c *Client) ShortenURL(url string) (string, error) {
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

func (c *Client) RetrieveURL(shortCode string) (string, error) {
	data, err := c.dataStore.RetrieveURL(shortCode)
	if err != nil {
		return "", err
	}

	return data.OriginalURL, nil
}
