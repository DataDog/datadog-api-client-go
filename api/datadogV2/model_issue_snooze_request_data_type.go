// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueSnoozeRequestDataType Type of the object.
type IssueSnoozeRequestDataType string

// List of IssueSnoozeRequestDataType.
const (
	ISSUESNOOZEREQUESTDATATYPE_ERROR_TRACKING_ISSUE IssueSnoozeRequestDataType = "error_tracking_issue"
)

var allowedIssueSnoozeRequestDataTypeEnumValues = []IssueSnoozeRequestDataType{
	ISSUESNOOZEREQUESTDATATYPE_ERROR_TRACKING_ISSUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueSnoozeRequestDataType) GetAllowedValues() []IssueSnoozeRequestDataType {
	return allowedIssueSnoozeRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueSnoozeRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueSnoozeRequestDataType(value)
	return nil
}

// NewIssueSnoozeRequestDataTypeFromValue returns a pointer to a valid IssueSnoozeRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueSnoozeRequestDataTypeFromValue(v string) (*IssueSnoozeRequestDataType, error) {
	ev := IssueSnoozeRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueSnoozeRequestDataType: valid values are %v", v, allowedIssueSnoozeRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueSnoozeRequestDataType) IsValid() bool {
	for _, existing := range allowedIssueSnoozeRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueSnoozeRequestDataType value.
func (v IssueSnoozeRequestDataType) Ptr() *IssueSnoozeRequestDataType {
	return &v
}
