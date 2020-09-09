package testdata

// GetOrderResponse example
const GetOrderResponse = `{
     "resource": "order",
     "id": "ord_kEn1PlbGa",
     "profileId": "pfl_URR55HPMGx",
     "method": "ideal",
     "amount": {
         "value": "1027.99",
         "currency": "EUR"
     },
     "status": "created",
     "isCancelable": true,
     "metadata": null,
     "createdAt": "2018-08-02T09:29:56+00:00",
     "expiresAt": "2018-08-30T09:29:56+00:00",
     "mode": "live",
     "locale": "nl_NL",
     "billingAddress": {
         "organizationName": "Mollie B.V.",
         "streetAndNumber": "Keizersgracht 313",
         "postalCode": "1016 EE",
         "city": "Amsterdam",
         "country": "nl",
         "givenName": "Luke",
         "familyName": "Skywalker",
         "email": "luke@skywalker.com"
     },
     "shopperCountryMustMatchBillingCountry": false,
     "consumerDateOfBirth": "1993-10-21",
     "orderNumber": "18475",
     "shippingAddress": {
         "organizationName": "Mollie B.V.",
         "streetAndNumber": "Keizersgracht 313",
         "postalCode": "1016 EE",
         "city": "Amsterdam",
         "country": "nl",
         "givenName": "Luke",
         "familyName": "Skywalker",
         "email": "luke@skywalker.com"
     },
     "redirectUrl": "https://example.org/redirect",
     "lines": [
         {
             "resource": "orderline",
             "id": "odl_dgtxyl",
             "orderId": "ord_pbjz8x",
             "name": "LEGO 42083 Bugatti Chiron",
             "sku": "5702016116977",
             "type": "physical",
             "status": "created",
             "metadata": null,
             "isCancelable": false,
             "quantity": 2,
             "quantityShipped": 0,
             "amountShipped": {
                 "value": "0.00",
                 "currency": "EUR"
             },
             "quantityRefunded": 0,
             "amountRefunded": {
                 "value": "0.00",
                 "currency": "EUR"
             },
             "quantityCanceled": 0,
             "amountCanceled": {
                 "value": "0.00",
                 "currency": "EUR"
             },
             "shippableQuantity": 0,
             "refundableQuantity": 0,
             "cancelableQuantity": 0,
             "unitPrice": {
                 "value": "399.00",
                 "currency": "EUR"
             },
             "vatRate": "21.00",
             "vatAmount": {
                 "value": "121.14",
                 "currency": "EUR"
             },
             "discountAmount": {
                 "value": "100.00",
                 "currency": "EUR"
             },
             "totalAmount": {
                 "value": "698.00",
                 "currency": "EUR"
             },
             "createdAt": "2018-08-02T09:29:56+00:00",
             "_links": {
                 "productUrl": {
                     "href": "https://shop.lego.com/nl-NL/Bugatti-Chiron-42083",
                     "type": "text/html"
                 },
                 "imageUrl": {
                     "href": "https://sh-s7-live-s.legocdn.com/is/image//LEGO/42083_alt1?$main$",
                     "type": "text/html"
                 }
             }
         },
         {
             "resource": "orderline",
             "id": "odl_jp31jz",
             "orderId": "ord_pbjz8x",
             "name": "LEGO 42056 Porsche 911 GT3 RS",
             "sku": "5702015594028",
             "type": "physical",
             "status": "created",
             "metadata": null,
             "isCancelable": false,
             "quantity": 1,
             "quantityShipped": 0,
             "amountShipped": {
                 "value": "0.00",
                 "currency": "EUR"
             },
             "quantityRefunded": 0,
             "amountRefunded": {
                 "value": "0.00",
                 "currency": "EUR"
             },
             "quantityCanceled": 0,
             "amountCanceled": {
                 "value": "0.00",
                 "currency": "EUR"
             },
             "shippableQuantity": 0,
             "refundableQuantity": 0,
             "cancelableQuantity": 0,
             "unitPrice": {
                 "value": "329.99",
                 "currency": "EUR"
             },
             "vatRate": "21.00",
             "vatAmount": {
                 "value": "57.27",
                 "currency": "EUR"
             },
             "totalAmount": {
                 "value": "329.99",
                 "currency": "EUR"
             },
             "createdAt": "2018-08-02T09:29:56+00:00",
             "_links": {
                 "productUrl": {
                     "href": "https://shop.lego.com/nl-NL/Porsche-911-GT3-RS-42056",
                     "type": "text/html"
                 },
                 "imageUrl": {
                     "href": "https://sh-s7-live-s.legocdn.com/is/image/LEGO/42056?$PDPDefault$",
                     "type": "text/html"
                 }
             }
         }
     ],
     "_embedded": {
         "payments": [
             {
                 "resource": "payment",
                 "id": "tr_ncaPcAhuUV",
                 "mode": "live",
                 "createdAt": "2018-09-07T12:00:05+00:00",
                 "amount": {
                     "value": "1027.99",
                     "currency": "EUR"
                 },
                 "description": "Order #1337 (Lego cars)",
                 "method": null,
                 "metadata": null,
                 "status": "open",
                 "isCancelable": false,
                 "locale": "nl_NL",
                 "profileId": "pfl_URR55HPMGx",
                 "orderId": "ord_kEn1PlbGa",
                 "sequenceType": "oneoff",
                 "redirectUrl": "https://example.org/redirect",
                 "_links": {
                     "self": {
                         "href": "https://api.mollie.com/v2/payments/tr_ncaPcAhuUV",
                         "type": "application/hal+json"
                     },
                     "checkout": {
                         "href": "https://www.mollie.com/payscreen/select-method/ncaPcAhuUV",
                         "type": "text/html"
                     },
                     "dashboard": {
                         "href": "https://www.mollie.com/dashboard/org_123456789/payments/tr_ncaPcAhuUV",
                         "type": "text/html"
                     },
                     "order": {
                         "href": "https://api.mollie.com/v2/orders/ord_kEn1PlbGa",
                         "type": "application/hal+json"
                     }
                 }
             }
         ],
         "refunds": [
             {
                 "resource": "refund",
                 "id": "re_vD3Jm32wQt",
                 "amount": {
                     "value": "329.99",
                     "currency": "EUR"
                 },
                 "status": "pending",
                 "createdAt": "2019-01-15T15:41:21+00:00",
                 "description": "Required quantity not in stock, refunding one photo book.",
                 "orderId": "ord_kEn1PlbGa",
                 "paymentId": "tr_mjvPwykz3x",
                 "settlementAmount": {
                     "value": "-329.99",
                     "currency": "EUR"
                 },
                 "lines": [
                     {
                         "resource": "orderline",
                         "id": "odl_dgtxyl",
                         "orderId": "ord_kEn1PlbGa",
                         "name": "LEGO 42056 Porsche 911 GT3 RS",
                         "sku": "5702015594028",
                         "type": "physical",
                         "status": "completed",
                         "isCancelable": false,
                         "quantity": 1,
                         "unitPrice": {
                             "value": "329.99",
                             "currency": "EUR"
                         },
                         "vatRate": "21.00",
                         "vatAmount": {
                             "value": "57.27",
                             "currency": "EUR"
                         },
                         "totalAmount": {
                             "value": "329.99",
                             "currency": "EUR"
                         },
                         "createdAt": "2019-01-15T15:22:45+00:00",
                         "_links": {
                             "productUrl": {
                                 "href": "https://shop.lego.com/nl-NL/Porsche-911-GT3-RS-42056",
                                 "type": "text/html"
                             },
                             "imageUrl": {
                                 "href": "https://sh-s7-live-s.legocdn.com/is/image/LEGO/42056?$PDPDefault$",
                                 "type": "text/html"
                             }
                         }
                     }
                 ],
                 "_links": {
                     "self": {
                         "href": "https://api.mollie.com/v2/payments/tr_mjvPwykz3x/refunds/re_vD3Jm32wQt",
                         "type": "application/hal+json"
                     },
                     "payment": {
                         "href": "https://api.mollie.com/v2/payments/tr_mjvPwykz3x",
                         "type": "application/hal+json"
                     },
                     "order": {
                         "href": "https://api.mollie.com/v2/orders/ord_kEn1PlbGa",
                         "type": "application/hal+json"
                     }
                 }
             }
         ]
     },
     "_links": {
         "self": {
             "href": "https://api.mollie.com/v2/orders/ord_pbjz8x",
             "type": "application/hal+json"
         },
         "checkout": {
             "href": "https://www.mollie.com/payscreen/order/checkout/pbjz8x",
             "type": "text/html"
         },
        "dashboard": {
            "href": "https://www.mollie.com/dashboard/org_123456789/orders/ord_pbjz8x",
            "type": "text/html"
        },
         "documentation": {
             "href": "https://docs.mollie.com/reference/v2/orders-api/get-order",
             "type": "text/html"
         }
     }
 }`

