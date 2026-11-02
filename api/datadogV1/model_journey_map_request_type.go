// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JourneyMapRequestType Request type for the Journey Paths widget.
type JourneyMapRequestType string

// List of JourneyMapRequestType.
const (
	JOURNEYMAPREQUESTTYPE_JOURNEY_MAP JourneyMapRequestType = "journey_map"
)

var allowedJourneyMapRequestTypeEnumValues = []JourneyMapRequestType{
	JOURNEYMAPREQUESTTYPE_JOURNEY_MAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *JourneyMapRequestType) GetAllowedValues() []JourneyMapRequestType {
	return allowedJourneyMapRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *JourneyMapRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = JourneyMapRequestType(value)
	return nil
}

// NewJourneyMapRequestTypeFromValue returns a pointer to a valid JourneyMapRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewJourneyMapRequestTypeFromValue(v string) (*JourneyMapRequestType, error) {
	ev := JourneyMapRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for JourneyMapRequestType: valid values are %v", v, allowedJourneyMapRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v JourneyMapRequestType) IsValid() bool {
	for _, existing := range allowedJourneyMapRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JourneyMapRequestType value.
func (v JourneyMapRequestType) Ptr() *JourneyMapRequestType {
	return &v
}
