package testdata

// GetShipmentsResponse example
const GetShipmentsResponse = `{
     "resource": "shipment",
     "id": "shp_3wmsgCJN4U",
     "orderId": "ord_kEn1PlbGa",
     "createdAt": "2018-08-09T14:33:54+00:00",
     "tracking": {
         "carrier": "PostNL",
         "code": "3SKABA000000000",
         "url": "http://postnl.nl/tracktrace/?B=3SKABA000000000&P=1016EE&D=NL&T=C"
     },
     "lines": [
         {
             "resource": "orderline",
             "id": "odl_dgtxyl",
             "orderId": "ord_pbjz8x",
             "name": "LEGO 42083 Bugatti Chiron",
             "sku": "5702016116977",
             "type": "physical",
             "status": "shipping",
             "metadata": null,
             "isCancelable": true,
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
         },
         {
             "resource": "orderline",
             "id": "odl_jp31jz",
             "orderId": "ord_pbjz8x",
             "name": "LEGO 42056 Porsche 911 GT3 RS",
             "sku": "5702015594028",
             "type": "physical",
             "status": "completed",
             "metadata": null,
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
             "href": "https://api.mollie.com/v2/order/ord_kEn1PlbGa/shipments/shp_3wmsgCJN4U",
             "type": "application/hal+json"
         },
         "order": {
             "href": "https://api.mollie.com/v2/orders/ord_kEn1PlbGa",
             "type": "application/hal+json"
         },
         "documentation": {
             "href": "https://docs.mollie.com/reference/v2/shipments-api/get-shipment",
             "type": "text/html"
         }
     }
 }`

// ListShipmentsResponse example
const ListShipmentsResponse = `{
    "count": 2,
    "_embedded": {
        "shipments": [
            {
                "resource": "shipment",
                "id": "shp_3wmsgCJN4U",
                "orderId": "ord_kEn1PlbGa",
                "createdAt": "2018-08-09T14:33:54+00:00",
                "tracking": {
                    "carrier": "PostNL",
                    "code": "3SKABA000000000",
                    "url": "http://postnl.nl/tracktrace/?B=3SKABA000000000&P=1016EE&D=NL&T=C"
                },
                "lines": [
                    {
                        "resource": "orderline",
                        "id": "odl_dgtxyl",
                        "orderId": "ord_pbjz8x",
                        "name": "LEGO 42083 Bugatti Chiron",
                        "sku": "5702016116977",
                        "type": "physical",
                        "status": "shipping",
                        "metadata": null,
                        "isCancelable": true,
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
                    },
                    { }
                ]
                "_links": {
                    "self": {
                        "href": "https://api.mollie.com/v2/order/ord_kEn1PlbGa/shipments/shp_3wmsgCJN4U",
                        "type": "application/hal+json"
                    },
                    "order": {
                        "href": "https://api.mollie.com/v2/orders/ord_kEn1PlbGa",
                        "type": "application/hal+json"
                    },
                    "documentation": {
                        "href": "https://docs.mollie.com/reference/v2/shipments-api/get-shipment",
                        "type": "text/html"
                    }
                }
            },
            { }
        ]
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/order/ord_kEn1PlbGa/shipments",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/shipments-api/list-shipments",
            "type": "text/html"
        }
    }
}`
