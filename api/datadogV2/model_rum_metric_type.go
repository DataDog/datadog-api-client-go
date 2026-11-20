// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMMetricType The type of the resource. The value must be `rum_metrics`.
type RUMMetricType string

// List of RUMMetricType.
const (
	RUMMETRICTYPE_RUM_METRICS RUMMetricType = "rum_metrics"
)

var allowedRUMMetricTypeEnumValues = []RUMMetricType{
	RUMMETRICTYPE_RUM_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMMetricType) GetAllowedValues() []RUMMetricType {
	return allowedRUMMetricTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMMetricType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMMetricType(value)
	return nil
}

// NewRUMMetricTypeFromValue returns a pointer to a valid RUMMetricType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMMetricTypeFromValue(v string) (*RUMMetricType, error) {
	ev := RUMMetricType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMMetricType: valid values are %v", v, allowedRUMMetricTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMMetricType) IsValid() bool {
	for _, existing := range allowedRUMMetricTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMMetricType value.
func (v RUMMetricType) Ptr() *RUMMetricType {
	return &v
}
