package testdata

// GetPaymentLink response sample.
const GetPaymentLinkResponse = `{
    "resource": "payment-link",
    "id": "pl_4Y0eZitmBnQ6IDoMqZQKh",
    "mode": "test",
    "profileId": "pfl_QkEhN94Ba",
    "createdAt": "2021-03-20T09:13:37+00:00",
    "paidAt": "2021-03-21T09:13:37+00:00",
    "updatedAt": "2021-03-21T09:13:37+00:00",
    "expiresAt": null,
    "amount": {
        "value": "24.95",
        "currency": "EUR"
    },
    "description": "Bicycle tires",
    "redirectUrl": "https://webshop.example.org/thanks",
    "webhookUrl": "https://webshop.example.org/payment-links/webhook/",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/payment-links/pl_4Y0eZitmBnQ6IDoMqZQKh",
            "type": "application/json"
        },
        "paymentLink": {
            "href": "https://paymentlink.mollie.com/payment/4Y0eZitmBnQ6IDoMqZQKh/",
            "type": "text/html"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/payment-links-api/get-payment-link",
            "type": "text/html"
        }
    }
}
`

const CreatePaymentLinkResponse = `{
    "resource": "payment-link",
    "id": "pl_4Y0eZitmBnQ6IDoMqZQKh",
    "mode": "test",
    "profileId": "pfl_QkEhN94Ba",
    "createdAt": "2021-03-20T09:13:37+00:00",
    "paidAt": null,
    "updatedAt": null,
    "expiresAt": "2021-06-06T11:00:00+00:00",
    "amount": {
        "value": "24.95",
        "currency": "EUR"
    },
    "description": "Bicycle tires",
    "redirectUrl": "https://webshop.example.org/thanks",
    "webhookUrl": "https://webshop.example.org/payment-links/webhook/",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/payment-links/pl_4Y0eZitmBnQ6IDoMqZQKh",
            "type": "application/json"
        },
        "paymentLink": {
            "href": "https://paymentlink.mollie.com/payment/4Y0eZitmBnQ6IDoMqZQKh/",
            "type": "text/html"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/payment-links-api/create-payment-link",
            "type": "text/html"
        }
    }
}`

const ListPaymentLinksResponse = `{
    "count": 1,
    "_embedded": {
        "payment_links": [
              {
                  "resource": "payment-link",
                  "id": "pl_4Y0eZitmBnQ6IDoMqZQKh",
                  "mode": "test",
                  "profileId": "pfl_QkEhN94Ba",
                  "createdAt": "2021-03-20T09:13:37+00:00",
                  "paidAt": "2021-03-21T09:13:37+00:00",
                  "updatedAt": "2021-03-21T09:13:37+00:00",
                  "expiresAt": null,
                  "amount": {
                      "value": "24.95",
                      "currency": "EUR"
                  },
                  "description": "Bicycle tires",
                  "redirectUrl": "https://webshop.example.org/thanks",
                  "webhookUrl": "https://webshop.example.org/payment-links/webhook/",
                  "_links": {
                      "self": {
                          "href": "https://api.mollie.com/v2/payment-links/pl_4Y0eZitmBnQ6IDoMqZQKh",
                          "type": "application/json"
                      },
                      "paymentLink": {
                          "href": "https://paymentlink.mollie.com/payment/4Y0eZitmBnQ6IDoMqZQKh/",
                          "type": "text/html"
                      },
                      "documentation": {
                          "href": "https://docs.mollie.com/reference/payment-links-api/get-payment-link",
                          "type": "text/html"
                      }
                  }
              }
        ]
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/payment-links?limit=5",
            "type": "application/hal+json"
        },
        "previous": null,
        "next": {
            "href": "https://api.mollie.com/v2/payment-links?from=pl_ER6aqfpXg6nZrJvcsxNsm&limit=5",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/payment-links-api/list-payment-links",
            "type": "text/html"
        }
    }
}`

// UpdatePaymentLinksResponse example.
const UpdatePaymentLinksResponse = `{
  "resource": "payment-link",
  "id": "pl_4Y0eZitmBnQ6IDoMqZQKh",
  "mode": "live",
  "description": "Bicycle tires",
  "amount": {
    "currency": "EUR",
    "value": "24.95"
  },
  "archived": true,
  "redirectUrl": "https://webshop.example.org/thanks",
  "webhookUrl": "https://webshop.example.org/payment-links/webhook",
  "profileId": "pfl_QkEhN94Ba",
  "createdAt": "2021-03-20T09:29:56+00:00",
  "expiresAt": "2023-06-06T11:00:00+00:00",
  "_links": {
    "self": {
      "href": "...",
      "type": "application/hal+json"
    },
    "paymentLink": {
      "href": "https://paymentlink.mollie.com/payment/4Y0eZitmBnQ6IDoMqZQKh",
      "type": "text/html"
    },
    "documentation": {
      "href": "...",
      "type": "text/html"
    }
  }
}`
