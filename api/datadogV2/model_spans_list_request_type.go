// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"
)

// SpansListRequestType The type of resource. The value should always be search_request.
type SpansListRequestType string

// List of SpansListRequestType.
const (
	SPANSLISTREQUESTTYPE_SEARCH_REQUEST SpansListRequestType = "search_request"
)

var allowedSpansListRequestTypeEnumValues = []SpansListRequestType{
	SPANSLISTREQUESTTYPE_SEARCH_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SpansListRequestType) GetAllowedValues() []SpansListRequestType {
	return allowedSpansListRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SpansListRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SpansListRequestType(value)
	return nil
}

// NewSpansListRequestTypeFromValue returns a pointer to a valid SpansListRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSpansListRequestTypeFromValue(v string) (*SpansListRequestType, error) {
	ev := SpansListRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SpansListRequestType: valid values are %v", v, allowedSpansListRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SpansListRequestType) IsValid() bool {
	for _, existing := range allowedSpansListRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SpansListRequestType value.
func (v SpansListRequestType) Ptr() *SpansListRequestType {
	return &v
}

// NullableSpansListRequestType handles when a null is used for SpansListRequestType.
type NullableSpansListRequestType struct {
	value *SpansListRequestType
	isSet bool
}

// Get returns the associated value.
func (v NullableSpansListRequestType) Get() *SpansListRequestType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableSpansListRequestType) Set(val *SpansListRequestType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableSpansListRequestType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableSpansListRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSpansListRequestType initializes the struct as if Set has been called.
func NewNullableSpansListRequestType(val *SpansListRequestType) *NullableSpansListRequestType {
	return &NullableSpansListRequestType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableSpansListRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableSpansListRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
