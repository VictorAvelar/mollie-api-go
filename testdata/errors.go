package testdata

// UnauthorizedErrorResponse example.
const UnauthorizedErrorResponse = `{
    "status": 401,
    "title": "Unauthorized Request",
    "detail": "Missing authentication, or failed to authenticate",
    "_links": {
        "documentation": {
            "href": "https://docs.mollie.com/overview/authentication",
            "type": "text/html"
        }
    }
}`

// NotFoundErrorResponse example.
const NotFoundErrorResponse = `{
    "status": 404,
    "title": "Not Found",
    "detail": "No payment exists with token tr_I_dont_exist.",
    "_links": {
        "documentation": {
            "href": "https://docs.mollie.com/errors",
            "type": "text/html"
        }
    }
}`

// UnprocessableEntityErrorResponse example.
const UnprocessableEntityErrorResponse = `{
    "status": 422,
    "title": "Unprocessable Entity",
    "detail": "The amount is higher than the maximum",
    "field": "amount",
    "_links": {
       "documentation": {
            "href": "https://docs.mollie.com/errors",
            "type": "text/html"
        }
    }
}`

// InternalServerErrorResponse example.
const InternalServerErrorResponse = `{
    "status": 500,
    "title": "Internal Server Error",
    "detail": "An internal server error occurred while processing your request",
    "_links": {
        "documentation": {
            "href": "https://docs.mollie.com/overview/authentication",
            "type": "text/html"
        }
    }
}`
