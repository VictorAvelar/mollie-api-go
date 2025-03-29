// Package idempotency contains the services in charge
// of generating a unique keys to be passed in
// POST requests to ensure operation uniqueness.
//
// See: https://docs.mollie.com/overview/api-idempotency
//
// The std generator uses google's uuid library to return a new uuid
// as unique idempotency key.
//
// You can build your own generator and pass it to the library by
// implementing the KeyGenerator interface.
package idempotency
