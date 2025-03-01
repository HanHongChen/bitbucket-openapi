/*
 * Ntsctsf_TimeSynchronization Service API
 *
 * TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains the filter conditions to match for notifying the event(s) of time synchronization capabilities. 
type EventFilter struct {
	InstanceTypes []InstanceType `json:"instanceTypes,omitempty" yaml:"instanceTypes" bson:"instanceTypes" mapstructure:"InstanceTypes"`
	TransProtocols []Protocol `json:"transProtocols,omitempty" yaml:"transProtocols" bson:"transProtocols" mapstructure:"TransProtocols"`
	PtpProfiles []string `json:"ptpProfiles,omitempty" yaml:"ptpProfiles" bson:"ptpProfiles" mapstructure:"PtpProfiles"`
}
