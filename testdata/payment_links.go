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

const ListPaymentLinkPaymentsResponse = `{
  "count": 1,
  "_embedded": {
    "payments": [
      {
        "resource": "payment",
        "id": "tr_5B8cwPMGnU6qLbRvo7qEZo",
        "mode": "live",
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
        "createdAt": "2024-02-12T11:58:35.0Z",
        "expiresAt": "2024-02-12T12:13:35.0Z",
        "reusable": false,
        "_links": {
          "self": {
            "href": "...",
            "type": "application/hal+json"
          },
          "checkout": {
            "href": "https://www.mollie.com/checkout/issuer/select/ideal/7UhSN1zuXS",
            "type": "text/html"
          },
          "dashboard": {
            "href": "https://www.mollie.com/dashboard/org_12345678/payments/tr_5B8cwPMGnU6qLbRvo7qEZo",
            "type": "text/html"
          }
        }
      }
    ]
  },
  "_links": {
    "self": {
      "href": "...",
      "type": "application/hal+json"
    },
    "previous": null,
    "next": {
      "href": "https://api.mollie.com/v2/payment-links/pl_4Y0eZitmBnQ6IDoMqZQKh/payments?from=tr_SDkzMggpvx&limit=5",
      "type": "application/hal+json"
    },
    "documentation": {
      "href": "...",
      "type": "text/html"
    }
  }
}`
