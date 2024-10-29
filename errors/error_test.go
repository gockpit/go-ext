package errors_test

import (
	"errors"
	"testing"

	eerrors "github.com/gockpit/goext/errors"
)

func TestError(t *testing.T) {
	type condition struct {
		name string
		err  *eerrors.Error
	}
	type expect struct {
		err string
	}

	cases := map[*condition]*expect{
		{
			name: "all non empty fields",
			err: &eerrors.Error{
				Inner:   errors.New("inner"),
				Package: "pkg",
				Kind:    "kind",
				Cause:   "cause.",
				Detail:  "detail.",
			},
		}: {
			err: "pkg: kind: cause. detail. [inner]",
		},
		{
			name: "empty inner error",
			err: &eerrors.Error{
				Inner:   nil,
				Package: "pkg",
				Kind:    "kind",
				Cause:   "cause.",
				Detail:  "detail.",
			},
		}: {
			err: "pkg: kind: cause. detail.",
		},
		{
			name: "empty package",
			err: &eerrors.Error{
				Inner:   errors.New("inner"),
				Package: "",
				Kind:    "kind",
				Cause:   "cause.",
				Detail:  "detail.",
			},
		}: {
			err: ": kind: cause. detail. [inner]",
		},
		{
			name: "empty kind",
			err: &eerrors.Error{
				Inner:   errors.New("inner"),
				Package: "pkg",
				Kind:    "",
				Cause:   "cause.",
				Detail:  "detail.",
			},
		}: {
			err: "pkg: cause. detail. [inner]",
		},
		{
			name: "empty cause",
			err: &eerrors.Error{
				Inner:   errors.New("inner"),
				Package: "pkg",
				Kind:    "kind",
				Cause:   "",
				Detail:  "detail.",
			},
		}: {
			err: "pkg: kind:  detail. [inner]",
		},
		{
			name: "all non empty fields",
			err: &eerrors.Error{
				Inner:   errors.New("inner"),
				Package: "pkg",
				Kind:    "kind",
				Cause:   "cause.",
				Detail:  "detail.",
			},
		}: {
			err: "pkg: kind: cause. detail. [inner]",
		},
	}

	for c, e := range cases {
		t.Run(c.name, func(t *testing.T) {
			want := e.err
			got := c.err.Error()
			if want != got {
				t.Errorf("incorrect: got:%s want:%s", got, want)
			}
		})
	}
}
