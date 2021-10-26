# JSON Cat

![Go](https://github.com/clarify/jsoncat/workflows/Go/badge.svg)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/clarify/jsoncat)][pkgref]

The `jsoncat` package provide functions for concatenating JSON entities of the
same JSON type (object, array or string) while preserving the order of elements.

See the [package reference][pkgref] for more details and examples.

[pkgref]: https://pkg.go.dev/github.com/clarify/jsoncat

## Example

The following example shows a particular use-case where this library can come in
handy:

1. You build an API client in Go.
2. You choose to rely on [embedded structs][embedstruct] to compose your
   models through defining some re-usable "field sets".
3. You realize that some of the fields in the shared field-sets are _read-only_
   and you get an error when writing the model back.
4. You realize that `json.Marshal` does not look for `json.Marshaler`
   implementation in embedded fields... Or at least not in the way you want, as
   the method is just inherited.

If only there was an easy way to let JSON Marshaling of struct consult
potential `json.Marshaler` implementations for embedded fields..

... well, no there is with `jsoncat`:

[embedstruct]: https://cwinters.com/2014/09/02/embedded_structs_in_go.html

```go
import (
    "encoding/json"
    "time"

    "github.com/clarify/jsoncat"
)

// Model is assumed a base client-side base model for resources from an API.
// Some fields are read-only.  Since unmarshaling is case-insensitive, and we
// implement our own MarshalJSON we don't need to specify any JSON tags
// for this example.
type Model struct {
    ID        string

    CreatedBy string    // read-only field
    UpdatedBy string    // read-only field
    CreatedAt time.Time // read-only field
    UpdatedAt time.Time // read-only field

}

func (m Model) MarshalJSON() ([]byte, error) {
    // We omit the read-only fields when encoding m.
    return json.Marshal(struct {
        ID string `json:"id"`
    }{
        ID: m.ID,
    })
}

// SoftDelete provides an optional field-set for models that can be
// soft-deleted. All fields are read-write, so we don't need a custom
// MarshalJSON implementation, and we define struct-tags to get the right case
// when marshaling.
type SoftDelete struct {
    DeletedAt *time.Time `json:"deletedAt"`
    ExpiresAt *time.Time `json:"expiresAt"`
}


// User is an example Model that rely on several predefined field-sets, as well
// as defining some fields of it's own.  Since unmarshaling is case-insensitive,
// and we implement our own MarshalJSON we don't need to specify any JSON tags
// for this example.
type User struct {
    // By relying on embedded structs, json.Unmarshal will be able to write
    // fields into the correct members from a flat JSON object.
    Model
    SoftDelete

    // Non embedded fields are also correctly decoded from the flat JSON.
    FirstName string
    LastName  string
}

func (a User) MarshalJSON() ([]byte, error) {
    // We overwrite MarshalJSON function, we can respect the json.Marshaler
    // implementation of our embedded fields.
    return jsoncat.MarshalObject(a.Model, a.SoftDelete, struct{
        FirstName string `json:"firstName"`
        LastName  string `json:"lastName"`
    }{
        FirstName: a.FirstName,
        LastName:  a.LastName,
    })
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
