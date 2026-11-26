// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ComputeTableColumnType Test-only generated schema for the dashboard versioning proof of concept.
type ComputeTableColumnType string

// List of ComputeTableColumnType.
const (
	COMPUTETABLECOLUMNTYPE_COMPUTE ComputeTableColumnType = "compute"
)

var allowedComputeTableColumnTypeEnumValues = []ComputeTableColumnType{
	COMPUTETABLECOLUMNTYPE_COMPUTE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ComputeTableColumnType) GetAllowedValues() []ComputeTableColumnType {
	return allowedComputeTableColumnTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ComputeTableColumnType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ComputeTableColumnType(value)
	return nil
}

// NewComputeTableColumnTypeFromValue returns a pointer to a valid ComputeTableColumnType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewComputeTableColumnTypeFromValue(v string) (*ComputeTableColumnType, error) {
	ev := ComputeTableColumnType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ComputeTableColumnType: valid values are %v", v, allowedComputeTableColumnTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ComputeTableColumnType) IsValid() bool {
	for _, existing := range allowedComputeTableColumnTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ComputeTableColumnType value.
func (v ComputeTableColumnType) Ptr() *ComputeTableColumnType {
	return &v
}