// CreateOrderRequest example of create order request
const CreateOrderRequest = `{
	"amount": {
	   "value": "1027.99",
	   "currency": "EUR"
	},
	"billingAddress": {
	   "organizationName": "Mollie B.V.",
	   "streetAndNumber": "Keizersgracht 313",
	   "city": "Amsterdam",
	   "region": "Noord-Holland",
	   "postalCode": "1234AB",
	   "country": "NL",
	   "title": "Dhr",
	   "givenName": "Piet",
	   "familyName": "Mondriaan",
	   "email": "piet@mondriaan.com",
	   "phone": "+31208202070"
	},
	"shippingAddress": {
	   "organizationName": "Mollie B.V.",
	   "streetAndNumber": "Prinsengracht 313",
	   "streetAdditional": "4th floor",
	   "city": "Haarlem",
	   "region": "Noord-Holland",
	   "postalCode": "5678AB",
	   "country": "NL",
	   "title": "Mr",
	   "givenName": "Chuck",
	   "familyName": "Norris",
	   "email": "norris@chucknorrisfacts.net"
	},
	"metadata": {
	   "order_id": "1337",
	   "description": "Lego cars"
	},
	"consumerDateOfBirth": "1958-01-31",
	"locale": "nl_NL",
	"orderNumber": "1337",
	"redirectUrl": "https://example.org/redirect",
	"webhookUrl": "https://example.org/webhook",
	"method": "klarnapaylater",
	"lines": [
	   {
			 "type": "physical",
			 "sku": "5702016116977",
			 "name": "LEGO 42083 Bugatti Chiron",
			 "productUrl": "https://shop.lego.com/nl-NL/Bugatti-Chiron-42083",
			 "imageUrl": "https://sh-s7-live-s.legocdn.com/is/image//LEGO/42083_alt1?$main$",
			 "metadata": {
				"order_id": "1337",
				"description": "Bugatti Chiron"
			 },
			 "quantity": 2,
			 "vatRate": "21.00",
			 "unitPrice": {
				"currency": "EUR",
				"value": "399.00"
			 },
			 "totalAmount": {
				"currency": "EUR",
				"value": "698.00"
			 },
			 "discountAmount": {
				"currency": "EUR",
				"value": "100.00"
			 },
			 "vatAmount": {
				"currency": "EUR",
				"value": "121.14"
			 }
	   },
	   {
			 "type": "physical",
			 "sku": "5702015594028",
			 "name": "LEGO 42056 Porsche 911 GT3 RS",
			 "productUrl": "https://shop.lego.com/nl-NL/Porsche-911-GT3-RS-42056",
			 "imageUrl": "https://sh-s7-live-s.legocdn.com/is/image/LEGO/42056?$PDPDefault$",
			 "quantity": 1,
			 "vatRate": "21.00",
			 "unitPrice": {
				"currency": "EUR",
				"value": "329.99"
			 },
			 "totalAmount": {
				"currency": "EUR",
				"value": "329.99"
			 },
			 "vatAmount": {
				"currency": "EUR",
				"value": "57.27"
			 }
	   }
	]
 }`

