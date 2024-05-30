/*
 * Ntsctsf_TimeSynchronization Service API
 *
 * TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains PTP instance configuration and activation requested by the AF.
type PtpInstance struct {
	InstanceType InstanceType `json:"instanceType" yaml:"instanceType" bson:"instanceType" mapstructure:"InstanceType"`
	Protocol Protocol `json:"protocol" yaml:"protocol" bson:"protocol" mapstructure:"Protocol"`
	PtpProfile string `json:"ptpProfile" yaml:"ptpProfile" bson:"ptpProfile" mapstructure:"PtpProfile"`
	PortConfigs []ConfigForPort `json:"portConfigs,omitempty" yaml:"portConfigs" bson:"portConfigs" mapstructure:"PortConfigs"`
}
