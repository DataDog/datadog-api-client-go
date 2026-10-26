// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueSampleType Type of the object.
type IssueSampleType string

// List of IssueSampleType.
const (
	ISSUESAMPLETYPE_ERROR_TRACKING_ISSUE_SAMPLE IssueSampleType = "error_tracking_issue_sample"
)

var allowedIssueSampleTypeEnumValues = []IssueSampleType{
	ISSUESAMPLETYPE_ERROR_TRACKING_ISSUE_SAMPLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueSampleType) GetAllowedValues() []IssueSampleType {
	return allowedIssueSampleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueSampleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueSampleType(value)
	return nil
}

// NewIssueSampleTypeFromValue returns a pointer to a valid IssueSampleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueSampleTypeFromValue(v string) (*IssueSampleType, error) {
	ev := IssueSampleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueSampleType: valid values are %v", v, allowedIssueSampleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueSampleType) IsValid() bool {
	for _, existing := range allowedIssueSampleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueSampleType value.
func (v IssueSampleType) Ptr() *IssueSampleType {
	return &v
}
