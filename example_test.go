package pass_test

import (
	"errors"
	"testing"

	"github.com/breml/pass"
)

func TestExample(t *testing.T) {
	tests := []struct {
		name string

		value string

		assertError func(t *testing.T, err error)
		want        string
	}{
		{
			name: "success",

			value: "hello",

			assertError: assertNoError,
			want:        "hello",
		},
		{
			name: "error - empty value not allowed for fooer",

			value: "",

			assertError: func(t *testing.T, err error) {
				assertError(t, err)

				// pass.New ends the test here as successful, since it is not save
				// to continue the test execution, since "fooer" is "nil" in the case
				// of an error and therefore f.String() would panic in the test.
				pass.Now(t)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := newFooer(tt.value)
			tt.assertError(t, err)

			got := f.String()

			assertEqual(t, tt.want, got)
		})
	}
}

type fooer struct {
	value string
}

func newFooer(v string) (*fooer, error) {
	if v == "" {
		return nil, errors.New("boom!")
	}

	return &fooer{
		value: v,
	}, nil
}

func (f *fooer) String() string {
	return f.value
}

// Your assertion package of choice starts here...

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("no error expected, got: %v", err)
	}
}

func assertError(t *testing.T, err error) {
	if err == nil {
		t.Fatal("error expected, got none")
	}
}

func assertEqual(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Fatalf("not equal, expected: %q, got: %q", expected, actual)
	}
}
