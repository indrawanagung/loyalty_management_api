package util

import (
	"strconv"
	"time"
)

func GetUnixTimestamp() string {
	return strconv.Itoa(int(time.Now().Unix()))
}
