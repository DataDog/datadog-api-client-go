// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendRequestAnyOfAnyOfItemResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendRequestAnyOfAnyOfItemResponseFormat string

// List of ToplistBackendRequestAnyOfAnyOfItemResponseFormat.
const (
	TOPLISTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT_TABULAR ToplistBackendRequestAnyOfAnyOfItemResponseFormat = "tabular"
)

var allowedToplistBackendRequestAnyOfAnyOfItemResponseFormatEnumValues = []ToplistBackendRequestAnyOfAnyOfItemResponseFormat{
	TOPLISTBACKENDREQUESTANYOFANYOFITEMRESPONSEFORMAT_TABULAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ToplistBackendRequestAnyOfAnyOfItemResponseFormat) GetAllowedValues() []ToplistBackendRequestAnyOfAnyOfItemResponseFormat {
	return allowedToplistBackendRequestAnyOfAnyOfItemResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ToplistBackendRequestAnyOfAnyOfItemResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ToplistBackendRequestAnyOfAnyOfItemResponseFormat(value)
	return nil
}

// NewToplistBackendRequestAnyOfAnyOfItemResponseFormatFromValue returns a pointer to a valid ToplistBackendRequestAnyOfAnyOfItemResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewToplistBackendRequestAnyOfAnyOfItemResponseFormatFromValue(v string) (*ToplistBackendRequestAnyOfAnyOfItemResponseFormat, error) {
	ev := ToplistBackendRequestAnyOfAnyOfItemResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ToplistBackendRequestAnyOfAnyOfItemResponseFormat: valid values are %v", v, allowedToplistBackendRequestAnyOfAnyOfItemResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ToplistBackendRequestAnyOfAnyOfItemResponseFormat) IsValid() bool {
	for _, existing := range allowedToplistBackendRequestAnyOfAnyOfItemResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ToplistBackendRequestAnyOfAnyOfItemResponseFormat value.
func (v ToplistBackendRequestAnyOfAnyOfItemResponseFormat) Ptr() *ToplistBackendRequestAnyOfAnyOfItemResponseFormat {
	return &v
}
