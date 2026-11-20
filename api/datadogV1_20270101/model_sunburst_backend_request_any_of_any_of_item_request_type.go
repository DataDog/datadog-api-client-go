// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstBackendRequestAnyOfAnyOfItemRequestType Test-only generated schema for the dashboard versioning proof of concept.
type SunburstBackendRequestAnyOfAnyOfItemRequestType string

// List of SunburstBackendRequestAnyOfAnyOfItemRequestType.
const (
	SUNBURSTBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE_ADVANCED_QUERY SunburstBackendRequestAnyOfAnyOfItemRequestType = "advanced_query"
)

var allowedSunburstBackendRequestAnyOfAnyOfItemRequestTypeEnumValues = []SunburstBackendRequestAnyOfAnyOfItemRequestType{
	SUNBURSTBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE_ADVANCED_QUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SunburstBackendRequestAnyOfAnyOfItemRequestType) GetAllowedValues() []SunburstBackendRequestAnyOfAnyOfItemRequestType {
	return allowedSunburstBackendRequestAnyOfAnyOfItemRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SunburstBackendRequestAnyOfAnyOfItemRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SunburstBackendRequestAnyOfAnyOfItemRequestType(value)
	return nil
}

// NewSunburstBackendRequestAnyOfAnyOfItemRequestTypeFromValue returns a pointer to a valid SunburstBackendRequestAnyOfAnyOfItemRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSunburstBackendRequestAnyOfAnyOfItemRequestTypeFromValue(v string) (*SunburstBackendRequestAnyOfAnyOfItemRequestType, error) {
	ev := SunburstBackendRequestAnyOfAnyOfItemRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SunburstBackendRequestAnyOfAnyOfItemRequestType: valid values are %v", v, allowedSunburstBackendRequestAnyOfAnyOfItemRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SunburstBackendRequestAnyOfAnyOfItemRequestType) IsValid() bool {
	for _, existing := range allowedSunburstBackendRequestAnyOfAnyOfItemRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SunburstBackendRequestAnyOfAnyOfItemRequestType value.
func (v SunburstBackendRequestAnyOfAnyOfItemRequestType) Ptr() *SunburstBackendRequestAnyOfAnyOfItemRequestType {
	return &v
}
