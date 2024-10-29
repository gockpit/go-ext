package errors

import (
	"errors"
	"strings"
)

// Error is an general error definition.
type Error struct {
	// Inner is the inner error if any.
	Inner error
	// Package is the package name that raise this error.
	// This should not be empty.
	// Package is compared in [Is] method.
	Package string
	// Kind is the kind of this error if any.
	// Operation name or
	// Kind is compared in [Is] method.
	Kind string
	// Cause is the cause of this error.
	// Cause should not be empty.
	// Cause is compared in [Is] method.
	Cause string
	// Details is the detail of this error.
	// Details is not compared in [Is] which means
	// different two errors with different detail are
	// considered the same.
	Detail string
}

func (e *Error) Error() string {
	var b strings.Builder
	b.WriteString(e.Package)
	b.WriteString(": ")
	if e.Kind != "" {
		b.WriteString(e.Kind)
		b.WriteString(": ")
	}
	b.WriteString(e.Cause)
	if e.Detail != "" {
		b.WriteString(" ")
		b.WriteString(e.Detail)
	}
	if e.Inner != nil {
		b.WriteString(" [")
		b.WriteString(e.Inner.Error())
		b.WriteString("]")
	}
	return b.String()
}

// Wrap sets the given error as the inner error of this error.
// This replaces the existing inner error if any.
// Use [Unwrap] to extract the inner error.
func (e *Error) Wrap(err error) error {
	e.Inner = err
	return e
}

// Unwrap returns inner error if any.
func (e *Error) Unwrap() error {
	return e.Inner
}

func (e *Error) Is(err error) bool {
	if err == nil || e == nil {
		return e == err
	}
	for {
		ee, ok := err.(*Error)
		if ok {
			return e.Package == ee.Package && e.Kind == ee.Kind && e.Cause == ee.Cause
		}
		err = errors.Unwrap(err)
		if err == nil {
			return false
		}
	}
}

// TODO: Implement.
func (e *Error) As(err any) bool {
	panic("This method has not been implemented.")
}
