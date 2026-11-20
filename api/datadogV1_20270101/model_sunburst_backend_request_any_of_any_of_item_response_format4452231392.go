// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392 Test-only generated schema for the dashboard versioning proof of concept.
type SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392 string

// List of SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392.
const (
	SUNBURSTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4452231392_TABULAR SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392 = "tabular"
)

var allowedSunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392EnumValues = []SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392{
	SUNBURSTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4452231392_TABULAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392) GetAllowedValues() []SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392 {
	return allowedSunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392(value)
	return nil
}

// NewSunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392FromValue returns a pointer to a valid SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392FromValue(v string) (*SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392, error) {
	ev := SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392: valid values are %v", v, allowedSunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392) IsValid() bool {
	for _, existing := range allowedSunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392 value.
func (v SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392) Ptr() *SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452231392 {
	return &v
}
