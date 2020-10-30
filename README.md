# JSON Cat

![Go](https://github.com/searis/jsoncat/workflows/Go/badge.svg)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/searis/jsoncat)][pkgref]

The `jsoncat` package provide functions for concatenating JSON entities of the
same JSON type (object, array or string) while preserving the order of elements.

See the [package reference][pkgref] for more details and examples.

[pkgref]: https://pkg.go.dev/github.com/searis/jsoncat

## Example

The following example shows a particular use-case where this library can come in
handy:

1. You build an API client in Go.
2. You choose to rely on [struct embedding] to compose your models through
   defining some re-usable field sets.
3. You realize that some of the fields in the shared field-sets are _read-only_
   and you get an error when writing the model back.
4. You realize that `json.Marshal` does not look for `json.Marshaler`
   implementation in embedded fields... Or at least not in the way you want, as
   the method is just inherited.

If only there was an easy way to let JSON Marshaling of struct consult
potential `json.Marshaler` implementations for embedded fields..

... well, no there is with `jsoncat`:

```go
import (
    "encoding/json"
    "time"

    "github.com/searis/jsoncat"
)

// Model assumed a base model for an API client.
type Model struct {
    ID        string    `json:"id"`
    CreatedAt time.Time `json:"createdAt"` // read-only field
    CreatedBy string    `json:"createdBy"` // read-only field
}

func (m Model) MarshalJSON() ([]byte, error) {
    // We omit the read-only fields when encoding m.
    return json.Marshal(struct {
        ID string `json:"id"`
    }{
        ID: m.ID,
    })
}

type AttrUser struct {
    FirstName string `json:"firstName"`
    LastName  string `json:"lastName"`
}

type User struct {
    Model
    AttrUser
}

func (a User) MarshalJSON() ([]byte, error) {
    // By overwriting the MarshalJSON function, we can marshal decode
    return jsoncat.MarshalObject(a.Model, a.AttrUser)
}
```

## Performance over features

The library is aimed at decent performance and preserving the order of elements.
The input to functions is thus not validated to contain valid JSON. The library
will be able to detect if the JSON elements does not start or end with the right
delimiters (`[]`, `{}`, or `""`), and leading and trailing whitespace will be
trimmed.

The library will not remove duplicated keys or entries; if you concatenate
`{"foo":"bar"}` and `{"foo":"baz"}` the result is `{"foo":"bar","foo":"baz"}`.
