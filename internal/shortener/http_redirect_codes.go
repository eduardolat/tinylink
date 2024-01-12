package shortener

// HTTPRedirectCode is the HTTP code that will be used to redirect the user
// to the original URL
type HTTPRedirectCode int

const (
	// HTTPRedirectCode301 indicates that the requested resource has been permanently moved to a new location.
	HTTPRedirectCode301 HTTPRedirectCode = 301

	// HTTPRedirectCode302 indicates that the requested resource resides temporarily under a different URI.
	HTTPRedirectCode302 HTTPRedirectCode = 302

	// HTTPRedirectCode303 indicates that the response to the request can be found under a different URI and should be retrieved using a GET method.
	HTTPRedirectCode303 HTTPRedirectCode = 303

	// HTTPRedirectCode307 indicates that the requested resource resides temporarily under a different URI and the request should be repeated with the same method and body.
	HTTPRedirectCode307 HTTPRedirectCode = 307

	// HTTPRedirectCode308 indicates that the requested resource has been permanently moved to a new location and the request should be repeated with the same method and body.
	HTTPRedirectCode308 HTTPRedirectCode = 308

	// HTTPRedirectCodePermanent is an alias for HTTPRedirectCode301.
	HTTPRedirectCodePermanent HTTPRedirectCode = HTTPRedirectCode301

	// HTTPRedirectCodeTemporary is an alias for HTTPRedirectCode302.
	HTTPRedirectCodeTemporary HTTPRedirectCode = HTTPRedirectCode302
)
