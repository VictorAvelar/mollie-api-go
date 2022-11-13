package testdata

// GetBalanceResponse example.
const GetBalanceResponse = `{
  "resource": "balance",
  "id": "bal_gVMhHKqSSRYJyPsuoPNFH",
  "mode": "live",
  "createdAt": "2019-01-10T10:23:41+00:00",
  "currency": "EUR",
  "status": "active",
  "availableAmount": {
    "value": "905.25",
    "currency": "EUR"
  },
  "pendingAmount": {
    "value": "0.00",
    "currency": "EUR"
  },
  "transferFrequency": "twice-a-month",
  "transferThreshold": {
    "value": "5.00",
    "currency": "EUR"
  },
  "transferReference": "Mollie payout",
  "transferDestination": {
    "type": "bank-account",
    "beneficiaryName": "Jack Bauer",
    "bankAccount": "NL53INGB0654422370",
    "bankAccountId": "bnk_jrty3f"
  },
  "_links": {
    "self": {
      "href": "https://api.mollie.com/v2/balances/bal_gVMhHKqSSRYJyPsuoPNFH",
      "type": "application/hal+json"
    },
    "documentation": {
      "href": "https://docs.mollie.com/reference/v2/balances-api/get-balance",
      "type": "text/html"
    }
  }
}`

const ListBalancesResponse = `{
  "count": 2,
  "_embedded": {
    "balances": [
       {
         "resource": "balance",
         "id": "bal_gVMhHKqSSRYJyPsuoPNFH",
         "mode": "live",
         "createdAt": "2019-01-10T12:06:28+00:00",
         "currency": "EUR",
         "status": "active",
         "availableAmount": {
           "value": "0.00",
           "currency": "EUR"
         },
         "pendingAmount": {
           "value": "0.00",
           "currency": "EUR"
         },
         "transferFrequency": "daily",
         "transferThreshold": {
           "value": "40.00",
           "currency": "EUR"
         },
         "transferReference": "Mollie payout",
         "transferDestination": {
           "type": "bank-account",
           "beneficiaryName": "Jack Bauer",
           "bankAccount": "NL53INGB0654422370",
           "bankAccountId": "bnk_jrty3f"
         },
         "_links": {
           "self": {
             "href": "https://api.mollie.com/v2/balances/bal_gVMhHKqSSRYJyPsuoPNFH",
             "type": "application/hal+json"
           }
         }
       },
       {
         "resource": "balance",
         "id": "bal_gVMhHKqSSRYJyPsuoPABC",
         "mode": "live",
         "createdAt": "2019-01-10T10:23:41+00:00",
         "status": "active",
         "currency": "EUR",
         "availableAmount": {
           "value": "0.00",
           "currency": "EUR"
         },
         "pendingAmount": {
           "value": "0.00",
           "currency": "EUR"
         },
         "transferFrequency": "twice-a-month",
         "transferThreshold": {
           "value": "5.00",
           "currency": "EUR"
         },
         "transferReference": "Mollie payout",
         "transferDestination": {
           "type": "bank-account",
           "beneficiaryName": "Jack Bauer",
           "bankAccount": "NL97MOLL6351480700",
           "bankAccountId": "bnk_jrty3e"
         },
         "_links": {
           "self": {
             "href": "https://api.mollie.com/v2/balances/bal_gVMhHKqSSRYJyPsuoPABC",
             "type": "application/hal+json"
           }
         }
       }
    ]
  },
  "_links": {
    "documentation": {
      "href": "https://docs.mollie.com/reference/v2/balances-api/list-balances",
      "type": "text/html"
    },
    "self": {
      "href": "https://api.mollie.com/v2/balances?limit=2",
      "type": "application/hal+json"
    },
    "previous": null,
    "next": {
      "href": "https://api.mollie.com/v2/balances?from=bal_gVMhHKqSSRYJyPsuoPABC&limit=2",
      "type": "application/hal+json"
    }
  }
}`
