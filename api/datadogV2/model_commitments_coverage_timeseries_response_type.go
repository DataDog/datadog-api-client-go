// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsCoverageTimeseriesResponseType Type of the timeseries coverage response resource.
type CommitmentsCoverageTimeseriesResponseType string

// List of CommitmentsCoverageTimeseriesResponseType.
const (
	COMMITMENTSCOVERAGETIMESERIESRESPONSETYPE_COVERAGE_TIMESERIES_RESPONSE CommitmentsCoverageTimeseriesResponseType = "coverage_timeseries_response"
)

var allowedCommitmentsCoverageTimeseriesResponseTypeEnumValues = []CommitmentsCoverageTimeseriesResponseType{
	COMMITMENTSCOVERAGETIMESERIESRESPONSETYPE_COVERAGE_TIMESERIES_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CommitmentsCoverageTimeseriesResponseType) GetAllowedValues() []CommitmentsCoverageTimeseriesResponseType {
	return allowedCommitmentsCoverageTimeseriesResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CommitmentsCoverageTimeseriesResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CommitmentsCoverageTimeseriesResponseType(value)
	return nil
}

// NewCommitmentsCoverageTimeseriesResponseTypeFromValue returns a pointer to a valid CommitmentsCoverageTimeseriesResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCommitmentsCoverageTimeseriesResponseTypeFromValue(v string) (*CommitmentsCoverageTimeseriesResponseType, error) {
	ev := CommitmentsCoverageTimeseriesResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CommitmentsCoverageTimeseriesResponseType: valid values are %v", v, allowedCommitmentsCoverageTimeseriesResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CommitmentsCoverageTimeseriesResponseType) IsValid() bool {
	for _, existing := range allowedCommitmentsCoverageTimeseriesResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommitmentsCoverageTimeseriesResponseType value.
func (v CommitmentsCoverageTimeseriesResponseType) Ptr() *CommitmentsCoverageTimeseriesResponseType {
	return &v
}
