// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type RumStreamQueryDataSource string

// List of RumStreamQueryDataSource.
const (
	RUMSTREAMQUERYDATASOURCE_RUM_STREAM RumStreamQueryDataSource = "rum_stream"
)

var allowedRumStreamQueryDataSourceEnumValues = []RumStreamQueryDataSource{
	RUMSTREAMQUERYDATASOURCE_RUM_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumStreamQueryDataSource) GetAllowedValues() []RumStreamQueryDataSource {
	return allowedRumStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumStreamQueryDataSource(value)
	return nil
}

// NewRumStreamQueryDataSourceFromValue returns a pointer to a valid RumStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumStreamQueryDataSourceFromValue(v string) (*RumStreamQueryDataSource, error) {
	ev := RumStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumStreamQueryDataSource: valid values are %v", v, allowedRumStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedRumStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumStreamQueryDataSource value.
func (v RumStreamQueryDataSource) Ptr() *RumStreamQueryDataSource {
	return &v
}
