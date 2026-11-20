// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationGroupInclude Related resource to add to `included`.
type RemediationGroupInclude string

// List of RemediationGroupInclude.
const (
	REMEDIATIONGROUPINCLUDE_CASE RemediationGroupInclude = "case"
)

var allowedRemediationGroupIncludeEnumValues = []RemediationGroupInclude{
	REMEDIATIONGROUPINCLUDE_CASE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RemediationGroupInclude) GetAllowedValues() []RemediationGroupInclude {
	return allowedRemediationGroupIncludeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RemediationGroupInclude) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RemediationGroupInclude(value)
	return nil
}

// NewRemediationGroupIncludeFromValue returns a pointer to a valid RemediationGroupInclude
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRemediationGroupIncludeFromValue(v string) (*RemediationGroupInclude, error) {
	ev := RemediationGroupInclude(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RemediationGroupInclude: valid values are %v", v, allowedRemediationGroupIncludeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RemediationGroupInclude) IsValid() bool {
	for _, existing := range allowedRemediationGroupIncludeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RemediationGroupInclude value.
func (v RemediationGroupInclude) Ptr() *RemediationGroupInclude {
	return &v
}
