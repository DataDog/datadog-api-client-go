// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableProjectionDimensionDimension Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableProjectionDimensionDimension string

// List of QueryTableProjectionDimensionDimension.
const (
	QUERYTABLEPROJECTIONDIMENSIONDIMENSION_COLUMN QueryTableProjectionDimensionDimension = "column"
)

var allowedQueryTableProjectionDimensionDimensionEnumValues = []QueryTableProjectionDimensionDimension{
	QUERYTABLEPROJECTIONDIMENSIONDIMENSION_COLUMN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryTableProjectionDimensionDimension) GetAllowedValues() []QueryTableProjectionDimensionDimension {
	return allowedQueryTableProjectionDimensionDimensionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryTableProjectionDimensionDimension) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryTableProjectionDimensionDimension(value)
	return nil
}

// NewQueryTableProjectionDimensionDimensionFromValue returns a pointer to a valid QueryTableProjectionDimensionDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryTableProjectionDimensionDimensionFromValue(v string) (*QueryTableProjectionDimensionDimension, error) {
	ev := QueryTableProjectionDimensionDimension(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryTableProjectionDimensionDimension: valid values are %v", v, allowedQueryTableProjectionDimensionDimensionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryTableProjectionDimensionDimension) IsValid() bool {
	for _, existing := range allowedQueryTableProjectionDimensionDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryTableProjectionDimensionDimension value.
func (v QueryTableProjectionDimensionDimension) Ptr() *QueryTableProjectionDimensionDimension {
	return &v
}
