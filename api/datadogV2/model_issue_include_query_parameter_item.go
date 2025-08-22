// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueIncludeQueryParameterItem Relationship object that should be included in the response.
type IssueIncludeQueryParameterItem string

// List of IssueIncludeQueryParameterItem.
const (
	ISSUEINCLUDEQUERYPARAMETERITEM_CASE  IssueIncludeQueryParameterItem = "case"
	ISSUEINCLUDEQUERYPARAMETERITEM_TEAMS IssueIncludeQueryParameterItem = "teams"
)

var allowedIssueIncludeQueryParameterItemEnumValues = []IssueIncludeQueryParameterItem{
	ISSUEINCLUDEQUERYPARAMETERITEM_CASE,
	ISSUEINCLUDEQUERYPARAMETERITEM_TEAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueIncludeQueryParameterItem) GetAllowedValues() []IssueIncludeQueryParameterItem {
	return allowedIssueIncludeQueryParameterItemEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueIncludeQueryParameterItem) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueIncludeQueryParameterItem(value)
	return nil
}

// NewIssueIncludeQueryParameterItemFromValue returns a pointer to a valid IssueIncludeQueryParameterItem
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueIncludeQueryParameterItemFromValue(v string) (*IssueIncludeQueryParameterItem, error) {
	ev := IssueIncludeQueryParameterItem(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueIncludeQueryParameterItem: valid values are %v", v, allowedIssueIncludeQueryParameterItemEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueIncludeQueryParameterItem) IsValid() bool {
	for _, existing := range allowedIssueIncludeQueryParameterItemEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueIncludeQueryParameterItem value.
func (v IssueIncludeQueryParameterItem) Ptr() *IssueIncludeQueryParameterItem {
	return &v
}
