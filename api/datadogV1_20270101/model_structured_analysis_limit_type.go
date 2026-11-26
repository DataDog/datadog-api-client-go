// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisLimitType Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisLimitType string

// List of StructuredAnalysisLimitType.
const (
	STRUCTUREDANALYSISLIMITTYPE_LIMIT StructuredAnalysisLimitType = "limit"
)

var allowedStructuredAnalysisLimitTypeEnumValues = []StructuredAnalysisLimitType{
	STRUCTUREDANALYSISLIMITTYPE_LIMIT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StructuredAnalysisLimitType) GetAllowedValues() []StructuredAnalysisLimitType {
	return allowedStructuredAnalysisLimitTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StructuredAnalysisLimitType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StructuredAnalysisLimitType(value)
	return nil
}

// NewStructuredAnalysisLimitTypeFromValue returns a pointer to a valid StructuredAnalysisLimitType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStructuredAnalysisLimitTypeFromValue(v string) (*StructuredAnalysisLimitType, error) {
	ev := StructuredAnalysisLimitType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StructuredAnalysisLimitType: valid values are %v", v, allowedStructuredAnalysisLimitTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StructuredAnalysisLimitType) IsValid() bool {
	for _, existing := range allowedStructuredAnalysisLimitTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StructuredAnalysisLimitType value.
func (v StructuredAnalysisLimitType) Ptr() *StructuredAnalysisLimitType {
	return &v
}
