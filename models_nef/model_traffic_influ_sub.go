/*
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models_nef

import (
	"github.com/HanHongChen/bitbucket-openapi/openapi/models"
)

type TrafficInfluSub struct {

	// Identifies a service on behalf of which the AF is issuing the request.
	AfServiceId string `json:"afServiceId,omitempty" bson:"afServiceId"`

	// Identifies an application.
	AfAppId string `json:"afAppId,omitempty" bson:"afAppId"`

	// Identifies an NEF Northbound interface transaction, generated by the AF.
	AfTransId string `json:"afTransId,omitempty" bson:"afTransId"`

	// Identifies whether an application can be relocated once a location of the application has been selected.
	AppReloInd bool `json:"appReloInd,omitempty" bson:"appReloInd"`

	Dnn string `json:"dnn,omitempty" bson:"dnn"`

	Snssai *models.Snssai `json:"snssai,omitempty" bson:"snssai"`

	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExternalGroupId string `json:"externalGroupId,omitempty" bson:"externalGroupId"`

	// Identifies whether the AF request applies to any UE. This attribute shall set to \"true\" if applicable for any UE, otherwise, set to \"false\".
	AnyUeInd bool `json:"anyUeInd,omitempty" bson:"anyUeInd"`

	// Identifies the requirement to be notified of the event(s).
	SubscribedEvents []models.SubscribedEvent `json:"subscribedEvents,omitempty" bson:"subscribedEvents"`

	Gpsi string `json:"gpsi,omitempty" bson:"gpsi"`

	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	Ipv4Addr string `json:"ipv4Addr,omitempty" bson:"ipv4Addr"`

	// string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used.
	Ipv6Addr string `json:"ipv6Addr,omitempty" bson:"ipv6Addr"`

	MacAddr string `json:"macAddr,omitempty" bson:"macAddr"`

	DnaiChgType models.DnaiChangeType `json:"dnaiChgType,omitempty" bson:"dnaiChgType"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination string `json:"notificationDestination,omitempty" bson:"notificationDestination"`

	// Set to true by the SCS/AS to request the NEF to send a test notification as defined in subclause 5.2.5.3. Set to false or omitted otherwise.
	RequestTestNotification bool `json:"requestTestNotification,omitempty" bson:"requestTestNotification"`

	WebsockNotifConfig *models.WebsockNotifConfig `json:"websockNotifConfig,omitempty" bson:"websockNotifConfig"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self string `json:"self,omitempty" bson:"self"`

	// Identifies IP packet filters.
	TrafficFilters []models.FlowInfo `json:"trafficFilters,omitempty" bson:"trafficFilters"`

	// Identifies Ethernet packet filters.
	EthTrafficFilters []models.EthFlowDescription `json:"ethTrafficFilters,omitempty" bson:"ethTrafficFilters"`

	// Identifies the N6 traffic routing requirement.
	TrafficRoutes []models.RouteToLocation `json:"trafficRoutes,omitempty" bson:"trafficRoutes"`

	TfcCorrInd bool `json:"tfcCorrInd,omitempty" bson:"tfcCorrInd"`

	TempValidities []models.TemporalValidity `json:"tempValidities,omitempty" bson:"tempValidities"`

	// Identifies a geographic zone that the AF request applies only to the traffic of UE(s) located in this specific zone.
	ValidGeoZoneIds []string `json:"validGeoZoneIds,omitempty" bson:"validGeoZoneIds"`

	AfAckInd bool `json:"afAckInd,omitempty" bson:"afAckInd"`

	AddrPreserInd bool `json:"addrPreserInd,omitempty" bson:"addrPreserInd"`

	SuppFeat string `json:"suppFeat,omitempty" bson:"suppFeat"`
}
