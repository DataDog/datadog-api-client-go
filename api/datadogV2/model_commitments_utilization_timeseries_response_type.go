// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsUtilizationTimeseriesResponseType Type of the timeseries utilization response resource.
type CommitmentsUtilizationTimeseriesResponseType string

// List of CommitmentsUtilizationTimeseriesResponseType.
const (
	COMMITMENTSUTILIZATIONTIMESERIESRESPONSETYPE_UTILIZATION_TIMESERIES_RESPONSE CommitmentsUtilizationTimeseriesResponseType = "utilization_timeseries_response"
)

var allowedCommitmentsUtilizationTimeseriesResponseTypeEnumValues = []CommitmentsUtilizationTimeseriesResponseType{
	COMMITMENTSUTILIZATIONTIMESERIESRESPONSETYPE_UTILIZATION_TIMESERIES_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CommitmentsUtilizationTimeseriesResponseType) GetAllowedValues() []CommitmentsUtilizationTimeseriesResponseType {
	return allowedCommitmentsUtilizationTimeseriesResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CommitmentsUtilizationTimeseriesResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CommitmentsUtilizationTimeseriesResponseType(value)
	return nil
}

// NewCommitmentsUtilizationTimeseriesResponseTypeFromValue returns a pointer to a valid CommitmentsUtilizationTimeseriesResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCommitmentsUtilizationTimeseriesResponseTypeFromValue(v string) (*CommitmentsUtilizationTimeseriesResponseType, error) {
	ev := CommitmentsUtilizationTimeseriesResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CommitmentsUtilizationTimeseriesResponseType: valid values are %v", v, allowedCommitmentsUtilizationTimeseriesResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CommitmentsUtilizationTimeseriesResponseType) IsValid() bool {
	for _, existing := range allowedCommitmentsUtilizationTimeseriesResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommitmentsUtilizationTimeseriesResponseType value.
func (v CommitmentsUtilizationTimeseriesResponseType) Ptr() *CommitmentsUtilizationTimeseriesResponseType {
	return &v
}
