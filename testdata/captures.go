package testdata

// ListCapturesResponse example
const ListCapturesResponse = `{
    "_embedded": {
        "captures": [
            {
                "resource": "capture",
                "id": "cpt_4qqhO89gsT",
                "mode": "live",
                "amount": {
                    "value": "1027.99",
                    "currency": "EUR"
                },
                "settlementAmount": {
                    "value": "399.00",
                    "currency": "EUR"
                },
                "paymentId": "tr_WDqYK6vllg",
                "shipmentId": "shp_3wmsgCJN4U",
                "settlementId": "stl_jDk30akdN",
                "createdAt": "2018-08-02T09:29:56+00:00",
                "_links": {
                    "self": {
                        "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg/captures/cpt_4qqhO89gsT",
                        "type": "application/hal+json"
                    },
                    "payment": {
                        "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg",
                        "type": "application/hal+json"
                    },
                    "shipment": {
                        "href": "https://api.mollie.com/v2/orders/ord_8wmqcHMN4U/shipments/shp_3wmsgCJN4U",
                        "type": "application/hal+json"
                    },
                    "settlement": {
                        "href": "https://api.mollie.com/v2/settlements/stl_jDk30akdN",
                        "type": "application/hal+json"
                    },
                    "documentation": {
                        "href": "https://docs.mollie.com/reference/v2/captures-api/get-capture",
                        "type": "text/html"
                    }
                }
            }
        ]
    },
    "count": 1,
    "_links": {
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/captures-api/list-captures",
            "type": "text/html"
        },
        "self": {
            "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg/captures?limit=50",
            "type": "application/hal+json"
        },
        "previous": null,
        "next": null
    }
}`

// GetCaptureResponse example
const GetCaptureResponse = `{
    "resource": "capture",
    "id": "cpt_4qqhO89gsT",
    "mode": "live",
    "amount": {
        "value": "1027.99",
        "currency": "EUR"
    },
    "settlementAmount": {
        "value": "399.00",
        "currency": "EUR"
    },
    "paymentId": "tr_WDqYK6vllg",
    "shipmentId": "shp_3wmsgCJN4U",
    "settlementId": "stl_jDk30akdN",
    "createdAt": "2018-08-02T09:29:56+00:00",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg/captures/cpt_4qqhO89gsT",
            "type": "application/hal+json"
        },
        "payment": {
            "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg",
            "type": "application/hal+json"
        },
        "shipment": {
            "href": "https://api.mollie.com/v2/orders/ord_8wmqcHMN4U/shipments/shp_3wmsgCJN4U",
            "type": "application/hal+json"
        },
        "settlement": {
            "href": "https://api.mollie.com/v2/settlements/stl_jDk30akdN",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/captures-api/get-capture",
            "type": "text/html"
        }
    }
}`

// CreateCaptureResponse example.
const CreateCaptureResponse = `{
    "resource": "capture",
    "id": "cpt_mNepDkEtco6ah3QNPUGYH",
    "mode": "live",
    "amount": {
        "value": "35.95",
        "currency": "EUR"
    },
    "paymentId": "tr_WDqYK6vllg",
    "createdAt": "2018-08-02T09:29:56+00:00",
    "description": "Capture for cart #12345",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg/captures/cpt_mNepDkEtco6ah3QNPUGYH",
            "type": "application/hal+json"
        },
        "payment": {
            "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/captures-api/create-capture",
            "type": "text/html"
        }
    }
}`

// CreateCaptureWithAccessTokenResponse example.
const CreateCaptureWithAccessTokenResponse = `{
    "resource": "capture",
    "id": "cpt_mNepDkEtco6ah3QNPUGYH",
    "mode": "live",
    "amount": {
        "value": "35.95",
        "currency": "EUR"
    },
    "paymentId": "tr_WDqYK6vllg",
    "createdAt": "2018-08-02T09:29:56+00:00",
    "description": "Capture for cart #12345",
    "testmode": true,
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg/captures/cpt_mNepDkEtco6ah3QNPUGYH",
            "type": "application/hal+json"
        },
        "payment": {
            "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/captures-api/create-capture",
            "type": "text/html"
        }
    }
}`
