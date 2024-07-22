package models

type PduSessionTsnBridge struct {
	TsnBridgeInfo       *TsnBridgeInformation
	TsnBridgeManCont    *BridgeManagementContainer `json:"tsnBridgeManCont,omitempty" yaml:"tsnBridgeManCont" bson:"tsnBridgeManCont" mapstructure:"TsnBridgeManCont"`
	TsnPortManContDstt  *PortManagementContainer
	TsnPortManContNwtts []PortManagementContainer
	UeIpv4Addr          string  `json:"ueIpv4Addr,omitempty" yaml:"ueIpv4Addr" bson:"UeIpv4Addr"`
	Dnn                 string  `json:"dnn,omitempty" yaml:"dnn" bson:"dnn" mapstructure:"Dnn"`
	Snssai              *Snssai `json:"sNssai,omitempty"`
	IpDomain            string  `json:"ipDomain,omitempty" yaml:"ipDomain" bson:"ipDomain" mapstructure:"IpDomain"`
	UeIpv6AddrPrefix    string  `json:"ueIpv6AddrPrefix,omitempty" yaml:"ueIpv6AddrPrefix" bson:"UeIpv6AddrPrefix"`
}
