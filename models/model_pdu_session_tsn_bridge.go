package models

type PduSessionTsnBridge struct {
	TsnBridgeInfo      *TsnBridgeInformation
	TsnPortManContDstt *PortManagementContainer
	//TsnPortManCont
	TsnPortManContNwtts []PortManagementContainer
}
