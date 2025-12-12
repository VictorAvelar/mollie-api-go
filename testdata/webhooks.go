package testdata

const CreateWebhookExample = `{
  "resource": "webhook",
  "id": "hook_B2EyhTH5N4KWUnoYPcgiH",
  "url": "https://mollie.com",
  "profileId": "pfl_8XcSdLtrNK",
  "createdAt": "2024-12-06T10:09:56+00:00",
  "name": "Webhook #1",
  "status": "enabled",
  "mode": "live",
  "webhookSecret": "VpQ3WukU6uSCGQ8TPTD3WPDpac3GyNEj",
  "eventTypes": [
    "payment-link.paid"
  ],
  "_links": {
    "self": {
      "href": "...",
      "type": "application/hal+json"
    },
    "documentation": {
      "href": "https://docs.mollie.com/reference/create-webhook",
      "type": "text/html"
    }
  }
}`

const ListWebhooksExample = `{
  "_embedded": {
    "webhooks": [
      {
        "resource": "webhook",
        "id": "hook_yjtBMWDCGw5YFSPQ3HPzH",
        "url": "https://mollie.com",
        "profileId": "pfl_8XcSdLtrNK",
        "createdAt": "2024-12-06T10:09:56+00:00",
        "name": "Webhook #2",
        "eventTypes": [
          "payment-link.paid",
          "sales-invoice.created"
        ],
        "status": "enabled",
        "mode": "live",
        "_links": {
          "self": {
            "href": "...",
            "type": "application/hal+json"
          },
          "documentation": {
            "href": "https://docs.mollie.com/reference/get-webhook",
            "type": "text/html"
          }
        }
      },
      {
        "resource": "webhook",
        "id": "hook_qHknfTxaDx6s8JNhuGPzH",
        "url": "https://mollie.com",
        "profileId": "pfl_8XcSdLtrNK",
        "createdAt": "2024-12-06T10:08:48+00:00",
        "name": "Webhook #1",
        "eventTypes": [
          "payment-link.paid"
        ],
        "status": "blocked",
        "mode": "live",
        "_links": {
          "self": {
            "href": "...",
            "type": "application/hal+json"
          },
          "documentation": {
            "href": "https://docs.mollie.com/reference/get-webhook",
            "type": "text/html"
          }
        }
      }
    ]
  },
  "count": 2,
  "_links": {
    "documentation": {
      "href": "https://docs.mollie.com/reference/list-webhook",
      "type": "text/html"
    },
    "self": {
      "href": "https://api.mollie.localhost/v2/webhooks?from=hook_yjtBMWDCGw5YFSPQ3HPzH&limit=2",
      "type": "application/hal+json"
    },
    "previous": {
      "href": "https://api.mollie.localhost/v2/webhooks?from=hook_5foxphpBru4xNPCDJJPzH&limit=2",
      "type": "application/hal+json"
    },
    "next": {
      "href": "https://api.mollie.localhost/v2/webhooks?from=hook_fTqARmWsfs9oXvKbZEPzH&limit=2",
      "type": "application/hal+json"
    }
  }
}`

const GetWebhookExample = `{
  "resource": "webhook",
  "id": "hook_B2EyhTH5N4KWUnoYPcgiH",
  "url": "https://mollie.com",
  "profileId": "pfl_8XcSdLtrNK",
  "createdAt": "2024-12-06T10:09:56+00:00",
  "name": "Webhook #1",
  "status": "enabled",
  "mode": "test",
  "eventTypes": [
    "payment-link.paid",
    "sales-invoice.paid"
  ],
  "_links": {
    "self": {
      "href": "...",
      "type": "application/hal+json"
    },
    "documentation": {
      "href": "https://docs.mollie.com/reference/get-webhook",
      "type": "text/html"
    }
  }
}`
