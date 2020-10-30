package jsoncat

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// MarshalObject takes multiple values that JSON encode to objects, and return
// JSON for a single concatenated object. Note that the order of elements is
// preserved and duplicated keys are not removed.
func MarshalObject(values ...interface{}) ([]byte, error) {
	entries := make([][]byte, 0, len(values))
	for _, v := range values {
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		entries = append(entries, b)
	}
	return Objects(entries...)
}

// MarshalArray takes multiple values that JSON encode to arrays, and return
// JSON for a single concatenated array. The order of elements is preserved.
func MarshalArray(values ...interface{}) ([]byte, error) {
	entries := make([][]byte, 0, len(values))
	for _, v := range values {
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		entries = append(entries, b)
	}
	return Arrays(entries...)
}

// MarshalString takes multiple values that JSON encode to strings, and return
// JSON for a single concatenated string. The order of elements is preserved.
func MarshalString(values ...interface{}) ([]byte, error) {
	entries := make([][]byte, 0, len(values))
	for _, v := range values {
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		entries = append(entries, b)
	}
	return Arrays(entries...)
}

// Objects returns a concatenation of all fields in the passed in JSON objects,
// or returns an error if one or more of the passed in byte-slice does not
// appear to contain a valid JSON object. Note that the order of elements is
// preserved and duplicated keys are not removed.
func Objects(objects ...[]byte) ([]byte, error) {
	return cat(ErrNotObject, '{', '}', ",", objects...)
}

// Arrays returns a concatenation of all entries in the passed in JSON arrays,
// or returns an error if one or more of the passed in byte-slice does not
// appear to contain a valid JSON array. The order of elements is preserved.
func Arrays(arrays ...[]byte) ([]byte, error) {
	return cat(ErrNotArray, '[', ']', ",", arrays...)
}

// Strings returns a concatenation of all entries in the passed in JSON strings,
// or returns an error if one or more of the passed in byte-slice is not appear
// to contain a valid JSON string. The order of elements is preserved.
func Strings(strings ...[]byte) ([]byte, error) {
	return cat(ErrNotString, '"', '"', "", strings...)
}

func cat(w error, start, end byte, sep string, entries ...[]byte) ([]byte, error) {
	// Validate and get a sub-slice of entry content excluding start, stop and
	// whitespace padding. Since we do not modify any bytes, there are no
	// observerabe changes to the entries slices outside the cat function.
	for i, entry := range entries {
		entry = bytes.TrimSpace(entry)
		if l := len(entry); l < 2 || entry[0] != start || entry[l-1] != end {
			return nil, fmt.Errorf("%d: %w", i, w)
		}
		entry = entry[1 : len(entry)-1]
		// When entries has a separated (arrays and objects), then we must trim
		// the entry again to correctly check for empty entries.
		if len(sep) > 0 {
			entry = bytes.TrimSpace(entry)
		}
		entries[i] = entry
	}

	// Allocate target slice.
	allocSize := 2
	for _, c := range entries {
		allocSize += len(c) + len(sep)
	}
	result := make([]byte, 0, allocSize)

	// Concatenate entries.
	result = append(result, start) // length might be 0
	_sep := []byte{}
	nextSep := []byte(sep)
	for _, obj := range entries {
		if len(obj) == 0 {
			continue
		}
		result = append(result, _sep...)
		result = append(result, obj...)
		_sep = nextSep
	}
	result = append(result, end)
	return result, nil
}
