// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListStreamBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type ListStreamBackendDefinitionType string

// List of ListStreamBackendDefinitionType.
const (
	LISTSTREAMBACKENDDEFINITIONTYPE_LIST_STREAM ListStreamBackendDefinitionType = "list_stream"
)

var allowedListStreamBackendDefinitionTypeEnumValues = []ListStreamBackendDefinitionType{
	LISTSTREAMBACKENDDEFINITIONTYPE_LIST_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ListStreamBackendDefinitionType) GetAllowedValues() []ListStreamBackendDefinitionType {
	return allowedListStreamBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListStreamBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListStreamBackendDefinitionType(value)
	return nil
}

// NewListStreamBackendDefinitionTypeFromValue returns a pointer to a valid ListStreamBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListStreamBackendDefinitionTypeFromValue(v string) (*ListStreamBackendDefinitionType, error) {
	ev := ListStreamBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListStreamBackendDefinitionType: valid values are %v", v, allowedListStreamBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListStreamBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedListStreamBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListStreamBackendDefinitionType value.
func (v ListStreamBackendDefinitionType) Ptr() *ListStreamBackendDefinitionType {
	return &v
}
