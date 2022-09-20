package models

type TsnQosContainer struct {
	MaxTscBurstSize uint
	TscPackDelay    int
	TscPrioLevel    uint
}
