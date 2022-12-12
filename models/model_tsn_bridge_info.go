package models

// TS29.512 5.6.2.41
type TsnBridgeInformation struct {
	DsttAddr      string
	DsttResidTime [8]uint8
	BridgeId      uint64
	NwttPortNums  []uint32
	DsttPortNum   uint32
}
