/*
 * Ntsctsf_TimeSynchronization Service API
 *
 * TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains the notification of time synchronization service.
type TimeSyncExposureSubsNotif struct {
	// Notification Correlation ID assigned by the NF service consumer.
	SubsNotifId string `json:"subsNotifId,omitempty" yaml:"subsNotifId" bson:"subsNotifId" mapstructure:"SubsNotifId"`
	EventNotifs []SubsEventNotification `json:"eventNotifs,omitempty" yaml:"eventNotifs" bson:"eventNotifs" mapstructure:"EventNotifs"`
}
