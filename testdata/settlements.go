package testdata

// GetSettlementsResponse example
const GetSettlementsResponse = `{
    "resource": "settlement",
    "id": "stl_jDk30akdN",
    "reference": "1234567.1804.03",
    "createdAt": "2018-04-06T06:00:01.0Z",
    "settledAt": "2018-04-06T09:41:44.0Z",
    "status": "paidout",
    "amount": {
        "currency": "EUR",
        "value": "39.75"
    },
    "periods": {
        "2018": {
            "4": {
                "revenue": [
                    {
                        "description": "iDEAL",
                        "method": "ideal",
                        "count": 6,
                        "amountNet": {
                            "currency": "EUR",
                            "value": "86.1000"
                        },
                        "amountVat": null,
                        "amountGross": {
                            "currency": "EUR",
                            "value": "86.1000"
                        }
                    },
                    {
                        "description": "Refunds iDEAL",
                        "method": "refund",
                        "count": 2,
                        "amountNet": {
                            "currency": "EUR",
                            "value": "-43.2000"
                        },
                        "amountVat": null,
                        "amountGross": {
                            "currency": "EUR",
                            "value": "43.2000"
                        }
                    }
                ],
                "costs": [
                    {
                        "description": "iDEAL",
                        "method": "ideal",
                        "count": 6,
                        "rate": {
                            "fixed": {
                                "currency": "EUR",
                                "value": "0.3500"
                            },
                            "percentage": null
                        },
                        "amountNet": {
                            "currency": "EUR",
                            "value": "2.1000"
                        },
                        "amountVat": {
                            "currency": "EUR",
                            "value": "0.4410"
                        },
                        "amountGross": {
                            "currency": "EUR",
                            "value": "2.5410"
                        }
                    },
                    {
                        "description": "Refunds iDEAL",
                        "method": "refund",
                        "count": 2,
                        "rate": {
                            "fixed": {
                                "currency": "EUR",
                                "value": "0.2500"
                            },
                            "percentage": null
                        },
                        "amountNet": {
                            "currency": "EUR",
                            "value": "0.5000"
                        },
                        "amountVat": {
                            "currency": "EUR",
                            "value": "0.1050"
                        },
                        "amountGross": {
                            "currency": "EUR",
                            "value": "0.6050"
                        }
                    }
                ]
            }
        }
    },
    "invoiceId": "inv_FrvewDA3Pr",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/settlements/next",
            "type": "application/hal+json"
        },
        "payments": {
            "href": "https://api.mollie.com/v2/settlements/stl_jDk30akdN/payments",
            "type": "application/hal+json"
        },
        "refunds": {
            "href": "https://api.mollie.com/v2/settlements/stl_jDk30akdN/refunds",
            "type": "application/hal+json"
        },
        "chargebacks": {
            "href": "https://api.mollie.com/v2/settlements/stl_jDk30akdN/chargebacks",
            "type": "application/hal+json"
        },
        "captures": {
            "href": "https://api.mollie.com/v2/settlements/stl_jDk30akdN/captures",
            "type": "application/hal+json"
        },
        "invoice": {
             "href": "https://api.mollie.com/v2/invoices/inv_FrvewDA3Pr",
             "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/settlements-api/get-settlement",
            "type": "text/html"
        }
    }
}`

// GetNextSettlementResponse example
const GetNextSettlementResponse = `{
    "resource": "settlement",
    "id": "next",
    "reference": null,
    "createdAt": "2018-04-06T06:00:01.0Z",
    "settledAt": null,
    "status": "open",
    "amount": {
        "currency": "EUR",
        "value": "39.75"
    },
    "periods": {
        "2018": {
            "4": {
                "revenue": [
                    {
                        "description": "iDEAL",
                        "method": "ideal",
                        "count": 6,
                        "amountNet": {
                            "currency": "EUR",
                            "value": "86.1000"
                        },
                        "amountVat": null,
                        "amountGross": {
                            "currency": "EUR",
                            "value": "86.1000"
                        }
                    },
                    {
                        "description": "Refunds iDEAL",
                        "method": "refund",
                        "count": 2,
                        "amountNet": {
                            "currency": "EUR",
                            "value": "-43.2000"
                        },
                        "amountVat": null,
                        "amountGross": {
                            "currency": "EUR",
                            "value": "43.2000"
                        }
                    }
                ],
                "costs": [
                    {
                        "description": "iDEAL",
                        "method": "ideal",
                        "count": 6,
                        "rate": {
                            "fixed": {
                                "currency": "EUR",
                                "value": "0.3500"
                            },
                            "percentage": null
                        },
                        "amountNet": {
                            "currency": "EUR",
                            "value": "2.1000"
                        },
                        "amountVat": {
                            "currency": "EUR",
                            "value": "0.4410"
                        },
                        "amountGross": {
                            "currency": "EUR",
                            "value": "2.5410"
                        }
                    },
                    {
                        "description": "Refunds iDEAL",
                        "method": "refund",
                        "count": 2,
                        "rate": {
                            "fixed": {
                                "currency": "EUR",
                                "value": "0.2500"
                            },
                            "percentage": null
                        },
                        "amountNet": {
                            "currency": "EUR",
                            "value": "0.5000"
                        },
                        "amountVat": {
                            "currency": "EUR",
                            "value": "0.1050"
                        },
                        "amountGross": {
                            "currency": "EUR",
                            "value": "0.6050"
                        }
                    }
                ]
            }
        }
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/settlements/next",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/settlements-api/get-next-settlement",
            "type": "text/html"
        }
    }
}`

