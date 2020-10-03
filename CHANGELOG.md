# Changelog

All notable changes to `mollie-api-go` will be documented in this file.

Updates should follow the [Keep a CHANGELOG](http://keepachangelog.com/) principles.

## master

## Added
- Code samples for Chargebacks API #79
- Playground ready example tests for GetChargeback #79

## Fixed
- golint style errors in `mandates.go` #
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