// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesWidgetAnomalyDetectionSensitivity Sensitivity level for anomaly detection. Use `never_detect` to disable anomaly detection.
type TimeseriesWidgetAnomalyDetectionSensitivity string

// List of TimeseriesWidgetAnomalyDetectionSensitivity.
const (
	TIMESERIESWIDGETANOMALYDETECTIONSENSITIVITY_NEVER_DETECT TimeseriesWidgetAnomalyDetectionSensitivity = "never_detect"
)

var allowedTimeseriesWidgetAnomalyDetectionSensitivityEnumValues = []TimeseriesWidgetAnomalyDetectionSensitivity{
	TIMESERIESWIDGETANOMALYDETECTIONSENSITIVITY_NEVER_DETECT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesWidgetAnomalyDetectionSensitivity) GetAllowedValues() []TimeseriesWidgetAnomalyDetectionSensitivity {
	return allowedTimeseriesWidgetAnomalyDetectionSensitivityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesWidgetAnomalyDetectionSensitivity) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesWidgetAnomalyDetectionSensitivity(value)
	return nil
}

// NewTimeseriesWidgetAnomalyDetectionSensitivityFromValue returns a pointer to a valid TimeseriesWidgetAnomalyDetectionSensitivity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesWidgetAnomalyDetectionSensitivityFromValue(v string) (*TimeseriesWidgetAnomalyDetectionSensitivity, error) {
	ev := TimeseriesWidgetAnomalyDetectionSensitivity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesWidgetAnomalyDetectionSensitivity: valid values are %v", v, allowedTimeseriesWidgetAnomalyDetectionSensitivityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesWidgetAnomalyDetectionSensitivity) IsValid() bool {
	for _, existing := range allowedTimeseriesWidgetAnomalyDetectionSensitivityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesWidgetAnomalyDetectionSensitivity value.
func (v TimeseriesWidgetAnomalyDetectionSensitivity) Ptr() *TimeseriesWidgetAnomalyDetectionSensitivity {
	return &v
}
