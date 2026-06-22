// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationConfidence The agent's self-rated confidence in a plan.
type RemediationConfidence string

// List of RemediationConfidence.
const (
	REMEDIATIONCONFIDENCE_LOW    RemediationConfidence = "low"
	REMEDIATIONCONFIDENCE_MEDIUM RemediationConfidence = "medium"
	REMEDIATIONCONFIDENCE_HIGH   RemediationConfidence = "high"
)

var allowedRemediationConfidenceEnumValues = []RemediationConfidence{
	REMEDIATIONCONFIDENCE_LOW,
	REMEDIATIONCONFIDENCE_MEDIUM,
	REMEDIATIONCONFIDENCE_HIGH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RemediationConfidence) GetAllowedValues() []RemediationConfidence {
	return allowedRemediationConfidenceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RemediationConfidence) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RemediationConfidence(value)
	return nil
}

// NewRemediationConfidenceFromValue returns a pointer to a valid RemediationConfidence
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRemediationConfidenceFromValue(v string) (*RemediationConfidence, error) {
	ev := RemediationConfidence(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RemediationConfidence: valid values are %v", v, allowedRemediationConfidenceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RemediationConfidence) IsValid() bool {
	for _, existing := range allowedRemediationConfidenceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RemediationConfidence value.
func (v RemediationConfidence) Ptr() *RemediationConfidence {
	return &v
}
