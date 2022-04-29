package refactor

import (
	"testing"
)

func TestNoxFactor(t *testing.T) {
	if err := Run("../../"); err != nil {
		t.Fatal(err)
	}
}
