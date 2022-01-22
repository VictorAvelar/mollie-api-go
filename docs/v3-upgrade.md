# Changes included in v3

## Breaking changes

- Minimum go version is now v1.17
- All service actions accept a `context.Context` as first parameter.
- Methods service actions now return pointers instead of struct literals.
- Invoice service actions now return pointers instead of struct literals.
- Mandates service actions now return pointers instead of struct literals.
- Payments service actions now return pointers instead of struct literals.
- Get permissions now accepts a direct reference to a PermissionGrant struct instead of a string.
- All services now return an instance of mollie.Response as first value changing the signature from (value, error) to (response, value, error).
- All structs are changed to the following naming patterns
  - For Lists: {Vaue}List
  - For Options: {Value}Options | {Value}ListOptions
  - For Links: {Value}Links | {Value}ListLinks
- Params now contain the full name of the value they contain. Ex. pID now is payment.
- Methods and its references is not PaymentMethods to make its purpose clear.
- Services that now return the full response object as part of the returned values (the response is always the first value of `n` returned):
  - Captures
  - Chargebacks
  - Customers
  - Invoices
  - Mandates
  - Miscellaneous
  - Onboarding
  - Partners
  - Orders
  - Organizations
  - PaymentLinks
  - PaymentMethods
  - Payments
  - Permissions
  - Profiles
  - Refunds
  - Settlements
  - Shipments
  - Subscriptions

## Other changes

- All the tests are not using testify.Suite
- Removed the examples tests as it was not accurate nor well implemented.
- Remove CHANGELOG.md as the releases now provide a more accurate report of changes.
- Some typos were fixed on several query serializable param tags
- Client now contains helpers for the used http actions (get, post, patch & delete) to simplify the way the requests are dispatched.
