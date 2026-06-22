// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationPlanSource How the plan was produced.
type RemediationPlanSource string

// List of RemediationPlanSource.
const (
	REMEDIATIONPLANSOURCE_DETERMINISTIC RemediationPlanSource = "deterministic"
	REMEDIATIONPLANSOURCE_AI            RemediationPlanSource = "ai"
)

var allowedRemediationPlanSourceEnumValues = []RemediationPlanSource{
	REMEDIATIONPLANSOURCE_DETERMINISTIC,
	REMEDIATIONPLANSOURCE_AI,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RemediationPlanSource) GetAllowedValues() []RemediationPlanSource {
	return allowedRemediationPlanSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RemediationPlanSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RemediationPlanSource(value)
	return nil
}

// NewRemediationPlanSourceFromValue returns a pointer to a valid RemediationPlanSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRemediationPlanSourceFromValue(v string) (*RemediationPlanSource, error) {
	ev := RemediationPlanSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RemediationPlanSource: valid values are %v", v, allowedRemediationPlanSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RemediationPlanSource) IsValid() bool {
	for _, existing := range allowedRemediationPlanSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RemediationPlanSource value.
func (v RemediationPlanSource) Ptr() *RemediationPlanSource {
	return &v
}
