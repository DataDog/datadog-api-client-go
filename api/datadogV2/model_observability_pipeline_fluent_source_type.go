// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineFluentSourceType The source type. The value should always be `fluent`.
type ObservabilityPipelineFluentSourceType string

// List of ObservabilityPipelineFluentSourceType.
const (
	OBSERVABILITYPIPELINEFLUENTSOURCETYPE_FLUENT ObservabilityPipelineFluentSourceType = "fluent"
)

var allowedObservabilityPipelineFluentSourceTypeEnumValues = []ObservabilityPipelineFluentSourceType{
	OBSERVABILITYPIPELINEFLUENTSOURCETYPE_FLUENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineFluentSourceType) GetAllowedValues() []ObservabilityPipelineFluentSourceType {
	return allowedObservabilityPipelineFluentSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineFluentSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineFluentSourceType(value)
	return nil
}

// NewObservabilityPipelineFluentSourceTypeFromValue returns a pointer to a valid ObservabilityPipelineFluentSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineFluentSourceTypeFromValue(v string) (*ObservabilityPipelineFluentSourceType, error) {
	ev := ObservabilityPipelineFluentSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineFluentSourceType: valid values are %v", v, allowedObservabilityPipelineFluentSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineFluentSourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineFluentSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineFluentSourceType value.
func (v ObservabilityPipelineFluentSourceType) Ptr() *ObservabilityPipelineFluentSourceType {
	return &v
}
