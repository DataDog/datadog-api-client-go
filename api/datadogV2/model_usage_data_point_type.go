// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"
)

// UsageDataPointType The type of shape of the data.
type UsageDataPointType string

// List of UsageDataPointType.
const (
	USAGEDATAPOINTTYPE_USAGE_DATA_POINT UsageDataPointType = "usage_data_point"
)

var allowedUsageDataPointTypeEnumValues = []UsageDataPointType{
	USAGEDATAPOINTTYPE_USAGE_DATA_POINT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UsageDataPointType) GetAllowedValues() []UsageDataPointType {
	return allowedUsageDataPointTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UsageDataPointType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UsageDataPointType(value)
	return nil
}

// NewUsageDataPointTypeFromValue returns a pointer to a valid UsageDataPointType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUsageDataPointTypeFromValue(v string) (*UsageDataPointType, error) {
	ev := UsageDataPointType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UsageDataPointType: valid values are %v", v, allowedUsageDataPointTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UsageDataPointType) IsValid() bool {
	for _, existing := range allowedUsageDataPointTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UsageDataPointType value.
func (v UsageDataPointType) Ptr() *UsageDataPointType {
	return &v
}
