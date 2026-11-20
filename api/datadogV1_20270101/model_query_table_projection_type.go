// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableProjectionType Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableProjectionType string

// List of QueryTableProjectionType.
const (
	QUERYTABLEPROJECTIONTYPE_QUERY_TABLE QueryTableProjectionType = "query_table"
)

var allowedQueryTableProjectionTypeEnumValues = []QueryTableProjectionType{
	QUERYTABLEPROJECTIONTYPE_QUERY_TABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryTableProjectionType) GetAllowedValues() []QueryTableProjectionType {
	return allowedQueryTableProjectionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryTableProjectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryTableProjectionType(value)
	return nil
}

// NewQueryTableProjectionTypeFromValue returns a pointer to a valid QueryTableProjectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryTableProjectionTypeFromValue(v string) (*QueryTableProjectionType, error) {
	ev := QueryTableProjectionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryTableProjectionType: valid values are %v", v, allowedQueryTableProjectionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryTableProjectionType) IsValid() bool {
	for _, existing := range allowedQueryTableProjectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryTableProjectionType value.
func (v QueryTableProjectionType) Ptr() *QueryTableProjectionType {
	return &v
}
