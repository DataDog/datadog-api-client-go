// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsSavingsTimeseriesResponseType Type of the timeseries savings response resource.
type CommitmentsSavingsTimeseriesResponseType string

// List of CommitmentsSavingsTimeseriesResponseType.
const (
	COMMITMENTSSAVINGSTIMESERIESRESPONSETYPE_SAVINGS_TIMESERIES_RESPONSE CommitmentsSavingsTimeseriesResponseType = "savings_timeseries_response"
)

var allowedCommitmentsSavingsTimeseriesResponseTypeEnumValues = []CommitmentsSavingsTimeseriesResponseType{
	COMMITMENTSSAVINGSTIMESERIESRESPONSETYPE_SAVINGS_TIMESERIES_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CommitmentsSavingsTimeseriesResponseType) GetAllowedValues() []CommitmentsSavingsTimeseriesResponseType {
	return allowedCommitmentsSavingsTimeseriesResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CommitmentsSavingsTimeseriesResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CommitmentsSavingsTimeseriesResponseType(value)
	return nil
}

// NewCommitmentsSavingsTimeseriesResponseTypeFromValue returns a pointer to a valid CommitmentsSavingsTimeseriesResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCommitmentsSavingsTimeseriesResponseTypeFromValue(v string) (*CommitmentsSavingsTimeseriesResponseType, error) {
	ev := CommitmentsSavingsTimeseriesResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CommitmentsSavingsTimeseriesResponseType: valid values are %v", v, allowedCommitmentsSavingsTimeseriesResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CommitmentsSavingsTimeseriesResponseType) IsValid() bool {
	for _, existing := range allowedCommitmentsSavingsTimeseriesResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommitmentsSavingsTimeseriesResponseType value.
func (v CommitmentsSavingsTimeseriesResponseType) Ptr() *CommitmentsSavingsTimeseriesResponseType {
	return &v
}
