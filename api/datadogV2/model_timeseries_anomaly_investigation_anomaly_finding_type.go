// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationAnomalyFindingType Finding category for an anomaly without a displayable influential tag.
type TimeseriesAnomalyInvestigationAnomalyFindingType string

// List of TimeseriesAnomalyInvestigationAnomalyFindingType.
const (
	TIMESERIESANOMALYINVESTIGATIONANOMALYFINDINGTYPE_ANOMALY TimeseriesAnomalyInvestigationAnomalyFindingType = "anomaly"
)

var allowedTimeseriesAnomalyInvestigationAnomalyFindingTypeEnumValues = []TimeseriesAnomalyInvestigationAnomalyFindingType{
	TIMESERIESANOMALYINVESTIGATIONANOMALYFINDINGTYPE_ANOMALY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesAnomalyInvestigationAnomalyFindingType) GetAllowedValues() []TimeseriesAnomalyInvestigationAnomalyFindingType {
	return allowedTimeseriesAnomalyInvestigationAnomalyFindingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesAnomalyInvestigationAnomalyFindingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesAnomalyInvestigationAnomalyFindingType(value)
	return nil
}

// NewTimeseriesAnomalyInvestigationAnomalyFindingTypeFromValue returns a pointer to a valid TimeseriesAnomalyInvestigationAnomalyFindingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesAnomalyInvestigationAnomalyFindingTypeFromValue(v string) (*TimeseriesAnomalyInvestigationAnomalyFindingType, error) {
	ev := TimeseriesAnomalyInvestigationAnomalyFindingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesAnomalyInvestigationAnomalyFindingType: valid values are %v", v, allowedTimeseriesAnomalyInvestigationAnomalyFindingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesAnomalyInvestigationAnomalyFindingType) IsValid() bool {
	for _, existing := range allowedTimeseriesAnomalyInvestigationAnomalyFindingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesAnomalyInvestigationAnomalyFindingType value.
func (v TimeseriesAnomalyInvestigationAnomalyFindingType) Ptr() *TimeseriesAnomalyInvestigationAnomalyFindingType {
	return &v
}
