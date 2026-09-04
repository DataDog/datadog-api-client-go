// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationInfluentialTagFindingType Finding category for an influential tag.
type TimeseriesAnomalyInvestigationInfluentialTagFindingType string

// List of TimeseriesAnomalyInvestigationInfluentialTagFindingType.
const (
	TIMESERIESANOMALYINVESTIGATIONINFLUENTIALTAGFINDINGTYPE_INFLUENTIAL_TAG TimeseriesAnomalyInvestigationInfluentialTagFindingType = "influential_tag"
)

var allowedTimeseriesAnomalyInvestigationInfluentialTagFindingTypeEnumValues = []TimeseriesAnomalyInvestigationInfluentialTagFindingType{
	TIMESERIESANOMALYINVESTIGATIONINFLUENTIALTAGFINDINGTYPE_INFLUENTIAL_TAG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesAnomalyInvestigationInfluentialTagFindingType) GetAllowedValues() []TimeseriesAnomalyInvestigationInfluentialTagFindingType {
	return allowedTimeseriesAnomalyInvestigationInfluentialTagFindingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesAnomalyInvestigationInfluentialTagFindingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesAnomalyInvestigationInfluentialTagFindingType(value)
	return nil
}

// NewTimeseriesAnomalyInvestigationInfluentialTagFindingTypeFromValue returns a pointer to a valid TimeseriesAnomalyInvestigationInfluentialTagFindingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesAnomalyInvestigationInfluentialTagFindingTypeFromValue(v string) (*TimeseriesAnomalyInvestigationInfluentialTagFindingType, error) {
	ev := TimeseriesAnomalyInvestigationInfluentialTagFindingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesAnomalyInvestigationInfluentialTagFindingType: valid values are %v", v, allowedTimeseriesAnomalyInvestigationInfluentialTagFindingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesAnomalyInvestigationInfluentialTagFindingType) IsValid() bool {
	for _, existing := range allowedTimeseriesAnomalyInvestigationInfluentialTagFindingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesAnomalyInvestigationInfluentialTagFindingType value.
func (v TimeseriesAnomalyInvestigationInfluentialTagFindingType) Ptr() *TimeseriesAnomalyInvestigationInfluentialTagFindingType {
	return &v
}
