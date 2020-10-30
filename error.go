package jsoncat

type strErr string

func (err strErr) Error() string {
	return string(err)
}

// Error predicates.
const (
	ErrNotObject strErr = "not a JSON object"
	ErrNotArray  strErr = "not a JSON array"
	ErrNotString strErr = "not a JSON string"
)
