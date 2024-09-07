package testdata

const ListPartnerClientsResponse = `{
   "count":1,
   "_embedded":{
      "clients":[
         {
            "resource":"client",
            "id":"org_1337",
            "organizationCreatedAt":"2018-03-21T13:13:37+00:00",
            "commission":{
               "count":200,
               "totalAmount":{
                  "currency":"EUR",
                  "value":"10.00"
               }
            },
            "_links":{
               "self":{
                  "href":"https://api.mollie.com/v2/clients/org_1337",
                  "type":"application/hal+json"
               },
               "organization":{
                  "href":"https://api.mollie.com/v2/organizations/org_1337",
                  "type":"application/hal+json"
               },
               "onboarding":{
                  "href":"https://api.mollie.com/v2/onboarding/org_1337",
                  "type":"application/hal+json"
               },
               "documentation":{
                  "href":"https://docs.mollie.com/reference/partners-api/get-client",
                  "type":"text/html"
               }
            }
         }
      ]
   },
   "_links":{
      "self":{
         "href":"https://api.mollie.com/v2/clients?limit=3",
         "type":"application/hal+json"
      },
      "previous":null,
      "next":{
         "href":"https://api.mollie.com/v2/clients?from=org_1379&limit=3",
         "type":"application/hal+json"
      },
      "documentation":{
         "href":"https://docs.mollie.com/reference/partners-api/list-clients",
         "type":"text/html"
      }
   }
}`

const GetPartnerClientResponse = `{
    "resource": "client",
    "id": "org_1337",
    "organizationCreatedAt": "2018-03-21T13:13:37+00:00",
    "commission": {
        "count": 200,
        "totalAmount": {
            "currency": "EUR",
            "value": "10.00"
        }
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/clients/org_1337",
            "type": "application/hal+json"
        },
        "organization": {
            "href": "https://api.mollie.com/v2/organizations/org_1337",
            "type": "application/hal+json"
        },
        "onboarding": {
            "href": "https://api.mollie.com/v2/onboarding/org_1337",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/partners-api/get-client",
            "type": "text/html"
        }
    }
}`
