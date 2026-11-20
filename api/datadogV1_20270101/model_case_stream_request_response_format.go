// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseStreamRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type CaseStreamRequestResponseFormat string

// List of CaseStreamRequestResponseFormat.
const (
	CASESTREAMREQUESTRESPONSEFORMAT_EVENT_LIST CaseStreamRequestResponseFormat = "event_list"
)

var allowedCaseStreamRequestResponseFormatEnumValues = []CaseStreamRequestResponseFormat{
	CASESTREAMREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseStreamRequestResponseFormat) GetAllowedValues() []CaseStreamRequestResponseFormat {
	return allowedCaseStreamRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseStreamRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseStreamRequestResponseFormat(value)
	return nil
}

// NewCaseStreamRequestResponseFormatFromValue returns a pointer to a valid CaseStreamRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseStreamRequestResponseFormatFromValue(v string) (*CaseStreamRequestResponseFormat, error) {
	ev := CaseStreamRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseStreamRequestResponseFormat: valid values are %v", v, allowedCaseStreamRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseStreamRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedCaseStreamRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseStreamRequestResponseFormat value.
func (v CaseStreamRequestResponseFormat) Ptr() *CaseStreamRequestResponseFormat {
	return &v
}
