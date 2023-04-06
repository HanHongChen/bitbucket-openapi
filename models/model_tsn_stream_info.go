package models

type TsnStreamContext struct {
	TsnStreamInfo []TsnStreamInfomation
}

type TsnStreamInfomation struct {
	StreamId         uint8
	Priority         uint8
	MaxFrameSize     uint32
	FramePerInterval uint32
	Periodicity      uint32
	MaxLatency       uint64
}
