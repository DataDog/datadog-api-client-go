// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableComputeColumnType
type GuidedTableComputeColumnType string

// List of GuidedTableComputeColumnType.
const (
	GUIDEDTABLECOMPUTECOLUMNTYPE_COMPUTE GuidedTableComputeColumnType = "compute"
)

var allowedGuidedTableComputeColumnTypeEnumValues = []GuidedTableComputeColumnType{
	GUIDEDTABLECOMPUTECOLUMNTYPE_COMPUTE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GuidedTableComputeColumnType) GetAllowedValues() []GuidedTableComputeColumnType {
	return allowedGuidedTableComputeColumnTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GuidedTableComputeColumnType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GuidedTableComputeColumnType(value)
	return nil
}

// NewGuidedTableComputeColumnTypeFromValue returns a pointer to a valid GuidedTableComputeColumnType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGuidedTableComputeColumnTypeFromValue(v string) (*GuidedTableComputeColumnType, error) {
	ev := GuidedTableComputeColumnType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GuidedTableComputeColumnType: valid values are %v", v, allowedGuidedTableComputeColumnTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GuidedTableComputeColumnType) IsValid() bool {
	for _, existing := range allowedGuidedTableComputeColumnTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GuidedTableComputeColumnType value.
func (v GuidedTableComputeColumnType) Ptr() *GuidedTableComputeColumnType {
	return &v
}
