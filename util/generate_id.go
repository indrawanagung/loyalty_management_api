package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func GenerateIDTransaction(id int) string {
	totalRepeatZero := 5 - (Akar10(id))
	now := time.Now()
	var day string
	if now.Day() < 10 {
		day = fmt.Sprintf("0%s", strconv.Itoa(now.Day()))
	} else {
		day = strconv.Itoa(now.Day())
	}

	var month string
	if now.Month() < 10 {
		month = fmt.Sprintf("0%d", now.Month())
	} else {
		month = fmt.Sprintf("%d", now.Month())
	}
	year := strconv.Itoa(now.Year())

	dateString := fmt.Sprintf("%s%s%s", day, month, year)
	return fmt.Sprintf("TRINV/%s/%s", strings.Repeat("0", totalRepeatZero)+strconv.Itoa(id), dateString) //"TRINV/000001/02022024"
}

func GenerateIDGetMember(id int) string {
	totalRepeatZero := 5 - (id / 10)
	now := time.Now()
	var day string
	if now.Day() < 10 {
		day = fmt.Sprintf("0%s", strconv.Itoa(now.Day()))
	} else {
		day = strconv.Itoa(now.Day())
	}
	var month string
	if now.Month() < 10 {
		day = fmt.Sprintf("0%d", now.Month())
	} else {
		day = fmt.Sprintf("%d", now.Month())
	}
	year := strconv.Itoa(now.Year())

	dateString := fmt.Sprintf("%s%s%s", day, month, year)
	return fmt.Sprintf("TRMGM-%s-%s", strings.Repeat("0", totalRepeatZero)+strconv.Itoa(id), dateString) //"TRINV/000001/02022024"
}

func GenerateIDActivityMember(id int) string {
	totalRepeatZero := 5 - (id / 10)
	now := time.Now()
	var day string
	if now.Day() < 10 {
		day = fmt.Sprintf("0%s", strconv.Itoa(now.Day()))
	} else {
		day = strconv.Itoa(now.Day())
	}
	var month string
	if now.Month() < 10 {
		day = fmt.Sprintf("0%d", now.Month())
	} else {
		day = fmt.Sprintf("%d", now.Month())
	}
	year := strconv.Itoa(now.Year())

	dateString := fmt.Sprintf("%s%s%s", day, month, year)
	return fmt.Sprintf("TRACT-%s-%s", strings.Repeat("0", totalRepeatZero)+strconv.Itoa(id), dateString) //"TRINV/000001/02022024"
}

func Akar10(num int) int {
	count := 0

	for {
		if num < 10 {
			break
		}
		num = num / 10
		count += 1
	}
	return count
}
