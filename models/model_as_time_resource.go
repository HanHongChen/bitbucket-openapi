/*
 * Ntsctsf_TimeSynchronization Service API
 *
 * TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models
type AsTimeResource string

// List of AsTimeResource
const (
	AsTimeResource_ATOMIC_CLOCK AsTimeResource = "ATOMIC_CLOCK"
	AsTimeResource_GNSS AsTimeResource = "GNSS"
	AsTimeResource_TERRESTRIAL_RADIO AsTimeResource = "TERRESTRIAL_RADIO"
	AsTimeResource_SERIAL_TIME_CODE AsTimeResource = "SERIAL_TIME_CODE"
	AsTimeResource_PTP AsTimeResource = "PTP"
	AsTimeResource_NTP AsTimeResource = "NTP"
	AsTimeResource_HAND_SET AsTimeResource = "HAND_SET"
	AsTimeResource_INTERNAL_OSCILLATOR AsTimeResource = "INTERNAL_OSCILLATOR"
	AsTimeResource_OTHER AsTimeResource = "OTHER"
)