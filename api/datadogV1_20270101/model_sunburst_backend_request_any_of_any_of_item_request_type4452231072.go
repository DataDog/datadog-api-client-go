// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072 Test-only generated schema for the dashboard versioning proof of concept.
type SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072 string

// List of SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072.
const (
	SUNBURSTBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4452231072_LOCAL_DATASET SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072 = "local_dataset"
)

var allowedSunburstBackendRequestAnyOfAnyOfItemRequestType4452231072EnumValues = []SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072{
	SUNBURSTBACKENDREQUESTANYOFANYOFITEMREQUESTTYPE4452231072_LOCAL_DATASET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072) GetAllowedValues() []SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072 {
	return allowedSunburstBackendRequestAnyOfAnyOfItemRequestType4452231072EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072(value)
	return nil
}

// NewSunburstBackendRequestAnyOfAnyOfItemRequestType4452231072FromValue returns a pointer to a valid SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSunburstBackendRequestAnyOfAnyOfItemRequestType4452231072FromValue(v string) (*SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072, error) {
	ev := SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072: valid values are %v", v, allowedSunburstBackendRequestAnyOfAnyOfItemRequestType4452231072EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072) IsValid() bool {
	for _, existing := range allowedSunburstBackendRequestAnyOfAnyOfItemRequestType4452231072EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072 value.
func (v SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072) Ptr() *SunburstBackendRequestAnyOfAnyOfItemRequestType4452231072 {
	return &v
}
