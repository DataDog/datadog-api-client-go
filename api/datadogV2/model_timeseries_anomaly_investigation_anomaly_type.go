// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationAnomalyType Direction of an anomaly relative to its expected range.
type TimeseriesAnomalyInvestigationAnomalyType string

// List of TimeseriesAnomalyInvestigationAnomalyType.
const (
	TIMESERIESANOMALYINVESTIGATIONANOMALYTYPE_SPIKE TimeseriesAnomalyInvestigationAnomalyType = "spike"
	TIMESERIESANOMALYINVESTIGATIONANOMALYTYPE_DIP   TimeseriesAnomalyInvestigationAnomalyType = "dip"
)

var allowedTimeseriesAnomalyInvestigationAnomalyTypeEnumValues = []TimeseriesAnomalyInvestigationAnomalyType{
	TIMESERIESANOMALYINVESTIGATIONANOMALYTYPE_SPIKE,
	TIMESERIESANOMALYINVESTIGATIONANOMALYTYPE_DIP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesAnomalyInvestigationAnomalyType) GetAllowedValues() []TimeseriesAnomalyInvestigationAnomalyType {
	return allowedTimeseriesAnomalyInvestigationAnomalyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesAnomalyInvestigationAnomalyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesAnomalyInvestigationAnomalyType(value)
	return nil
}

// NewTimeseriesAnomalyInvestigationAnomalyTypeFromValue returns a pointer to a valid TimeseriesAnomalyInvestigationAnomalyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesAnomalyInvestigationAnomalyTypeFromValue(v string) (*TimeseriesAnomalyInvestigationAnomalyType, error) {
	ev := TimeseriesAnomalyInvestigationAnomalyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesAnomalyInvestigationAnomalyType: valid values are %v", v, allowedTimeseriesAnomalyInvestigationAnomalyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesAnomalyInvestigationAnomalyType) IsValid() bool {
	for _, existing := range allowedTimeseriesAnomalyInvestigationAnomalyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesAnomalyInvestigationAnomalyType value.
func (v TimeseriesAnomalyInvestigationAnomalyType) Ptr() *TimeseriesAnomalyInvestigationAnomalyType {
	return &v
}
