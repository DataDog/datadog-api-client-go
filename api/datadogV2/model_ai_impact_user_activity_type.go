// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIImpactUserActivityType JSON:API type for AI Impact user activity entries.
type AIImpactUserActivityType string

// List of AIImpactUserActivityType.
const (
	AIIMPACTUSERACTIVITYTYPE_AI_IMPACT_USER_ACTIVITY AIImpactUserActivityType = "ai_impact_user_activity"
)

var allowedAIImpactUserActivityTypeEnumValues = []AIImpactUserActivityType{
	AIIMPACTUSERACTIVITYTYPE_AI_IMPACT_USER_ACTIVITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AIImpactUserActivityType) GetAllowedValues() []AIImpactUserActivityType {
	return allowedAIImpactUserActivityTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AIImpactUserActivityType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AIImpactUserActivityType(value)
	return nil
}

// NewAIImpactUserActivityTypeFromValue returns a pointer to a valid AIImpactUserActivityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAIImpactUserActivityTypeFromValue(v string) (*AIImpactUserActivityType, error) {
	ev := AIImpactUserActivityType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AIImpactUserActivityType: valid values are %v", v, allowedAIImpactUserActivityTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AIImpactUserActivityType) IsValid() bool {
	for _, existing := range allowedAIImpactUserActivityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AIImpactUserActivityType value.
func (v AIImpactUserActivityType) Ptr() *AIImpactUserActivityType {
	return &v
}
