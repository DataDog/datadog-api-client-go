// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableBackendDefinitionType string

// List of QueryTableBackendDefinitionType.
const (
	QUERYTABLEBACKENDDEFINITIONTYPE_QUERY_TABLE QueryTableBackendDefinitionType = "query_table"
)

var allowedQueryTableBackendDefinitionTypeEnumValues = []QueryTableBackendDefinitionType{
	QUERYTABLEBACKENDDEFINITIONTYPE_QUERY_TABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryTableBackendDefinitionType) GetAllowedValues() []QueryTableBackendDefinitionType {
	return allowedQueryTableBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryTableBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryTableBackendDefinitionType(value)
	return nil
}

// NewQueryTableBackendDefinitionTypeFromValue returns a pointer to a valid QueryTableBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryTableBackendDefinitionTypeFromValue(v string) (*QueryTableBackendDefinitionType, error) {
	ev := QueryTableBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryTableBackendDefinitionType: valid values are %v", v, allowedQueryTableBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryTableBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedQueryTableBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryTableBackendDefinitionType value.
func (v QueryTableBackendDefinitionType) Ptr() *QueryTableBackendDefinitionType {
	return &v
}
