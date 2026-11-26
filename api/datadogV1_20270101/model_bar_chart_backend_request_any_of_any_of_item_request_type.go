// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartBackendRequestAnyOfAnyOfItemRequestType Test-only generated schema for the dashboard versioning proof of concept.
type BarChartBackendRequestAnyOfAnyOfItemRequestType string

// List of BarChartBackendRequestAnyOfAnyOfItemRequestType.
const (
	BARCHARTBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE_ADVANCED_QUERY BarChartBackendRequestAnyOfAnyOfItemRequestType = "advanced_query"
)

var allowedBarChartBackendRequestAnyOfAnyOfItemRequestTypeEnumValues = []BarChartBackendRequestAnyOfAnyOfItemRequestType{
	BARCHARTBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE_ADVANCED_QUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BarChartBackendRequestAnyOfAnyOfItemRequestType) GetAllowedValues() []BarChartBackendRequestAnyOfAnyOfItemRequestType {
	return allowedBarChartBackendRequestAnyOfAnyOfItemRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BarChartBackendRequestAnyOfAnyOfItemRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BarChartBackendRequestAnyOfAnyOfItemRequestType(value)
	return nil
}

// NewBarChartBackendRequestAnyOfAnyOfItemRequestTypeFromValue returns a pointer to a valid BarChartBackendRequestAnyOfAnyOfItemRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBarChartBackendRequestAnyOfAnyOfItemRequestTypeFromValue(v string) (*BarChartBackendRequestAnyOfAnyOfItemRequestType, error) {
	ev := BarChartBackendRequestAnyOfAnyOfItemRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BarChartBackendRequestAnyOfAnyOfItemRequestType: valid values are %v", v, allowedBarChartBackendRequestAnyOfAnyOfItemRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BarChartBackendRequestAnyOfAnyOfItemRequestType) IsValid() bool {
	for _, existing := range allowedBarChartBackendRequestAnyOfAnyOfItemRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BarChartBackendRequestAnyOfAnyOfItemRequestType value.
func (v BarChartBackendRequestAnyOfAnyOfItemRequestType) Ptr() *BarChartBackendRequestAnyOfAnyOfItemRequestType {
	return &v
}