// CreateOrderResponse example response of create order
const CreateOrderResponse = `{
    "resource": "order",
    "id": "ord_pbjz8x",
    "profileId": "pfl_URR55HPMGx",
    "method": "klarnapaylater",
    "amount": {
        "value": "1027.99",
        "currency": "EUR"
    },
    "status": "created",
    "isCancelable": true,
    "metadata": {
        "order_id": "1337",
        "description": "Lego cars"
    },
    "createdAt": "2018-08-02T09:29:56+00:00",
    "expiresAt": "2018-08-30T09:29:56+00:00",
    "mode": "test",
    "locale": "nl_NL",
    "billingAddress": {
        "organizationName": "Mollie B.V.",
        "streetAndNumber": "Keizersgracht 313",
        "city": "Amsterdam",
        "region": "Noord-Holland",
        "postalCode": "1234AB",
        "country": "NL",
        "title": "Dhr.",
        "givenName": "Piet",
        "familyName": "Mondriaan",
        "email": "piet@mondriaan.com",
        "phone": "+31309202070"
    },
    "consumerDateOfBirth": "1958-01-31",
    "orderNumber": "1337",
    "shippingAddress": {
        "organizationName": "Mollie B.V.",
        "streetAndNumber": "Keizersgracht 313",
        "streetAdditional": "4th floor",
        "city": "Haarlem",
        "region": "Noord-Holland",
        "postalCode": "5678AB",
        "country": "NL",
        "title": "Mr.",
        "givenName": "Chuck",
        "familyName": "Norris",
        "email": "norris@chucknorrisfacts.net"
    },
    "redirectUrl": "https://example.org/redirect",
    "webhookUrl": "https://example.org/webhook",
    "lines": [
        {
            "resource": "orderline",
            "id": "odl_dgtxyl",
            "orderId": "ord_pbjz8x",
            "name": "LEGO 42083 Bugatti Chiron",
            "sku": "5702016116977",
            "type": "physical",
            "status": "created",
            "metadata": {
               "order_id": "1337",
               "description": "Bugatti Chiron"
            },
            "isCancelable": false,
            "quantity": 2,
            "quantityShipped": 0,
            "amountShipped": {
                "value": "0.00",
                "currency": "EUR"
            },
            "quantityRefunded": 0,
            "amountRefunded": {
                "value": "0.00",
                "currency": "EUR"
            },
            "quantityCanceled": 0,
            "amountCanceled": {
                "value": "0.00",
                "currency": "EUR"
            },
            "shippableQuantity": 0,
            "refundableQuantity": 0,
            "cancelableQuantity": 0,
            "unitPrice": {
                "value": "399.00",
                "currency": "EUR"
            },
            "vatRate": "21.00",
            "vatAmount": {
                "value": "121.14",
                "currency": "EUR"
            },
            "discountAmount": {
                "value": "100.00",
                "currency": "EUR"
            },
            "totalAmount": {
                "value": "698.00",
                "currency": "EUR"
            },
            "createdAt": "2018-08-02T09:29:56+00:00",
            "_links": {
                "productUrl": {
                    "href": "https://shop.lego.com/nl-NL/Bugatti-Chiron-42083",
                    "type": "text/html"
                },
                "imageUrl": {
                    "href": "https://sh-s7-live-s.legocdn.com/is/image//LEGO/42083_alt1?$main$",
                    "type": "text/html"
                }
            }
        },
        {
            "resource": "orderline",
            "id": "odl_jp31jz",
            "orderId": "ord_pbjz8x",
            "name": "LEGO 42056 Porsche 911 GT3 RS",
            "sku": "5702015594028",
            "type": "physical",
            "status": "created",
            "metadata": null,
            "isCancelable": false,
            "quantity": 1,
            "quantityShipped": 0,
            "amountShipped": {
                "value": "0.00",
                "currency": "EUR"
            },
            "quantityRefunded": 0,
            "amountRefunded": {
                "value": "0.00",
                "currency": "EUR"
            },
            "quantityCanceled": 0,
            "amountCanceled": {
                "value": "0.00",
                "currency": "EUR"
            },
            "shippableQuantity": 0,
            "refundableQuantity": 0,
            "cancelableQuantity": 0,
            "unitPrice": {
                "value": "329.99",
                "currency": "EUR"
            },
            "vatRate": "21.00",
            "vatAmount": {
                "value": "57.27",
                "currency": "EUR"
            },
            "totalAmount": {
                "value": "329.99",
                "currency": "EUR"
            },
            "createdAt": "2018-08-02T09:29:56+00:00",
            "_links": {
                "productUrl": {
                    "href": "https://shop.lego.com/nl-NL/Porsche-911-GT3-RS-42056",
                    "type": "text/html"
                },
                "imageUrl": {
                    "href": "https://sh-s7-live-s.legocdn.com/is/image/LEGO/42056?$PDPDefault$",
                    "type": "text/html"
                }
            }
        }
    ],
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/orders/ord_pbjz8x",
            "type": "application/hal+json"
        },
        "checkout": {
            "href": "https://www.mollie.com/payscreen/order/checkout/pbjz8x",
            "type": "text/html"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/orders-api/get-order",
            "type": "text/html"
        }
    }
}`

