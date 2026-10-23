// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageScheduledActionAttributesType The kind of action that is scheduled.
type OnCallPageScheduledActionAttributesType string

// List of OnCallPageScheduledActionAttributesType.
const (
	ONCALLPAGESCHEDULEDACTIONATTRIBUTESTYPE_ESCALATE  OnCallPageScheduledActionAttributesType = "escalate"
	ONCALLPAGESCHEDULEDACTIONATTRIBUTESTYPE_RESOLVE   OnCallPageScheduledActionAttributesType = "resolve"
	ONCALLPAGESCHEDULEDACTIONATTRIBUTESTYPE_RELEASE   OnCallPageScheduledActionAttributesType = "release"
	ONCALLPAGESCHEDULEDACTIONATTRIBUTESTYPE_RETRIGGER OnCallPageScheduledActionAttributesType = "retrigger"
)

var allowedOnCallPageScheduledActionAttributesTypeEnumValues = []OnCallPageScheduledActionAttributesType{
	ONCALLPAGESCHEDULEDACTIONATTRIBUTESTYPE_ESCALATE,
	ONCALLPAGESCHEDULEDACTIONATTRIBUTESTYPE_RESOLVE,
	ONCALLPAGESCHEDULEDACTIONATTRIBUTESTYPE_RELEASE,
	ONCALLPAGESCHEDULEDACTIONATTRIBUTESTYPE_RETRIGGER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnCallPageScheduledActionAttributesType) GetAllowedValues() []OnCallPageScheduledActionAttributesType {
	return allowedOnCallPageScheduledActionAttributesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnCallPageScheduledActionAttributesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnCallPageScheduledActionAttributesType(value)
	return nil
}

// NewOnCallPageScheduledActionAttributesTypeFromValue returns a pointer to a valid OnCallPageScheduledActionAttributesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnCallPageScheduledActionAttributesTypeFromValue(v string) (*OnCallPageScheduledActionAttributesType, error) {
	ev := OnCallPageScheduledActionAttributesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnCallPageScheduledActionAttributesType: valid values are %v", v, allowedOnCallPageScheduledActionAttributesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnCallPageScheduledActionAttributesType) IsValid() bool {
	for _, existing := range allowedOnCallPageScheduledActionAttributesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnCallPageScheduledActionAttributesType value.
func (v OnCallPageScheduledActionAttributesType) Ptr() *OnCallPageScheduledActionAttributesType {
	return &v
}
