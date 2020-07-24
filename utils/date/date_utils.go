package date

import (
	"errors"
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

func GetFmtDate(time time.Time) string {
	return time.Format(apiDateLayout)
}

func ConvertToDate(timeStr string) (time.Time, error) {
	date, err := time.Parse(apiDateLayout, timeStr)
	if err != nil {
		return time.Time{}, errors.New("not a valid date format")
	}
	return date, nil
}