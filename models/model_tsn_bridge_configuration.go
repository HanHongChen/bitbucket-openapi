package models

type TsnBridgeConfiguration struct {
	BridgeId       uint64 `json:"bridgeid"`
	Traffic_class  uint8  `json:"traffic_class"`
	IngressPortNum uint32 `json:"ingressPortnum"`
	EgressPortNum  uint32 `json:"egressPortnum"`
}
