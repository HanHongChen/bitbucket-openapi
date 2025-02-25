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

type EllipsoidArcAllOf struct {
	Point models.GeographicalCoordinates `json:"point"`

	InnerRadius int32 `json:"innerRadius"`

	UncertaintyRadius float32 `json:"uncertaintyRadius"`

	OffsetAngle int32 `json:"offsetAngle"`

	IncludedAngle int32 `json:"includedAngle"`

	Confidence int32 `json:"confidence"`
}
