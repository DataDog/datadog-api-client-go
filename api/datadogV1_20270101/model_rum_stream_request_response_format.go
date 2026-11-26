// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumStreamRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type RumStreamRequestResponseFormat string

// List of RumStreamRequestResponseFormat.
const (
	RUMSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST RumStreamRequestResponseFormat = "event_list"
)

var allowedRumStreamRequestResponseFormatEnumValues = []RumStreamRequestResponseFormat{
	RUMSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumStreamRequestResponseFormat) GetAllowedValues() []RumStreamRequestResponseFormat {
	return allowedRumStreamRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumStreamRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumStreamRequestResponseFormat(value)
	return nil
}

// NewRumStreamRequestResponseFormatFromValue returns a pointer to a valid RumStreamRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumStreamRequestResponseFormatFromValue(v string) (*RumStreamRequestResponseFormat, error) {
	ev := RumStreamRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumStreamRequestResponseFormat: valid values are %v", v, allowedRumStreamRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumStreamRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedRumStreamRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumStreamRequestResponseFormat value.
func (v RumStreamRequestResponseFormat) Ptr() *RumStreamRequestResponseFormat {
	return &v
}
