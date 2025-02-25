/*
 * 3gpp-pfd-management
 *
 * API for PFD management. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models_nef

import (
	"github.com/HanHongChen/bitbucket-openapi/models"
)

type LocationArea5G struct {

	// Identifies a list of geographic area of the user where the UE is located.
	GeographicAreas []models.GeographicArea `json:"geographicAreas,omitempty" bson:"geographicAreas"`

	// Identifies a list of civic addresses of the user where the UE is located.
	CivicAddresses []models.CivicAddress `json:"civicAddresses,omitempty" bson:"civicAddresses"`

	NwAreaInfo *models.NetworkAreaInfo `json:"nwAreaInfo,omitempty" bson:"nwAreaInfo"`
}
