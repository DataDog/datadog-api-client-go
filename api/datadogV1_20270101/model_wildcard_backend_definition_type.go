// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WildcardBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type WildcardBackendDefinitionType string

// List of WildcardBackendDefinitionType.
const (
	WILDCARDBACKENDDEFINITIONTYPE_WILDCARD WildcardBackendDefinitionType = "wildcard"
)

var allowedWildcardBackendDefinitionTypeEnumValues = []WildcardBackendDefinitionType{
	WILDCARDBACKENDDEFINITIONTYPE_WILDCARD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WildcardBackendDefinitionType) GetAllowedValues() []WildcardBackendDefinitionType {
	return allowedWildcardBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WildcardBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WildcardBackendDefinitionType(value)
	return nil
}

// NewWildcardBackendDefinitionTypeFromValue returns a pointer to a valid WildcardBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWildcardBackendDefinitionTypeFromValue(v string) (*WildcardBackendDefinitionType, error) {
	ev := WildcardBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WildcardBackendDefinitionType: valid values are %v", v, allowedWildcardBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WildcardBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedWildcardBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WildcardBackendDefinitionType value.
func (v WildcardBackendDefinitionType) Ptr() *WildcardBackendDefinitionType {
	return &v
}
