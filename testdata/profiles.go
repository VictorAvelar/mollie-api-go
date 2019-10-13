package testdata

// GetProfileResponse example
const GetProfileResponse = `{
    "resource": "profile",
    "id": "pfl_v9hTwCvYqw",
    "mode": "live",
    "name": "My website name",
    "website": "https://www.mywebsite.com",
    "email": "info@mywebsite.com",
    "phone": "+31208202070",
    "categoryCode": 5399,
    "status": "verified",
    "review": {
        "status": "pending"
    },
    "createdAt": "2018-03-20T09:28:37+00:00",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/profiles/pfl_v9hTwCvYqw",
            "type": "application/hal+json"
        },
        "chargebacks": {
            "href": "https://api.mollie.com/v2/chargebacks?profileId=pfl_v9hTwCvYqw",
            "type": "application/hal+json"
        },
        "methods": {
            "href": "https://api.mollie.com/v2/methods?profileId=pfl_v9hTwCvYqw",
            "type": "application/hal+json"
        },
        "payments": {
            "href": "https://api.mollie.com/v2/payments?profileId=pfl_v9hTwCvYqw",
            "type": "application/hal+json"
        },
        "refunds": {
            "href": "https://api.mollie.com/v2/refunds?profileId=pfl_v9hTwCvYqw",
            "type": "application/hal+json"
        },
        "checkoutPreviewUrl": {
            "href": "https://www.mollie.com/payscreen/preview/pfl_v9hTwCvYqw",
            "type": "text/html"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/profiles-api/create-profile",
            "type": "text/html"
        }
    }
}`

// GetProfilesListResponse example
const GetProfilesListResponse = `{
    "_embedded": {
        "profiles": [
            {
                "resource": "profiles",
                "id": "pfl_v9hTwCvYqw",
                "mode": "live",
                "name": "My website name",
                "website": "https://www.mywebsite.com",
                "email": "info@mywebsite.com",
                "phone": "+31208202070",
                "categoryCode": 5399,
                "status": "verified",
                "review": {
                    "status": "pending"
                },
                "createdAt": "2018-03-20T09:28:37+00:00",
                "_links": {
                    "self": {
                        "href": "https://api.mollie.com/v2/profiles/pfl_v9hTwCvYqw",
                        "type": "application/hal+json"
                    },
                    "chargebacks": {
                        "href": "https://api.mollie.com/v2/chargebacks?profileId=pfl_v9hTwCvYqw",
                        "type": "application/hal+json"
                    },
                    "methods": {
                        "href": "https://api.mollie.com/v2/methods?profileId=pfl_v9hTwCvYqw",
                        "type": "application/hal+json"
                    },
                    "payments": {
                        "href": "https://api.mollie.com/v2/payments?profileId=pfl_v9hTwCvYqw",
                        "type": "application/hal+json"
                    },
                    "refunds": {
                        "href": "https://api.mollie.com/v2/refunds?profileId=pfl_v9hTwCvYqw",
                        "type": "application/hal+json"
                    },
                    "checkoutPreviewUrl": {
                        "href": "https://www.mollie.com/payscreen/preview/pfl_v9hTwCvYqw",
                        "type": "text/html"
                    },
                    "documentation": {
                        "href": "https://docs.mollie.com/reference/v2/profiles-api/create-profile",
                        "type": "text/html"
                    }
                }
            }
        ]
    },
    "count": 5,
    "_links": {
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/profiles-api/list-profiles",
            "type": "text/html"
        },
        "self": {
            "href": "https://api.mollie.com/v2/profiles?limit=5",
            "type": "application/hal+json"
        },
        "previous": null,
        "next": {
            "href": "https://api.mollie.com/v2/profiles?from=pfl_3RkSN1zuPE&limit=5",
            "type": "application/hal+json"
        }
    }
}`

// CreateProfileResponse example
const CreateProfileResponse = `
{
    "resource": "profile",
    "id": "pfl_v9hTwCvYqw",
    "mode": "live",
    "name": "My website name",
    "website": "https://www.mywebsite.com",
    "email": "info@mywebsite.com",
    "phone": "+31208202070",
    "categoryCode": 5399,
    "status": "unverified",
    "createdAt": "2018-03-20T09:28:37+00:00",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/profiles/pfl_v9hTwCvYqw",
            "type": "application/hal+json"
        },
        "chargebacks": {
            "href": "https://api.mollie.com/v2/chargebacks?profileId=pfl_v9hTwCvYqw",
            "type": "application/hal+json"
        },
        "methods": {
            "href": "https://api.mollie.com/v2/methods?profileId=pfl_v9hTwCvYqw",
            "type": "application/hal+json"
        },
        "payments": {
            "href": "https://api.mollie.com/v2/payments?profileId=pfl_v9hTwCvYqw",
            "type": "application/hal+json"
        },
        "refunds": {
            "href": "https://api.mollie.com/v2/refunds?profileId=pfl_v9hTwCvYqw",
            "type": "application/hal+json"
        },
        "checkoutPreviewUrl": {
            "href": "https://www.mollie.com/payscreen/preview/pfl_v9hTwCvYqw",
            "type": "text/html"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/profiles-api/create-profile",
            "type": "text/html"
        }
    }
}`

const UpdateProfileResponse = `{
    "resource": "profile",
    "id": "pfl_v9hTwCvYqw",
    "mode": "live",
    "name": "My website name - Update 1",
    "website": "https://www.mywebsite2.com",
    "email": "info@mywebsite2.com",
    "phone": "+31208202070",
    "categoryCode": 5399,
    "status": "verified",
    "review": {
        "status": "pending"
    },
    "createdAt": "2018-03-20T09:28:37+00:00",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/profiles/pfl_v9hTwCvYqw",
            "type": "application/hal+json"
        },
        "chargebacks": {
            "href": "https://api.mollie.com/v2/chargebacks?profileId=pfl_v9hTwCvYqw",
            "type": "application/hal+json"
        },
        "methods": {
            "href": "https://api.mollie.com/v2/methods?profileId=pfl_v9hTwCvYqw",
            "type": "application/hal+json"
        },
        "payments": {
            "href": "https://api.mollie.com/v2/payments?profileId=pfl_v9hTwCvYqw",
            "type": "application/hal+json"
        },
        "refunds": {
            "href": "https://api.mollie.com/v2/refunds?profileId=pfl_v9hTwCvYqw",
            "type": "application/hal+json"
        },
        "checkoutPreviewUrl": {
            "href": "https://www.mollie.com/payscreen/preview/pfl_v9hTwCvYqw",
            "type": "text/html"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/profiles-api/create-profile",
            "type": "text/html"
        }
    }
}`

// EnablePaymentMethodResponse example
const EnablePaymentMethodResponse = `{
    "resource": "method",
    "id": "bancontact",
    "description": "Bancontact",
    "image": {
        "size1x": "https://www.mollie.com/external/icons/payment-methods/bancontact.png",
        "size2x": "https://www.mollie.com/external/icons/payment-methods/bancontact%402x.png",
        "svg": "https://www.mollie.com/external/icons/payment-methods/bancontact.svg"
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/methods/bancontact",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/profiles-api/activate-method",
            "type": "text/html"
        }
    }
}
`
