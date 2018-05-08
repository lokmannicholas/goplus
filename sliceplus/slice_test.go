package sliceplus

import "testing"

func TestContains(t *testing.T) {
	s := []string{"a", "N"}
	t.Logf("Test for success: %t", Contains(s, "a"))
	t.Logf("Test for success: %t", Contains(s, []string{"a", "N"}))
	t.Logf("Test for fail: %t", Contains(s, "c"))

	i := []int{123, 111}
	t.Logf("Test for success: %t", Contains(i, 123))
	t.Logf("Test for fail: %t", Contains(i, 456))

	i64 := []int64{123111, 123456789112356789}
	t.Logf("Test for success: %t", Contains(i64, int64(123456789112356789)))
	t.Logf("Test for fail: %t", Contains(i64, 456))

}
