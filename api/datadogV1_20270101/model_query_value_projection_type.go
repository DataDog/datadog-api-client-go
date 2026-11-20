// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryValueProjectionType Test-only generated schema for the dashboard versioning proof of concept.
type QueryValueProjectionType string

// List of QueryValueProjectionType.
const (
	QUERYVALUEPROJECTIONTYPE_QUERY_VALUE QueryValueProjectionType = "query_value"
)

var allowedQueryValueProjectionTypeEnumValues = []QueryValueProjectionType{
	QUERYVALUEPROJECTIONTYPE_QUERY_VALUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryValueProjectionType) GetAllowedValues() []QueryValueProjectionType {
	return allowedQueryValueProjectionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryValueProjectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryValueProjectionType(value)
	return nil
}

// NewQueryValueProjectionTypeFromValue returns a pointer to a valid QueryValueProjectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryValueProjectionTypeFromValue(v string) (*QueryValueProjectionType, error) {
	ev := QueryValueProjectionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryValueProjectionType: valid values are %v", v, allowedQueryValueProjectionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryValueProjectionType) IsValid() bool {
	for _, existing := range allowedQueryValueProjectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryValueProjectionType value.
func (v QueryValueProjectionType) Ptr() *QueryValueProjectionType {
	return &v
}