// UpdateOrderRequest is a request example of updating an order
const UpdateOrderRequest = `{
	"billingAddress": {
	   "organizationName": "Mollie B.V.",
	   "streetAndNumber": "Keizersgracht 313",
	   "city": "Amsterdam",
	   "region": "Noord-Holland",
	   "postalCode": "1234AB",
	   "country": "NL",
	   "title": "Dhr",
	   "givenName": "Piet",
	   "familyName": "Mondriaan",
	   "email": "piet@mondriaan.com",
	   "phone": "+31208202070"
	}
}`

// UpdateOrderResponse is a response example of updating an order
const UpdateOrderResponse = `{
	"resource": "order",
	"id": "ord_kEn1PlbGa",
	"profileId": "pfl_URR55HPMGx",
	"method": "ideal",
	"amount": {
		"value": "1027.99",
		"currency": "EUR"
	},
	"status": "created",
	"isCancelable": true,
	"metadata": null,
	"createdAt": "2018-08-02T09:29:56+00:00",
	"expiresAt": "2018-08-30T09:29:56+00:00",
	"mode": "live",
	"locale": "nl_NL",
	"billingAddress": {
		"organizationName": "Mollie B.V.",
		"streetAndNumber": "Keizersgracht 313",
		"city": "Amsterdam",
		"region": "Noord-Holland",
		"postalCode": "1234AB",
		"country": "NL",
		"title": "Dhr",
		"givenName": "Piet",
		"familyName": "Mondriaan",
		"email": "piet@mondriaan.com",
		"phone": "+31208202070"
	},
	"orderNumber": "18475",
	"shippingAddress": {
		"organizationName": "Mollie B.V.",
		"streetAndNumber": "Keizersgracht 313",
		"postalCode": "1016 EE",
		"city": "Amsterdam",
		"country": "nl",
		"givenName": "Luke",
		"familyName": "Skywalker",
		"email": "luke@skywalker.com"
	},
   "redirectUrl": "https://example.org/redirect",
	"lines": [
		{
			"resource": "orderline",
			"id": "odl_dgtxyl",
			"orderId": "ord_pbjz8x",
			"name": "LEGO 42083 Bugatti Chiron",
			"sku": "5702016116977",
			"type": "physical",
			"status": "created",
			"metadata": null,
			"isCancelable": false,
			"quantity": 2,
			"quantityShipped": 0,
			"amountShipped": {
				"value": "0.00",
				"currency": "EUR"
			},
			"quantityRefunded": 0,
			"amountRefunded": {
				"value": "0.00",
				"currency": "EUR"
			},
			"quantityCanceled": 0,
			"amountCanceled": {
				"value": "0.00",
				"currency": "EUR"
			},
		   "shippableQuantity": 0,
		   "refundableQuantity": 0,
		   "cancelableQuantity": 0,
			"unitPrice": {
				"value": "399.00",
				"currency": "EUR"
			},
			"vatRate": "21.00",
			"vatAmount": {
				"value": "121.14",
				"currency": "EUR"
			},
			"discountAmount": {
				"value": "100.00",
				"currency": "EUR"
			},
			"totalAmount": {
				"value": "698.00",
				"currency": "EUR"
			},
			"createdAt": "2018-08-02T09:29:56+00:00",
			"_links": {
				"productUrl": {
					"href": "https://shop.lego.com/nl-NL/Bugatti-Chiron-42083",
					"type": "text/html"
				},
				"imageUrl": {
					"href": "https://sh-s7-live-s.legocdn.com/is/image//LEGO/42083_alt1?$main$",
					"type": "text/html"
				}
			}
		},
		{
			"resource": "orderline",
			"id": "odl_jp31jz",
			"orderId": "ord_pbjz8x",
			"name": "LEGO 42056 Porsche 911 GT3 RS",
			"sku": "5702015594028",
			"type": "physical",
			"status": "created",
			"metadata": null,
			"isCancelable": false,
			"quantity": 1,
			"quantityShipped": 0,
			"amountShipped": {
				"value": "0.00",
				"currency": "EUR"
			},
			"quantityRefunded": 0,
			"amountRefunded": {
				"value": "0.00",
				"currency": "EUR"
			},
			"quantityCanceled": 0,
			"amountCanceled": {
				"value": "0.00",
				"currency": "EUR"
			},
		   "shippableQuantity": 0,
		   "refundableQuantity": 0,
		   "cancelableQuantity": 0,
			"unitPrice": {
				"value": "329.99",
				"currency": "EUR"
			},
			"vatRate": "21.00",
			"vatAmount": {
				"value": "57.27",
				"currency": "EUR"
			},
			"totalAmount": {
				"value": "329.99",
				"currency": "EUR"
			},
			"createdAt": "2018-08-02T09:29:56+00:00",
			"_links": {
				"productUrl": {
					"href": "https://shop.lego.com/nl-NL/Porsche-911-GT3-RS-42056",
					"type": "text/html"
				},
				"imageUrl": {
					"href": "https://sh-s7-live-s.legocdn.com/is/image/LEGO/42056?$PDPDefault$",
					"type": "text/html"
				}
			}
		}
	],
	"_links": {
		"self": {
			"href": "https://api.mollie.com/v2/orders/ord_pbjz8x",
			"type": "application/hal+json"
		},
		"checkout": {
			"href": "https://www.mollie.com/payscreen/order/checkout/pbjz8x",
			"type": "text/html"
		},
		"documentation": {
			"href": "https://docs.mollie.com/reference/v2/orders-api/get-order",
			"type": "text/html"
		}
	}
}`

