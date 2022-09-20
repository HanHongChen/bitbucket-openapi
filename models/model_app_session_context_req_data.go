/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Identifies the service requirements of an Individual Application Session Context.
type AppSessionContextReqData struct {
	// Contains an AF application identifier.
	AfAppId   string                `json:"afAppId,omitempty" yaml:"afAppId" bson:"afAppId" mapstructure:"AfAppId"`
	AfRoutReq *AfRoutingRequirement `json:"afRoutReq,omitempty" yaml:"afRoutReq" bson:"afRoutReq" mapstructure:"AfRoutReq"`
	// Contains an identity of an application service provider.
	AspId string `json:"aspId,omitempty" yaml:"aspId" bson:"aspId" mapstructure:"AspId"`
	// string identifying a BDT Reference ID as defined in subclause 5.3.3 of 3GPP TS 29.154.
	BdtRefId      string                    `json:"bdtRefId,omitempty" yaml:"bdtRefId" bson:"bdtRefId" mapstructure:"BdtRefId"`
	Dnn           string                    `json:"dnn,omitempty" yaml:"dnn" bson:"dnn" mapstructure:"Dnn"`
	EvSubsc       *EventsSubscReqData       `json:"evSubsc,omitempty" yaml:"evSubsc" bson:"evSubsc" mapstructure:"EvSubsc"`
	MedComponents map[string]MediaComponent `json:"medComponents,omitempty" yaml:"medComponents" bson:"medComponents" mapstructure:"MedComponents"`
	IpDomain      string                    `json:"ipDomain,omitempty" yaml:"ipDomain" bson:"ipDomain" mapstructure:"IpDomain"`
	// indication of MPS service request
	MpsId string `json:"mpsId,omitempty" yaml:"mpsId" bson:"mpsId" mapstructure:"MpsId"`
	// string providing an URI formatted according to IETF RFC 3986.
	NotifUri  string  `json:"notifUri" yaml:"notifUri" bson:"notifUri" mapstructure:"NotifUri"`
	SliceInfo *Snssai `json:"sliceInfo,omitempty" yaml:"sliceInfo" bson:"sliceInfo" mapstructure:"SliceInfo"`
	// Contains an identity of a sponsor.
	SponId             string           `json:"sponId,omitempty" yaml:"sponId" bson:"sponId" mapstructure:"SponId"`
	SponStatus         SponsoringStatus `json:"sponStatus,omitempty" yaml:"sponStatus" bson:"sponStatus" mapstructure:"SponStatus"`
	Supi               string           `json:"supi,omitempty" yaml:"supi" bson:"supi" mapstructure:"Supi"`
	Gpsi               string           `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi" mapstructure:"Gpsi"`
	SuppFeat           string           `json:"suppFeat" yaml:"suppFeat" bson:"suppFeat" mapstructure:"SuppFeat"`
	UeIpv4             string           `json:"ueIpv4,omitempty" yaml:"ueIpv4" bson:"ueIpv4" mapstructure:"UeIpv4"`
	UeIpv6             string           `json:"ueIpv6,omitempty" yaml:"ueIpv6" bson:"ueIpv6" mapstructure:"UeIpv6"`
	UeMac              string           `json:"ueMac,omitempty" yaml:"ueMac" bson:"ueMac" mapstructure:"UeMac"`
	TsnPortManContDstt *PortMangementContainer
	TsnPortManContNwtt *PortMangementContainer
}
