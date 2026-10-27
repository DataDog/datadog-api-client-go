// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HistoricalJobFailureReason The reason why the job failed.
type HistoricalJobFailureReason string

// List of HistoricalJobFailureReason.
const (
	HISTORICALJOBFAILUREREASON_RULE_CONFIGURATION  HistoricalJobFailureReason = "rule_configuration"
	HISTORICALJOBFAILUREREASON_QUERY_TOO_EXPENSIVE HistoricalJobFailureReason = "query_too_expensive"
	HISTORICALJOBFAILUREREASON_INTERNAL_ERROR      HistoricalJobFailureReason = "internal_error"
)

var allowedHistoricalJobFailureReasonEnumValues = []HistoricalJobFailureReason{
	HISTORICALJOBFAILUREREASON_RULE_CONFIGURATION,
	HISTORICALJOBFAILUREREASON_QUERY_TOO_EXPENSIVE,
	HISTORICALJOBFAILUREREASON_INTERNAL_ERROR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HistoricalJobFailureReason) GetAllowedValues() []HistoricalJobFailureReason {
	return allowedHistoricalJobFailureReasonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HistoricalJobFailureReason) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HistoricalJobFailureReason(value)
	return nil
}

// NewHistoricalJobFailureReasonFromValue returns a pointer to a valid HistoricalJobFailureReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHistoricalJobFailureReasonFromValue(v string) (*HistoricalJobFailureReason, error) {
	ev := HistoricalJobFailureReason(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HistoricalJobFailureReason: valid values are %v", v, allowedHistoricalJobFailureReasonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HistoricalJobFailureReason) IsValid() bool {
	for _, existing := range allowedHistoricalJobFailureReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HistoricalJobFailureReason value.
func (v HistoricalJobFailureReason) Ptr() *HistoricalJobFailureReason {
	return &v
}
