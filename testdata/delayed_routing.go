package testdata

const GetDelayedRoutingExample = `{
  "resource": "route",
  "id": "crt_dyARQ3JzCgtPDhU2Pbq3J",
  "paymentId": "tr_sva3o24VKn",
  "amount": {
    "value": "10.00",
    "currency": "EUR"
  },
  "description": "Payment for Order #12345",
  "destination": {
    "type": "organization",
    "organizationId": "org_123"
  },
  "_links": {
    "self": {
      "href": "...",
      "type": "application/hal+json"
    },
    "documentation": {
      "href": "...",
      "type": "text/html"
    }
  }
}`

const ListDelayedRoutingsExample = `{
  "count": 3,
  "_embedded": {
    "routes": [
      {
        "resource": "route",
        "id": "crt_kSq9CDTa6FzW7c6gpYq3J",
        "paymentId": "tr_sva3o24VKn",
        "createdAt": "2024-05-22T15:11:19+02:00",
        "description": "Payment for Order #12345",
        "amount": {
          "value": "1.00",
          "currency": "EUR"
        },
        "destination": {
          "type": "organization",
          "organizationId": "org_123"
        },
        "_links": {
          "self": {
            "href": "...",
            "type": "application/hal+json"
          },
          "documentation": {
            "href": "...",
            "type": "text/html"
          }
        }
      },
      {
        "resource": "route",
        "id": "crt_dyARQ3JzCgtPDhU2Pbq3J",
        "paymentId": "tr_sva3o24VKn",
        "createdAt": "2024-05-22T15:11:19+02:00",
        "description": "Payment for Order #12346",
        "amount": {
          "value": "2.00",
          "currency": "EUR"
        },
        "destination": {
          "type": "organization",
          "organizationId": "org_456"
        },
        "_links": {
          "self": {
            "href": "...",
            "type": "application/hal+json"
          },
          "documentation": {
            "href": "...",
            "type": "text/html"
          }
        }
      },
      {
        "resource": "route",
        "id": "crt_tntKsr6tffuVdqnEvhq3J",
        "paymentId": "tr_sva3o24VKn",
        "createdAt": "2024-05-22T15:11:19+02:00",
        "description": "Payment for Order #12347",
        "amount": {
          "value": "3.00",
          "currency": "EUR"
        },
        "destination": {
          "type": "organization",
          "organizationId": "org_789"
        },
        "_links": {
          "self": {
            "href": "...",
            "type": "application/hal+json"
          },
          "documentation": {
            "href": "...",
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
    "documentation": {
      "href": "...",
      "type": "text/html"
    }
  }
}`
