/*
 * Nbsf_Management
 *
 * Binding Support Management Service API
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type PcfBinding struct {
	Supi       string `json:"supi,omitempty" yaml:"supi" bson:"supi" mapstructure:"Supi"`
	Gpsi       string `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi" mapstructure:"Gpsi"`
	Ipv4Addr   string `json:"ipv4Addr,omitempty" yaml:"ipv4Addr" bson:"ipv4Addr" mapstructure:"Ipv4Addr"`
	Ipv6Prefix string `json:"ipv6Prefix,omitempty" yaml:"ipv6Prefix" bson:"ipv6Prefix" mapstructure:"Ipv6Prefix"`
	IpDomain   string `json:"ipDomain,omitempty" yaml:"ipDomain" bson:"ipDomain" mapstructure:"IpDomain"`
	MacAddr48  string `json:"macAddr48,omitempty" yaml:"macAddr48" bson:"macAddr48" mapstructure:"MacAddr48"`
	Dnn        string `json:"dnn" yaml:"dnn" bson:"dnn" mapstructure:"Dnn"`
	PcfFqdn    string `json:"pcfFqdn,omitempty" yaml:"pcfFqdn" bson:"pcfFqdn" mapstructure:"PcfFqdn"`
	// IP end points of the PCF hosting the Npcf_PolicyAuthorization service. At least one of pcfFqdn or pcfIpEndPoints shall be included if the PCF supports the Npcf_PolicyAuthorization service.
	PcfIpEndPoints []IpEndPoint `json:"pcfIpEndPoints,omitempty" yaml:"pcfIpEndPoints" bson:"pcfIpEndPoints" mapstructure:"PcfIpEndPoints"`
	PcfDiamHost    string       `json:"pcfDiamHost,omitempty" yaml:"pcfDiamHost" bson:"pcfDiamHost" mapstructure:"PcfDiamHost"`
	PcfDiamRealm   string       `json:"pcfDiamRealm,omitempty" yaml:"pcfDiamRealm" bson:"pcfDiamRealm" mapstructure:"PcfDiamRealm"`
	Snssai         *Snssai      `json:"snssai" yaml:"snssai" bson:"snssai" mapstructure:"Snssai"`
	SuppFeat       string       `json:"suppFeat,omitempty" yaml:"suppFeat" bson:"suppFeat" mapstructure:"SuppFeat"`
	PcfId          string       `json:"pcfId,omitempty" yaml:"pcfId" bson:"pcfId" mapstructure:"PcfId"`
	RecoveryTime   *time.Time   `json:"recoveryTime,omitempty" yaml:"recoveryTime" bson:"recoveryTime" mapstructure:"RecoveryTime"`
}
