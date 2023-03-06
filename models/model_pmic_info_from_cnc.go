package models

type PMICInfoFromCNC struct {
	BridgeId uint64  `json:"bridgeid"`
	PortNum  uint32  `json:"Portnum"`
	Pmic     []uint8 `json:"pmic"`
}
