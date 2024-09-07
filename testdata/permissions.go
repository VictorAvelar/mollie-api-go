package testdata

// GetPermissionsResponse example.
const GetPermissionsResponse = `{
    "resource": "permission",
    "id": "payments.read",
    "description": "View your payments",
    "granted": true,
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/permissions/payments.read",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/permissions-api/get-permission",
            "type": "text/html"
        }
    }
}
`

// ListPermissionsResponse example.
const ListPermissionsResponse = `{
    "_embedded": {
        "permissions": [
            {
                "resource": "permission",
                "id": "payments.write",
                "description": "Create new payments",
                "granted": false,
                "_links": {
                    "self": {
                        "href": "https://api.mollie.com/v2/permissions/payments.write",
                        "type": "application/hal+json"
                    }
                }
            },
            {
                "resource": "permission",
                "id": "payments.read",
                "description": "View your payments",
                "granted": true,
                "_links": {
                    "self": {
                        "href": "https://api.mollie.com/v2/permissions/payments.read",
                        "type": "application/hal+json"
                    }
                }
            }
       ]
    },
    "count": 15,
    "_links": {
        "documentation": {
            "href": "https://docs.mollie.com/reference/permissions-api/list-permissions",
            "type": "text/html"
        },
        "self": {
            "href": "https://api.mollie.com/v2/permissions",
            "type": "application/hal+json"
        }
    }
}
`
