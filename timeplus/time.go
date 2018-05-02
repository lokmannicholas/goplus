package timeplus

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

const (
	DDMMYYw       = "02/01/2006"
	E8601DAw      = "2006-01-02"
	E8601DZwd     = "2006-01-02T15:04:05+00:00"
	E8601DTwd     = "2006-01-02T15:04:05"
	B8601DAw      = "20060102"
	E8601DTwdZone = "2006-01-02T15:04:05.000Z"
)

//MatchWithE8601DAw ... Match with "2006-01-02"
func MatchWithDDMMYYw(date string) bool {
	match, mErr := regexp.MatchString("^\\d{2}/\\d{2}/\\d{4}$", date)
	if mErr != nil {
		panic(mErr.Error())
		return false
	}
	return match
}

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

//MatchWithE8601DTwdZone ... Match with "2006-01-02T15:04:05.000Z"
func MatchWithE8601DTwdZone(date string) bool {
	match, mErr := regexp.MatchString("^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}.000Z$", date)
	if mErr != nil {
		panic(mErr.Error())
		return false
	}
	return match
}
func dateFormatVaild(qDate string) (*string, string) {
	var datetimeFormat *string
	if qDate != "" {
		var df = time.RFC3339
		//"2006-01-02T00:00:00"
		if MatchWithE8601DTwd(qDate) {
			df = E8601DTwd
		} else
		//"2006-01-02T15:04:05.000Z"
		if MatchWithE8601DTwdZone(qDate) {
			df = E8601DTwdZone
		} else
		//"2006-01-02"
		if MatchWithE8601DAw(qDate) {
			df = E8601DAw
		} else
		//"20060102"
		if MatchWithB8601DAw(qDate) {
			qDate = qDate[0:4] + "-" + qDate[4:6] + "-" + qDate[6:8]
			df = E8601DAw
		} else
		//"2006-01-02T00:00:00+08:00"
		if MatchWithE8601DZwd(qDate) {
			df = time.RFC3339
		} else
		//"02/01/2006"
		if MatchWithDDMMYYw(qDate) {
			ddmmyyyArray := strings.Split(qDate, "/")
			if len(ddmmyyyArray) != 3 {
				return nil, qDate
			}
			df = E8601DAw
			qDate = ddmmyyyArray[2] + "-" + ddmmyyyArray[1] + "-" + ddmmyyyArray[0]
		}
		datetimeFormat = &df
	}
	return datetimeFormat, qDate
}
func Parse(qDate string) (*time.Time, error) {
	if len(qDate) == 0 {
		return nil, errors.New("Input data error")
	}
	df, qDate := dateFormatVaild(qDate)
	if df == nil {
		return nil, errors.New("Input data error")
	}
	t, tErr := time.Parse(*df, qDate)
	if tErr != nil {
		return nil, tErr
	} else {
		return &t, nil
	}

}
func ParseWithLocation(qDate string, location *time.Location) (*time.Time, error) {

	if len(qDate) == 0 {
		return nil, errors.New("Input data error")
	}
	df, qDate := dateFormatVaild(qDate)
	if df == nil {
		return nil, errors.New("Input data error")
	}
	t, tErr := time.ParseInLocation(*df, qDate, location)
	if tErr != nil {
		return nil, tErr
	} else {
		return &t, nil
	}
}
