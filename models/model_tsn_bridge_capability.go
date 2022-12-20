package models

// TS23.501 5.28.1 P267, 5GS Bridge delay per port pair per traffic class
type TsnBridgeCapability struct {
	IngressPortNum      uint32
	EgressPortNum       uint32
	Traffic_class       uint8
	DependentDelayMin   float64
	DependentDelayMax   float64
	IndependentDelayMin float64
	IndependentDelayMax float64
}
