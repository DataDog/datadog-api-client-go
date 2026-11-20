// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864 Test-only generated schema for the dashboard versioning proof of concept.
type BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864 string

// List of BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864.
const (
	BARCHARTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4443924864_TABULAR BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864 = "tabular"
)

var allowedBarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864EnumValues = []BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864{
	BARCHARTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4443924864_TABULAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864) GetAllowedValues() []BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864 {
	return allowedBarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864(value)
	return nil
}

// NewBarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864FromValue returns a pointer to a valid BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864FromValue(v string) (*BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864, error) {
	ev := BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864: valid values are %v", v, allowedBarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864) IsValid() bool {
	for _, existing := range allowedBarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864 value.
func (v BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864) Ptr() *BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443924864 {
	return &v
}
