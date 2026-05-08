// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReplayAnalysisIssueSessionDataType RUM replay analysis issue session resource type.
type ReplayAnalysisIssueSessionDataType string

// List of ReplayAnalysisIssueSessionDataType.
const (
	REPLAYANALYSISISSUESESSIONDATATYPE_SESSIONS ReplayAnalysisIssueSessionDataType = "sessions"
)

var allowedReplayAnalysisIssueSessionDataTypeEnumValues = []ReplayAnalysisIssueSessionDataType{
	REPLAYANALYSISISSUESESSIONDATATYPE_SESSIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReplayAnalysisIssueSessionDataType) GetAllowedValues() []ReplayAnalysisIssueSessionDataType {
	return allowedReplayAnalysisIssueSessionDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReplayAnalysisIssueSessionDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReplayAnalysisIssueSessionDataType(value)
	return nil
}

// NewReplayAnalysisIssueSessionDataTypeFromValue returns a pointer to a valid ReplayAnalysisIssueSessionDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReplayAnalysisIssueSessionDataTypeFromValue(v string) (*ReplayAnalysisIssueSessionDataType, error) {
	ev := ReplayAnalysisIssueSessionDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReplayAnalysisIssueSessionDataType: valid values are %v", v, allowedReplayAnalysisIssueSessionDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReplayAnalysisIssueSessionDataType) IsValid() bool {
	for _, existing := range allowedReplayAnalysisIssueSessionDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReplayAnalysisIssueSessionDataType value.
func (v ReplayAnalysisIssueSessionDataType) Ptr() *ReplayAnalysisIssueSessionDataType {
	return &v
}
