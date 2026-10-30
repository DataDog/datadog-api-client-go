// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardContentPartType The type of the content part.
type AIGuardContentPartType string

// List of AIGuardContentPartType.
const (
	AIGUARDCONTENTPARTTYPE_TEXT      AIGuardContentPartType = "text"
	AIGUARDCONTENTPARTTYPE_IMAGE_URL AIGuardContentPartType = "image_url"
)

var allowedAIGuardContentPartTypeEnumValues = []AIGuardContentPartType{
	AIGUARDCONTENTPARTTYPE_TEXT,
	AIGUARDCONTENTPARTTYPE_IMAGE_URL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AIGuardContentPartType) GetAllowedValues() []AIGuardContentPartType {
	return allowedAIGuardContentPartTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AIGuardContentPartType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AIGuardContentPartType(value)
	return nil
}

// NewAIGuardContentPartTypeFromValue returns a pointer to a valid AIGuardContentPartType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAIGuardContentPartTypeFromValue(v string) (*AIGuardContentPartType, error) {
	ev := AIGuardContentPartType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AIGuardContentPartType: valid values are %v", v, allowedAIGuardContentPartTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AIGuardContentPartType) IsValid() bool {
	for _, existing := range allowedAIGuardContentPartTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AIGuardContentPartType value.
func (v AIGuardContentPartType) Ptr() *AIGuardContentPartType {
	return &v
}
