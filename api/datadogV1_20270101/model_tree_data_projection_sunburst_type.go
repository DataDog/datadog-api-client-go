// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreeDataProjectionSunburstType Test-only generated schema for the dashboard versioning proof of concept.
type TreeDataProjectionSunburstType string

// List of TreeDataProjectionSunburstType.
const (
	TREEDATAPROJECTIONSUNBURSTTYPE_SUNBURST TreeDataProjectionSunburstType = "sunburst"
)

var allowedTreeDataProjectionSunburstTypeEnumValues = []TreeDataProjectionSunburstType{
	TREEDATAPROJECTIONSUNBURSTTYPE_SUNBURST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TreeDataProjectionSunburstType) GetAllowedValues() []TreeDataProjectionSunburstType {
	return allowedTreeDataProjectionSunburstTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TreeDataProjectionSunburstType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TreeDataProjectionSunburstType(value)
	return nil
}

// NewTreeDataProjectionSunburstTypeFromValue returns a pointer to a valid TreeDataProjectionSunburstType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTreeDataProjectionSunburstTypeFromValue(v string) (*TreeDataProjectionSunburstType, error) {
	ev := TreeDataProjectionSunburstType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TreeDataProjectionSunburstType: valid values are %v", v, allowedTreeDataProjectionSunburstTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TreeDataProjectionSunburstType) IsValid() bool {
	for _, existing := range allowedTreeDataProjectionSunburstTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TreeDataProjectionSunburstType value.
func (v TreeDataProjectionSunburstType) Ptr() *TreeDataProjectionSunburstType {
	return &v
}
