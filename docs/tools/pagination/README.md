# pagination

--
    import "github.com/VictorAvelar/mollie-api-go/v3/mollie/tools/pagination"

Package pagination provides utilities to handle pagination in API responses.

Pagination is a common feature in APIs that allows retrieval of large datasets
in smaller chunks, enhancing performance and resource usage. This package aims
to simplify pagination-related tasks by providing helpful functions.

## Usage

#### func  ExtractFromQueryParam

```go
func ExtractFromQueryParam(uri string) (lastID string, err error)
```

ExtractFromQueryParam extracts the lastID from the given URI, which is assumed
to be a URL with query parameters. It specifically looks for a query parameter
named 'from' and returns its value as a string. If the URI cannot be parsed or
the query parameter is not found, it returns an empty string and the encountered
error.
