package dateutil

import "time"

/**
获取当前毫秒时间
*/
func CurrentTimeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

/**
获取当前纳秒时间
*/
func CurrentTimeNanos() int64 {
	return time.Now().UnixNano()
}
