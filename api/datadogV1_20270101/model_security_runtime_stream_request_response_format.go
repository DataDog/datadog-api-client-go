// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityRuntimeStreamRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type SecurityRuntimeStreamRequestResponseFormat string

// List of SecurityRuntimeStreamRequestResponseFormat.
const (
	SECURITYRUNTIMESTREAMREQUESTRESPONSEFORMAT_EVENT_LIST SecurityRuntimeStreamRequestResponseFormat = "event_list"
)

var allowedSecurityRuntimeStreamRequestResponseFormatEnumValues = []SecurityRuntimeStreamRequestResponseFormat{
	SECURITYRUNTIMESTREAMREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityRuntimeStreamRequestResponseFormat) GetAllowedValues() []SecurityRuntimeStreamRequestResponseFormat {
	return allowedSecurityRuntimeStreamRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityRuntimeStreamRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityRuntimeStreamRequestResponseFormat(value)
	return nil
}

// NewSecurityRuntimeStreamRequestResponseFormatFromValue returns a pointer to a valid SecurityRuntimeStreamRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityRuntimeStreamRequestResponseFormatFromValue(v string) (*SecurityRuntimeStreamRequestResponseFormat, error) {
	ev := SecurityRuntimeStreamRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityRuntimeStreamRequestResponseFormat: valid values are %v", v, allowedSecurityRuntimeStreamRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityRuntimeStreamRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedSecurityRuntimeStreamRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityRuntimeStreamRequestResponseFormat value.
func (v SecurityRuntimeStreamRequestResponseFormat) Ptr() *SecurityRuntimeStreamRequestResponseFormat {
	return &v
}
