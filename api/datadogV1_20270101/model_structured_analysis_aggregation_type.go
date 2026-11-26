// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisAggregationType Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisAggregationType string

// List of StructuredAnalysisAggregationType.
const (
	STRUCTUREDANALYSISAGGREGATIONTYPE_AGGREGATION StructuredAnalysisAggregationType = "aggregation"
)

var allowedStructuredAnalysisAggregationTypeEnumValues = []StructuredAnalysisAggregationType{
	STRUCTUREDANALYSISAGGREGATIONTYPE_AGGREGATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StructuredAnalysisAggregationType) GetAllowedValues() []StructuredAnalysisAggregationType {
	return allowedStructuredAnalysisAggregationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StructuredAnalysisAggregationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StructuredAnalysisAggregationType(value)
	return nil
}

// NewStructuredAnalysisAggregationTypeFromValue returns a pointer to a valid StructuredAnalysisAggregationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStructuredAnalysisAggregationTypeFromValue(v string) (*StructuredAnalysisAggregationType, error) {
	ev := StructuredAnalysisAggregationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StructuredAnalysisAggregationType: valid values are %v", v, allowedStructuredAnalysisAggregationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StructuredAnalysisAggregationType) IsValid() bool {
	for _, existing := range allowedStructuredAnalysisAggregationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StructuredAnalysisAggregationType value.
func (v StructuredAnalysisAggregationType) Ptr() *StructuredAnalysisAggregationType {
	return &v
}
