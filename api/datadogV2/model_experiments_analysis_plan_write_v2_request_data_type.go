// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsAnalysisPlanWriteV2RequestDataType Analysis plans resource type.
type ExperimentsAnalysisPlanWriteV2RequestDataType string

// List of ExperimentsAnalysisPlanWriteV2RequestDataType.
const (
	EXPERIMENTSANALYSISPLANWRITEV2REQUESTDATATYPE_ANALYSIS_PLANS ExperimentsAnalysisPlanWriteV2RequestDataType = "analysis-plans"
)

var allowedExperimentsAnalysisPlanWriteV2RequestDataTypeEnumValues = []ExperimentsAnalysisPlanWriteV2RequestDataType{
	EXPERIMENTSANALYSISPLANWRITEV2REQUESTDATATYPE_ANALYSIS_PLANS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsAnalysisPlanWriteV2RequestDataType) GetAllowedValues() []ExperimentsAnalysisPlanWriteV2RequestDataType {
	return allowedExperimentsAnalysisPlanWriteV2RequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsAnalysisPlanWriteV2RequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsAnalysisPlanWriteV2RequestDataType(value)
	return nil
}

// NewExperimentsAnalysisPlanWriteV2RequestDataTypeFromValue returns a pointer to a valid ExperimentsAnalysisPlanWriteV2RequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsAnalysisPlanWriteV2RequestDataTypeFromValue(v string) (*ExperimentsAnalysisPlanWriteV2RequestDataType, error) {
	ev := ExperimentsAnalysisPlanWriteV2RequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsAnalysisPlanWriteV2RequestDataType: valid values are %v", v, allowedExperimentsAnalysisPlanWriteV2RequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsAnalysisPlanWriteV2RequestDataType) IsValid() bool {
	for _, existing := range allowedExperimentsAnalysisPlanWriteV2RequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsAnalysisPlanWriteV2RequestDataType value.
func (v ExperimentsAnalysisPlanWriteV2RequestDataType) Ptr() *ExperimentsAnalysisPlanWriteV2RequestDataType {
	return &v
}
