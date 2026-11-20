// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisColumnExtractionType Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisColumnExtractionType string

// List of StructuredAnalysisColumnExtractionType.
const (
	STRUCTUREDANALYSISCOLUMNEXTRACTIONTYPE_GROK StructuredAnalysisColumnExtractionType = "grok"
)

var allowedStructuredAnalysisColumnExtractionTypeEnumValues = []StructuredAnalysisColumnExtractionType{
	STRUCTUREDANALYSISCOLUMNEXTRACTIONTYPE_GROK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StructuredAnalysisColumnExtractionType) GetAllowedValues() []StructuredAnalysisColumnExtractionType {
	return allowedStructuredAnalysisColumnExtractionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StructuredAnalysisColumnExtractionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StructuredAnalysisColumnExtractionType(value)
	return nil
}

// NewStructuredAnalysisColumnExtractionTypeFromValue returns a pointer to a valid StructuredAnalysisColumnExtractionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStructuredAnalysisColumnExtractionTypeFromValue(v string) (*StructuredAnalysisColumnExtractionType, error) {
	ev := StructuredAnalysisColumnExtractionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StructuredAnalysisColumnExtractionType: valid values are %v", v, allowedStructuredAnalysisColumnExtractionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StructuredAnalysisColumnExtractionType) IsValid() bool {
	for _, existing := range allowedStructuredAnalysisColumnExtractionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StructuredAnalysisColumnExtractionType value.
func (v StructuredAnalysisColumnExtractionType) Ptr() *StructuredAnalysisColumnExtractionType {
	return &v
}
