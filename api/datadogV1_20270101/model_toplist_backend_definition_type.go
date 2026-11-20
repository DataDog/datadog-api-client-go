// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendDefinitionType string

// List of ToplistBackendDefinitionType.
const (
	TOPLISTBACKENDDEFINITIONTYPE_TOPLIST ToplistBackendDefinitionType = "toplist"
)

var allowedToplistBackendDefinitionTypeEnumValues = []ToplistBackendDefinitionType{
	TOPLISTBACKENDDEFINITIONTYPE_TOPLIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ToplistBackendDefinitionType) GetAllowedValues() []ToplistBackendDefinitionType {
	return allowedToplistBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ToplistBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ToplistBackendDefinitionType(value)
	return nil
}

// NewToplistBackendDefinitionTypeFromValue returns a pointer to a valid ToplistBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewToplistBackendDefinitionTypeFromValue(v string) (*ToplistBackendDefinitionType, error) {
	ev := ToplistBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ToplistBackendDefinitionType: valid values are %v", v, allowedToplistBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ToplistBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedToplistBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ToplistBackendDefinitionType value.
func (v ToplistBackendDefinitionType) Ptr() *ToplistBackendDefinitionType {
	return &v
}
