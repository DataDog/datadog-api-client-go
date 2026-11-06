// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionMetricQueryDataSource Data source for metrics queries.
type MonitorFormulaAndFunctionMetricQueryDataSource string

// List of MonitorFormulaAndFunctionMetricQueryDataSource.
const (
	MONITORFORMULAANDFUNCTIONMETRICQUERYDATASOURCE_METRICS MonitorFormulaAndFunctionMetricQueryDataSource = "metrics"
)

var allowedMonitorFormulaAndFunctionMetricQueryDataSourceEnumValues = []MonitorFormulaAndFunctionMetricQueryDataSource{
	MONITORFORMULAANDFUNCTIONMETRICQUERYDATASOURCE_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorFormulaAndFunctionMetricQueryDataSource) GetAllowedValues() []MonitorFormulaAndFunctionMetricQueryDataSource {
	return allowedMonitorFormulaAndFunctionMetricQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorFormulaAndFunctionMetricQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorFormulaAndFunctionMetricQueryDataSource(value)
	return nil
}

// NewMonitorFormulaAndFunctionMetricQueryDataSourceFromValue returns a pointer to a valid MonitorFormulaAndFunctionMetricQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorFormulaAndFunctionMetricQueryDataSourceFromValue(v string) (*MonitorFormulaAndFunctionMetricQueryDataSource, error) {
	ev := MonitorFormulaAndFunctionMetricQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorFormulaAndFunctionMetricQueryDataSource: valid values are %v", v, allowedMonitorFormulaAndFunctionMetricQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorFormulaAndFunctionMetricQueryDataSource) IsValid() bool {
	for _, existing := range allowedMonitorFormulaAndFunctionMetricQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorFormulaAndFunctionMetricQueryDataSource value.
func (v MonitorFormulaAndFunctionMetricQueryDataSource) Ptr() *MonitorFormulaAndFunctionMetricQueryDataSource {
	return &v
}
