// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationMetaType Response metadata type for a timeseries anomaly investigation.
type TimeseriesAnomalyInvestigationMetaType string

// List of TimeseriesAnomalyInvestigationMetaType.
const (
	TIMESERIESANOMALYINVESTIGATIONMETATYPE_TIMESERIES_ANOMALY_INVESTIGATION TimeseriesAnomalyInvestigationMetaType = "timeseries_anomaly_investigation"
)

var allowedTimeseriesAnomalyInvestigationMetaTypeEnumValues = []TimeseriesAnomalyInvestigationMetaType{
	TIMESERIESANOMALYINVESTIGATIONMETATYPE_TIMESERIES_ANOMALY_INVESTIGATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesAnomalyInvestigationMetaType) GetAllowedValues() []TimeseriesAnomalyInvestigationMetaType {
	return allowedTimeseriesAnomalyInvestigationMetaTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesAnomalyInvestigationMetaType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesAnomalyInvestigationMetaType(value)
	return nil
}

// NewTimeseriesAnomalyInvestigationMetaTypeFromValue returns a pointer to a valid TimeseriesAnomalyInvestigationMetaType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesAnomalyInvestigationMetaTypeFromValue(v string) (*TimeseriesAnomalyInvestigationMetaType, error) {
	ev := TimeseriesAnomalyInvestigationMetaType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesAnomalyInvestigationMetaType: valid values are %v", v, allowedTimeseriesAnomalyInvestigationMetaTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesAnomalyInvestigationMetaType) IsValid() bool {
	for _, existing := range allowedTimeseriesAnomalyInvestigationMetaTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesAnomalyInvestigationMetaType value.
func (v TimeseriesAnomalyInvestigationMetaType) Ptr() *TimeseriesAnomalyInvestigationMetaType {
	return &v
}