// GetOpenSettlementResponse example
const GetOpenSettlementResponse = `{
    "resource": "settlement",
    "id": "open",
    "reference": null,
    "createdAt": "2018-04-06T06:00:01.0Z",
    "settledAt": null,
    "status": "open",
    "amount": {
        "currency": "EUR",
        "value": "39.75"
    },
    "periods": {
        "2018": {
            "4": {
                "revenue": [
                    {
                        "description": "iDEAL",
                        "method": "ideal",
                        "count": 6,
                        "amountNet": {
                            "currency": "EUR",
                            "value": "86.1000"
                        },
                        "amountVat": null,
                        "amountGross": {
                            "currency": "EUR",
                            "value": "86.1000"
                        }
                    },
                    {
                        "description": "Refunds iDEAL",
                        "method": "refund",
                        "count": 2,
                        "amountNet": {
                            "currency": "EUR",
                            "value": "-43.2000"
                        },
                        "amountVat": null,
                        "amountGross": {
                            "currency": "EUR",
                            "value": "43.2000"
                        }
                    }
                ],
                "costs": [
                    {
                        "description": "iDEAL",
                        "method": "ideal",
                        "count": 6,
                        "rate": {
                            "fixed": {
                                "currency": "EUR",
                                "value": "0.3500"
                            },
                            "percentage": null
                        },
                        "amountNet": {
                            "currency": "EUR",
                            "value": "2.1000"
                        },
                        "amountVat": {
                            "currency": "EUR",
                            "value": "0.4410"
                        },
                        "amountGross": {
                            "currency": "EUR",
                            "value": "2.5410"
                        }
                    },
                    {
                        "description": "Refunds iDEAL",
                        "method": "refund",
                        "count": 2,
                        "rate": {
                            "fixed": {
                                "currency": "EUR",
                                "value": "0.2500"
                            },
                            "percentage": null
                        },
                        "amountNet": {
                            "currency": "EUR",
                            "value": "0.5000"
                        },
                        "amountVat": {
                            "currency": "EUR",
                            "value": "0.1050"
                        },
                        "amountGross": {
                            "currency": "EUR",
                            "value": "0.6050"
                        }
                    }
                ]
            }
        }
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/settlements/next",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/settlements-api/get-next-settlement",
            "type": "text/html"
        }
    }
}`

const ListSettlementsResponse = `
{
    "count": 5,
    "_embedded": {
        "settlements": [
            {
                "resource": "settlement",
                "id": "stl_jDk30akdN",
                "reference": "1234567.1804.03",
                "createdAt": "2018-04-06T06:00:01.0Z",
                "settledAt": "2018-04-06T09:41:44.0Z",
                "status": "paidout",
                "amount": {
                    "currency": "EUR",
                    "value": "39.75"
                },
                "periods": {
                    "2018": {
                        "4": {
                            "revenue": [ ],
                            "costs": [ ]
                        }
                    }
                },
                "_links": {
                    "self": {
                        "href": "https://api.mollie.com/v2/settlements/next",
                        "type": "application/hal+json"
                    },
                    "payments": {
                        "href": "https://api.mollie.com/v2/settlements/stl_jDk30akdN/payments",
                        "type": "application/hal+json"
                    },
                    "refunds": {
                        "href": "https://api.mollie.com/v2/settlements/stl_jDk30akdN/refunds",
                        "type": "application/hal+json"
                    },
                    "chargebacks": {
                        "href": "https://api.mollie.com/v2/settlements/stl_jDk30akdN/chargebacks",
                        "type": "application/hal+json"
                    },
                    "captures": {
                        "href": "https://api.mollie.com/v2/settlements/stl_jDk30akdN/captures",
                        "type": "application/hal+json"
                    },
                    "documentation": {
                        "href": "https://docs.mollie.com/reference/settlements-api/get-settlement",
                        "type": "text/html"
                    }
                }
            }
        ]
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/settlements?limit=5",
            "type": "application/hal+json"
        },
        "previous": null,
        "next": {
            "href": "https://api.mollie.com/v2/settlements?from=stl_QM8w7JDEhU&limit=5",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/settlements-api/list-settlements",
            "type": "text/html"
        }
    }
}`
