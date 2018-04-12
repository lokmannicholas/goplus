package timeplus

import (
	"errors"
	"regexp"
	"time"
)

const (
	E8601DAw      = "2006-01-02"
	E8601DZwd     = "2006-01-02T00:00:00+00:00"
	E8601DTwd     = "2006-01-02T00:00:00"
	B8601DAw      = "20060102"
	E8601DTwdZone = "2006-01-02T15:04:05Z"
)

//MatchWithE8601DAw ... Match with "2006-01-02"
func MatchWithE8601DAw(date string) bool {
	match, mErr := regexp.MatchString("^\\d{4}-\\d{2}-\\d{2}$", date)
	if mErr != nil {
		panic(mErr.Error())
		return false
	}
	return match
}

//MatchWithE8601DZwd ... Match with "2006-01-02T00:00:00+00:00"
func MatchWithE8601DZwd(date string) bool {
	match, mErr := regexp.MatchString("^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}\\+\\d{2}:\\d{2}$", date)
	if mErr != nil {
		panic(mErr.Error())
		return false
	}
	return match
}

//MatchWithE8601DTwd ... Match with "2006-01-02T00:00:00"
func MatchWithE8601DTwd(date string) bool {
	match, mErr := regexp.MatchString("^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}$", date)
	if mErr != nil {
		panic(mErr.Error())
		return false
	}
	return match
}

//MatchWithB8601DAw ... Match with "20060102"
func MatchWithB8601DAw(date string) bool {
	match, mErr := regexp.MatchString("^([0-9]){8}$", date)
	if mErr != nil {
		panic(mErr.Error())
		return false
	}
	return match
}

//MatchWithE8601DTwdZone ... Match with "2006-01-02T15:04:05Z"
func MatchWithE8601DTwdZone(date string) bool {
	match, mErr := regexp.MatchString("^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}Z$", date)
	if mErr != nil {
		panic(mErr.Error())
		return false
	}
	return match
}

func Prase(qDate string) (*time.Time, error) {

	if qDate != "" {
		var datetimeFormat = time.RFC3339
		//"2006-01-02T00:00:00"
		if MatchWithE8601DTwd(qDate) {
			datetimeFormat = E8601DTwd
		}
		//"2006-01-02T15:04:05Z"
		if MatchWithE8601DTwdZone(qDate) {
			datetimeFormat = E8601DTwdZone
		}
		//"2006-01-02"
		if MatchWithE8601DAw(qDate) {
			datetimeFormat = E8601DAw
		}
		//"20060102"
		if MatchWithB8601DAw(qDate) {
			qDate = qDate[0:4] + "-" + qDate[4:6] + "-" + qDate[6:8]
			datetimeFormat = E8601DAw
		}
		//"2006-01-02T00:00:00+08:00"
		if MatchWithE8601DZwd(qDate) {
			datetimeFormat = E8601DZwd
		}
		t, tErr := time.Parse(datetimeFormat, qDate)
		if tErr != nil {
			return nil, tErr
		} else {
			return &t, nil
		}
	}
	return nil, errors.New("Input data error")
}
