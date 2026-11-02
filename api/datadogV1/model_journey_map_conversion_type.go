// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JourneyMapConversionType Whether to show the paths that convert to the end step or drop off before it.
type JourneyMapConversionType string

// List of JourneyMapConversionType.
const (
	JOURNEYMAPCONVERSIONTYPE_CONVERSION JourneyMapConversionType = "conversion"
	JOURNEYMAPCONVERSIONTYPE_DROP_OFF   JourneyMapConversionType = "drop-off"
)

var allowedJourneyMapConversionTypeEnumValues = []JourneyMapConversionType{
	JOURNEYMAPCONVERSIONTYPE_CONVERSION,
	JOURNEYMAPCONVERSIONTYPE_DROP_OFF,
}

// GetAllowedValues reeturns the list of possible values.
func (v *JourneyMapConversionType) GetAllowedValues() []JourneyMapConversionType {
	return allowedJourneyMapConversionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *JourneyMapConversionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = JourneyMapConversionType(value)
	return nil
}

// NewJourneyMapConversionTypeFromValue returns a pointer to a valid JourneyMapConversionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewJourneyMapConversionTypeFromValue(v string) (*JourneyMapConversionType, error) {
	ev := JourneyMapConversionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for JourneyMapConversionType: valid values are %v", v, allowedJourneyMapConversionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v JourneyMapConversionType) IsValid() bool {
	for _, existing := range allowedJourneyMapConversionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JourneyMapConversionType value.
func (v JourneyMapConversionType) Ptr() *JourneyMapConversionType {
	return &v
}
