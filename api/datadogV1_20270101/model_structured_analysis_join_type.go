// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisJoinType Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisJoinType string

// List of StructuredAnalysisJoinType.
const (
	STRUCTUREDANALYSISJOINTYPE_JOIN StructuredAnalysisJoinType = "join"
)

var allowedStructuredAnalysisJoinTypeEnumValues = []StructuredAnalysisJoinType{
	STRUCTUREDANALYSISJOINTYPE_JOIN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StructuredAnalysisJoinType) GetAllowedValues() []StructuredAnalysisJoinType {
	return allowedStructuredAnalysisJoinTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StructuredAnalysisJoinType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StructuredAnalysisJoinType(value)
	return nil
}

// NewStructuredAnalysisJoinTypeFromValue returns a pointer to a valid StructuredAnalysisJoinType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStructuredAnalysisJoinTypeFromValue(v string) (*StructuredAnalysisJoinType, error) {
	ev := StructuredAnalysisJoinType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StructuredAnalysisJoinType: valid values are %v", v, allowedStructuredAnalysisJoinTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StructuredAnalysisJoinType) IsValid() bool {
	for _, existing := range allowedStructuredAnalysisJoinTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StructuredAnalysisJoinType value.
func (v StructuredAnalysisJoinType) Ptr() *StructuredAnalysisJoinType {
	return &v
}
