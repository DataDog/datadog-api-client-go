// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeEventCreateRequestType Entity type.
type ChangeEventCreateRequestType string

// List of ChangeEventCreateRequestType.
const (
	CHANGEEVENTCREATEREQUESTTYPE_EVENT ChangeEventCreateRequestType = "event"
)

var allowedChangeEventCreateRequestTypeEnumValues = []ChangeEventCreateRequestType{
	CHANGEEVENTCREATEREQUESTTYPE_EVENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChangeEventCreateRequestType) GetAllowedValues() []ChangeEventCreateRequestType {
	return allowedChangeEventCreateRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChangeEventCreateRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeEventCreateRequestType(value)
	return nil
}

// NewChangeEventCreateRequestTypeFromValue returns a pointer to a valid ChangeEventCreateRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChangeEventCreateRequestTypeFromValue(v string) (*ChangeEventCreateRequestType, error) {
	ev := ChangeEventCreateRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChangeEventCreateRequestType: valid values are %v", v, allowedChangeEventCreateRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChangeEventCreateRequestType) IsValid() bool {
	for _, existing := range allowedChangeEventCreateRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeEventCreateRequestType value.
func (v ChangeEventCreateRequestType) Ptr() *ChangeEventCreateRequestType {
	return &v
}
