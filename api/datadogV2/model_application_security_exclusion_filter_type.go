// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityExclusionFilterType Type of the resource. The value should always be `exclusion_filter`.
type ApplicationSecurityExclusionFilterType string

// List of ApplicationSecurityExclusionFilterType.
const (
	APPLICATIONSECURITYEXCLUSIONFILTERTYPE_EXCLUSION_FILTER ApplicationSecurityExclusionFilterType = "exclusion_filter"
)

var allowedApplicationSecurityExclusionFilterTypeEnumValues = []ApplicationSecurityExclusionFilterType{
	APPLICATIONSECURITYEXCLUSIONFILTERTYPE_EXCLUSION_FILTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApplicationSecurityExclusionFilterType) GetAllowedValues() []ApplicationSecurityExclusionFilterType {
	return allowedApplicationSecurityExclusionFilterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApplicationSecurityExclusionFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApplicationSecurityExclusionFilterType(value)
	return nil
}

// NewApplicationSecurityExclusionFilterTypeFromValue returns a pointer to a valid ApplicationSecurityExclusionFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApplicationSecurityExclusionFilterTypeFromValue(v string) (*ApplicationSecurityExclusionFilterType, error) {
	ev := ApplicationSecurityExclusionFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApplicationSecurityExclusionFilterType: valid values are %v", v, allowedApplicationSecurityExclusionFilterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApplicationSecurityExclusionFilterType) IsValid() bool {
	for _, existing := range allowedApplicationSecurityExclusionFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApplicationSecurityExclusionFilterType value.
func (v ApplicationSecurityExclusionFilterType) Ptr() *ApplicationSecurityExclusionFilterType {
	return &v
}
