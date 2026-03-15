// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScoreType The JSON:API type for scores.
type ScoreType string

// List of ScoreType.
const (
	SCORETYPE_SCORE ScoreType = "score"
)

var allowedScoreTypeEnumValues = []ScoreType{
	SCORETYPE_SCORE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScoreType) GetAllowedValues() []ScoreType {
	return allowedScoreTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScoreType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScoreType(value)
	return nil
}

// NewScoreTypeFromValue returns a pointer to a valid ScoreType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScoreTypeFromValue(v string) (*ScoreType, error) {
	ev := ScoreType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScoreType: valid values are %v", v, allowedScoreTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScoreType) IsValid() bool {
	for _, existing := range allowedScoreTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScoreType value.
func (v ScoreType) Ptr() *ScoreType {
	return &v
}
