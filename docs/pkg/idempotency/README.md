# idempotency

--
    import "github.com/VictorAvelar/mollie-api-go/v4/pkg/idempotency"

package idempotency contains the services in charge of generating a unique keys
to be passed in POST requests to ensure operation uniqueness.

See: <https://docs.mollie.com/overview/api-idempotency>

The std generator uses google's uuid library to return a new uuid as unique
idempotency key.

You can build your own generator and pass it to the library by implementing the
KeyGenerator interface.

## Usage

```go
const (
 TestKeyExpected = "test_ikg_key"
)
```

#### type KeyGenerator

```go
type KeyGenerator interface {
 // Generate encapsulates the logic to return a string representation of
 // a unique idempotency key.
 Generate() string
}
```

KeyGenerator describes the service in charge of generating a unique idempotency
key to be passed in POST requests to ensure operation uniqueness.

See: <https://docs.mollie.com/overview/api-idempotency>

#### func  NewNopGenerator

```go
func NewNopGenerator(exp string) KeyGenerator
```

NewNopGenerator returns a dummy implementation of the IdempotencyKeyGenerator
interface.

Good for testing or when a predictable result is required.

If exp is an empty string, then TestKeyExpected is used as default value for the
NOpGenerator.

#### func  NewStdGenerator

```go
func NewStdGenerator() KeyGenerator
```

NewStdGenerator returns an standard and common way of generating idempotency
unique keys.
