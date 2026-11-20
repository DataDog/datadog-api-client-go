// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueStreamSortOrder Test-only generated schema for the dashboard versioning proof of concept.
type IssueStreamSortOrder string

// List of IssueStreamSortOrder.
const (
	ISSUESTREAMSORTORDER_DESC IssueStreamSortOrder = "desc"
)

var allowedIssueStreamSortOrderEnumValues = []IssueStreamSortOrder{
	ISSUESTREAMSORTORDER_DESC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueStreamSortOrder) GetAllowedValues() []IssueStreamSortOrder {
	return allowedIssueStreamSortOrderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueStreamSortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueStreamSortOrder(value)
	return nil
}

// NewIssueStreamSortOrderFromValue returns a pointer to a valid IssueStreamSortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueStreamSortOrderFromValue(v string) (*IssueStreamSortOrder, error) {
	ev := IssueStreamSortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueStreamSortOrder: valid values are %v", v, allowedIssueStreamSortOrderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueStreamSortOrder) IsValid() bool {
	for _, existing := range allowedIssueStreamSortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueStreamSortOrder value.
func (v IssueStreamSortOrder) Ptr() *IssueStreamSortOrder {
	return &v
}
