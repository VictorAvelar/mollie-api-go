package testdata

// ListMethodsResponse example
var ListMethodsResponse = `{
    "count": 13,
    "_embedded": {
        "methods": [
            {
                 "resource": "method",
                 "id": "ideal",
                 "description": "iDEAL",
                 "minimumAmount": {
                     "value": "0.01",
                     "currency": "EUR"
                 },
                 "maximumAmount": {
                     "value": "50000.00",
                     "currency": "EUR"
                 },
                 "image": {
                     "size1x": "https://mollie.com/external/icons/payment-methods/ideal.png",
                     "size2x": "https://mollie.com/external/icons/payment-methods/ideal%402x.png",
                     "svg": "https://mollie.com/external/icons/payment-methods/ideal.svg"
                 },
                 "pricing": [
                     {
                         "description": "Netherlands",
                         "fixed": {
                             "value": "0.29",
                             "currency": "EUR"
                         },
                         "variable": "0"
                     }
                 ],
                 "_links": {
                     "self": {
                         "href": "https://api.mollie.com/v2/methods/ideal",
                         "type": "application/hal+json"
                     }
                 }
            },
            {
                 "resource": "method",
                 "id": "creditcard",
                 "description": "Credit card",
                 "minimumAmount": {
                     "value": "0.01",
                     "currency": "EUR"
                 },
                 "maximumAmount": {
                     "value": "2000.00",
                     "currency": "EUR"
                 },
                 "image": {
                     "size1x": "https://mollie.com/external/icons/payment-methods/creditcard.png",
                     "size2x": "https://mollie.com/external/icons/payment-methods/creditcard%402x.png",
                     "svg": "https://mollie.com/external/icons/payment-methods/creditcard.svg"
                 },
                 "pricing": [
                     {
                         "description": "Commercial & non-European cards",
                         "fixed": {
                             "value": "0.25",
                             "currency": "EUR"
                         },
                         "variable": "2.8"
                     },
                     {
                         "description": "European cards",
                         "fixed": {
                             "value": "0.25",
                             "currency": "EUR"
                         },
                         "variable": "1.8"
                     },
                     {
                         "description": "American Express",
                         "fixed": {
                             "value": "0.25",
                             "currency": "EUR"
                         },
                         "variable": "2.8"
                     }
                 ],
                 "_links": {
                     "self": {
                         "href": "https://api.mollie.com/v2/methods/creditcard",
                         "type": "application/hal+json"
                     }
                 }
            }
        ]
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/methods",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/methods-api/list-methods",
            "type": "text/html"
        }
    }
}`

// GetMethodResponse example
var GetMethodResponse = `{
     "resource": "method",
     "id": "ideal",
     "description": "iDEAL",
     "minimumAmount": {
         "value": "0.01",
         "currency": "EUR"
     },
     "maximumAmount": {
         "value": "50000.00",
         "currency": "EUR"
     },
     "image": {
         "size1x": "https://www.mollie.com/external/icons/payment-methods/ideal.png",
         "size2x": "https://www.mollie.com/external/icons/payment-methods/ideal%402x.png",
         "svg": "https://www.mollie.com/external/icons/payment-methods/ideal.svg"
     },
     "issuers": [
         {
             "resource": "issuer",
             "id": "ideal_ABNANL2A",
             "name": "ABN AMRO",
             "image": {
                 "size1x": "https://www.mollie.com/external/icons/ideal-issuers/ABNANL2A.png",
                 "size2x": "https://www.mollie.com/external/icons/ideal-issuers/ABNANL2A%402x.png",
                 "svg": "https://www.mollie.com/external/icons/ideal-issuers/ABNANL2A.svg"
             }
         },
         {
             "resource": "issuer",
             "id": "ideal_ASNBNL21",
             "name": "ASN Bank",
             "image": {
                 "size1x": "https://www.mollie.com/external/icons/ideal-issuers/ASNBNL21.png",
                 "size2x": "https://www.mollie.com/external/icons/ideal-issuers/ASNBNL21%402x.png",
                 "svg": "https://www.mollie.com/external/icons/ideal-issuers/ASNBNL21.svg"
             }
         },
         { },
         { }
     ],
     "pricing": [
         {
             "description": "The Netherlands",
             "fixed": {
                 "value": "0.29",
                 "currency": "EUR"
             },
             "variable": "0"
         }
     ],
     "_links": {
         "self": {
             "href": "https://api.mollie.com/v2/methods/ideal",
             "type": "application/hal+json"
         },
         "documentation": {
             "href": "https://docs.mollie.com/reference/v2/methods-api/get-method",
             "type": "text/html"
         }
     }
 }`
