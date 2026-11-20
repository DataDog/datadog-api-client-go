// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type SunburstBackendDefinitionType string

// List of SunburstBackendDefinitionType.
const (
	SUNBURSTBACKENDDEFINITIONTYPE_SUNBURST SunburstBackendDefinitionType = "sunburst"
)

var allowedSunburstBackendDefinitionTypeEnumValues = []SunburstBackendDefinitionType{
	SUNBURSTBACKENDDEFINITIONTYPE_SUNBURST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SunburstBackendDefinitionType) GetAllowedValues() []SunburstBackendDefinitionType {
	return allowedSunburstBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SunburstBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SunburstBackendDefinitionType(value)
	return nil
}

// NewSunburstBackendDefinitionTypeFromValue returns a pointer to a valid SunburstBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSunburstBackendDefinitionTypeFromValue(v string) (*SunburstBackendDefinitionType, error) {
	ev := SunburstBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SunburstBackendDefinitionType: valid values are %v", v, allowedSunburstBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SunburstBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedSunburstBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SunburstBackendDefinitionType value.
func (v SunburstBackendDefinitionType) Ptr() *SunburstBackendDefinitionType {
	return &v
}
