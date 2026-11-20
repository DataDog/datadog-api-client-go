// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineCrowdStrikeFdrSourceType The source type. Always `crowdstrike_fdr`.
type ObservabilityPipelineCrowdStrikeFdrSourceType string

// List of ObservabilityPipelineCrowdStrikeFdrSourceType.
const (
	OBSERVABILITYPIPELINECROWDSTRIKEFDRSOURCETYPE_CROWDSTRIKE_FDR ObservabilityPipelineCrowdStrikeFdrSourceType = "crowdstrike_fdr"
)

var allowedObservabilityPipelineCrowdStrikeFdrSourceTypeEnumValues = []ObservabilityPipelineCrowdStrikeFdrSourceType{
	OBSERVABILITYPIPELINECROWDSTRIKEFDRSOURCETYPE_CROWDSTRIKE_FDR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineCrowdStrikeFdrSourceType) GetAllowedValues() []ObservabilityPipelineCrowdStrikeFdrSourceType {
	return allowedObservabilityPipelineCrowdStrikeFdrSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineCrowdStrikeFdrSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineCrowdStrikeFdrSourceType(value)
	return nil
}

// NewObservabilityPipelineCrowdStrikeFdrSourceTypeFromValue returns a pointer to a valid ObservabilityPipelineCrowdStrikeFdrSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineCrowdStrikeFdrSourceTypeFromValue(v string) (*ObservabilityPipelineCrowdStrikeFdrSourceType, error) {
	ev := ObservabilityPipelineCrowdStrikeFdrSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineCrowdStrikeFdrSourceType: valid values are %v", v, allowedObservabilityPipelineCrowdStrikeFdrSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineCrowdStrikeFdrSourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineCrowdStrikeFdrSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineCrowdStrikeFdrSourceType value.
func (v ObservabilityPipelineCrowdStrikeFdrSourceType) Ptr() *ObservabilityPipelineCrowdStrikeFdrSourceType {
	return &v
}
