// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageDataType Indicates that the resource is of type `pages`.
type OnCallPageDataType string

// List of OnCallPageDataType.
const (
	ONCALLPAGEDATATYPE_PAGES OnCallPageDataType = "pages"
)

var allowedOnCallPageDataTypeEnumValues = []OnCallPageDataType{
	ONCALLPAGEDATATYPE_PAGES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnCallPageDataType) GetAllowedValues() []OnCallPageDataType {
	return allowedOnCallPageDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnCallPageDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnCallPageDataType(value)
	return nil
}

// NewOnCallPageDataTypeFromValue returns a pointer to a valid OnCallPageDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnCallPageDataTypeFromValue(v string) (*OnCallPageDataType, error) {
	ev := OnCallPageDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnCallPageDataType: valid values are %v", v, allowedOnCallPageDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnCallPageDataType) IsValid() bool {
	for _, existing := range allowedOnCallPageDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnCallPageDataType value.
func (v OnCallPageDataType) Ptr() *OnCallPageDataType {
	return &v
}
