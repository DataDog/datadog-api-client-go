// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// STIXPatternType The supported STIX pattern language.
type STIXPatternType string

// List of STIXPatternType.
const (
	STIXPATTERNTYPE_STIX STIXPatternType = "stix"
)

var allowedSTIXPatternTypeEnumValues = []STIXPatternType{
	STIXPATTERNTYPE_STIX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *STIXPatternType) GetAllowedValues() []STIXPatternType {
	return allowedSTIXPatternTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *STIXPatternType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = STIXPatternType(value)
	return nil
}

// NewSTIXPatternTypeFromValue returns a pointer to a valid STIXPatternType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSTIXPatternTypeFromValue(v string) (*STIXPatternType, error) {
	ev := STIXPatternType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for STIXPatternType: valid values are %v", v, allowedSTIXPatternTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v STIXPatternType) IsValid() bool {
	for _, existing := range allowedSTIXPatternTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to STIXPatternType value.
func (v STIXPatternType) Ptr() *STIXPatternType {
	return &v
}
