# Notable changes included in v4

## Breaking changes

- `idempotency` package has moved from the `mollie` directory to the `pkg` directory.
- `pagination` package has moved from the`mollie` directory to the `pkg` directory.
- `connect` package has moved from the `mollie` directory to the `pkg` directory.
- root namespace is not `github.com/VictorAvelar/mollie-api-go/v4/`.
- Changes in all resources:
  - Data structures to describe request content in create, update and get operations are no longer the same, e.g. for `payments` there is a `CreatePayment` struct, a `UpdatePayment` struct and a `Payment` struct. This enables future extensions and modifications without having to check for cross request compatibility.
  - Data structures that describe available url parameters are consistently names `List{ResourceName}Options`.
  - Data structures that describe lists responses are consistently named as follows: e.g. for payments: `PaymentsList`
- API aliases now use the parent objects, e.g. for settlements when listing payments the options passed to the request are using the `ListPaymentsOptions` object and not a local object.
- All resources were checked for API consistency and parity, optional resources with custom types are now pointers to ensure proper json encoding and decoding to avoid issues as the one mentioned un #271
- All resources embed a struct containing all the fields specific to access tokens, following this pattern the same happens for fields specific to Mollie connect

## Other changes

- `testify.Suite` was removed from all testing.
- Improvements for devcontainer files
- Major versions of multiple github actions updated
- Base `Dockerfile` using Go 1.22.x
- Tests were update to use the new types.
- Test coverage was slightly improved.
