// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ErrorsStreamRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type ErrorsStreamRequestResponseFormat string

// List of ErrorsStreamRequestResponseFormat.
const (
	ERRORSSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST ErrorsStreamRequestResponseFormat = "event_list"
)

var allowedErrorsStreamRequestResponseFormatEnumValues = []ErrorsStreamRequestResponseFormat{
	ERRORSSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ErrorsStreamRequestResponseFormat) GetAllowedValues() []ErrorsStreamRequestResponseFormat {
	return allowedErrorsStreamRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ErrorsStreamRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ErrorsStreamRequestResponseFormat(value)
	return nil
}

// NewErrorsStreamRequestResponseFormatFromValue returns a pointer to a valid ErrorsStreamRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewErrorsStreamRequestResponseFormatFromValue(v string) (*ErrorsStreamRequestResponseFormat, error) {
	ev := ErrorsStreamRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ErrorsStreamRequestResponseFormat: valid values are %v", v, allowedErrorsStreamRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ErrorsStreamRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedErrorsStreamRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ErrorsStreamRequestResponseFormat value.
func (v ErrorsStreamRequestResponseFormat) Ptr() *ErrorsStreamRequestResponseFormat {
	return &v
}
