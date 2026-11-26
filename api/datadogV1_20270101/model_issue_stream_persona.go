// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueStreamPersona Test-only generated schema for the dashboard versioning proof of concept.
type IssueStreamPersona string

// List of IssueStreamPersona.
const (
	ISSUESTREAMPERSONA_ALL     IssueStreamPersona = "all"
	ISSUESTREAMPERSONA_BROWSER IssueStreamPersona = "browser"
	ISSUESTREAMPERSONA_MOBILE  IssueStreamPersona = "mobile"
	ISSUESTREAMPERSONA_BACKEND IssueStreamPersona = "backend"
	ISSUESTREAMPERSONA_AGENT   IssueStreamPersona = "agent"
)

var allowedIssueStreamPersonaEnumValues = []IssueStreamPersona{
	ISSUESTREAMPERSONA_ALL,
	ISSUESTREAMPERSONA_BROWSER,
	ISSUESTREAMPERSONA_MOBILE,
	ISSUESTREAMPERSONA_BACKEND,
	ISSUESTREAMPERSONA_AGENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueStreamPersona) GetAllowedValues() []IssueStreamPersona {
	return allowedIssueStreamPersonaEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueStreamPersona) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueStreamPersona(value)
	return nil
}

// NewIssueStreamPersonaFromValue returns a pointer to a valid IssueStreamPersona
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueStreamPersonaFromValue(v string) (*IssueStreamPersona, error) {
	ev := IssueStreamPersona(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueStreamPersona: valid values are %v", v, allowedIssueStreamPersonaEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueStreamPersona) IsValid() bool {
	for _, existing := range allowedIssueStreamPersonaEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueStreamPersona value.
func (v IssueStreamPersona) Ptr() *IssueStreamPersona {
	return &v
}
