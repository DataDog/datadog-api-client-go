// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DetectionRuleStreamRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type DetectionRuleStreamRequestResponseFormat string

// List of DetectionRuleStreamRequestResponseFormat.
const (
	DETECTIONRULESTREAMREQUESTRESPONSEFORMAT_EVENT_LIST DetectionRuleStreamRequestResponseFormat = "event_list"
)

var allowedDetectionRuleStreamRequestResponseFormatEnumValues = []DetectionRuleStreamRequestResponseFormat{
	DETECTIONRULESTREAMREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DetectionRuleStreamRequestResponseFormat) GetAllowedValues() []DetectionRuleStreamRequestResponseFormat {
	return allowedDetectionRuleStreamRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DetectionRuleStreamRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DetectionRuleStreamRequestResponseFormat(value)
	return nil
}

// NewDetectionRuleStreamRequestResponseFormatFromValue returns a pointer to a valid DetectionRuleStreamRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDetectionRuleStreamRequestResponseFormatFromValue(v string) (*DetectionRuleStreamRequestResponseFormat, error) {
	ev := DetectionRuleStreamRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DetectionRuleStreamRequestResponseFormat: valid values are %v", v, allowedDetectionRuleStreamRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DetectionRuleStreamRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedDetectionRuleStreamRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DetectionRuleStreamRequestResponseFormat value.
func (v DetectionRuleStreamRequestResponseFormat) Ptr() *DetectionRuleStreamRequestResponseFormat {
	return &v
}
