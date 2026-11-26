// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisDestructureType Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisDestructureType string

// List of StructuredAnalysisDestructureType.
const (
	STRUCTUREDANALYSISDESTRUCTURETYPE_DESTRUCTURE StructuredAnalysisDestructureType = "destructure"
)

var allowedStructuredAnalysisDestructureTypeEnumValues = []StructuredAnalysisDestructureType{
	STRUCTUREDANALYSISDESTRUCTURETYPE_DESTRUCTURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StructuredAnalysisDestructureType) GetAllowedValues() []StructuredAnalysisDestructureType {
	return allowedStructuredAnalysisDestructureTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StructuredAnalysisDestructureType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StructuredAnalysisDestructureType(value)
	return nil
}

// NewStructuredAnalysisDestructureTypeFromValue returns a pointer to a valid StructuredAnalysisDestructureType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStructuredAnalysisDestructureTypeFromValue(v string) (*StructuredAnalysisDestructureType, error) {
	ev := StructuredAnalysisDestructureType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StructuredAnalysisDestructureType: valid values are %v", v, allowedStructuredAnalysisDestructureTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StructuredAnalysisDestructureType) IsValid() bool {
	for _, existing := range allowedStructuredAnalysisDestructureTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StructuredAnalysisDestructureType value.
func (v StructuredAnalysisDestructureType) Ptr() *StructuredAnalysisDestructureType {
	return &v
}
