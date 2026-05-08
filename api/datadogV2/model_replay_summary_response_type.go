// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReplaySummaryResponseType RUM replay summary response resource type.
type ReplaySummaryResponseType string

// List of ReplaySummaryResponseType.
const (
	REPLAYSUMMARYRESPONSETYPE_SUMMARY_RESPONSE ReplaySummaryResponseType = "summary_response"
)

var allowedReplaySummaryResponseTypeEnumValues = []ReplaySummaryResponseType{
	REPLAYSUMMARYRESPONSETYPE_SUMMARY_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReplaySummaryResponseType) GetAllowedValues() []ReplaySummaryResponseType {
	return allowedReplaySummaryResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReplaySummaryResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReplaySummaryResponseType(value)
	return nil
}

// NewReplaySummaryResponseTypeFromValue returns a pointer to a valid ReplaySummaryResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReplaySummaryResponseTypeFromValue(v string) (*ReplaySummaryResponseType, error) {
	ev := ReplaySummaryResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReplaySummaryResponseType: valid values are %v", v, allowedReplaySummaryResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReplaySummaryResponseType) IsValid() bool {
	for _, existing := range allowedReplaySummaryResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReplaySummaryResponseType value.
func (v ReplaySummaryResponseType) Ptr() *ReplaySummaryResponseType {
	return &v
}
