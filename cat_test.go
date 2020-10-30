package jsoncat_test

import (
	"testing"

	"github.com/searis/subtest"

	"github.com/searis/jsoncat"
)

func TestArrays(t *testing.T) {
	const (
		// Added whitespace padding to make it hard.
		i0 = ` [ ] `
		i1 = `["foo","bar"  ]   ` + "\n\t"
		i2 = `  [   "foo","baz","baz","foobar"]`

		// Concat results.
		r0   = `[]`
		r1   = `["foo","bar"]`
		r2   = `["foo","bar"]`
		r12  = `["foo","bar","foo","baz","baz","foobar"]`
		r121 = `["foo","bar","foo","baz","baz","foobar","foo","bar"]`
	)

	t.Run("0 arrays",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Arrays()
		}).DeepEqual([]byte(r0)))
	t.Run("1 array",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Arrays([]byte(i1))
		}).DeepEqual([]byte(r1)))
	t.Run("2 arrays",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Arrays(
				[]byte(i1),
				[]byte(i2),
			)
		}).DeepEqual([]byte(r12)))
	t.Run("3 arrays",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Arrays(
				[]byte(i1),
				[]byte(i2),
				[]byte(i1),
			)
		}).DeepEqual([]byte(r121)))
	t.Run("1 empty array",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Arrays(
				[]byte(i0),
			)
		}).DeepEqual([]byte(r0)))
	t.Run("2 empty arrays",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Arrays(
				[]byte(i0),
				[]byte(i0),
			)
		}).DeepEqual([]byte(r0)))
	t.Run("3 arrays + 2 empty arrays",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Arrays(
				[]byte(i1),
				[]byte(i0),
				[]byte(i2),
				[]byte(i1),
				[]byte(i0),
			)
		}).DeepEqual(
			[]byte(r121),
		),
	)
}

func TestObjects(t *testing.T) {
	const (
		// Added whitespace padding to make it hard.
		i0 = ` { } `
		i1 = "\n\t" + ` {  "foo":"bar"}`
		i2 = `{"foo":"baz","baz":"foobar"  }  `

		// Concat results.
		r0   = `{}`
		r1   = `{"foo":"bar"}`
		r2   = `{"foo":"baz","baz":"foobar"}`
		r12  = `{"foo":"bar","foo":"baz","baz":"foobar"}`
		r121 = `{"foo":"bar","foo":"baz","baz":"foobar","foo":"bar"}`
	)
	t.Run("0 objects",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Objects()
		}).DeepEqual([]byte(r0)))
	t.Run("1 object",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Objects([]byte(i1))
		}).DeepEqual([]byte(r1)))
	t.Run("2 objects",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Objects(
				[]byte(i1),
				[]byte(i2),
			)
		}).DeepEqual([]byte(r12)))
	t.Run("3 objects",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Objects(
				[]byte(i1),
				[]byte(i2),
				[]byte(i1),
			)
		}).DeepEqual([]byte(r121)))
	t.Run("1 empty object",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Objects(
				[]byte(i0),
			)
		}).DeepEqual([]byte(r0)))
	t.Run("2 empty objects",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Objects(
				[]byte(i0),
				[]byte(i0),
			)
		}).DeepEqual([]byte(r0)))
	t.Run("3 objects + 2 empty objects",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Objects(
				[]byte(i1),
				[]byte(i0),
				[]byte(i2),
				[]byte(i1),
				[]byte(i0),
			)
		}).DeepEqual(
			[]byte(r121),
		),
	)
}

func TestStrings(t *testing.T) {
	const (
		// Added whitespace padding to make it hard.
		i0 = ` "" `
		i1 = "\n\t" + `   "foo\" bar"`
		i2 = `"foobar \"baz\" foobar}  "    `

		// Concat results.
		r0   = `""`
		r1   = `"foo\" bar"`
		r2   = `"foobar \"baz\" foobar}  "`
		r12  = `"foo\" barfoobar \"baz\" foobar}  "`
		r121 = `"foo\" barfoobar \"baz\" foobar}  foo\" bar"`
	)
	t.Run("0 strings",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Strings()
		}).DeepEqual([]byte(r0)))
	t.Run("1 string",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Strings([]byte(i1))
		}).DeepEqual([]byte(r1)))
	t.Run("2 strings",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Strings(
				[]byte(i1),
				[]byte(i2),
			)
		}).DeepEqual([]byte(r12)))
	t.Run("3 strings",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Strings(
				[]byte(i1),
				[]byte(i2),
				[]byte(i1),
			)
		}).DeepEqual([]byte(r121)))
	t.Run("1 empty string",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Strings(
				[]byte(i0),
			)
		}).DeepEqual([]byte(r0)))
	t.Run("2 empty strings",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Strings(
				[]byte(i0),
				[]byte(i0),
			)
		}).DeepEqual([]byte(r0)))
	t.Run("3 strings + 2 empty strings",
		subtest.ValueFunc(func() (interface{}, error) {
			return jsoncat.Strings(
				[]byte(i1),
				[]byte(i0),
				[]byte(i2),
				[]byte(i1),
				[]byte(i0),
			)
		}).DeepEqual(
			[]byte(r121),
		),
	)
}
