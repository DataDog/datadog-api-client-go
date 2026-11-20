// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatmapBackendRequestAnyOfAnyOfItemRequestType Test-only generated schema for the dashboard versioning proof of concept.
type HeatmapBackendRequestAnyOfAnyOfItemRequestType string

// List of HeatmapBackendRequestAnyOfAnyOfItemRequestType.
const (
	HEATMAPBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE_HISTOGRAM HeatmapBackendRequestAnyOfAnyOfItemRequestType = "histogram"
)

var allowedHeatmapBackendRequestAnyOfAnyOfItemRequestTypeEnumValues = []HeatmapBackendRequestAnyOfAnyOfItemRequestType{
	HEATMAPBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE_HISTOGRAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatmapBackendRequestAnyOfAnyOfItemRequestType) GetAllowedValues() []HeatmapBackendRequestAnyOfAnyOfItemRequestType {
	return allowedHeatmapBackendRequestAnyOfAnyOfItemRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatmapBackendRequestAnyOfAnyOfItemRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatmapBackendRequestAnyOfAnyOfItemRequestType(value)
	return nil
}

// NewHeatmapBackendRequestAnyOfAnyOfItemRequestTypeFromValue returns a pointer to a valid HeatmapBackendRequestAnyOfAnyOfItemRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatmapBackendRequestAnyOfAnyOfItemRequestTypeFromValue(v string) (*HeatmapBackendRequestAnyOfAnyOfItemRequestType, error) {
	ev := HeatmapBackendRequestAnyOfAnyOfItemRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatmapBackendRequestAnyOfAnyOfItemRequestType: valid values are %v", v, allowedHeatmapBackendRequestAnyOfAnyOfItemRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatmapBackendRequestAnyOfAnyOfItemRequestType) IsValid() bool {
	for _, existing := range allowedHeatmapBackendRequestAnyOfAnyOfItemRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatmapBackendRequestAnyOfAnyOfItemRequestType value.
func (v HeatmapBackendRequestAnyOfAnyOfItemRequestType) Ptr() *HeatmapBackendRequestAnyOfAnyOfItemRequestType {
	return &v
}
