// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelSuggestionResponseDataType
type FunnelSuggestionResponseDataType string

// List of FunnelSuggestionResponseDataType.
const (
	FUNNELSUGGESTIONRESPONSEDATATYPE_FUNNEL_SUGGESTION_RESPONSE FunnelSuggestionResponseDataType = "funnel_suggestion_response"
)

var allowedFunnelSuggestionResponseDataTypeEnumValues = []FunnelSuggestionResponseDataType{
	FUNNELSUGGESTIONRESPONSEDATATYPE_FUNNEL_SUGGESTION_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FunnelSuggestionResponseDataType) GetAllowedValues() []FunnelSuggestionResponseDataType {
	return allowedFunnelSuggestionResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FunnelSuggestionResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FunnelSuggestionResponseDataType(value)
	return nil
}

// NewFunnelSuggestionResponseDataTypeFromValue returns a pointer to a valid FunnelSuggestionResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFunnelSuggestionResponseDataTypeFromValue(v string) (*FunnelSuggestionResponseDataType, error) {
	ev := FunnelSuggestionResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FunnelSuggestionResponseDataType: valid values are %v", v, allowedFunnelSuggestionResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FunnelSuggestionResponseDataType) IsValid() bool {
	for _, existing := range allowedFunnelSuggestionResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FunnelSuggestionResponseDataType value.
func (v FunnelSuggestionResponseDataType) Ptr() *FunnelSuggestionResponseDataType {
	return &v
}
