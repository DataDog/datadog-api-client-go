// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationDataSource Data source for an anomaly investigation query.
type TimeseriesAnomalyInvestigationDataSource string

// List of TimeseriesAnomalyInvestigationDataSource.
const (
	TIMESERIESANOMALYINVESTIGATIONDATASOURCE_METRICS TimeseriesAnomalyInvestigationDataSource = "metrics"
)

var allowedTimeseriesAnomalyInvestigationDataSourceEnumValues = []TimeseriesAnomalyInvestigationDataSource{
	TIMESERIESANOMALYINVESTIGATIONDATASOURCE_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesAnomalyInvestigationDataSource) GetAllowedValues() []TimeseriesAnomalyInvestigationDataSource {
	return allowedTimeseriesAnomalyInvestigationDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesAnomalyInvestigationDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesAnomalyInvestigationDataSource(value)
	return nil
}

// NewTimeseriesAnomalyInvestigationDataSourceFromValue returns a pointer to a valid TimeseriesAnomalyInvestigationDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesAnomalyInvestigationDataSourceFromValue(v string) (*TimeseriesAnomalyInvestigationDataSource, error) {
	ev := TimeseriesAnomalyInvestigationDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesAnomalyInvestigationDataSource: valid values are %v", v, allowedTimeseriesAnomalyInvestigationDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesAnomalyInvestigationDataSource) IsValid() bool {
	for _, existing := range allowedTimeseriesAnomalyInvestigationDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesAnomalyInvestigationDataSource value.
func (v TimeseriesAnomalyInvestigationDataSource) Ptr() *TimeseriesAnomalyInvestigationDataSource {
	return &v
}
