// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionDataQualityDetectionType How a `data-jobs alert` job duration monitor detects a breach. Must match the
// `.detection()` call in the monitor query. Required for job duration monitors.
type MonitorFormulaAndFunctionDataQualityDetectionType string

// List of MonitorFormulaAndFunctionDataQualityDetectionType.
const (
	MONITORFORMULAANDFUNCTIONDATAQUALITYDETECTIONTYPE_STATIC  MonitorFormulaAndFunctionDataQualityDetectionType = "static"
	MONITORFORMULAANDFUNCTIONDATAQUALITYDETECTIONTYPE_ANOMALY MonitorFormulaAndFunctionDataQualityDetectionType = "anomaly"
)

var allowedMonitorFormulaAndFunctionDataQualityDetectionTypeEnumValues = []MonitorFormulaAndFunctionDataQualityDetectionType{
	MONITORFORMULAANDFUNCTIONDATAQUALITYDETECTIONTYPE_STATIC,
	MONITORFORMULAANDFUNCTIONDATAQUALITYDETECTIONTYPE_ANOMALY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorFormulaAndFunctionDataQualityDetectionType) GetAllowedValues() []MonitorFormulaAndFunctionDataQualityDetectionType {
	return allowedMonitorFormulaAndFunctionDataQualityDetectionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorFormulaAndFunctionDataQualityDetectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorFormulaAndFunctionDataQualityDetectionType(value)
	return nil
}

// NewMonitorFormulaAndFunctionDataQualityDetectionTypeFromValue returns a pointer to a valid MonitorFormulaAndFunctionDataQualityDetectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorFormulaAndFunctionDataQualityDetectionTypeFromValue(v string) (*MonitorFormulaAndFunctionDataQualityDetectionType, error) {
	ev := MonitorFormulaAndFunctionDataQualityDetectionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorFormulaAndFunctionDataQualityDetectionType: valid values are %v", v, allowedMonitorFormulaAndFunctionDataQualityDetectionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorFormulaAndFunctionDataQualityDetectionType) IsValid() bool {
	for _, existing := range allowedMonitorFormulaAndFunctionDataQualityDetectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorFormulaAndFunctionDataQualityDetectionType value.
func (v MonitorFormulaAndFunctionDataQualityDetectionType) Ptr() *MonitorFormulaAndFunctionDataQualityDetectionType {
	return &v
}
