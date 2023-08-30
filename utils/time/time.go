package time

import (
	"time"
)

const (
	Layout1  = "2006-01-02 15:04:05"
	Layout2  = "2006.01.02 15:04:05"
	Layout3  = "2006-01-02 15:04"
	Layout4  = "20060102150405"
	Layout5  = "2006年01月17日 15:04"
	Layout6  = "20060102"
	Layout7  = "200601021504"
	TimeADay = 86400
)

func GetCurrentTimestampByMill() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetCost(t int64) int64 {
	return time.Now().UnixNano()/1e6 - t
}

func IntToUnix(i int64, layout string) (res string) {
	return time.Unix(i, 0).Format(layout)
}

func StrToUnix(s string, layout string) (res int64, err error) {
	if s == "" {
		return
	}
	t, err := time.ParseInLocation(layout, s, time.Local)
	if err != nil {
		return
	}
	res = t.Unix()
	return
}
