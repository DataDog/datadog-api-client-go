// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormVersionListType The resource type for a list of form versions.
type FormVersionListType string

// List of FormVersionListType.
const (
	FORMVERSIONLISTTYPE_FORM_VERSION_LISTS FormVersionListType = "form_version_lists"
)

var allowedFormVersionListTypeEnumValues = []FormVersionListType{
	FORMVERSIONLISTTYPE_FORM_VERSION_LISTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormVersionListType) GetAllowedValues() []FormVersionListType {
	return allowedFormVersionListTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormVersionListType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormVersionListType(value)
	return nil
}

// NewFormVersionListTypeFromValue returns a pointer to a valid FormVersionListType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormVersionListTypeFromValue(v string) (*FormVersionListType, error) {
	ev := FormVersionListType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormVersionListType: valid values are %v", v, allowedFormVersionListTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormVersionListType) IsValid() bool {
	for _, existing := range allowedFormVersionListTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormVersionListType value.
func (v FormVersionListType) Ptr() *FormVersionListType {
	return &v
}
