package models

// TS23.501 5.28.1 P267
type TsnBridgePortInfo struct {
	Traffic_class_table []uint8
	TxPropagationDelay  uint32
}
