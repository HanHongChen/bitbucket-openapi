/*
 * Ntsctsf_TimeSynchronization Service API
 *
 * TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models
import (
	"time"
)

// Contains the parameters for the subscription to notification of capability of time  synchronization service. 
type TimeSyncExposureSubsc struct {
	Supis []string `json:"supis,omitempty" yaml:"supis" bson:"supis" mapstructure:"Supis"`
	Gpsis []string `json:"gpsis,omitempty" yaml:"gpsis" bson:"gpsis" mapstructure:"Gpsis"`
	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.  
	InterGrpId string `json:"interGrpId,omitempty" yaml:"interGrpId" bson:"interGrpId" mapstructure:"InterGrpId"`
	// String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.  
	ExterGrpId string `json:"exterGrpId,omitempty" yaml:"exterGrpId" bson:"exterGrpId" mapstructure:"ExterGrpId"`
	// Identifies whether the request applies to any UE. This attribute shall set to \"true\" if  applicable for any UE, otherwise, set to \"false\". 
	AnyUeInd bool `json:"anyUeInd,omitempty" yaml:"anyUeInd" bson:"anyUeInd" mapstructure:"AnyUeInd"`
	NotifMethod NotificationMethod `json:"notifMethod,omitempty" yaml:"notifMethod" bson:"notifMethod" mapstructure:"NotifMethod"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn" yaml:"dnn" bson:"dnn" mapstructure:"Dnn"`
	Snssai *Snssai `json:"snssai" yaml:"snssai" bson:"snssai" mapstructure:"Snssai"`
	SubscribedEvents []SubscribedEvent `json:"subscribedEvents" yaml:"subscribedEvents" bson:"subscribedEvents" mapstructure:"SubscribedEvents"`
	EventFilters []EventFilter `json:"eventFilters,omitempty" yaml:"eventFilters" bson:"eventFilters" mapstructure:"EventFilters"`
	// String providing an URI formatted according to RFC 3986.
	SubsNotifUri string `json:"subsNotifUri" yaml:"subsNotifUri" bson:"subsNotifUri" mapstructure:"SubsNotifUri"`
	// Notification Correlation ID assigned by the NF service consumer.
	SubsNotifId string `json:"subsNotifId" yaml:"subsNotifId" bson:"subsNotifId" mapstructure:"SubsNotifId"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxReportNbr int32 `json:"maxReportNbr,omitempty" yaml:"maxReportNbr" bson:"maxReportNbr" mapstructure:"MaxReportNbr"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty" yaml:"expiry" bson:"expiry" mapstructure:"Expiry"`
	// indicating a time in seconds.
	RepPeriod int32 `json:"repPeriod,omitempty" yaml:"repPeriod" bson:"repPeriod" mapstructure:"RepPeriod"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat string `json:"suppFeat,omitempty" yaml:"suppFeat" bson:"suppFeat" mapstructure:"SuppFeat"`
}
