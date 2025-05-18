package setime

import "time"

func GetSETimeCurrentDay() (time.Time, time.Time) {
	now := time.Now()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	endTime := startTime.AddDate(0, 0, 1)
	return startTime, endTime
}

func GetSETimeCurrentWeek() (time.Time, time.Time) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	startTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	endTime := startTime.AddDate(0, 0, 7)
	return startTime, endTime
}

func GetSETimeCurrentMonth() (time.Time, time.Time) {
	now := time.Now()
	firstDateTime := now.AddDate(0, 0, -now.Day()+1)
	startTime := time.Date(firstDateTime.Year(), firstDateTime.Month(), firstDateTime.Day(), 0, 0, 0, 0, time.Local)

	lastDateTime := firstDateTime.AddDate(0, 1, 0)
	endTime := time.Date(lastDateTime.Year(), lastDateTime.Month(), lastDateTime.Day(), 0, 0, 0, 0, time.Local)

	return startTime, endTime
}

func GetSETimeCurrentYear() (time.Time, time.Time) {
	now := time.Now()
	currentYear, _, _ := now.Date()
	startTime := time.Date(currentYear, time.January, 1, 0, 0, 0, 0, time.Local)

	endTime := startTime.AddDate(1, 0, 0)
	return startTime, endTime
}
