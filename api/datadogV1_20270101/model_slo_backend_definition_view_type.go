// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SloBackendDefinitionViewType Test-only generated schema for the dashboard versioning proof of concept.
type SloBackendDefinitionViewType string

// List of SloBackendDefinitionViewType.
const (
	SLOBACKENDDEFINITIONVIEWTYPE_DETAIL SloBackendDefinitionViewType = "detail"
)

var allowedSloBackendDefinitionViewTypeEnumValues = []SloBackendDefinitionViewType{
	SLOBACKENDDEFINITIONVIEWTYPE_DETAIL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SloBackendDefinitionViewType) GetAllowedValues() []SloBackendDefinitionViewType {
	return allowedSloBackendDefinitionViewTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SloBackendDefinitionViewType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SloBackendDefinitionViewType(value)
	return nil
}

// NewSloBackendDefinitionViewTypeFromValue returns a pointer to a valid SloBackendDefinitionViewType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSloBackendDefinitionViewTypeFromValue(v string) (*SloBackendDefinitionViewType, error) {
	ev := SloBackendDefinitionViewType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SloBackendDefinitionViewType: valid values are %v", v, allowedSloBackendDefinitionViewTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SloBackendDefinitionViewType) IsValid() bool {
	for _, existing := range allowedSloBackendDefinitionViewTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SloBackendDefinitionViewType value.
func (v SloBackendDefinitionViewType) Ptr() *SloBackendDefinitionViewType {
	return &v
}
