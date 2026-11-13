// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueImpactType Type of the object.
type IssueImpactType string

// List of IssueImpactType.
const (
	ISSUEIMPACTTYPE_ERROR_TRACKING_ISSUE_IMPACT IssueImpactType = "error_tracking_issue_impact"
)

var allowedIssueImpactTypeEnumValues = []IssueImpactType{
	ISSUEIMPACTTYPE_ERROR_TRACKING_ISSUE_IMPACT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueImpactType) GetAllowedValues() []IssueImpactType {
	return allowedIssueImpactTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueImpactType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueImpactType(value)
	return nil
}

// NewIssueImpactTypeFromValue returns a pointer to a valid IssueImpactType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueImpactTypeFromValue(v string) (*IssueImpactType, error) {
	ev := IssueImpactType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueImpactType: valid values are %v", v, allowedIssueImpactTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueImpactType) IsValid() bool {
	for _, existing := range allowedIssueImpactTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueImpactType value.
func (v IssueImpactType) Ptr() *IssueImpactType {
	return &v
}
