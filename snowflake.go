// snowflake
package cfsequence

import (
	"sync"
	"time"
)

type _snowflake struct {
	//最后一次的时间戳
	lastTimestamp int64
	//序列号
	index int64
	//节点ID
	nodeId int64
	//
	lock *sync.Mutex
}

func NewSnowflake(nodeId int64) Snowflake {
	if nodeId > maxNodeId {
		return nil
	}
	return &_snowflake{0, 0, int64(nodeId), new(sync.Mutex)}
}

func (this *_snowflake) GetEpoch() int64 {
	return epoch
}

func (this *_snowflake) GetNodeId() int64 {
	return this.nodeId
}

func (this *_snowflake) NextId() int64 {
	this.lock.Lock()
	defer this.lock.Unlock()
	timestamp := getMillisecond()
	if this.lastTimestamp == timestamp {
		this.index = (this.index + 1) & indexMask
		if this.index == 0 {
			//当前毫秒内计数满了，则等待下一秒
			timestamp = tilNextMillis(this.lastTimestamp)
			this.lastTimestamp = timestamp
		}
	} else {
		this.index = 0
		this.lastTimestamp = timestamp
	}
	return (timestamp-epoch)<<timestampLeftShift | this.index<<indexShift | this.nodeId
}

func (this *_snowflake) ParseSequence(sequence int64) Sequence {
	index := sequence >> indexShift & indexMask
	unixTime := (epoch + (sequence >> timestampLeftShift)) * millisecond
	createTime := time.Unix(0, unixTime)
	nodeId := sequence & nodeMask
	return Sequence{sequence, createTime, nodeId, index}
}