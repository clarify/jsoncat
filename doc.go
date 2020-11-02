// Package jsoncat provide functions for concatenating JSON entities of the same
// JSON type (object, array or string) while preserving the order of elements.
// You can choose to conactinate JSON []byte slices directly with the Strings,
// Arrrays ar Objects functions or you can pass in Go types to the Marshall...
// functions as a convenience to first marshal and then concatenate.
//
// Be aware that the package mostly expects the input JSON to be valid. We strip
// leading and trailing white-space, and check for the right delimiters in each
// end of the JSON bytes, but we do not validate anything in between.
package jsoncat
