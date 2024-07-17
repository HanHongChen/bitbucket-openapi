package models

// TS29.512 5.6.2.41
type TsnBridgeInformation struct {
	BridgeId      uint64
	DsttAddr      string
	DsttResidTime [8]uint8
	DsttPortNum   uint32
}
