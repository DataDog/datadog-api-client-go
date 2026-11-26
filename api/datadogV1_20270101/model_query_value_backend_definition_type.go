// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryValueBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type QueryValueBackendDefinitionType string

// List of QueryValueBackendDefinitionType.
const (
	QUERYVALUEBACKENDDEFINITIONTYPE_QUERY_VALUE QueryValueBackendDefinitionType = "query_value"
)

var allowedQueryValueBackendDefinitionTypeEnumValues = []QueryValueBackendDefinitionType{
	QUERYVALUEBACKENDDEFINITIONTYPE_QUERY_VALUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryValueBackendDefinitionType) GetAllowedValues() []QueryValueBackendDefinitionType {
	return allowedQueryValueBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryValueBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryValueBackendDefinitionType(value)
	return nil
}

// NewQueryValueBackendDefinitionTypeFromValue returns a pointer to a valid QueryValueBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryValueBackendDefinitionTypeFromValue(v string) (*QueryValueBackendDefinitionType, error) {
	ev := QueryValueBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryValueBackendDefinitionType: valid values are %v", v, allowedQueryValueBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryValueBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedQueryValueBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryValueBackendDefinitionType value.
func (v QueryValueBackendDefinitionType) Ptr() *QueryValueBackendDefinitionType {
	return &v
}
