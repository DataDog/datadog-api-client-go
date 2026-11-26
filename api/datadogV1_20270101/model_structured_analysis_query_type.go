// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisQueryType Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisQueryType string

// List of StructuredAnalysisQueryType.
const (
	STRUCTUREDANALYSISQUERYTYPE_STRUCTURED_ANALYSIS StructuredAnalysisQueryType = "structured_analysis"
)

var allowedStructuredAnalysisQueryTypeEnumValues = []StructuredAnalysisQueryType{
	STRUCTUREDANALYSISQUERYTYPE_STRUCTURED_ANALYSIS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StructuredAnalysisQueryType) GetAllowedValues() []StructuredAnalysisQueryType {
	return allowedStructuredAnalysisQueryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StructuredAnalysisQueryType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StructuredAnalysisQueryType(value)
	return nil
}

// NewStructuredAnalysisQueryTypeFromValue returns a pointer to a valid StructuredAnalysisQueryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStructuredAnalysisQueryTypeFromValue(v string) (*StructuredAnalysisQueryType, error) {
	ev := StructuredAnalysisQueryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StructuredAnalysisQueryType: valid values are %v", v, allowedStructuredAnalysisQueryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StructuredAnalysisQueryType) IsValid() bool {
	for _, existing := range allowedStructuredAnalysisQueryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StructuredAnalysisQueryType value.
func (v StructuredAnalysisQueryType) Ptr() *StructuredAnalysisQueryType {
	return &v
}
