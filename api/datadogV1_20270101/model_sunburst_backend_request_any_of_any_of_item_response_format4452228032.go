// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032 Test-only generated schema for the dashboard versioning proof of concept.
type SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032 string

// List of SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032.
const (
	SUNBURSTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4452228032_SCALAR SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032 = "scalar"
)

var allowedSunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032EnumValues = []SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032{
	SUNBURSTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT4452228032_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032) GetAllowedValues() []SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032 {
	return allowedSunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032(value)
	return nil
}

// NewSunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032FromValue returns a pointer to a valid SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032FromValue(v string) (*SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032, error) {
	ev := SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032: valid values are %v", v, allowedSunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032) IsValid() bool {
	for _, existing := range allowedSunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032 value.
func (v SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032) Ptr() *SunburstBackendRequestAnyOfAnyOfItemResponseFormat4452228032 {
	return &v
}
