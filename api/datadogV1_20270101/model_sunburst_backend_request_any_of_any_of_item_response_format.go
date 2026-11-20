// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstBackendRequestAnyOfAnyOfItemResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type SunburstBackendRequestAnyOfAnyOfItemResponseFormat string

// List of SunburstBackendRequestAnyOfAnyOfItemResponseFormat.
const (
	SUNBURSTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT_TABULAR SunburstBackendRequestAnyOfAnyOfItemResponseFormat = "tabular"
)

var allowedSunburstBackendRequestAnyOfAnyOfItemResponseFormatEnumValues = []SunburstBackendRequestAnyOfAnyOfItemResponseFormat{
	SUNBURSTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT_TABULAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SunburstBackendRequestAnyOfAnyOfItemResponseFormat) GetAllowedValues() []SunburstBackendRequestAnyOfAnyOfItemResponseFormat {
	return allowedSunburstBackendRequestAnyOfAnyOfItemResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SunburstBackendRequestAnyOfAnyOfItemResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SunburstBackendRequestAnyOfAnyOfItemResponseFormat(value)
	return nil
}

// NewSunburstBackendRequestAnyOfAnyOfItemResponseFormatFromValue returns a pointer to a valid SunburstBackendRequestAnyOfAnyOfItemResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSunburstBackendRequestAnyOfAnyOfItemResponseFormatFromValue(v string) (*SunburstBackendRequestAnyOfAnyOfItemResponseFormat, error) {
	ev := SunburstBackendRequestAnyOfAnyOfItemResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SunburstBackendRequestAnyOfAnyOfItemResponseFormat: valid values are %v", v, allowedSunburstBackendRequestAnyOfAnyOfItemResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SunburstBackendRequestAnyOfAnyOfItemResponseFormat) IsValid() bool {
	for _, existing := range allowedSunburstBackendRequestAnyOfAnyOfItemResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SunburstBackendRequestAnyOfAnyOfItemResponseFormat value.
func (v SunburstBackendRequestAnyOfAnyOfItemResponseFormat) Ptr() *SunburstBackendRequestAnyOfAnyOfItemResponseFormat {
	return &v
}
