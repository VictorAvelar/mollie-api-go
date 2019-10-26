package testdata

// GetCustomerResponse example
const GetCustomerResponse = `{
    "resource": "customer",
    "id": "cst_kEn1PlbGa",
    "mode": "test",
    "name": "Customer A",
    "email": "customer@example.org",
    "locale": "nl_NL",
    "metadata": null,
    "createdAt": "2018-04-06T13:23:21.0Z",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/customers/cst_kEn1PlbGa",
            "type": "application/hal+json"
        },
        "mandates": {
            "href": "https://api.mollie.com/v2/customers/cst_kEn1PlbGa/mandates",
            "type": "application/hal+json"
        },
        "subscriptions": {
            "href": "https://api.mollie.com/v2/customers/cst_kEn1PlbGa/subscriptions",
            "type": "application/hal+json"
        },
        "payments": {
            "href": "https://api.mollie.com/v2/customers/cst_kEn1PlbGa/payments",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/customers-api/get-customer",
            "type": "text/html"
        }
    }
}`

// CreateCustomerResponse example
const CreateCustomerResponse = `{
    "resource": "customer",
    "id": "cst_8wmqcHMN4U",
    "mode": "test",
    "name": "Customer A",
    "email": "customer@example.org",
    "locale": null,
    "metadata": null,
    "createdAt": "2018-04-06T13:10:19.0Z",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/customers/cst_8wmqcHMN4U",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/customers-api/create-customer",
            "type": "text/html"
        }
    }
}`

// UpdateCustomerResponse example
const UpdateCustomerResponse = `{
    "resource": "customer",
    "id": "cst_8wmqcHMN4U",
    "mode": "test",
    "name": "Updated Customer A",
    "email": "updated-customer@example.org",
    "locale": "nl_NL",
    "metadata": null,
    "createdAt": "2018-04-06T13:23:21.0Z",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/customers/cst_8wmqcHMN4U",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/customers-api/get-customer",
            "type": "text/html"
        }
    }
}`

// ListCustomersResponse example
const ListCustomersResponse = `{
    "count": 3,
    "_embedded": {
        "customers": [
            {
                "resource": "customer",
                "id": "cst_kEn1PlbGa",
                "mode": "test",
                "name": "Customer A",
                "email": "customer@example.org",
                "locale": "nl_NL",
                "metadata": null,
                "createdAt": "2018-04-06T13:23:21.0Z",
                "_links": {
                    "self": {
                        "href": "https://api.mollie.com/v2/customers/cst_kEn1PlbGa",
                        "type": "application/hal+json"
                    },
                    "documentation": {
                        "href": "https://docs.mollie.com/reference/v2/customers-api/get-customer",
                        "type": "text/html"
                    }
                }
            },
            { },
            { }
        ]
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/customers",
            "type": "application/hal+json"
        },
        "previous": null,
        "next": {
            "href": "https://api.mollie.com/v2/customers?from=cst_stTC2WHAuS",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/customers-api/list-customers",
            "type": "text/html"
        }
    }
}`
