// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScaleType Test-only generated schema for the dashboard versioning proof of concept.
type ScaleType string

// List of ScaleType.
const (
	SCALETYPE_LINEAR ScaleType = "linear"
	SCALETYPE_LOG    ScaleType = "log"
	SCALETYPE_POW    ScaleType = "pow"
	SCALETYPE_SQRT   ScaleType = "sqrt"
)

var allowedScaleTypeEnumValues = []ScaleType{
	SCALETYPE_LINEAR,
	SCALETYPE_LOG,
	SCALETYPE_POW,
	SCALETYPE_SQRT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScaleType) GetAllowedValues() []ScaleType {
	return allowedScaleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScaleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScaleType(value)
	return nil
}

// NewScaleTypeFromValue returns a pointer to a valid ScaleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScaleTypeFromValue(v string) (*ScaleType, error) {
	ev := ScaleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScaleType: valid values are %v", v, allowedScaleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScaleType) IsValid() bool {
	for _, existing := range allowedScaleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScaleType value.
func (v ScaleType) Ptr() *ScaleType {
	return &v
}
