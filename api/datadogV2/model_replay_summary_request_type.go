// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReplaySummaryRequestType RUM replay summary request resource type.
type ReplaySummaryRequestType string

// List of ReplaySummaryRequestType.
const (
	REPLAYSUMMARYREQUESTTYPE_REPLAY_SUMMARY_REQUEST ReplaySummaryRequestType = "replay_summary_request"
)

var allowedReplaySummaryRequestTypeEnumValues = []ReplaySummaryRequestType{
	REPLAYSUMMARYREQUESTTYPE_REPLAY_SUMMARY_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReplaySummaryRequestType) GetAllowedValues() []ReplaySummaryRequestType {
	return allowedReplaySummaryRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReplaySummaryRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReplaySummaryRequestType(value)
	return nil
}

// NewReplaySummaryRequestTypeFromValue returns a pointer to a valid ReplaySummaryRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReplaySummaryRequestTypeFromValue(v string) (*ReplaySummaryRequestType, error) {
	ev := ReplaySummaryRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReplaySummaryRequestType: valid values are %v", v, allowedReplaySummaryRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReplaySummaryRequestType) IsValid() bool {
	for _, existing := range allowedReplaySummaryRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReplaySummaryRequestType value.
func (v ReplaySummaryRequestType) Ptr() *ReplaySummaryRequestType {
	return &v
}
