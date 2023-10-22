package pagination

import (
	"net/url"
)

// ExtractFromQueryParam extracts the lastID from the given URI, which is assumed to be a URL with query parameters.
// It specifically looks for a query parameter named 'from' and returns its value as a string.
// If the URI cannot be parsed or the query parameter is not found, it returns an empty string and the encountered
// error.
func ExtractFromQueryParam(uri string) (lastID string, err error) {
	const from = "from"

	return parseURIAndReturnQueryParam(uri, from)
}

func parseURIAndReturnQueryParam(uri string, param string) (val string, err error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}

	v := u.Query().Get(param)

	return v, nil
}
