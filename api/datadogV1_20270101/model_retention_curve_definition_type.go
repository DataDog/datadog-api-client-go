// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionCurveDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type RetentionCurveDefinitionType string

// List of RetentionCurveDefinitionType.
const (
	RETENTIONCURVEDEFINITIONTYPE_RETENTION_CURVE RetentionCurveDefinitionType = "retention_curve"
)

var allowedRetentionCurveDefinitionTypeEnumValues = []RetentionCurveDefinitionType{
	RETENTIONCURVEDEFINITIONTYPE_RETENTION_CURVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetentionCurveDefinitionType) GetAllowedValues() []RetentionCurveDefinitionType {
	return allowedRetentionCurveDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetentionCurveDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetentionCurveDefinitionType(value)
	return nil
}

// NewRetentionCurveDefinitionTypeFromValue returns a pointer to a valid RetentionCurveDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetentionCurveDefinitionTypeFromValue(v string) (*RetentionCurveDefinitionType, error) {
	ev := RetentionCurveDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetentionCurveDefinitionType: valid values are %v", v, allowedRetentionCurveDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetentionCurveDefinitionType) IsValid() bool {
	for _, existing := range allowedRetentionCurveDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionCurveDefinitionType value.
func (v RetentionCurveDefinitionType) Ptr() *RetentionCurveDefinitionType {
	return &v
}
