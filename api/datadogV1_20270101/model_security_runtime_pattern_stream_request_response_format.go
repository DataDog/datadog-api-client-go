// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityRuntimePatternStreamRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type SecurityRuntimePatternStreamRequestResponseFormat string

// List of SecurityRuntimePatternStreamRequestResponseFormat.
const (
	SECURITYRUNTIMEPATTERNSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST SecurityRuntimePatternStreamRequestResponseFormat = "event_list"
)

var allowedSecurityRuntimePatternStreamRequestResponseFormatEnumValues = []SecurityRuntimePatternStreamRequestResponseFormat{
	SECURITYRUNTIMEPATTERNSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityRuntimePatternStreamRequestResponseFormat) GetAllowedValues() []SecurityRuntimePatternStreamRequestResponseFormat {
	return allowedSecurityRuntimePatternStreamRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityRuntimePatternStreamRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityRuntimePatternStreamRequestResponseFormat(value)
	return nil
}

// NewSecurityRuntimePatternStreamRequestResponseFormatFromValue returns a pointer to a valid SecurityRuntimePatternStreamRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityRuntimePatternStreamRequestResponseFormatFromValue(v string) (*SecurityRuntimePatternStreamRequestResponseFormat, error) {
	ev := SecurityRuntimePatternStreamRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityRuntimePatternStreamRequestResponseFormat: valid values are %v", v, allowedSecurityRuntimePatternStreamRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityRuntimePatternStreamRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedSecurityRuntimePatternStreamRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityRuntimePatternStreamRequestResponseFormat value.
func (v SecurityRuntimePatternStreamRequestResponseFormat) Ptr() *SecurityRuntimePatternStreamRequestResponseFormat {
	return &v
}
