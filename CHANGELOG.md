# Changelog

All notable changes to `mollie-api-go` will be documented in this file.

Updates should follow the [Keep a CHANGELOG](http://keepachangelog.com/) principles.

## master

## v2.1.1

### Added
- Payment details address for supported payment methods. #109

## v2.1.0

## Changed
- Rollback to use pointers inside struct tags as structs have no empty value
and thus the usage of values instead of pointers breaks most of the post requests.

## v2.0.2

### Fixed
- Normalized wording across structs

## v2.0.1

### Fixed
- Path for v2 support

### Deprecated
- v1 support

## v2.0.0

### Added
- Missing payment list url options

### Changed
- Replaced pointers with values in the whole library (BREAKING)
- All the documentation links name unified to `Documentation`

## v1.6.5

### Added
- Goreleaser workflow
- Examples for subscriptions #96
- Examples or payment methods

### Changed
- Renamed ChargebackID to PaymentID in Chargeback struct #98

## v1.6.4

### Added
- Full http error content when the response is within the http error range #95

## v1.6.3

### Fixed
- Error when parsing amount values on payment methods #94

## v1.6.2

### Fixed
- Wrong tag in `MethodOptions`struct #92

## v1.6.1

### Fixed
- api token parsing from env variable #89
- go fmt errors #89

### Changed
- Updated documentation

## v1.6.0

### Changed
- Go mod version upgraded to 1.15 #86 

### Removed
- Unused variable in refunds.go #87 

## v1.5.0

### Added
- Code samples for Chargebacks API #79
- Playground ready example tests for GetChargeback #79
- Dockerfile for test execution #81
- Makefile with used commands #81
- Code samples for profiles #82
- Methods for enabling/disabling payment methods #84

### Fixed
- golint style errors in `mandates.go` #80
- ignored error in `newResponse` method is now properly handled #66
- yaml files were linted #79

## v1.4.1

### Added
- Dashboard URL for payments, profiles, orders, customers and organizations

### Changed
- Response samples for the aforementioned resources in this tag

## v1.4.0

### Added
- ShortDate struct JSON Marshaler
- Test coverage for the added code

## v1.3.0

### Added
- General project housekeeping
- Added chargebacks as possible embeded value

## v1.2.1

### Added
- MandatesID added to Subscriptions resource as follow up to the previous release.

## v1.2.0

### Added
- Mandates API resource

## v1.1.0

### Added
- Added changes to provide API parity according to mollie's API changelog

## v1.0.1

### Fixed
- Fixed error reported when trying to access the `mollie.Error.Response.Body` content (#47).

## v1.0.0

### Added
- Base HTTP Client for interacting with Mollie REST API
- Tests for the added code
- 8e3130d Added scrutinizer check for code style and coverage
- dbaf84f Added Mollie standard types
- 1b53ea9-payments resource implementation
- f658c63-fixed go linting errors
- 8893236-methods resource implementation


### Deprecated
- Nothing

### Changed
- ClientContext for Ctx in APIClient
- 1a220e4 Improved code comments
- a089626 Increased code coverage and testing doc examples

### Fixed
- Missing types for MOLLIE env variables
- Named variable reference on NewClient method
- Typos and spelling errors on method's comments
- Refactored orders
- Multiple typos and consistency errors

### Removed
- Nothing

### Security
- Nothing
