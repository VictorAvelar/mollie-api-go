package testdata

// CreateMandateResponse
const CreateMandateResponse = `{
    "resource": "mandate",
    "id": "mdt_h3gAaD5zP",
    "mode": "test",
    "status": "valid",
    "method": "directdebit",
    "details": {
        "consumerName": "John Doe",
        "consumerAccount": "NL55INGB0000000000",
        "consumerBic": "INGBNL2A"
    },
    "mandateReference": "YOUR-COMPANY-MD13804",
    "signatureDate": "2018-05-07",
    "createdAt": "2018-05-07T10:49:08+00:00",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/customers/cst_4qqhO89gsT/mandates/mdt_h3gAaD5zP",
            "type": "application/hal+json"
        },
        "customer": {
            "href": "https://api.mollie.com/v2/customers/cst_4qqhO89gsT",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/mandates-api/create-mandate",
            "type": "text/html"
        }
    }
}`

// GetMandateResponse
const GetMandateResponse = `{
    "resource": "mandate",
    "id": "mdt_h3gAaD5zP",
    "mode": "test",
    "status": "valid",
    "method": "directdebit",
    "details": {
        "consumerName": "John Doe",
        "consumerAccount": "NL55INGB0000000000",
        "consumerBic": "INGBNL2A"
    },
    "mandateReference": "YOUR-COMPANY-MD1380",
    "signatureDate": "2018-05-07",
    "createdAt": "2018-05-07T10:49:08+00:00",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/customers/cst_4qqhO89gsT/mandates/mdt_h3gAaD5zP",
            "type": "application/hal+json"
        },
        "customer": {
            "href": "https://api.mollie.com/v2/customers/cst_4qqhO89gsT",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/mandates-api/get-mandate",
            "type": "text/html"
        }
    }
}`

const ListMandatesResponse = `
{
    "count": 2,
    "_embedded": {
        "mandates": [
            {
                "resource": "mandate",
                "id": "mdt_AcQl5fdL4h",
                "mode": "test",
                "status": "valid",
                "method": "directdebit",
                "details": {
                    "consumerName": "John Doe",
                    "consumerAccount": "NL55INGB0000000000",
                    "consumerBic": "INGBNL2A"
                },
                "mandateReference": null,
                "signatureDate": "2018-05-07",
                "createdAt": "2018-05-07T10:49:08+00:00",
                "_links": {
                    "self": {
                        "href": "https://api.mollie.com/v2/customers/cst_8wmqcHMN4U/mandates/mdt_AcQl5fdL4h",
                        "type": "application/hal+json"
                    },
                    "customer": {
                        "href": "https://api.mollie.com/v2/customers/cst_8wmqcHMN4U",
                        "type": "application/hal+json"
                    },
                    "documentation": {
                        "href": "https://mollie.com/en/docs/reference/customers/create-mandate",
                        "type": "text/html"
                    }
                }
            },
{
                "resource": "mandate",
                "id": "mdt_AcQl5fdL4h",
                "mode": "test",
                "status": "valid",
                "method": "directdebit",
                "details": {
                    "consumerName": "John Doe",
                    "consumerAccount": "NL55INGB0000000000",
                    "consumerBic": "INGBNL2A"
                },
                "mandateReference": null,
                "signatureDate": "2018-05-07",
                "createdAt": "2018-05-07T10:49:08+00:00",
                "_links": {
                    "self": {
                        "href": "https://api.mollie.com/v2/customers/cst_8wmqcHMN4U/mandates/mdt_AcQl5fdL4h",
                        "type": "application/hal+json"
                    },
                    "customer": {
                        "href": "https://api.mollie.com/v2/customers/cst_8wmqcHMN4U",
                        "type": "application/hal+json"
                    },
                    "documentation": {
                        "href": "https://mollie.com/en/docs/reference/customers/create-mandate",
                        "type": "text/html"
                    }
                }
            }
        ]
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/customers/cst_8wmqcHMN4U/mandates?limit=5",
            "type": "application/hal+json"
        },
        "previous": null,
        "next": {
            "href": "https://api.mollie.com/v2/customers/cst_8wmqcHMN4U/mandates?from=mdt_AcQl5fdL4h&limit=5",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/mandates-api/revoke-mandate",
            "type": "text/html"
        }
    }
}`
