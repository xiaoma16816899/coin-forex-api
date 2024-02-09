package utils

import "time"

func GetCurrentTimeStamp() int64 {
	return time.Now().Unix()
}
