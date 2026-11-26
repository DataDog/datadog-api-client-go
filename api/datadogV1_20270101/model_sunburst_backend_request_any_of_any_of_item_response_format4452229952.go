// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952 Test-only generated schema for the dashboard versioning proof of concept.
type SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952 string

// List of SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952.
const (
	SUNBURSTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4452229952_SCALAR SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952 = "scalar"
)

var allowedSunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952EnumValues = []SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952{
	SUNBURSTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4452229952_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952) GetAllowedValues() []SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952 {
	return allowedSunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952(value)
	return nil
}

// NewSunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952FromValue returns a pointer to a valid SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952FromValue(v string) (*SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952, error) {
	ev := SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952: valid values are %v", v, allowedSunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952) IsValid() bool {
	for _, existing := range allowedSunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952 value.
func (v SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952) Ptr() *SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452229952 {
	return &v
}
