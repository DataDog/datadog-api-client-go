// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UnifiedIssueStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type UnifiedIssueStreamQueryDataSource string

// List of UnifiedIssueStreamQueryDataSource.
const (
	UNIFIEDISSUESTREAMQUERYDATASOURCE_ISSUE_STREAM UnifiedIssueStreamQueryDataSource = "issue_stream"
)

var allowedUnifiedIssueStreamQueryDataSourceEnumValues = []UnifiedIssueStreamQueryDataSource{
	UNIFIEDISSUESTREAMQUERYDATASOURCE_ISSUE_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UnifiedIssueStreamQueryDataSource) GetAllowedValues() []UnifiedIssueStreamQueryDataSource {
	return allowedUnifiedIssueStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UnifiedIssueStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UnifiedIssueStreamQueryDataSource(value)
	return nil
}

// NewUnifiedIssueStreamQueryDataSourceFromValue returns a pointer to a valid UnifiedIssueStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUnifiedIssueStreamQueryDataSourceFromValue(v string) (*UnifiedIssueStreamQueryDataSource, error) {
	ev := UnifiedIssueStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UnifiedIssueStreamQueryDataSource: valid values are %v", v, allowedUnifiedIssueStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UnifiedIssueStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedUnifiedIssueStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UnifiedIssueStreamQueryDataSource value.
func (v UnifiedIssueStreamQueryDataSource) Ptr() *UnifiedIssueStreamQueryDataSource {
	return &v
}
