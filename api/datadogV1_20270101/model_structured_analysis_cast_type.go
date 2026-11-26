// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisCastType Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisCastType string

// List of StructuredAnalysisCastType.
const (
	STRUCTUREDANALYSISCASTTYPE_CAST_COLUMN_TYPE StructuredAnalysisCastType = "cast_column_type"
)

var allowedStructuredAnalysisCastTypeEnumValues = []StructuredAnalysisCastType{
	STRUCTUREDANALYSISCASTTYPE_CAST_COLUMN_TYPE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StructuredAnalysisCastType) GetAllowedValues() []StructuredAnalysisCastType {
	return allowedStructuredAnalysisCastTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StructuredAnalysisCastType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StructuredAnalysisCastType(value)
	return nil
}

// NewStructuredAnalysisCastTypeFromValue returns a pointer to a valid StructuredAnalysisCastType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStructuredAnalysisCastTypeFromValue(v string) (*StructuredAnalysisCastType, error) {
	ev := StructuredAnalysisCastType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StructuredAnalysisCastType: valid values are %v", v, allowedStructuredAnalysisCastTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StructuredAnalysisCastType) IsValid() bool {
	for _, existing := range allowedStructuredAnalysisCastTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StructuredAnalysisCastType value.
func (v StructuredAnalysisCastType) Ptr() *StructuredAnalysisCastType {
	return &v
}
