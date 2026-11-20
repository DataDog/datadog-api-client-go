// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SampledAnalysisQueryType Test-only generated schema for the dashboard versioning proof of concept.
type SampledAnalysisQueryType string

// List of SampledAnalysisQueryType.
const (
	SAMPLEDANALYSISQUERYTYPE_SAMPLED_ANALYSIS SampledAnalysisQueryType = "sampled_analysis"
)

var allowedSampledAnalysisQueryTypeEnumValues = []SampledAnalysisQueryType{
	SAMPLEDANALYSISQUERYTYPE_SAMPLED_ANALYSIS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SampledAnalysisQueryType) GetAllowedValues() []SampledAnalysisQueryType {
	return allowedSampledAnalysisQueryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SampledAnalysisQueryType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SampledAnalysisQueryType(value)
	return nil
}

// NewSampledAnalysisQueryTypeFromValue returns a pointer to a valid SampledAnalysisQueryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSampledAnalysisQueryTypeFromValue(v string) (*SampledAnalysisQueryType, error) {
	ev := SampledAnalysisQueryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SampledAnalysisQueryType: valid values are %v", v, allowedSampledAnalysisQueryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SampledAnalysisQueryType) IsValid() bool {
	for _, existing := range allowedSampledAnalysisQueryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SampledAnalysisQueryType value.
func (v SampledAnalysisQueryType) Ptr() *SampledAnalysisQueryType {
	return &v
}
