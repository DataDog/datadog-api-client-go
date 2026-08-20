// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardAction The action recommendation from the AI Guard evaluation.
type AIGuardAction string

// List of AIGuardAction.
const (
	AIGUARDACTION_ALLOW AIGuardAction = "ALLOW"
	AIGUARDACTION_DENY  AIGuardAction = "DENY"
	AIGUARDACTION_ABORT AIGuardAction = "ABORT"
)

var allowedAIGuardActionEnumValues = []AIGuardAction{
	AIGUARDACTION_ALLOW,
	AIGUARDACTION_DENY,
	AIGUARDACTION_ABORT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AIGuardAction) GetAllowedValues() []AIGuardAction {
	return allowedAIGuardActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AIGuardAction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AIGuardAction(value)
	return nil
}

// NewAIGuardActionFromValue returns a pointer to a valid AIGuardAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAIGuardActionFromValue(v string) (*AIGuardAction, error) {
	ev := AIGuardAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AIGuardAction: valid values are %v", v, allowedAIGuardActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AIGuardAction) IsValid() bool {
	for _, existing := range allowedAIGuardActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AIGuardAction value.
func (v AIGuardAction) Ptr() *AIGuardAction {
	return &v
}
