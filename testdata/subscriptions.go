package testdata

// GetSubscriptionResponse example
const GetSubscriptionResponse = `{
    "resource": "subscription",
    "id": "sub_rVKGtNd6s3",
    "mode": "live",
    "createdAt": "2016-06-01T12:23:34+00:00",
    "status": "active",
    "amount": {
        "value": "25.00",
        "currency": "EUR"
    },
    "times": 4,
    "timesRemaining": 4,
    "interval": "3 months",
    "startDate": "2016-06-01",
    "nextPaymentDate": "2016-09-01",
    "description": "Quarterly payment",
    "method": null,
    "mandateId": "mdt_38HS4fsS",
    "webhookUrl": "https://webshop.example.org/payments/webhook",
    "metadata": {
        "plan": "small"
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/customers/cst_stTC2WHAuS/subscriptions/sub_rVKGtNd6s3",
            "type": "application/hal+json"
        },
        "customer": {
            "href": "https://api.mollie.com/v2/customers/cst_stTC2WHAuS",
            "type": "application/hal+json"
        },
        "profile": {
            "href": "https://api.mollie.com/v2/profiles/pfl_URR55HPMGx",
            "type": "application/hal+json"
        },
       "payments": {
            "href": "https://api.mollie.com/v2/customers/cst_stTC2WHAuS/subscriptions/sub_rVKGtNd6s3/payments",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/subscriptions-api/get-subscription",
            "type": "text/html"
        }
    }
}`

// DeleteSubscriptionResponse example
const DeleteSubscriptionResponse = `{
    "resource": "subscription",
    "id": "sub_rVKGtNd6s3",
    "mode": "live",
    "createdAt": "2018-06-01T12:23:34+00:00",
    "status": "canceled",
    "amount": {
        "value": "25.00",
        "currency": "EUR"
    },
    "times": 4,
    "interval": "3 months",
    "nextPaymentDate": null,
    "description": "Quarterly payment",
    "method": null,
    "startDate": "2016-06-01",
    "webhookUrl": "https://webshop.example.org/payments/webhook",
    "canceledAt": "2018-08-01T11:04:21+00:00",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/customers/cst_stTC2WHAuS/subscriptions/sub_rVKGtNd6s3",
            "type": "application/hal+json"
        },
        "customer": {
            "href": "https://api.mollie.com/v2/customers/cst_stTC2WHAuS",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/subscriptions-api/cancel-subscription",
            "type": "text/html"
        }
    }
}`

// ListAllSubscriptionsResponse example
const ListAllSubscriptionsResponse = `{
    "count": 3,
    "_embedded": {
        "subscriptions": [
            {
                "resource": "subscription",
                "id": "sub_rVKGtNd6s3",
                "mode": "live",
                "createdAt": "2018-06-01T12:23:34+00:00",
                "status": "active",
                "amount": {
                    "value": "25.00",
                    "currency": "EUR"
                },
                "times": 4,
                "timesRemaining": 3,
                "interval": "3 months",
                "startDate": "2016-06-01",
                "nextPaymentDate": "2016-09-01",
                "description": "Quarterly payment",
                "method": null,
                "webhookUrl": "https://webshop.example.org/subscriptions/webhook",
                "_links": {
                    "self": {
                        "href": "https://api.mollie.com/v2/customers/cst_stTC2WHAuS/subscriptions/sub_rVKGtNd6s3",
                        "type": "application/hal+json"
                    },
                    "profile": {
                        "href": "https://api.mollie.com/v2/profiles/pfl_URR55HPMGx",
                        "type": "application/hal+json"
                    },
                    "customer": {
                        "href": "https://api.mollie.com/v2/customers/cst_stTC2WHAuS",
                        "type": "application/hal+json"
                    }
                }
            }
        ]
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/subscriptions",
            "type": "application/hal+json"
        },
        "previous": null,
        "next": {
            "href": "https://api.mollie.com/v2/subscriptions?from=sub_mnfbwhMfvo",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/subscriptions-api/list-all-subscriptions",
            "type": "text/html"
        }
    }
}`
