// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DoraListStreamRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type DoraListStreamRequestResponseFormat string

// List of DoraListStreamRequestResponseFormat.
const (
	DORALISTSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST DoraListStreamRequestResponseFormat = "event_list"
)

var allowedDoraListStreamRequestResponseFormatEnumValues = []DoraListStreamRequestResponseFormat{
	DORALISTSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DoraListStreamRequestResponseFormat) GetAllowedValues() []DoraListStreamRequestResponseFormat {
	return allowedDoraListStreamRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DoraListStreamRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DoraListStreamRequestResponseFormat(value)
	return nil
}

// NewDoraListStreamRequestResponseFormatFromValue returns a pointer to a valid DoraListStreamRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDoraListStreamRequestResponseFormatFromValue(v string) (*DoraListStreamRequestResponseFormat, error) {
	ev := DoraListStreamRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DoraListStreamRequestResponseFormat: valid values are %v", v, allowedDoraListStreamRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DoraListStreamRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedDoraListStreamRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DoraListStreamRequestResponseFormat value.
func (v DoraListStreamRequestResponseFormat) Ptr() *DoraListStreamRequestResponseFormat {
	return &v
}
