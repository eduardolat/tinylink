package shortener

type Client struct {
	dataStore          DataStore
	shortCodeGenerator ShortCodeGenerator
}

func NewClient(
	shortCodeGenerator ShortCodeGenerator,
	dataStore DataStore,
) *Client {
	return &Client{
		shortCodeGenerator: shortCodeGenerator,
		dataStore:          dataStore,
	}
}

func (c *Client) ShortenURL(url string) (string, error) {
	shortCode, err := c.shortCodeGenerator.Generate()
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
