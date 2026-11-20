// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ErrorsStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type ErrorsStreamQueryDataSource string

// List of ErrorsStreamQueryDataSource.
const (
	ERRORSSTREAMQUERYDATASOURCE_ERRORS_STREAM ErrorsStreamQueryDataSource = "errors_stream"
)

var allowedErrorsStreamQueryDataSourceEnumValues = []ErrorsStreamQueryDataSource{
	ERRORSSTREAMQUERYDATASOURCE_ERRORS_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ErrorsStreamQueryDataSource) GetAllowedValues() []ErrorsStreamQueryDataSource {
	return allowedErrorsStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ErrorsStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ErrorsStreamQueryDataSource(value)
	return nil
}

// NewErrorsStreamQueryDataSourceFromValue returns a pointer to a valid ErrorsStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewErrorsStreamQueryDataSourceFromValue(v string) (*ErrorsStreamQueryDataSource, error) {
	ev := ErrorsStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ErrorsStreamQueryDataSource: valid values are %v", v, allowedErrorsStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ErrorsStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedErrorsStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ErrorsStreamQueryDataSource value.
func (v ErrorsStreamQueryDataSource) Ptr() *ErrorsStreamQueryDataSource {
	return &v
}
