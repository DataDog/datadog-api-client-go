// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationInfluenceType Kind of influence a tag has on a series.
type TimeseriesAnomalyInvestigationInfluenceType string

// List of TimeseriesAnomalyInvestigationInfluenceType.
const (
	TIMESERIESANOMALYINVESTIGATIONINFLUENCETYPE_SHAPE TimeseriesAnomalyInvestigationInfluenceType = "shape"
	TIMESERIESANOMALYINVESTIGATIONINFLUENCETYPE_VALUE TimeseriesAnomalyInvestigationInfluenceType = "value"
)

var allowedTimeseriesAnomalyInvestigationInfluenceTypeEnumValues = []TimeseriesAnomalyInvestigationInfluenceType{
	TIMESERIESANOMALYINVESTIGATIONINFLUENCETYPE_SHAPE,
	TIMESERIESANOMALYINVESTIGATIONINFLUENCETYPE_VALUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesAnomalyInvestigationInfluenceType) GetAllowedValues() []TimeseriesAnomalyInvestigationInfluenceType {
	return allowedTimeseriesAnomalyInvestigationInfluenceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesAnomalyInvestigationInfluenceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesAnomalyInvestigationInfluenceType(value)
	return nil
}

// NewTimeseriesAnomalyInvestigationInfluenceTypeFromValue returns a pointer to a valid TimeseriesAnomalyInvestigationInfluenceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesAnomalyInvestigationInfluenceTypeFromValue(v string) (*TimeseriesAnomalyInvestigationInfluenceType, error) {
	ev := TimeseriesAnomalyInvestigationInfluenceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesAnomalyInvestigationInfluenceType: valid values are %v", v, allowedTimeseriesAnomalyInvestigationInfluenceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesAnomalyInvestigationInfluenceType) IsValid() bool {
	for _, existing := range allowedTimeseriesAnomalyInvestigationInfluenceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesAnomalyInvestigationInfluenceType value.
func (v TimeseriesAnomalyInvestigationInfluenceType) Ptr() *TimeseriesAnomalyInvestigationInfluenceType {
	return &v
}