// CancelOrderResponse example response of cancel order
const CancelOrderResponse = `{
	"resource": "order",
	"id": "ord_8wmqcHMN4U",
	"profileId": "pfl_URR55HPMGx",
	"amount": {
		"value": "1027.99",
		"currency": "EUR"
	},
	"status": "canceled",
	"isCancelable": false,
	"metadata": null,
	"createdAt": "2018-08-02T09:29:56+00:00",
	"mode": "live",
	"locale": "nl_NL",
	"billingAddress": {
		"organizationName": "Mollie B.V.",
		"streetAndNumber": "Keizersgracht 313",
		"postalCode": "1016 EE",
		"city": "Amsterdam",
		"country": "nl",
		"givenName": "Luke",
		"familyName": "Skywalker",
		"email": "luke@skywalker.com"
	},
	"orderNumber": "18475",
	"shippingAddress": {
		"organizationName": "Mollie B.V.",
		"streetAndNumber": "Keizersgracht 313",
		"postalCode": "1016 EE",
		"city": "Amsterdam",
		"country": "nl",
		"givenName": "Luke",
		"familyName": "Skywalker",
		"email": "luke@skywalker.com"
	},
	"canceledAt": "2018-08-03T09:29:56+00:00",
	"redirectUrl": "https://example.org/redirect",
	"lines": [
		{
			"resource": "orderline",
			"id": "odl_dgtxyl",
			"orderId": "ord_pbjz8x",
			"name": "LEGO 42083 Bugatti Chiron",
			"sku": "5702016116977",
			"type": "physical",
			"status": "canceled",
			"metadata": null,
			"isCancelable": false,
			"quantity": 2,
			"quantityShipped": 0,
			"amountShipped": {
				"value": "0.00",
				"currency": "EUR"
			},
			"quantityRefunded": 0,
			"amountRefunded": {
				"value": "0.00",
				"currency": "EUR"
			},
			"quantityCanceled": 2,
			"amountCanceled": {
				"value": "698.00",
				"currency": "EUR"
			},
			"shippableQuantity": 0,
			"refundableQuantity": 0,
			"cancelableQuantity": 0,
			"unitPrice": {
				"value": "399.00",
				"currency": "EUR"
			},
			"vatRate": "21.00",
			"vatAmount": {
				"value": "121.14",
				"currency": "EUR"
			},
			"discountAmount": {
				"value": "100.00",
				"currency": "EUR"
			},
			"totalAmount": {
				"value": "698.00",
				"currency": "EUR"
			},
			"createdAt": "2018-08-02T09:29:56+00:00",
			"_links": {
				"productUrl": {
					"href": "https://shop.lego.com/nl-NL/Bugatti-Chiron-42083",
					"type": "text/html"
				},
				"imageUrl": {
					"href": "https://sh-s7-live-s.legocdn.com/is/image//LEGO/42083_alt1?$main$",
					"type": "text/html"
				}
			}
		},
		{
			"resource": "orderline",
			"id": "odl_jp31jz",
			"orderId": "ord_pbjz8x",
			"name": "LEGO 42056 Porsche 911 GT3 RS",
			"sku": "5702015594028",
			"type": "physical",
			"status": "canceled",
			"metadata": null,
			"isCancelable": false,
			"quantity": 1,
			"quantityShipped": 0,
			"amountShipped": {
				"value": "0.00",
				"currency": "EUR"
			},
			"quantityRefunded": 0,
			"amountRefunded": {
				"value": "0.00",
				"currency": "EUR"
			},
			"quantityCanceled": 1,
			"amountCanceled": {
				"value": "329.99",
				"currency": "EUR"
			},
			"shippableQuantity": 0,
			"refundableQuantity": 0,
			"cancelableQuantity": 0,
			"unitPrice": {
				"value": "329.99",
				"currency": "EUR"
			},
			"vatRate": "21.00",
			"vatAmount": {
				"value": "57.27",
				"currency": "EUR"
			},
			"totalAmount": {
				"value": "329.99",
				"currency": "EUR"
			},
			"createdAt": "2018-08-02T09:29:56+00:00",
			"_links": {
				"productUrl": {
					"href": "https://shop.lego.com/nl-NL/Porsche-911-GT3-RS-42056",
					"type": "text/html"
				},
				"imageUrl": {
					"href": "https://sh-s7-live-s.legocdn.com/is/image/LEGO/42056?$PDPDefault$",
					"type": "text/html"
				}
			}
		}
	],
	"_links": {
		"self": {
			"href": "https://api.mollie.com/v2/orders/ord_pbjz8x",
			"type": "application/hal+json"
		},
		"checkout": {
			"href": "https://www.mollie.com/payscreen/order/checkout/pbjz8x",
			"type": "text/html"
		},
		"documentation": {
			"href": "https://docs.mollie.com/reference/v2/orders-api/get-order",
			"type": "text/html"
		}
	}
}`

