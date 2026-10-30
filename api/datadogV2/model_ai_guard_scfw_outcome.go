// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardScfwOutcome The Supply Chain Firewall outcome for a package, a matched rule, or the whole request.
type AIGuardScfwOutcome string

// List of AIGuardScfwOutcome.
const (
	AIGUARDSCFWOUTCOME_ALLOW AIGuardScfwOutcome = "ALLOW"
	AIGUARDSCFWOUTCOME_WARN  AIGuardScfwOutcome = "WARN"
	AIGUARDSCFWOUTCOME_BLOCK AIGuardScfwOutcome = "BLOCK"
)

var allowedAIGuardScfwOutcomeEnumValues = []AIGuardScfwOutcome{
	AIGUARDSCFWOUTCOME_ALLOW,
	AIGUARDSCFWOUTCOME_WARN,
	AIGUARDSCFWOUTCOME_BLOCK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AIGuardScfwOutcome) GetAllowedValues() []AIGuardScfwOutcome {
	return allowedAIGuardScfwOutcomeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AIGuardScfwOutcome) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AIGuardScfwOutcome(value)
	return nil
}

// NewAIGuardScfwOutcomeFromValue returns a pointer to a valid AIGuardScfwOutcome
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAIGuardScfwOutcomeFromValue(v string) (*AIGuardScfwOutcome, error) {
	ev := AIGuardScfwOutcome(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AIGuardScfwOutcome: valid values are %v", v, allowedAIGuardScfwOutcomeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AIGuardScfwOutcome) IsValid() bool {
	for _, existing := range allowedAIGuardScfwOutcomeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AIGuardScfwOutcome value.
func (v AIGuardScfwOutcome) Ptr() *AIGuardScfwOutcome {
	return &v
}
