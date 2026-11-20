// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LegacyIssueStreamRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type LegacyIssueStreamRequestResponseFormat string

// List of LegacyIssueStreamRequestResponseFormat.
const (
	LEGACYISSUESTREAMREQUESTRESPONSEFORMAT_EVENT_LIST LegacyIssueStreamRequestResponseFormat = "event_list"
)

var allowedLegacyIssueStreamRequestResponseFormatEnumValues = []LegacyIssueStreamRequestResponseFormat{
	LEGACYISSUESTREAMREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LegacyIssueStreamRequestResponseFormat) GetAllowedValues() []LegacyIssueStreamRequestResponseFormat {
	return allowedLegacyIssueStreamRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LegacyIssueStreamRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LegacyIssueStreamRequestResponseFormat(value)
	return nil
}

// NewLegacyIssueStreamRequestResponseFormatFromValue returns a pointer to a valid LegacyIssueStreamRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLegacyIssueStreamRequestResponseFormatFromValue(v string) (*LegacyIssueStreamRequestResponseFormat, error) {
	ev := LegacyIssueStreamRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LegacyIssueStreamRequestResponseFormat: valid values are %v", v, allowedLegacyIssueStreamRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LegacyIssueStreamRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedLegacyIssueStreamRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LegacyIssueStreamRequestResponseFormat value.
func (v LegacyIssueStreamRequestResponseFormat) Ptr() *LegacyIssueStreamRequestResponseFormat {
	return &v
}