// OrderListResponse example response of list orders
const OrderListResponse = `{
    "count": 3,
    "_embedded": {
        "orders": [
            {
                "resource": "order",
                "id": "ord_kEn1PlbGa",
                "...": "..."
            },
            { },
            { }
        ]
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/orders",
            "type": "application/hal+json"
        },
        "previous": null,
        "next": {
            "href": "https://api.mollie.com/v2/orders?from=ord_stTC2WHAuS",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/orders-api/list-orders",
            "type": "text/html"
        }
    }
}`

// UpdateOrderlineRequest example request of update orderline
const UpdateOrderlineRequest = `{
	"name": "LEGO 71043 Hogwarts™ Castle",
	"productUrl": "https://shop.lego.com/en-GB/product/Hogwarts-Castle-71043",
	"imageUrl": "https://sh-s7-live-s.legocdn.com/is/image//LEGO/71043_alt1?$main$",
	"quantity": 2,
	"vatRate": "21.00",
	"unitPrice": {
	   "currency": "EUR",
	   "value": "349.00"
	},
	"totalAmount": {
	   "currency": "EUR",
	   "value": "598.00"
	},
	"discountAmount": {
	   "currency": "EUR",
	   "value": "100.00"
	},
	"vatAmount": {
	   "currency": "EUR",
	   "value": "103.79"
	}
}`

// UpdateOrderlineResponse example response of update orderline
const UpdateOrderlineResponse = `{
	"resource":"order",
	"id":"ord_pbjz8x",
	"profileId":"pfl_URR55HPMGx",
	"method":"ideal",
	"amount":{
	   "value":"927.99",
	   "currency":"EUR"
	},
	"status":"created",
	"isCancelable":true,
	"metadata":null,
	"createdAt":"2018-08-02T09:29:56+00:00",
	"expiresAt":"2018-08-30T09:29:56+00:00",
	"mode":"live",
	"locale":"nl_NL",
	"billingAddress":{
	   "organizationName":"Mollie B.V.",
	   "streetAndNumber":"Keizersgracht 313",
	   "city":"Amsterdam",
	   "region":"Noord-Holland",
	   "postalCode":"1234AB",
	   "country":"NL",
	   "title":"Dhr",
	   "givenName":"Piet",
	   "familyName":"Mondriaan",
	   "email":"piet@mondriaan.com",
	   "phone":"+31208202070"
	},
	"orderNumber":"18475",
	"shippingAddress":{
	   "organizationName":"Mollie B.V.",
	   "streetAndNumber":"Keizersgracht 313",
	   "postalCode":"1016 EE",
	   "city":"Amsterdam",
	   "country":"nl",
	   "givenName":"Luke",
	   "familyName":"Skywalker",
	   "email":"luke@skywalker.com"
	},
	"redirectUrl":"https://example.org/redirect",
	"lines":[
	   {
		  "resource":"orderline",
		  "id":"odl_dgtxyl",
		  "orderId":"ord_pbjz8x",
		  "name":"LEGO 71043 Hogwarts™ Castle",
		  "sku":"5702016116977",
		  "type":"physical",
		  "status":"created",
		  "metadata":null,
		  "isCancelable":false,
		  "quantity":2,
		  "quantityShipped":0,
		  "amountShipped":{
			 "value":"0.00",
			 "currency":"EUR"
		  },
		  "quantityRefunded":0,
		  "amountRefunded":{
			 "value":"0.00",
			 "currency":"EUR"
		  },
		  "quantityCanceled":0,
		  "amountCanceled":{
			 "value":"0.00",
			 "currency":"EUR"
		  },
		  "shippableQuantity":0,
		  "refundableQuantity":0,
		  "cancelableQuantity":0,
		  "unitPrice":{
			 "currency":"EUR",
			 "value":"349.00"
		  },
		  "totalAmount":{
			 "currency":"EUR",
			 "value":"598.00"
		  },
		  "discountAmount":{
			 "currency":"EUR",
			 "value":"100.00"
		  },
		  "vatAmount":{
			 "currency":"EUR",
			 "value":"103.79"
		  },
		  "createdAt":"2018-08-02T09:29:56+00:00",
		  "_links":{
			 "productUrl":{
				"href":"https://shop.lego.com/en-GB/product/Hogwarts-Castle-71043",
				"type":"text/html"
			 },
			 "imageUrl":{
				"href":"https://sh-s7-live-s.legocdn.com/is/image//LEGO/71043_alt1?$main$",
				"type":"text/html"
			 }
		  }
	   },
	   {
		  "resource":"orderline",
		  "id":"odl_jp31jz",
		  "orderId":"ord_pbjz8x",
		  "name":"LEGO 42056 Porsche 911 GT3 RS",
		  "sku":"5702015594028",
		  "type":"physical",
		  "status":"created",
		  "metadata":null,
		  "isCancelable":false,
		  "quantity":1,
		  "quantityShipped":0,
		  "amountShipped":{
			 "value":"0.00",
			 "currency":"EUR"
		  },
		  "quantityRefunded":0,
		  "amountRefunded":{
			 "value":"0.00",
			 "currency":"EUR"
		  },
		  "quantityCanceled":0,
		  "amountCanceled":{
			 "value":"0.00",
			 "currency":"EUR"
		  },
		  "shippableQuantity":0,
		  "refundableQuantity":0,
		  "cancelableQuantity":0,
		  "unitPrice":{
			 "value":"329.99",
			 "currency":"EUR"
		  },
		  "vatRate":"21.00",
		  "vatAmount":{
			 "value":"57.27",
			 "currency":"EUR"
		  },
		  "totalAmount":{
			 "value":"329.99",
			 "currency":"EUR"
		  },
		  "createdAt":"2018-08-02T09:29:56+00:00",
		  "_links":{
			 "productUrl":{
				"href":"https://shop.lego.com/nl-NL/Porsche-911-GT3-RS-42056",
				"type":"text/html"
			 },
			 "imageUrl":{
				"href":"https://sh-s7-live-s.legocdn.com/is/image/LEGO/42056?$PDPDefault$",
				"type":"text/html"
			 }
		  }
	   }
	],
	"_links":{
	   "self":{
		  "href":"https://api.mollie.com/v2/orders/ord_pbjz8x",
		  "type":"application/hal+json"
	   },
	   "checkout":{
		  "href":"https://www.mollie.com/payscreen/order/checkout/pbjz8x",
		  "type":"text/html"
	   },
	   "documentation":{
		  "href":"https://docs.mollie.com/reference/v2/orders-api/get-order",
		  "type":"text/html"
	   }
	}
 }`

