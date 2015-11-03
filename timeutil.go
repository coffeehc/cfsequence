// timeutil
package cfsequence

import "time"

//获取当前毫秒时间
func getMillisecond() int64 {
	return time.Now().UnixNano() / millisecond
}

//等待下一毫秒
func tilNextMillis(lastTimestamp int64) int64 {
	timestamp := getMillisecond()
	for timestamp <= lastTimestamp {
		timestamp = getMillisecond()
	}
	return timestamp
}
