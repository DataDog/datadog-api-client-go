// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisCastTargetType Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisCastTargetType string

// List of StructuredAnalysisCastTargetType.
const (
	STRUCTUREDANALYSISCASTTARGETTYPE_STRING    StructuredAnalysisCastTargetType = "string"
	STRUCTUREDANALYSISCASTTARGETTYPE_INT64     StructuredAnalysisCastTargetType = "int64"
	STRUCTUREDANALYSISCASTTARGETTYPE_FLOAT64   StructuredAnalysisCastTargetType = "float64"
	STRUCTUREDANALYSISCASTTARGETTYPE_TIMESTAMP StructuredAnalysisCastTargetType = "timestamp"
)

var allowedStructuredAnalysisCastTargetTypeEnumValues = []StructuredAnalysisCastTargetType{
	STRUCTUREDANALYSISCASTTARGETTYPE_STRING,
	STRUCTUREDANALYSISCASTTARGETTYPE_INT64,
	STRUCTUREDANALYSISCASTTARGETTYPE_FLOAT64,
	STRUCTUREDANALYSISCASTTARGETTYPE_TIMESTAMP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StructuredAnalysisCastTargetType) GetAllowedValues() []StructuredAnalysisCastTargetType {
	return allowedStructuredAnalysisCastTargetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StructuredAnalysisCastTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StructuredAnalysisCastTargetType(value)
	return nil
}

// NewStructuredAnalysisCastTargetTypeFromValue returns a pointer to a valid StructuredAnalysisCastTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStructuredAnalysisCastTargetTypeFromValue(v string) (*StructuredAnalysisCastTargetType, error) {
	ev := StructuredAnalysisCastTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StructuredAnalysisCastTargetType: valid values are %v", v, allowedStructuredAnalysisCastTargetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StructuredAnalysisCastTargetType) IsValid() bool {
	for _, existing := range allowedStructuredAnalysisCastTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StructuredAnalysisCastTargetType value.
func (v StructuredAnalysisCastTargetType) Ptr() *StructuredAnalysisCastTargetType {
	return &v
}
