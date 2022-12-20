package models

// TS23.501 5.28.1 P267
type TsnBridgeRegistrationInfo struct {
	BridgeId        uint64
	BridgeAddress   string
	Total_port_nums int
	Port_list       []uint32
	Port_pair_info  []TsnBridgeCapability
	Port_info       map[uint32]TsnBridgePortInfo
}
