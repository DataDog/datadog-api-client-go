// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationConfigurationSource Source of the anomaly detection configuration.
type TimeseriesAnomalyInvestigationConfigurationSource string

// List of TimeseriesAnomalyInvestigationConfigurationSource.
const (
	TIMESERIESANOMALYINVESTIGATIONCONFIGURATIONSOURCE_REQUEST_FORMULA           TimeseriesAnomalyInvestigationConfigurationSource = "request_formula"
	TIMESERIESANOMALYINVESTIGATIONCONFIGURATIONSOURCE_WATCHDOG_EXPLAINS_DEFAULT TimeseriesAnomalyInvestigationConfigurationSource = "watchdog_explains_default"
)

var allowedTimeseriesAnomalyInvestigationConfigurationSourceEnumValues = []TimeseriesAnomalyInvestigationConfigurationSource{
	TIMESERIESANOMALYINVESTIGATIONCONFIGURATIONSOURCE_REQUEST_FORMULA,
	TIMESERIESANOMALYINVESTIGATIONCONFIGURATIONSOURCE_WATCHDOG_EXPLAINS_DEFAULT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesAnomalyInvestigationConfigurationSource) GetAllowedValues() []TimeseriesAnomalyInvestigationConfigurationSource {
	return allowedTimeseriesAnomalyInvestigationConfigurationSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesAnomalyInvestigationConfigurationSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesAnomalyInvestigationConfigurationSource(value)
	return nil
}

// NewTimeseriesAnomalyInvestigationConfigurationSourceFromValue returns a pointer to a valid TimeseriesAnomalyInvestigationConfigurationSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesAnomalyInvestigationConfigurationSourceFromValue(v string) (*TimeseriesAnomalyInvestigationConfigurationSource, error) {
	ev := TimeseriesAnomalyInvestigationConfigurationSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesAnomalyInvestigationConfigurationSource: valid values are %v", v, allowedTimeseriesAnomalyInvestigationConfigurationSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesAnomalyInvestigationConfigurationSource) IsValid() bool {
	for _, existing := range allowedTimeseriesAnomalyInvestigationConfigurationSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesAnomalyInvestigationConfigurationSource value.
func (v TimeseriesAnomalyInvestigationConfigurationSource) Ptr() *TimeseriesAnomalyInvestigationConfigurationSource {
	return &v
}
