//REF: https://pkg.go.dev/github.com/matryer/is#section-readme
package main

import (
	"errors"
	"testing"

	"github.com/matryer/is"
)

func Test(t *testing.T) {
	// is := is.New(t)
	is := is.NewRelaxed(t)

	a, b := 1, 2

	is.Equal(a, b) // expect to be the same

	is.Fail() // TODO: write this test

	t.Run("sub", func(t *testing.T) {
		is := is.NewRelaxed(t)
		a, b := 2, 4
		is.Equal(a, b) // expect to be the same again
		// TODO: test
	})

	val, err := "", errors.New("can't read")
	is.NoErr(err)          // getVal error
	is.True(len(val) > 10) // val cannot be short

	/*
		signedin, err := isSignedIn(ctx)
		is.NoErr(err)            // isSignedIn error
		is.Equal(signedin, true) // must be signed in

		body := readBody(r)
		is.True(strings.Contains(body, "Hi there"))
	*/
}
