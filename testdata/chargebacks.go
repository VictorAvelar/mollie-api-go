package testdata

const (
	GetChargebackResponse = `{
    "resource": "chargeback",
    "id": "chb_n9z0tp",
    "amount": {
        "currency": "USD",
        "value": "43.38"
    },
    "settlementAmount": {
        "currency": "EUR",
        "value": "-35.07"
    },
    "createdAt": "2018-03-14T17:00:52.0Z",
     "reason": {
       "code": "AC01",
       "description": "Account identifier incorrect (i.e. invalid IBAN)"
     },
    "reversedAt": null,
    "paymentId": "tr_WDqYK6vllg",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg/chargebacks/chb_n9z0tp",
            "type": "application/hal+json"
        },
        "payment": {
            "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/chargebacks-api/get-payment-chargeback",
            "type": "text/html"
        }
    }
}`
	ListChargebacksResponse = `{
    "count": 3,
    "_embedded": {
        "chargebacks": [
            {
                "resource": "chargeback",
                "id": "chb_n9z0tp",
                "amount": {
                    "currency": "USD",
                    "value": "43.38"
                },
                "settlementAmount": {
                    "currency": "EUR",
                    "value": "35.07"
                },
                "createdAt": "2018-03-14T17:00:52.0Z",
                "reversedAt": null,
                "paymentId": "tr_WDqYK6vllg",
                "_links": {
                    "self": {
                        "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg/chargebacks/chb_n9z0tp",
                        "type": "application/hal+json"
                    },
                    "payment": {
                        "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg",
                        "type": "application/hal+json"
                    },
                    "documentation": {
                        "href": "https://docs.mollie.com/reference/chargebacks-api/get-chargeback",
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
            "href": "https://api.mollie.com/v2/payments/tr_7UhSN1zuXS/chargebacks",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/chargebacks-api/list-chargebacks",
            "type": "text/html"
        }
    }
}`
)
