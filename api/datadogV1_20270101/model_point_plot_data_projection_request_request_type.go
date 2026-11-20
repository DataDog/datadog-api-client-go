// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PointPlotDataProjectionRequestRequestType Test-only generated schema for the dashboard versioning proof of concept.
type PointPlotDataProjectionRequestRequestType string

// List of PointPlotDataProjectionRequestRequestType.
const (
	POINTPLOTDATAPROJECTIONREQUESTREQUESTTYPE_DATA_PROJECTION PointPlotDataProjectionRequestRequestType = "data_projection"
)

var allowedPointPlotDataProjectionRequestRequestTypeEnumValues = []PointPlotDataProjectionRequestRequestType{
	POINTPLOTDATAPROJECTIONREQUESTREQUESTTYPE_DATA_PROJECTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PointPlotDataProjectionRequestRequestType) GetAllowedValues() []PointPlotDataProjectionRequestRequestType {
	return allowedPointPlotDataProjectionRequestRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PointPlotDataProjectionRequestRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PointPlotDataProjectionRequestRequestType(value)
	return nil
}

// NewPointPlotDataProjectionRequestRequestTypeFromValue returns a pointer to a valid PointPlotDataProjectionRequestRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPointPlotDataProjectionRequestRequestTypeFromValue(v string) (*PointPlotDataProjectionRequestRequestType, error) {
	ev := PointPlotDataProjectionRequestRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PointPlotDataProjectionRequestRequestType: valid values are %v", v, allowedPointPlotDataProjectionRequestRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PointPlotDataProjectionRequestRequestType) IsValid() bool {
	for _, existing := range allowedPointPlotDataProjectionRequestRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PointPlotDataProjectionRequestRequestType value.
func (v PointPlotDataProjectionRequestRequestType) Ptr() *PointPlotDataProjectionRequestRequestType {
	return &v
}
