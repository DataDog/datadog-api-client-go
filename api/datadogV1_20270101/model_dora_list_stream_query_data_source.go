// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DoraListStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type DoraListStreamQueryDataSource string

// List of DoraListStreamQueryDataSource.
const (
	DORALISTSTREAMQUERYDATASOURCE_DORA_LIST_STREAM DoraListStreamQueryDataSource = "dora_list_stream"
)

var allowedDoraListStreamQueryDataSourceEnumValues = []DoraListStreamQueryDataSource{
	DORALISTSTREAMQUERYDATASOURCE_DORA_LIST_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DoraListStreamQueryDataSource) GetAllowedValues() []DoraListStreamQueryDataSource {
	return allowedDoraListStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DoraListStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DoraListStreamQueryDataSource(value)
	return nil
}

// NewDoraListStreamQueryDataSourceFromValue returns a pointer to a valid DoraListStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDoraListStreamQueryDataSourceFromValue(v string) (*DoraListStreamQueryDataSource, error) {
	ev := DoraListStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DoraListStreamQueryDataSource: valid values are %v", v, allowedDoraListStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DoraListStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedDoraListStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DoraListStreamQueryDataSource value.
func (v DoraListStreamQueryDataSource) Ptr() *DoraListStreamQueryDataSource {
	return &v
}