// CancelOrderLinesRequest example
const CancelOrderLinesRequest = `{
	"lines": [
	   {
		  "id": "odl_dgtxyl",
		  "quantity": 1
	   },
	   {
		  "id": "odl_jp31jz"
	   }
	]
}`

// CancelOrderLinesRequest example
const CancelOrderLinesRejectRequest = `{
	"lines": [
	   {
		  "id": "odl_dgtwkn",
		  "quantity": 1
	   },
	   {
		  "id": "odl_jp31jz"
	   }
	]
}`

// CancelOrderLinesResponseAmountRequired example
const CancelOrderLinesResponseAmountRequired = `{
	"status": 422,
	"title": "Unprocessable Entity",
	"detail": "Line 0 contains invalid data. An amount is required for this API call. The amount must be between €0.00 and €50.00.",
	"field": "lines.0.amount",
	"extra": {
		"minimumAmount": {
			"value": "0.00",
			"currency": "EUR"
		},
		"maximumAmount": {
			"value": "50.00",
			"currency": "EUR"
		}
	},
	"_links": {
		"documentation": {
			"href": "https://docs.mollie.com/reference/v2/orders-api/cancel-order-lines",
			"type": "text/html"
		}
	}
}`

// CancelOrderLinesResponseCancelReject example
const CancelOrderLinesResponseCancelReject = `{
    "status": 422,
    "title": "Unprocessable Entity",
    "detail": "Update authorization not allowed. Decision is based on order state and outcome of risk assessment.",
    "_links": {
        "documentation": {
            "href": "https://docs.mollie.com/guides/handling-errors",
            "type": "text/html"
        }
    }
}`

// CreateOrderPaymentRequest example request
const CreateOrderPaymentRequest = `{
	"method": "banktransfer"
}`

// CreateOrderPaymentResponse example response
const CreateOrderPaymentResponse = `{
    "resource": "payment",
    "id": "tr_WDqYK6vllg",
    "mode": "test",
    "amount": {
        "currency": "EUR",
        "value": "698.00"
    },
    "status": "open",
    "description": "Order #1337 (Lego cars)",
    "createdAt": "2018-12-01T17:09:02+00:00",
    "method": "banktransfer",
    "metadata": null,
    "orderId": "ord_stTC2WHAuS",
    "isCancelable": true,
    "locale": "nl_NL",
    "profileId": "pfl_URR55HPMGx",
    "sequenceType": "oneoff",
    "settlementAmount": {
        "value": "698.00",
        "currency": "EUR"
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg",
            "type": "application/hal+json"
        },
        "order": {
            "href": "https://api.mollie.com/v2/orders/ord_stTC2WHAuS",
            "type": "application/hal+json"
        },
        "checkout": {
            "href": "https://www.mollie.com/paymentscreen/testmode/?method=banktransfer&token=fgnwdh",
            "type": "text/html"
        },
        "status": {
            "href": "https://www.mollie.com/paymentscreen/banktransfer/status/fgnwdh",
            "type": "text/html"
        },
        "payOnline": {
            "href": "https://www.mollie.com/paymentscreen/banktransfer/pay-online/fgnwdh",
            "type": "text/html"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/orders-api/create-order-payment",
            "type": "text/html"
        }
    }
}`

// CreateOrderPaymentResponseFailed example response
const CreateOrderPaymentResponseFailed = `{
	"status": 422,
	"title": "Unprocessable Entity",
	"detail": "Cannot create a new payment for order ord_stTC2WHAuS when the order status is \"completed\".",
	"_links": {
		"documentation": {
			"href": "https://docs.mollie.com/reference/v2/orders-api/create-order-payment",
			"type": "text/html"
		}
	}
}`

// CreateOrderRefundRequest example
const CreateOrderRefundRequest = `{
	"lines": [
	   {
		  "id": "odl_dgtxyl",
		  "quantity": 1
	   }
	],
	"description": "Required quantity not in stock, refunding one photo book.",
	"metadata": {
	   "bookkeeping_id": 12345
	}
}`

