// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544 Test-only generated schema for the dashboard versioning proof of concept.
type BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544 string

// List of BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544.
const (
	BARCHARTBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4443924544_LOCAL_DATASET BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544 = "local_dataset"
)

var allowedBarChartBackendRequestAnyOfAnyOfItemRequestType4443924544EnumValues = []BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544{
	BARCHARTBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4443924544_LOCAL_DATASET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544) GetAllowedValues() []BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544 {
	return allowedBarChartBackendRequestAnyOfAnyOfItemRequestType4443924544EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544(value)
	return nil
}

// NewBarChartBackendRequestAnyOfAnyOfItemRequestType4443924544FromValue returns a pointer to a valid BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBarChartBackendRequestAnyOfAnyOfItemRequestType4443924544FromValue(v string) (*BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544, error) {
	ev := BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544: valid values are %v", v, allowedBarChartBackendRequestAnyOfAnyOfItemRequestType4443924544EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544) IsValid() bool {
	for _, existing := range allowedBarChartBackendRequestAnyOfAnyOfItemRequestType4443924544EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544 value.
func (v BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544) Ptr() *BarChartBackendRequestAnyOfAnyOfItemRequestType4443924544 {
	return &v
}
