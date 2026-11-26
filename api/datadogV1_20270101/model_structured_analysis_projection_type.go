// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisProjectionType Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisProjectionType string

// List of StructuredAnalysisProjectionType.
const (
	STRUCTUREDANALYSISPROJECTIONTYPE_PROJECTION StructuredAnalysisProjectionType = "projection"
)

var allowedStructuredAnalysisProjectionTypeEnumValues = []StructuredAnalysisProjectionType{
	STRUCTUREDANALYSISPROJECTIONTYPE_PROJECTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StructuredAnalysisProjectionType) GetAllowedValues() []StructuredAnalysisProjectionType {
	return allowedStructuredAnalysisProjectionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StructuredAnalysisProjectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StructuredAnalysisProjectionType(value)
	return nil
}

// NewStructuredAnalysisProjectionTypeFromValue returns a pointer to a valid StructuredAnalysisProjectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStructuredAnalysisProjectionTypeFromValue(v string) (*StructuredAnalysisProjectionType, error) {
	ev := StructuredAnalysisProjectionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StructuredAnalysisProjectionType: valid values are %v", v, allowedStructuredAnalysisProjectionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StructuredAnalysisProjectionType) IsValid() bool {
	for _, existing := range allowedStructuredAnalysisProjectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StructuredAnalysisProjectionType value.
func (v StructuredAnalysisProjectionType) Ptr() *StructuredAnalysisProjectionType {
	return &v
}
