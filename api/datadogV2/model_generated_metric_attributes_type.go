// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeneratedMetricAttributesType The generated metric attributes resource type.
type GeneratedMetricAttributesType string

// List of GeneratedMetricAttributesType.
const (
	GENERATEDMETRICATTRIBUTESTYPE_GENERATED_METRIC_ATTRIBUTES GeneratedMetricAttributesType = "generated_metric_attributes"
)

var allowedGeneratedMetricAttributesTypeEnumValues = []GeneratedMetricAttributesType{
	GENERATEDMETRICATTRIBUTESTYPE_GENERATED_METRIC_ATTRIBUTES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GeneratedMetricAttributesType) GetAllowedValues() []GeneratedMetricAttributesType {
	return allowedGeneratedMetricAttributesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GeneratedMetricAttributesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GeneratedMetricAttributesType(value)
	return nil
}

// NewGeneratedMetricAttributesTypeFromValue returns a pointer to a valid GeneratedMetricAttributesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGeneratedMetricAttributesTypeFromValue(v string) (*GeneratedMetricAttributesType, error) {
	ev := GeneratedMetricAttributesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GeneratedMetricAttributesType: valid values are %v", v, allowedGeneratedMetricAttributesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GeneratedMetricAttributesType) IsValid() bool {
	for _, existing := range allowedGeneratedMetricAttributesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeneratedMetricAttributesType value.
func (v GeneratedMetricAttributesType) Ptr() *GeneratedMetricAttributesType {
	return &v
}
