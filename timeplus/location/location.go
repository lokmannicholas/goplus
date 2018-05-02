package location

import "time"

func HongKong() *time.Location {
	loca, err := time.LoadLocation("Asia/Hong_Kong")
	if err == nil {
		return nil
	}
	return loca
}
