package testdata

const GetWebhookEventExample = `{
  "resource": "event",
  "id": "event_GvJ8WHrp5isUdRub9CJyH",
  "type": "payment-link.paid",
  "entityId": "pl_qng5gbbv8NAZ5gpM5ZYgx",
  "createdAt": "2024-12-16T15:57:04+00:00",
  "_embedded": {
    "entity": {
      "resource": "payment-link",
      "id": "pl_4Y0eZitmBnQ6IDoMqZQKh",
      "mode": "live",
      "description": "Bicycle tires",
      "amount": {
        "currency": "EUR",
        "value": "24.95"
      },
      "archived": false,
      "redirectUrl": "https://webshop.example.org/thanks",
      "webhookUrl": "https://webshop.example.org/payment-links/webhook",
      "profileId": "pfl_QkEhN94Ba",
      "createdAt": "2021-03-20T09:29:56+00:00",
      "paidAt": "2022-03-20T09:29:56+00:00",
      "expiresAt": "2023-06-06T11:00:00+00:00",
      "reusable": false,
      "allowedMethods": [
        "ideal"
      ],
      "sequenceType": "oneoff",
      "customerId": null,
      "_links": {
        "self": {
          "href": "https://api.mollie.com/v2/payment-links/pl_qng5gbbv8NAZ5gpM5ZYgx",
          "type": "application/hal+json"
        },
        "paymentLink": {
          "href": "https://www.mollie.com/paymentscreen/example",
          "type": "text/html"
        },
        "documentation": {
          "href": "https://docs.mollie.com/reference/v2/payment-links-api/get-payment-link",
          "type": "text/html"
        }
      }
    }
  },
  "_links": {
    "self": {
      "href": "https://api.mollie.com/v2/events/event_GvJ8WHrp5isUdRub9CJyH",
      "type": "application/hal+json"
    },
    "documentation": {
      "href": "https://docs.mollie.com/guides/webhooks",
      "type": "text/html"
    },
    "entity": {
      "href": "/v2/payment-links/pl_qng5gbbv8NAZ5gpM5ZYgx",
      "type": "application/hal+json"
    }
  }
}`
