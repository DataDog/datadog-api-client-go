// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueSampleDataSource Product that reported the error. The value is `NONE` when the error has no known origin.
type IssueSampleDataSource string

// List of IssueSampleDataSource.
const (
	ISSUESAMPLEDATASOURCE_APM        IssueSampleDataSource = "APM"
	ISSUESAMPLEDATASOURCE_LOGS       IssueSampleDataSource = "LOGS"
	ISSUESAMPLEDATASOURCE_RUM        IssueSampleDataSource = "RUM"
	ISSUESAMPLEDATASOURCE_LLMOBS     IssueSampleDataSource = "LLMOBS"
	ISSUESAMPLEDATASOURCE_SYNTHETICS IssueSampleDataSource = "SYNTHETICS"
	ISSUESAMPLEDATASOURCE_NONE       IssueSampleDataSource = "NONE"
)

var allowedIssueSampleDataSourceEnumValues = []IssueSampleDataSource{
	ISSUESAMPLEDATASOURCE_APM,
	ISSUESAMPLEDATASOURCE_LOGS,
	ISSUESAMPLEDATASOURCE_RUM,
	ISSUESAMPLEDATASOURCE_LLMOBS,
	ISSUESAMPLEDATASOURCE_SYNTHETICS,
	ISSUESAMPLEDATASOURCE_NONE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueSampleDataSource) GetAllowedValues() []IssueSampleDataSource {
	return allowedIssueSampleDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueSampleDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueSampleDataSource(value)
	return nil
}

// NewIssueSampleDataSourceFromValue returns a pointer to a valid IssueSampleDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueSampleDataSourceFromValue(v string) (*IssueSampleDataSource, error) {
	ev := IssueSampleDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueSampleDataSource: valid values are %v", v, allowedIssueSampleDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueSampleDataSource) IsValid() bool {
	for _, existing := range allowedIssueSampleDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueSampleDataSource value.
func (v IssueSampleDataSource) Ptr() *IssueSampleDataSource {
	return &v
}
