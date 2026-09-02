// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionDataQualityDiffType How the difference between the source and target measures is computed.
// `absolute` subtracts the two values, `diff_percent` expresses the difference
// as a percentage of the source value.
type MonitorFormulaAndFunctionDataQualityDiffType string

// List of MonitorFormulaAndFunctionDataQualityDiffType.
const (
	MONITORFORMULAANDFUNCTIONDATAQUALITYDIFFTYPE_ABSOLUTE     MonitorFormulaAndFunctionDataQualityDiffType = "absolute"
	MONITORFORMULAANDFUNCTIONDATAQUALITYDIFFTYPE_DIFF_PERCENT MonitorFormulaAndFunctionDataQualityDiffType = "diff_percent"
)

var allowedMonitorFormulaAndFunctionDataQualityDiffTypeEnumValues = []MonitorFormulaAndFunctionDataQualityDiffType{
	MONITORFORMULAANDFUNCTIONDATAQUALITYDIFFTYPE_ABSOLUTE,
	MONITORFORMULAANDFUNCTIONDATAQUALITYDIFFTYPE_DIFF_PERCENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorFormulaAndFunctionDataQualityDiffType) GetAllowedValues() []MonitorFormulaAndFunctionDataQualityDiffType {
	return allowedMonitorFormulaAndFunctionDataQualityDiffTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorFormulaAndFunctionDataQualityDiffType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorFormulaAndFunctionDataQualityDiffType(value)
	return nil
}

// NewMonitorFormulaAndFunctionDataQualityDiffTypeFromValue returns a pointer to a valid MonitorFormulaAndFunctionDataQualityDiffType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorFormulaAndFunctionDataQualityDiffTypeFromValue(v string) (*MonitorFormulaAndFunctionDataQualityDiffType, error) {
	ev := MonitorFormulaAndFunctionDataQualityDiffType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorFormulaAndFunctionDataQualityDiffType: valid values are %v", v, allowedMonitorFormulaAndFunctionDataQualityDiffTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorFormulaAndFunctionDataQualityDiffType) IsValid() bool {
	for _, existing := range allowedMonitorFormulaAndFunctionDataQualityDiffTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorFormulaAndFunctionDataQualityDiffType value.
func (v MonitorFormulaAndFunctionDataQualityDiffType) Ptr() *MonitorFormulaAndFunctionDataQualityDiffType {
	return &v
}
