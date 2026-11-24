// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeneratedMetricType The generated metric resource type.
type GeneratedMetricType string

// List of GeneratedMetricType.
const (
	GENERATEDMETRICTYPE_GENERATED_METRIC GeneratedMetricType = "generated_metric"
)

var allowedGeneratedMetricTypeEnumValues = []GeneratedMetricType{
	GENERATEDMETRICTYPE_GENERATED_METRIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GeneratedMetricType) GetAllowedValues() []GeneratedMetricType {
	return allowedGeneratedMetricTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GeneratedMetricType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GeneratedMetricType(value)
	return nil
}

// NewGeneratedMetricTypeFromValue returns a pointer to a valid GeneratedMetricType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGeneratedMetricTypeFromValue(v string) (*GeneratedMetricType, error) {
	ev := GeneratedMetricType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GeneratedMetricType: valid values are %v", v, allowedGeneratedMetricTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GeneratedMetricType) IsValid() bool {
	for _, existing := range allowedGeneratedMetricTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeneratedMetricType value.
func (v GeneratedMetricType) Ptr() *GeneratedMetricType {
	return &v
}
