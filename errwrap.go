package errwrap

import (
	"errors"
	"fmt"
	"strings"
)

func Errorf(format string, args ...interface{}) error {
	if format == "" || len(strings.TrimSpace(format)) == 0 {
		return nil
	}

	if args == nil || len(args) == 0 {
		return errors.New(format)
	}

	return errors.New(fmt.Sprintf(format, args...))
}

func WrapString(msg string) error {
	return errors.New(msg)
}

// wrappedError is an implementation of error that has both the
// outer and inner errors.
type wrappedError struct {
	Outer error
	Inner error
}

func (w *wrappedError) Error() string {
	return w.Outer.Error()
}

// Wrap defines that outer wraps inner, returning an error type that
// can be cleanly used with the other methods in this package, such as
// Contains, GetAll, etc.
//
// This function won't modify the error message at all (the outer message
// will be used).
func Wrap(outer, inner error) error {
	return &wrappedError{
		Outer: outer,
		Inner: inner,
	}
}

// Wrapf wraps an error with a formatting message. This is similar to using
// `fmt.Errorf` to wrap an error. If you're using `fmt.Errorf` to wrap
// errors, you should replace it with this.
//
// format is the format of the error message. The string '{{err}}' will
// be replaced with the original error message.
func Wrapf(format string, err error) error {
	outerMsg := "<nil>"
	if err != nil {
		outerMsg = err.Error()
	}

	outer := errors.New(strings.Replace(
		format, "{{err}}", outerMsg, -1))

	return Wrap(outer, err)
}
