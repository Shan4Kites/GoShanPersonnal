package unittesting

import "time"

func GetTimeDiffInMins(t1 time.Time, t2 time.Time) int64 {
	return int64(t1.Sub(t2).Minutes())
}

