// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionDataQualityDiffFunction Function applied to the measure before it is compared against the predicted bounds.
type MonitorFormulaAndFunctionDataQualityDiffFunction string

// List of MonitorFormulaAndFunctionDataQualityDiffFunction.
const (
	MONITORFORMULAANDFUNCTIONDATAQUALITYDIFFFUNCTION_DIFF         MonitorFormulaAndFunctionDataQualityDiffFunction = "DIFF"
	MONITORFORMULAANDFUNCTIONDATAQUALITYDIFFFUNCTION_DIFF_PERCENT MonitorFormulaAndFunctionDataQualityDiffFunction = "DIFF_PERCENT"
)

var allowedMonitorFormulaAndFunctionDataQualityDiffFunctionEnumValues = []MonitorFormulaAndFunctionDataQualityDiffFunction{
	MONITORFORMULAANDFUNCTIONDATAQUALITYDIFFFUNCTION_DIFF,
	MONITORFORMULAANDFUNCTIONDATAQUALITYDIFFFUNCTION_DIFF_PERCENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorFormulaAndFunctionDataQualityDiffFunction) GetAllowedValues() []MonitorFormulaAndFunctionDataQualityDiffFunction {
	return allowedMonitorFormulaAndFunctionDataQualityDiffFunctionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorFormulaAndFunctionDataQualityDiffFunction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorFormulaAndFunctionDataQualityDiffFunction(value)
	return nil
}

// NewMonitorFormulaAndFunctionDataQualityDiffFunctionFromValue returns a pointer to a valid MonitorFormulaAndFunctionDataQualityDiffFunction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorFormulaAndFunctionDataQualityDiffFunctionFromValue(v string) (*MonitorFormulaAndFunctionDataQualityDiffFunction, error) {
	ev := MonitorFormulaAndFunctionDataQualityDiffFunction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorFormulaAndFunctionDataQualityDiffFunction: valid values are %v", v, allowedMonitorFormulaAndFunctionDataQualityDiffFunctionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorFormulaAndFunctionDataQualityDiffFunction) IsValid() bool {
	for _, existing := range allowedMonitorFormulaAndFunctionDataQualityDiffFunctionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorFormulaAndFunctionDataQualityDiffFunction value.
func (v MonitorFormulaAndFunctionDataQualityDiffFunction) Ptr() *MonitorFormulaAndFunctionDataQualityDiffFunction {
	return &v
}
