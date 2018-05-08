package sliceplus

import "testing"

func TestContains(t *testing.T) {
	s := []string{"a", "N"}
	t.Logf("Test for success: %t", Contains(s, "a"))
	t.Logf("Test for success: %t", Contains(s, []string{"a", "N"}))
	t.Logf("Test for fail: %t", Contains(s, "c"))

	i := []int{123, 111}
	t.Logf("Test for success: %t", Contains(i, 123))
	t.Logf("Test for success: %t", Contains(i, 111))
	t.Logf("Test for fail: %t", Contains(i, 456))

	i64 := []int64{123111, 123456789112356789}
	t.Logf("Test for success: %t", Contains(i64, int64(123456789112356789)))
	t.Logf("Test for success: %t", Contains(i64, 123111))
	t.Logf("Test for fail: %t", Contains(i64, 456))

	i64 = append(i64, int64(1233333))
	t.Logf("Test for success: %t", Contains(i64, int64(1233333)))
	t.Logf("Test for fail: %t", Contains(i64, 1233333))

}
func TestFind(t *testing.T) {
	s := []string{"a", "N"}
	t.Logf("Test for success: %d", Find(s, "a"))
	t.Logf("Test for success: %d", Find(s, "N"))
	t.Logf("Test for success: %d", Find(s, []string{"a", "N"}))
	t.Logf("Test for fail: %d", Find(s, "c"))

	i := []int{123, 111}
	t.Logf("Test for success: %d", Find(i, 111))
	t.Logf("Test for success: %d", Find(i, 123))
	t.Logf("Test for fail: %d", Find(i, 456))

	i64 := []int64{123111, 123456789112356789}
	t.Logf("Test for success: %d", Find(i64, int64(123456789112356789)))
	t.Logf("Test for success: %d", Find(i64, int64(123111)))
	t.Logf("Test for success: %d", Find(i64, 123111))
	t.Logf("Test for fail: %d", Find(i64, 456))

}

func TestRemove(t *testing.T) {
	s := []string{"a", "N", "b", "c", "S", "F", "z"}
	t.Logf("Test for success: %v", Remove(s, 0))
	t.Logf("Test for fail: %v", Remove(s, -1))
}
func TestDelete(t *testing.T) {
	s := []string{"a", "N", "b", "c", "S", "F", "z"}
	t.Logf("Test data: %v", s)
	t.Logf("Test for success: %v", Delete(s, "a"))
	t.Logf("Test for success: %v", Delete(s, "b"))
	t.Logf("Test for success: %v", Delete(s, []string{"b", "N"}))
	t.Logf("Test for fail: %v", Delete(s, "h"))
	t.Logf("Test for fail: %v", Delete(s, []string{"H", "a"}))
}
