package testdata

const CreateSalesInvoicesResponse = `{
  "resource": "sales-invoice",
  "id": "invoice_4Y0eZitmBnQ6IDoMqZQKh",
  "mode": "live",
  "profileId": "pfl_QkEhN94Ba",
  "currency": "EUR",
  "status": "draft",
  "vatScheme": "standard",
  "paymentTerm": "30 days",
  "recipientIdentifier": "123532354",
  "recipient": {
    "type": "consumer",
    "givenName": "Given",
    "familyName": "Family",
    "email": "given.family@mollie.com",
    "streetAndNumber": "Street 1",
    "postalCode": "1000 AA",
    "city": "Amsterdam",
    "country": "NL",
    "locale": "nl_NL"
  },
  "lines": [
    {
      "description": "LEGO 4440 Forest Police Station",
      "quantity": 1,
      "vatRate": "21.00",
      "unitPrice": {
        "value": "89.00",
        "currency": "EUR"
      }
    }
  ],
  "amountDue": {
    "value": "107.69",
    "currency": "EUR"
  },
  "subtotalAmount": {
    "value": "89.00",
    "currency": "EUR"
  },
  "totalAmount": {
    "value": "107.69",
    "currency": "EUR"
  },
  "totalVatAmount": {
    "value": "18.69",
    "currency": "EUR"
  },
  "discountedSubtotalAmount": {
    "value": "89.00",
    "currency": "EUR"
  },
  "createdAt": "2024-10-03T10:47:38+00:00",
  "_links": {
    "self": {
      "href": "...",
      "type": "application/hal+json"
    },
    "invoicePayment": {
      "href": "...",
      "type": "application/hal+json"
    },
    "pdfLink": {
      "href": "...",
      "type": "application/hal+json"
    },
    "documentation": {
      "href": "...",
      "type": "text/html"
    }
  }
}`

const GetSalesInvoicesResponse = `{
  "resource": "sales-invoice",
  "id": "invoice_4Y0eZitmBnQ6IDoMqZQKh",
  "mode": "live",
  "profileId": "pfl_QkEhN94Ba",
  "currency": "EUR",
  "status": "draft",
  "vatScheme": "standard",
  "paymentTerm": "30 days",
  "recipientIdentifier": "123532354",
  "recipient": {
    "type": "consumer",
    "givenName": "Given",
    "familyName": "Family",
    "email": "given.family@mollie.com",
    "streetAndNumber": "Street 1",
    "postalCode": "1000 AA",
    "city": "Amsterdam",
    "country": "NL",
    "locale": "nl_NL"
  },
  "lines": [
    {
      "description": "LEGO 4440 Forest Police Station",
      "quantity": 1,
      "vatRate": "21.00",
      "unitPrice": {
        "value": "89.00",
        "currency": "EUR"
      }
    }
  ],
  "amountDue": {
    "value": "107.69",
    "currency": "EUR"
  },
  "subtotalAmount": {
    "value": "89.00",
    "currency": "EUR"
  },
  "totalAmount": {
    "value": "107.69",
    "currency": "EUR"
  },
  "totalVatAmount": {
    "value": "18.69",
    "currency": "EUR"
  },
  "discountedSubtotalAmount": {
    "value": "89.00",
    "currency": "EUR"
  },
  "createdAt": "2024-10-03T10:47:38+00:00",
  "_links": {
    "self": {
      "href": "...",
      "type": "application/hal+json"
    },
    "invoicePayment": {
      "href": "...",
      "type": "application/hal+json"
    },
    "pdfLink": {
      "href": "...",
      "type": "application/hal+json"
    },
    "documentation": {
      "href": "...",
      "type": "text/html"
    }
  }
}`

const ListSalesInvoicesResponse = `{
  "count": 1,
  "_embedded": {
    "sales_invoices": [
      {
        "resource": "sales-invoice",
        "id": "invoice_4Y0eZitmBnQ6IDoMqZQKh",
        "mode": "live",
        "profileId": "pfl_QkEhN94Ba",
        "currency": "EUR",
        "status": "draft",
        "vatScheme": "standard",
        "paymentTerm": "30 days",
        "recipientIdentifier": "123532354",
        "lines": [
          {
            "description": "LEGO 4440 Forest Police Station",
            "quantity": 1,
            "vatRate": "21.00",
            "unitPrice": {
              "value": "89.00",
              "currency": "EUR"
            }
          }
        ],
        "amountDue": {
          "value": "107.69",
          "currency": "EUR"
        },
        "subtotalAmount": {
          "value": "89.00",
          "currency": "EUR"
        },
        "totalAmount": {
          "value": "107.69",
          "currency": "EUR"
        },
        "totalVatAmount": {
          "value": "18.69",
          "currency": "EUR"
        },
        "discountedSubtotalAmount": {
          "value": "89.00",
          "currency": "EUR"
        },
        "createdAt": "2024-10-03T10:47:38+00:00"
      }
    ]
  },
  "_links": {
    "self": {
      "href": "...",
      "type": "application/hal+json"
    },
    "next": {
      "href": "https://api.mollie.com/v2/sales/invoices?from=invoice_4yUfQpbKnd2DUTouUdUwH&limit=5",
      "type": "application/hal+json"
    },
    "documentation": {
      "href": "...",
      "type": "text/html"
    },
    "previous": null
  }
}`
