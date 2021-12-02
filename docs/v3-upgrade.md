# Upgrade to v3

## Breaking changes

- Minimum go version is now v1.17
- All service actions accept a `context.Context` as first parameter.
- Methods service actions now return pointers instead of struct literals.
- Invoice service actions now return pointers instead of struct literals.
- Mandates service actions now return pointers instead of struct literals.
- Get permissions now accepts a direct reference to a PermissionGrant struct instead of a string.

## Other changes

- All the tests are not using testify.Suite
- Removed the examples tests as it was not accurate nor well implemented.
- Remove CHANGELOG.md as the releases now provide a more accurate report of changes.
- Some typos were fixed on several query serializable param tags
