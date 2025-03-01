/*
 * Ntsctsf_TimeSynchronization Service API
 *
 * TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains the capability of time synchronization service.
type TimeSyncCapability struct {
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	UpNodeId int32 `json:"upNodeId" yaml:"upNodeId" bson:"upNodeId" mapstructure:"UpNodeId"`
	GmCapables []GmCapable `json:"gmCapables,omitempty" yaml:"gmCapables" bson:"gmCapables" mapstructure:"GmCapables"`
	AsTimeRes AsTimeResource `json:"asTimeRes,omitempty" yaml:"asTimeRes" bson:"asTimeRes" mapstructure:"AsTimeRes"`
	// Contains the PTP capabilities supported by each of the SUPI(s). The key of the map is the SUPI. 
	PtpCapForUes map[string]PtpCapabilitiesPerUe `json:"ptpCapForUes,omitempty" yaml:"ptpCapForUes" bson:"ptpCapForUes" mapstructure:"PtpCapForUes"`
	// Contains the PTP capabilities supported by each of the GPSI(s). The key of the map is the GPSI. 
	PtpCapForGpsis map[string]PtpCapabilitiesPerUe `json:"ptpCapForGpsis,omitempty" yaml:"ptpCapForGpsis" bson:"ptpCapForGpsis" mapstructure:"PtpCapForGpsis"`
}
