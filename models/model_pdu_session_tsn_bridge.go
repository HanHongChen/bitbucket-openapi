package models

type PduSessionTsnBridge struct {
	TsnBridgeInfo      *TsnBridgeInformation
	TsnPortManContDstt *PortMangementContainer
	//TsnPortManCont
	TsnPortManContNwtts []PortMangementContainer
}
