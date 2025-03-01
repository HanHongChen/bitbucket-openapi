/*
 * Ntsctsf_TimeSynchronization Service API
 *
 * TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains the notification of capability of time synchronization for a list of UEs. 
type SubsEventNotification struct {
	Event SubscribedEvent `json:"event" yaml:"event" bson:"event" mapstructure:"Event"`
	TimeSyncCapas []TimeSyncCapability `json:"timeSyncCapas,omitempty" yaml:"timeSyncCapas" bson:"timeSyncCapas" mapstructure:"TimeSyncCapas"`
}
