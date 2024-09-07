package testdata

// GetPaymentResponse example
const GetPaymentResponse = `{
    "resource": "payment",
    "id": "tr_WDqYK6vllg",
    "mode": "test",
    "createdAt": "2018-03-20T13:13:37+00:00",
    "amount": {
        "value": "10.00",
        "currency": "EUR"
    },
    "description": "Order #12345",
    "method": null,
    "metadata": {
        "order_id": "12345"
    },
    "status": "open",
    "isCancelable": false,
    "locale": "nl_NL",
    "restrictPaymentMethodsToCountry": "NL",
    "expiresAt": "2018-03-20T13:28:37+00:00",
    "details": null,
    "profileId": "pfl_QkEhN94Ba",
    "sequenceType": "oneoff",
    "redirectUrl": "https://webshop.example.org/order/12345/",
    "webhookUrl": "https://webshop.example.org/payments/webhook/",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg",
            "type": "application/hal+json"
        },
        "checkout": {
            "href": "https://www.mollie.com/payscreen/select-method/WDqYK6vllg",
            "type": "text/html"
        },
        "dashboard": {
            "href": "https://www.mollie.com/dashboard/org_12345678/payments/tr_WDqYK6vllg",
            "type": "application/json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/payments-api/get-payment",
            "type": "text/html"
        }
    }
}`

// CancelPaymentResponse example
const CancelPaymentResponse = `{
    "resource": "payment",
    "id": "tr_WDqYK6vllg",
    "mode": "live",
    "createdAt": "2018-03-19T10:18:33+00:00",
    "amount": {
        "value": "35.07",
        "currency": "EUR"
    },
    "description": "Order 33",
    "method": "banktransfer",
    "metadata": null,
    "status": "canceled",
    "canceledAt": "2018-03-19T10:19:15+00:00",
    "details": {
        "bankName": "Stichting Mollie Payments",
        "bankAccount": "NL53ABNA0627535577",
        "bankBic": "ABNANL2A",
        "transferReference": "RF12-3456-7890-1234"
    },
    "profileId": "pfl_QkEhN94Ba",
    "sequenceType": "oneoff",
    "redirectUrl": "https://webshop.example.org/order/33/",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/payments-api/cancel-payment",
            "type": "text/html"
        }
    }
}`

// UpdatePaymentResponse example
const UpdatePaymentResponse = `{
    "resource": "payment",
    "id": "tr_7UhSN1zuXS",
    "mode": "test",
    "createdAt": "2018-03-20T09:13:37+00:00",
    "amount": {
        "value": "10.00",
        "currency": "EUR"
    },
    "description": "Order #98765",
    "method": null,
    "metadata": {
        "order_id": "98765"
    },
    "status": "open",
    "isCancelable": false,
    "expiresAt": "2018-03-20T09:28:37+00:00",
    "details": null,
    "profileId": "pfl_QkEhN94Ba",
    "sequenceType": "oneoff",
    "redirectUrl": "https://example.org/webshop/order/98765/",
    "webhookUrl": "https://example.org/webshop/payments/webhook/",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/payments/tr_7UhSN1zuXS",
            "type": "application/json"
        },
        "checkout": {
            "href": "https://www.mollie.com/payscreen/select-method/7UhSN1zuXS",
            "type": "text/html"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/payments-api/update-payment",
            "type": "text/html"
        }
    }
}
`

// ListPaymentsResponse example
const ListPaymentsResponse = `{
    "count": 5,
    "_embedded": {
        "payments": [
            {
                "resource": "payment",
                "id": "tr_7UhSN1zuXS",
                "mode": "test",
                "createdAt": "2018-02-12T11:58:35.0Z",
                "expiresAt": "2018-02-12T12:13:35.0Z",
                "status": "open",
                "isCancelable": false,
                "amount": {
                    "value": "75.00",
                    "currency": "GBP"
                },
                "description": "Order #12345",
                "method": "ideal",
                "metadata": null,
                "details": null,
                "profileId": "pfl_QkEhN94Ba",
                "redirectUrl": "https://webshop.example.org/order/12345/",
                "_links": {
                    "checkout": {
                        "href": "https://www.mollie.com/paymentscreen/issuer/select/ideal/7UhSN1zuXS",
                        "type": "text/html"
                    },
                    "self": {
                        "href": "https://api.mollie.com/v2/payments/tr_7UhSN1zuXS",
                        "type": "application/hal+json"
                    }
                }
            }
        ]
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/payments?limit=5",
            "type": "application/hal+json"
        },
        "previous": null,
        "next": {
            "href": "https://api.mollie.com/v2/payments?from=tr_SDkzMggpvx&limit=5",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/payments-api/list-payments",
            "type": "text/html"
        }
    }
}`
