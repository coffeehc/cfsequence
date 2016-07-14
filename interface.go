// interface
package cfsequence

import "time"

const (
	//纪元时间
	epoch = int64(1444281954363)
	//节点Id位数
	nodeIdBits = 8
	//最大节点Id
	maxNodeId = -1 ^ (-1 << nodeIdBits)
	nodeMask  = -1 ^ (-1 << nodeIdBits)
	//序列号位数
	indexBits = 10
	//序列号偏移量
	indexShift = nodeIdBits
	indexMask  = -1 ^ (-1 << indexBits)
	//时间戳偏移量
	timestampLeftShift = indexBits + nodeIdBits

	millisecond = int64(time.Millisecond)
)

//Snowflake接口定义
type SequenceService interface {
	//元年时间,使用默认配置epoch
	GetEpoch() int64
	//获取节点Id
	GetNodeId() int64
	//获取下一个Id
	NextId() int64
	//解析sequence
	ParseSequence(sequence int64) *Sequence
	//给出指定时间戳内最小的 SequenceId, 时间戳单位毫秒
	MinId(timestemp int64) int64
}

type Sequence struct {
	//实际的序列号
	Id int64 `json:"sequence"`
	//创建时间
	CreateTime time.Time `json:"createTime"`
	//节点Id
	NodeId int64 `json:"nodeId"`
	//同一毫秒下的索引号
	Index int64 `json:"index"`
}
