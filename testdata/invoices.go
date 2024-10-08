package testdata

// GetInvoiceResponse example
const GetInvoiceResponse = `{
    "resource": "invoice",
    "id": "inv_xBEbP9rvAq",
    "reference": "2016.10000",
    "vatNumber": "NL001234567B01",
    "status": "open",
    "issuedAt": "2016-08-31",
    "dueAt": "2016-09-14",
    "netAmount": {
        "value": "45.00",
        "currency": "EUR"
    },
    "vatAmount": {
        "value": "9.45",
        "currency": "EUR"
    },
    "grossAmount": {
        "value": "54.45",
        "currency": "EUR"
    },
    "lines":[
        {
            "period": "2016-09",
            "description": "iDEAL transactiekosten",
            "count": 100,
            "vatPercentage": 21,
            "amount": {
                "value": "45.00",
                "currency": "EUR"
            }
        }
    ],
    "_links": {
        "self": {
             "href": "https://api.mollie.com/v2/invoices/inv_xBEbP9rvAq",
             "type": "application/hal+json"
        },
        "pdf": {
             "href": "https://www.mollie.com/merchant/download/invoice/xBEbP9rvAq/2ab44d60b35b1d06090bba955fa2c602",
             "type": "application/pdf",
             "expiresAt": "2018-11-09T14:10:36+00:00"
        }
    }
}`

// ListInvoicesResponse example
const ListInvoicesResponse = `{
    "count": 5,
    "_embedded": {
        "invoices": [
            {
                "resource": "invoice",
                "id": "inv_xBEbP9rvAq",
                "reference": "2016.10000",
                "vatNumber": "NL001234567B01",
                "status": "open",
                "issuedAt": "2016-08-31",
                "dueAt": "2016-09-14",
                "netAmount": {
                    "value": "45.00",
                    "currency": "EUR"
                },
                "vatAmount": {
                    "value": "9.45",
                    "currency": "EUR"
                },
                "grossAmount": {
                    "value": "54.45",
                    "currency": "EUR"
                },
                "lines":[
                    {
                        "period": "2016-09",
                        "description": "iDEAL transactiekosten",
                        "count": 100,
                        "vatPercentage": 21,
                        "amount": {
                            "value": "45.00",
                            "currency": "EUR"
                        }
                    }
                ],
                "_links": {
                    "self": {
                         "href": "https://api.mollie.com/v2/invoices/inv_xBEbP9rvAq",
                         "type": "application/hal+json"
                    },
                    "pdf": {
                         "href": "https://www.mollie.com/merchant/download/invoice/xBEbP9rvAq/2ab44d60b35955fa2c602",
                         "type": "application/pdf",
                         "expiresAt": "2018-11-09T14:10:36+00:00"
                    }
                }
            }
        ]
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.nl/v2/invoices?limit=5",
            "type": "application/hal+json"
        },
        "previous": null,
        "next": {
            "href": "https://api.mollie.nl/v2/invoices?from=inv_xBEbP9rvAq&limit=5",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/invoices-api/list-invoices",
            "type": "text/html"
        }
    }
}`
