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

type PointAltitudeUncertaintyAllOf struct {
	Point models.GeographicalCoordinates `json:"point"`

	Altitude float64 `json:"altitude"`

	UncertaintyEllipse models.UncertaintyEllipse `json:"uncertaintyEllipse"`

	UncertaintyAltitude float32 `json:"uncertaintyAltitude"`

	Confidence int32 `json:"confidence"`
}
