# Notable changes included in v4

## Breaking changes

- `idempotency` package has moved from the `mollie` directory to the `pkg` directory.
- `pagination` package has moved from the`mollie` directory to the `pkg` directory.
- `connect` package has moved from the `mollie` directory to the `pkg` directory.
- root namespace is not `github.com/VictorAvelar/mollie-api-go/v4/`.
- Changes for payments resource
  - Added `CreatePayment` type to use when creating a payment.
  - Added `UpdatePayment` type to use when updating a payment.
  - Ensured `Payment` object (used for read operations) contains all the available fields.
  - Created types specific for fields that are conditional given the payment values, both to create new payments and to list existing payments.

## Other changes

- `testify.Suite` was removed from all testing.
- Improvements for devcontainer files
- Major versions of multiple github actions updated
- Base `Dockerfile` using Go 1.22.x
- Tests related to payments were update to use the new types.
