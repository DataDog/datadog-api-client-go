// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstBackendRequestAnyOfAnyOfItemRequestType4452229632 Test-only generated schema for the dashboard versioning proof of concept.
type SunburstBackendRequestAnyOfAnyOfItemRequestType4452229632 string

// List of SunburstBackendRequestAnyOfAnyOfItemRequestType4452229632.
const (
	SUNBURSTBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4452229632_DDSQL SunburstBackendRequestAnyOfAnyOfItemRequestType4452229632 = "ddsql"
)

var allowedSunburstBackendRequestAnyOfAnyOfItemRequestType4452229632EnumValues = []SunburstBackendRequestAnyOfAnyOfItemRequestType4452229632{
	SUNBURSTBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4452229632_DDSQL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SunburstBackendRequestAnyOfAnyOfItemRequestType4452229632) GetAllowedValues() []SunburstBackendRequestAnyOfAnyOfItemRequestType4452229632 {
	return allowedSunburstBackendRequestAnyOfAnyOfItemRequestType4452229632EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SunburstBackendRequestAnyOfAnyOfItemRequestType4452229632) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SunburstBackendRequestAnyOfAnyOfItemRequestType4452229632(value)
	return nil
}

// NewSunburstBackendRequestAnyOfAnyOfItemRequestType4452229632FromValue returns a pointer to a valid SunburstBackendRequestAnyOfAnyOfItemRequestType4452229632
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSunburstBackendRequestAnyOfAnyOfItemRequestType4452229632FromValue(v string) (*SunburstBackendRequestAnyOfAnyOfItemRequestType4452229632, error) {
	ev := SunburstBackendRequestAnyOfAnyOfItemRequestType4452229632(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SunburstBackendRequestAnyOfAnyOfItemRequestType4452229632: valid values are %v", v, allowedSunburstBackendRequestAnyOfAnyOfItemRequestType4452229632EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SunburstBackendRequestAnyOfAnyOfItemRequestType4452229632) IsValid() bool {
	for _, existing := range allowedSunburstBackendRequestAnyOfAnyOfItemRequestType4452229632EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SunburstBackendRequestAnyOfAnyOfItemRequestType4452229632 value.
func (v SunburstBackendRequestAnyOfAnyOfItemRequestType4452229632) Ptr() *SunburstBackendRequestAnyOfAnyOfItemRequestType4452229632 {
	return &v
}