// CreateOrderRefundResponse example
const CreateOrderRefundResponse = `{
    "resource": "refund",
    "id": "re_4qqhO89gsT",
    "amount": {
        "currency": "EUR",
        "value": "698.00"
    },
    "status": "pending",
    "createdAt": "2018-03-14T17:09:02.0Z",
    "description": "Required quantity not in stock, refunding one photo book.",
    "metadata": {
         "bookkeeping_id": 12345
    },
    "paymentId": "tr_WDqYK6vllg",
    "orderId": "ord_stTC2WHAuS",
    "lines": [
        {
            "resource": "orderline",
            "id": "odl_dgtxyl",
            "orderId": "ord_stTC2WHAuS",
            "name": "LEGO 42083 Bugatti Chiron",
            "sku": "5702016116977",
            "type": "physical",
            "status": "paid",
            "metadata": null,
            "quantity": 1,
            "unitPrice": {
                "value": "399.00",
                "currency": "EUR"
            },
            "vatRate": "21.00",
            "vatAmount": {
                "value": "51.89",
                "currency": "EUR"
            },
            "discountAmount": {
                "value": "100.00",
                "currency": "EUR"
            },
            "totalAmount": {
                "value": "299.00",
                "currency": "EUR"
            },
            "createdAt": "2018-08-02T09:29:56+00:00",
            "_links": {
                "productUrl": {
                    "href": "https://shop.lego.com/nl-NL/Bugatti-Chiron-42083",
                    "type": "text/html"
                },
                "imageUrl": {
                    "href": "https://sh-s7-live-s.legocdn.com/is/image//LEGO/42083_alt1?$main$",
                    "type": "text/html"
                }
            }
        }
    ],
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg/refunds/re_4qqhO89gsT",
            "type": "application/hal+json"
        },
        "payment": {
            "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg",
            "type": "application/hal+json"
        },
        "order": {
            "href": "https://api.mollie.com/v2/orders/ord_stTC2WHAuS",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/orders-api/create-order-refund",
            "type": "text/html"
        }
    }
}`

// CreateOrderRefundResponseFailed example
const CreateOrderRefundResponseFailed = `{
	"status": 422,
	"title": "Unprocessable Entity",
	"detail": "Line 0 contains invalid data. An amount is required for this API call. The amount must be between €0.00 and €50.00.",
	"field": "lines.0.amount",
	"extra": {
		"minimumAmount": {
			"value": "0.00",
			"currency": "EUR"
		},
		"maximumAmount": {
			"value": "50.00",
			"currency": "EUR"
		}
	},
	"_links": {
		"documentation": {
			"href": "https://docs.mollie.com/reference/v2/orders-api/create-order-refund",
			"type": "text/html"
		}
	}
}`

// ListOrderRefundResponse example
const ListOrderRefundResponse = `{
    "count": 1,
    "_embedded": {
        "refunds": [
            {
                "resource": "refund",
                "id": "re_4qqhO89gsT",
                "amount": {
                    "currency": "EUR",
                    "value": "698.00"
                },
                "status": "processing",
                "createdAt": "2018-03-14T17:09:02.0Z",
                "description": "Required quantity not in stock, refunding one photo book.",
                "metadata": {
                     "bookkeeping_id": 12345
                },
                "paymentId": "tr_WDqYK6vllg",
                "orderId": "ord_stTC2WHAuS",
                "lines": [
                    {
                        "resource": "orderline",
                        "id": "odl_dgtxyl",
                        "orderId": "ord_stTC2WHAuS",
                        "name": "LEGO 42083 Bugatti Chiron",
                        "sku": "5702016116977",
                        "type": "physical",
                        "status": "paid",
                        "metadata": null,
                        "quantity": 1,
                        "unitPrice": {
                            "value": "399.00",
                            "currency": "EUR"
                        },
                        "vatRate": "21.00",
                        "vatAmount": {
                            "value": "51.89",
                            "currency": "EUR"
                        },
                        "discountAmount": {
                            "value": "100.00",
                            "currency": "EUR"
                        },
                        "totalAmount": {
                            "value": "299.00",
                            "currency": "EUR"
                        },
                        "createdAt": "2018-08-02T09:29:56+00:00",
                        "_links": {
                            "productUrl": {
                                "href": "https://shop.lego.com/nl-NL/Bugatti-Chiron-42083",
                                "type": "text/html"
                            },
                            "imageUrl": {
                                "href": "https://sh-s7-live-s.legocdn.com/is/image//LEGO/42083_alt1?$main$",
                                "type": "text/html"
                            }
                        }
                    }
                ],
                "_links": {
                    "self": {
                        "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg/refunds/re_4qqhO89gsT",
                        "type": "application/hal+json"
                    },
                    "payment": {
                        "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg",
                        "type": "application/hal+json"
                    },
                    "order": {
                        "href": "https://api.mollie.com/v2/orders/ord_stTC2WHAuS",
                        "type": "application/hal+json"
                    },
                    "documentation": {
                        "href": "https://docs.mollie.com/reference/v2/refunds-api/get-refund",
                        "type": "text/html"
                    }
                }
            }
        ]
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/payments/tr_7UhSN1zuXS/refunds?limit=5",
            "type": "application/hal+json"
        },
        "previous": null,
        "next": {
            "href": "https://api.mollie.com/v2/payments/tr_7UhSN1zuXS/refunds?from=re_APBiGPH2vV&limit=5",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/orders-api/list-order-refunds",
            "type": "text/html"
        }
    }
}`
