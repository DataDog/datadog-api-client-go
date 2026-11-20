// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisFilterType Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisFilterType string

// List of StructuredAnalysisFilterType.
const (
	STRUCTUREDANALYSISFILTERTYPE_FILTER StructuredAnalysisFilterType = "filter"
)

var allowedStructuredAnalysisFilterTypeEnumValues = []StructuredAnalysisFilterType{
	STRUCTUREDANALYSISFILTERTYPE_FILTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StructuredAnalysisFilterType) GetAllowedValues() []StructuredAnalysisFilterType {
	return allowedStructuredAnalysisFilterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StructuredAnalysisFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StructuredAnalysisFilterType(value)
	return nil
}

// NewStructuredAnalysisFilterTypeFromValue returns a pointer to a valid StructuredAnalysisFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStructuredAnalysisFilterTypeFromValue(v string) (*StructuredAnalysisFilterType, error) {
	ev := StructuredAnalysisFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StructuredAnalysisFilterType: valid values are %v", v, allowedStructuredAnalysisFilterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StructuredAnalysisFilterType) IsValid() bool {
	for _, existing := range allowedStructuredAnalysisFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StructuredAnalysisFilterType value.
func (v StructuredAnalysisFilterType) Ptr() *StructuredAnalysisFilterType {
	return &v
}
