package models

// TS23.501 5.28.1 P267, 5GS Bridge delay per port pair per traffic class
type TsnBridgeCapability struct {
	IngressPortNum      uint32
	EgressPortNum       uint32
	Traffic_class       uint8
	DependentDelayMin   uint64
	DependentDelayMax   uint64
	IndependentDelayMin uint64
	IndependentDelayMax uint64
	Active              bool
}
