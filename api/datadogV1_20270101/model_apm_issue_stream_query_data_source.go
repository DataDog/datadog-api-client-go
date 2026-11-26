// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// APMIssueStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type APMIssueStreamQueryDataSource string

// List of APMIssueStreamQueryDataSource.
const (
	APMISSUESTREAMQUERYDATASOURCE_APM_ISSUE_STREAM APMIssueStreamQueryDataSource = "apm_issue_stream"
)

var allowedAPMIssueStreamQueryDataSourceEnumValues = []APMIssueStreamQueryDataSource{
	APMISSUESTREAMQUERYDATASOURCE_APM_ISSUE_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *APMIssueStreamQueryDataSource) GetAllowedValues() []APMIssueStreamQueryDataSource {
	return allowedAPMIssueStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *APMIssueStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = APMIssueStreamQueryDataSource(value)
	return nil
}

// NewAPMIssueStreamQueryDataSourceFromValue returns a pointer to a valid APMIssueStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAPMIssueStreamQueryDataSourceFromValue(v string) (*APMIssueStreamQueryDataSource, error) {
	ev := APMIssueStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for APMIssueStreamQueryDataSource: valid values are %v", v, allowedAPMIssueStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v APMIssueStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedAPMIssueStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to APMIssueStreamQueryDataSource value.
func (v APMIssueStreamQueryDataSource) Ptr() *APMIssueStreamQueryDataSource {
	return &v
}
