package shortener

const (
	// HTTPRedirectCode301 indicates that the requested resource has been permanently moved to a new location.
	HTTPRedirectCode301 int = 301

	// HTTPRedirectCode302 indicates that the requested resource resides temporarily under a different URI.
	HTTPRedirectCode302 int = 302

	// HTTPRedirectCode303 indicates that the response to the request can be found under a different URI and should be retrieved using a GET method.
	HTTPRedirectCode303 int = 303

	// HTTPRedirectCode307 indicates that the requested resource resides temporarily under a different URI and the request should be repeated with the same method and body.
	HTTPRedirectCode307 int = 307

	// HTTPRedirectCode308 indicates that the requested resource has been permanently moved to a new location and the request should be repeated with the same method and body.
	HTTPRedirectCode308 int = 308

	// HTTPRedirectCodePermanent is an alias for HTTPRedirectCode301.
	HTTPRedirectCodePermanent int = HTTPRedirectCode301

	// HTTPRedirectCodeTemporary is an alias for HTTPRedirectCode302.
	HTTPRedirectCodeTemporary int = HTTPRedirectCode302
)
