// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityExclusionFilterOnMatch The action taken when the exclusion filter matches. When set to `monitor`, security traces are emitted but the requests are not blocked. By default, security traces are not emitted and the requests are not blocked.
type ApplicationSecurityExclusionFilterOnMatch string

// List of ApplicationSecurityExclusionFilterOnMatch.
const (
	APPLICATIONSECURITYEXCLUSIONFILTERONMATCH_MONITOR ApplicationSecurityExclusionFilterOnMatch = "monitor"
)

var allowedApplicationSecurityExclusionFilterOnMatchEnumValues = []ApplicationSecurityExclusionFilterOnMatch{
	APPLICATIONSECURITYEXCLUSIONFILTERONMATCH_MONITOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApplicationSecurityExclusionFilterOnMatch) GetAllowedValues() []ApplicationSecurityExclusionFilterOnMatch {
	return allowedApplicationSecurityExclusionFilterOnMatchEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApplicationSecurityExclusionFilterOnMatch) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApplicationSecurityExclusionFilterOnMatch(value)
	return nil
}

// NewApplicationSecurityExclusionFilterOnMatchFromValue returns a pointer to a valid ApplicationSecurityExclusionFilterOnMatch
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApplicationSecurityExclusionFilterOnMatchFromValue(v string) (*ApplicationSecurityExclusionFilterOnMatch, error) {
	ev := ApplicationSecurityExclusionFilterOnMatch(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApplicationSecurityExclusionFilterOnMatch: valid values are %v", v, allowedApplicationSecurityExclusionFilterOnMatchEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApplicationSecurityExclusionFilterOnMatch) IsValid() bool {
	for _, existing := range allowedApplicationSecurityExclusionFilterOnMatchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApplicationSecurityExclusionFilterOnMatch value.
func (v ApplicationSecurityExclusionFilterOnMatch) Ptr() *ApplicationSecurityExclusionFilterOnMatch {
	return &v
}
