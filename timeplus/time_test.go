package timeplus

import (
	"testing"
	"time"
)

func TestFormate(t *testing.T) {
	t.Logf("Test for success: %t", MatchWithDDMMYYw("01/02/2006"))
	t.Logf("Test for success: %t", MatchWithE8601DAw("2006-01-02"))
	t.Logf("Test for success: %t", MatchWithE8601DZwd("2006-01-02T00:00:00+00:00"))
	t.Logf("Test for success: %t", MatchWithE8601DTwd("2006-01-02T00:00:00"))
	t.Logf("Test for success: %t", MatchWithB8601DAw("20060102"))
	t.Logf("Test for success: %t", MatchWithE8601DTwdZone("2006-01-02T15:04:05.000Z"))
	parseTime, err := Parse("2006-01-02T15:04:05.000Z")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Test for success: %t", parseTime.String())
	s := time.Now().Format(B8601DAw)
	t.Log(s)
}
