// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type CaseStreamQueryDataSource string

// List of CaseStreamQueryDataSource.
const (
	CASESTREAMQUERYDATASOURCE_CASE_STREAM CaseStreamQueryDataSource = "case_stream"
)

var allowedCaseStreamQueryDataSourceEnumValues = []CaseStreamQueryDataSource{
	CASESTREAMQUERYDATASOURCE_CASE_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseStreamQueryDataSource) GetAllowedValues() []CaseStreamQueryDataSource {
	return allowedCaseStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseStreamQueryDataSource(value)
	return nil
}

// NewCaseStreamQueryDataSourceFromValue returns a pointer to a valid CaseStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseStreamQueryDataSourceFromValue(v string) (*CaseStreamQueryDataSource, error) {
	ev := CaseStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseStreamQueryDataSource: valid values are %v", v, allowedCaseStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedCaseStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseStreamQueryDataSource value.
func (v CaseStreamQueryDataSource) Ptr() *CaseStreamQueryDataSource {
	return &v
}
