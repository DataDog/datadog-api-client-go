// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueStreamOrderingKey Test-only generated schema for the dashboard versioning proof of concept.
type IssueStreamOrderingKey string

// List of IssueStreamOrderingKey.
const (
	ISSUESTREAMORDERINGKEY_TOTAL_COUNT       IssueStreamOrderingKey = "TOTAL_COUNT"
	ISSUESTREAMORDERINGKEY_FIRST_SEEN        IssueStreamOrderingKey = "FIRST_SEEN"
	ISSUESTREAMORDERINGKEY_IMPACTED_SESSIONS IssueStreamOrderingKey = "IMPACTED_SESSIONS"
	ISSUESTREAMORDERINGKEY_IMPACTED_ACCOUNTS IssueStreamOrderingKey = "IMPACTED_ACCOUNTS"
	ISSUESTREAMORDERINGKEY_PRIORITY          IssueStreamOrderingKey = "PRIORITY"
)

var allowedIssueStreamOrderingKeyEnumValues = []IssueStreamOrderingKey{
	ISSUESTREAMORDERINGKEY_TOTAL_COUNT,
	ISSUESTREAMORDERINGKEY_FIRST_SEEN,
	ISSUESTREAMORDERINGKEY_IMPACTED_SESSIONS,
	ISSUESTREAMORDERINGKEY_IMPACTED_ACCOUNTS,
	ISSUESTREAMORDERINGKEY_PRIORITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueStreamOrderingKey) GetAllowedValues() []IssueStreamOrderingKey {
	return allowedIssueStreamOrderingKeyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueStreamOrderingKey) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueStreamOrderingKey(value)
	return nil
}

// NewIssueStreamOrderingKeyFromValue returns a pointer to a valid IssueStreamOrderingKey
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueStreamOrderingKeyFromValue(v string) (*IssueStreamOrderingKey, error) {
	ev := IssueStreamOrderingKey(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueStreamOrderingKey: valid values are %v", v, allowedIssueStreamOrderingKeyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueStreamOrderingKey) IsValid() bool {
	for _, existing := range allowedIssueStreamOrderingKeyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueStreamOrderingKey value.
func (v IssueStreamOrderingKey) Ptr() *IssueStreamOrderingKey {
	return &v
}
