package testdata

// CreatePaymentRefundResponse example.
const CreatePaymentRefundResponse = `{
    "resource": "refund",
    "id": "re_4qqhO89gsT",
    "amount": {
        "currency": "EUR",
        "value": "5.95"
    },
    "status": "pending",
    "createdAt": "2018-03-14T17:09:02.0Z",
    "description": "Order #33",
    "metadata": {
         "bookkeeping_id": 12345
    },
    "paymentId": "tr_WDqYK6vllg",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg/refunds/re_4qqhO89gsT",
            "type": "application/hal+json"
        },
        "payment": {
            "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/refunds-api/create-payment-refund",
            "type": "text/html"
        }
    }
}`

// GetPaymentRefundResponse example.
const GetPaymentRefundResponse = `{
    "resource": "refund",
    "id": "re_4qqhO89gsT",
    "amount": {
        "currency": "EUR",
        "value": "5.95"
    },
    "status": "pending",
    "createdAt": "2018-03-14T17:09:02.0Z",
    "description": "Order #33",
    "metadata": {
         "bookkeeping_id": 12345
    },
    "paymentId": "tr_WDqYK6vllg",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg/refunds/re_4qqhO89gsT",
            "type": "application/hal+json"
        },
        "payment": {
            "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/refunds-api/get-payment-refund",
            "type": "text/html"
        }
    }
}`

// ListPaymentRefundsResponse example.
const ListPaymentRefundsResponse = `{
    "count": 1,
    "_embedded": {
        "refunds": [
            {
                "resource": "refund",
                "id": "re_4qqhO89gsT",
                "amount": {
                    "currency": "EUR",
                    "value": "5.95"
                },
                "status": "pending",
                "createdAt": "2018-03-14T17:09:02.0Z",
                "description": "Order",
                "metadata": {
                     "bookkeeping_id": 12345
                },
                "paymentId": "tr_WDqYK6vllg",
                "_links": {
                    "self": {
                        "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg/refunds/re_4qqhO89gsT",
                        "type": "application/hal+json"
                    },
                    "payment": {
                        "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg",
                        "type": "application/hal+json"
                    },
                    "documentation": {
                        "href": "https://docs.mollie.com/reference/v2/refunds-api/get-payment-refund",
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
            "href": "https://docs.mollie.com/reference/v2/refunds-api/list-payment-refunds",
            "type": "text/html"
        }
    }
}`

// ListOrderRefundsResponse example.
const ListOrderRefundsResponse = `
{
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
                        "href": "https://docs.mollie.com/reference/v2/refunds-api/get-order-refund",
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
            "href": "https://docs.mollie.com/reference/v2/refunds-api/list-order-refunds",
            "type": "text/html"
        }
    }
}`

// ListRefundsResponse example.
const ListRefundsResponse = `{
    "count": 1,
    "_embedded": {
        "refunds": [
            {
                "resource": "refund",
                "id": "re_4qqhO89gsT",
                "amount": {
                    "currency": "EUR",
                    "value": "5.95"
                },
                "status": "pending",
                "createdAt": "2018-03-14T17:09:02.0Z",
                "description": "Order",
                "metadata": {
                     "bookkeeping_id": 12345
                },
                "paymentId": "tr_WDqYK6vllg",
                "_links": {
                    "self": {
                        "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg/refunds/re_4qqhO89gsT",
                        "type": "application/hal+json"
                    },
                    "payment": {
                        "href": "https://api.mollie.com/v2/payments/tr_WDqYK6vllg",
                        "type": "application/hal+json"
                    },
                    "documentation": {
                        "href": "https://docs.mollie.com/reference/v2/refunds-api/get-payment-refund",
                        "type": "text/html"
                    }
                }
            }
        ]
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/refunds?limit=5",
            "type": "application/hal+json"
        },
        "previous": null,
        "next": {
            "href": "https://api.mollie.com/v2/refunds?from=re_APBiGPH2vV&limit=5",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/refunds-api/list-refunds",
            "type": "text/html"
        }
    }
}`
