// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationGuardrailVerdict The verdict a guardrail applied to a plan or investigation.
type RemediationGuardrailVerdict string

// List of RemediationGuardrailVerdict.
const (
	REMEDIATIONGUARDRAILVERDICT_ALLOWED           RemediationGuardrailVerdict = "allowed"
	REMEDIATIONGUARDRAILVERDICT_APPROVAL_REQUIRED RemediationGuardrailVerdict = "approval_required"
	REMEDIATIONGUARDRAILVERDICT_DENIED            RemediationGuardrailVerdict = "denied"
)

var allowedRemediationGuardrailVerdictEnumValues = []RemediationGuardrailVerdict{
	REMEDIATIONGUARDRAILVERDICT_ALLOWED,
	REMEDIATIONGUARDRAILVERDICT_APPROVAL_REQUIRED,
	REMEDIATIONGUARDRAILVERDICT_DENIED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RemediationGuardrailVerdict) GetAllowedValues() []RemediationGuardrailVerdict {
	return allowedRemediationGuardrailVerdictEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RemediationGuardrailVerdict) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RemediationGuardrailVerdict(value)
	return nil
}

// NewRemediationGuardrailVerdictFromValue returns a pointer to a valid RemediationGuardrailVerdict
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRemediationGuardrailVerdictFromValue(v string) (*RemediationGuardrailVerdict, error) {
	ev := RemediationGuardrailVerdict(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RemediationGuardrailVerdict: valid values are %v", v, allowedRemediationGuardrailVerdictEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RemediationGuardrailVerdict) IsValid() bool {
	for _, existing := range allowedRemediationGuardrailVerdictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RemediationGuardrailVerdict value.
func (v RemediationGuardrailVerdict) Ptr() *RemediationGuardrailVerdict {
	return &v
}
