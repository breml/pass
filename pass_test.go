package pass_test

import (
	"testing"

	"github.com/breml/pass"
)

func TestNow(t *testing.T) {
	pass.Now(t)

	t.Fatal("this code must not be reached")
}

func TestSkipNow(t *testing.T) {
	t.SkipNow()

	t.Fatal("this code must not be reached")
}
