package strplus

import "testing"

func TestIsEmpty(t *testing.T) {
	t.Logf("Test for success: %t", IsEmpty("           "))
	t.Logf("Test for success: %t", IsEmpty(""))
	t.Logf("Test for success: %t", IsEmpty(" "))
	t.Logf("Test for fail: %t", IsEmpty(" c  sd "))
	t.Logf("Test for fail: %t", IsEmpty(" c "))
	t.Logf("Test for fail: %t", IsEmpty(" s d "))

}

func TestRemoveAllSymbol(t *testing.T) {
	t.Logf("Test for success: %t", RemoveAllSymbol("adsa=12~~cvxxz!@#$asfs%^&( "))
	t.Logf("Test for fail: %t", RemoveAllSymbol(" c  asdasdsd "))
}
