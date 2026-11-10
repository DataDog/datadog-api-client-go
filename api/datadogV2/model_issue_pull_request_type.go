// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssuePullRequestType Type of the object.
type IssuePullRequestType string

// List of IssuePullRequestType.
const (
	ISSUEPULLREQUESTTYPE_ERROR_TRACKING_PULL_REQUEST IssuePullRequestType = "error_tracking_pull_request"
)

var allowedIssuePullRequestTypeEnumValues = []IssuePullRequestType{
	ISSUEPULLREQUESTTYPE_ERROR_TRACKING_PULL_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssuePullRequestType) GetAllowedValues() []IssuePullRequestType {
	return allowedIssuePullRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssuePullRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssuePullRequestType(value)
	return nil
}

// NewIssuePullRequestTypeFromValue returns a pointer to a valid IssuePullRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssuePullRequestTypeFromValue(v string) (*IssuePullRequestType, error) {
	ev := IssuePullRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssuePullRequestType: valid values are %v", v, allowedIssuePullRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssuePullRequestType) IsValid() bool {
	for _, existing := range allowedIssuePullRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssuePullRequestType value.
func (v IssuePullRequestType) Ptr() *IssuePullRequestType {
	return &v
}
