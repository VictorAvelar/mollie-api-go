package testdata

// ListTerminalsResponse example.
const ListTerminalsResponse = `{
    "count": 5,
    "_embedded": {
        "terminals": [
            {
                "id": "term_7MgL4wea46qkRcoTZjWEH",
                "profileId": "pfl_QkEhN94Ba",
                "status": "active",
                "brand": "PAX",
                "model": "A920",
                "serialNumber": "1234567890",
                "currency": "EUR",
                "description": "Terminal #12341",
                "createdAt": "2022-02-12T11:58:35.0Z",
                "updatedAt": "2022-11-15T13:32:11.0Z",
                "_links": {
                    "self": {
                        "href": "https://api.mollie.com/v2/terminals/term_7MgL4wea46qkRcoTZjWEH",
                        "type": "application/hal+json"
                    }
                }
            },
            {
                "id": "term_7sgL4wea46qkRcoysdiWEH",
                "profileId": "pfl_QkEhN94Ba",
                "status": "active",
                "brand": "PAX",
                "model": "A920",
                "serialNumber": "1234567890",
                "currency": "MEX",
                "description": "Terminal #12342",
                "createdAt": "2022-02-12T11:58:35.0Z",
                "updatedAt": "2022-11-15T13:32:11.0Z",
                "_links": {
                    "self": {
                        "href": "https://api.mollie.com/v2/terminals/term_7sgL4wea46qkRcoysdiWEH",
                        "type": "application/hal+json"
                    }
                }
            },
            {
                "id": "term_7MgLsdD*b3asDayWEH",
                "profileId": "pfl_QkEhN94Ba",
                "status": "active",
                "brand": "PAX",
                "model": "A920",
                "serialNumber": "1234567890",
                "currency": "GBP",
                "description": "Terminal #12343",
                "createdAt": "2022-02-12T11:58:35.0Z",
                "updatedAt": "2022-11-15T13:32:11.0Z",
                "_links": {
                    "self": {
                        "href": "https://api.mollie.com/v2/terminals/term_7MgLsdD*b3asDayWEH",
                        "type": "application/hal+json"
                    }
                }
            },
            {
                "id": "term_7MgL4j5jAowWqkRcoTZjWEH",
                "profileId": "pfl_QkEhN94Ba",
                "status": "active",
                "brand": "PAX",
                "model": "A920",
                "serialNumber": "1234567890",
                "currency": "DLS",
                "description": "Terminal #12344",
                "createdAt": "2022-02-12T11:58:35.0Z",
                "updatedAt": "2022-11-15T13:32:11.0Z",
                "_links": {
                    "self": {
                        "href": "https://api.mollie.com/v2/terminals/term_7MgL4j5jAowWqkRcoTZjWEH",
                        "type": "application/hal+json"
                    }
                }
            },
            {
                "id": "term_7MgL4we02ujSeRcoTZjWEH",
                "profileId": "pfl_QkEhN94Ba",
                "status": "active",
                "brand": "PAX",
                "model": "A920",
                "serialNumber": "1234567890",
                "currency": "COP",
                "description": "Terminal #12345",
                "createdAt": "2022-02-12T11:58:35.0Z",
                "updatedAt": "2022-11-15T13:32:11.0Z",
                "_links": {
                    "self": {
                        "href": "https://api.mollie.com/v2/terminals/term_7MgL4we02ujSeRcoTZjWEH",
                        "type": "application/hal+json"
                    }
                }
            }
        ]
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/terminals?limit=5",
            "type": "application/hal+json"
        },
        "previous": null,
        "next": {
            "href": "https://api.mollie.com/v2/terminals?from=term_7MgL4we02ujSeRcoTZjWEH&limit=5",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/terminals-api/list-terminals",
            "type": "text/html"
        }
    }
}`

// GetTerminalResponse example.
const GetTerminalResponse = `{
    "id": "term_7MgL4wea46qkRcoTZjWEH",
    "profileId": "pfl_QkEhN94Ba",
    "status": "active",
    "brand": "PAX",
    "model": "A920",
    "serialNumber": "1234567890",
    "currency": "EUR",
    "description": "Terminal #12345",
    "createdAt": "2022-02-12T11:58:35.0Z",
    "updatedAt": "2022-11-15T13:32:11.0Z",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/terminals/term_7MgL4wea46qkRcoTZjWEH",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/terminals-api/get-terminal",
            "type": "text/html"
        }
    }
}`
