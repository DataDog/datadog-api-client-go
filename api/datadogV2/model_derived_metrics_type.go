// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DerivedMetricsType The derived metric resource type.
type DerivedMetricsType string

// List of DerivedMetricsType.
const (
	DERIVEDMETRICSTYPE_DERIVED_METRICS DerivedMetricsType = "derived_metrics"
)

var allowedDerivedMetricsTypeEnumValues = []DerivedMetricsType{
	DERIVEDMETRICSTYPE_DERIVED_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DerivedMetricsType) GetAllowedValues() []DerivedMetricsType {
	return allowedDerivedMetricsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DerivedMetricsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DerivedMetricsType(value)
	return nil
}

// NewDerivedMetricsTypeFromValue returns a pointer to a valid DerivedMetricsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDerivedMetricsTypeFromValue(v string) (*DerivedMetricsType, error) {
	ev := DerivedMetricsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DerivedMetricsType: valid values are %v", v, allowedDerivedMetricsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DerivedMetricsType) IsValid() bool {
	for _, existing := range allowedDerivedMetricsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DerivedMetricsType value.
func (v DerivedMetricsType) Ptr() *DerivedMetricsType {
	return &v
}
