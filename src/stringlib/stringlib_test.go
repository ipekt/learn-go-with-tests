package stringlib

import "testing"

func TestHasP(t *testing.T) {

	assert := func(t testing.TB, result, expected bool) {

		t.Helper()
		if result != expected {
			t.Errorf("expected %t but got %t", result, expected)
		}
	}

	t.Run("does not have p", func(t *testing.T) {
		result := HasP("beck")
		expected := false
		assert(t, result, expected)
	})

	t.Run("has p", func(t *testing.T) {
		result := HasP("peckles")
		expected := true
		assert(t, result, expected)
	})

}
