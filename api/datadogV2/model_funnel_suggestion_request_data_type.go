// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelSuggestionRequestDataType
type FunnelSuggestionRequestDataType string

// List of FunnelSuggestionRequestDataType.
const (
	FUNNELSUGGESTIONREQUESTDATATYPE_FUNNEL_SUGGESTION_REQUEST FunnelSuggestionRequestDataType = "funnel_suggestion_request"
)

var allowedFunnelSuggestionRequestDataTypeEnumValues = []FunnelSuggestionRequestDataType{
	FUNNELSUGGESTIONREQUESTDATATYPE_FUNNEL_SUGGESTION_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FunnelSuggestionRequestDataType) GetAllowedValues() []FunnelSuggestionRequestDataType {
	return allowedFunnelSuggestionRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FunnelSuggestionRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FunnelSuggestionRequestDataType(value)
	return nil
}

// NewFunnelSuggestionRequestDataTypeFromValue returns a pointer to a valid FunnelSuggestionRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFunnelSuggestionRequestDataTypeFromValue(v string) (*FunnelSuggestionRequestDataType, error) {
	ev := FunnelSuggestionRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FunnelSuggestionRequestDataType: valid values are %v", v, allowedFunnelSuggestionRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FunnelSuggestionRequestDataType) IsValid() bool {
	for _, existing := range allowedFunnelSuggestionRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FunnelSuggestionRequestDataType value.
func (v FunnelSuggestionRequestDataType) Ptr() *FunnelSuggestionRequestDataType {
	return &v
}
