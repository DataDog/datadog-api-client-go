// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944 Test-only generated schema for the dashboard versioning proof of concept.
type BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944 string

// List of BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944.
const (
	BARCHARTBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4443922944_DDSQL BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944 = "ddsql"
)

var allowedBarChartBackendRequestAnyOfAnyOfItemRequestType4443922944EnumValues = []BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944{
	BARCHARTBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4443922944_DDSQL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944) GetAllowedValues() []BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944 {
	return allowedBarChartBackendRequestAnyOfAnyOfItemRequestType4443922944EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944(value)
	return nil
}

// NewBarChartBackendRequestAnyOfAnyOfItemRequestType4443922944FromValue returns a pointer to a valid BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBarChartBackendRequestAnyOfAnyOfItemRequestType4443922944FromValue(v string) (*BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944, error) {
	ev := BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944: valid values are %v", v, allowedBarChartBackendRequestAnyOfAnyOfItemRequestType4443922944EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944) IsValid() bool {
	for _, existing := range allowedBarChartBackendRequestAnyOfAnyOfItemRequestType4443922944EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944 value.
func (v BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944) Ptr() *BarChartBackendRequestAnyOfAnyOfItemRequestType4443922944 {
	return &v
}
