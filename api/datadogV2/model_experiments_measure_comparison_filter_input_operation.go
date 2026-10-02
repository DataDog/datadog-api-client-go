// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsMeasureComparisonFilterInputOperation Comparison applied by this filter.
type ExperimentsMeasureComparisonFilterInputOperation string

// List of ExperimentsMeasureComparisonFilterInputOperation.
const (
	EXPERIMENTSMEASURECOMPARISONFILTERINPUTOPERATION_EQ    ExperimentsMeasureComparisonFilterInputOperation = "="
	EXPERIMENTSMEASURECOMPARISONFILTERINPUTOPERATION_NEQ   ExperimentsMeasureComparisonFilterInputOperation = "!="
	EXPERIMENTSMEASURECOMPARISONFILTERINPUTOPERATION_GT    ExperimentsMeasureComparisonFilterInputOperation = ">"
	EXPERIMENTSMEASURECOMPARISONFILTERINPUTOPERATION_GT_EQ ExperimentsMeasureComparisonFilterInputOperation = ">="
	EXPERIMENTSMEASURECOMPARISONFILTERINPUTOPERATION_LT    ExperimentsMeasureComparisonFilterInputOperation = "<"
	EXPERIMENTSMEASURECOMPARISONFILTERINPUTOPERATION_LT_EQ ExperimentsMeasureComparisonFilterInputOperation = "<="
)

var allowedExperimentsMeasureComparisonFilterInputOperationEnumValues = []ExperimentsMeasureComparisonFilterInputOperation{
	EXPERIMENTSMEASURECOMPARISONFILTERINPUTOPERATION_EQ,
	EXPERIMENTSMEASURECOMPARISONFILTERINPUTOPERATION_NEQ,
	EXPERIMENTSMEASURECOMPARISONFILTERINPUTOPERATION_GT,
	EXPERIMENTSMEASURECOMPARISONFILTERINPUTOPERATION_GT_EQ,
	EXPERIMENTSMEASURECOMPARISONFILTERINPUTOPERATION_LT,
	EXPERIMENTSMEASURECOMPARISONFILTERINPUTOPERATION_LT_EQ,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsMeasureComparisonFilterInputOperation) GetAllowedValues() []ExperimentsMeasureComparisonFilterInputOperation {
	return allowedExperimentsMeasureComparisonFilterInputOperationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsMeasureComparisonFilterInputOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsMeasureComparisonFilterInputOperation(value)
	return nil
}

// NewExperimentsMeasureComparisonFilterInputOperationFromValue returns a pointer to a valid ExperimentsMeasureComparisonFilterInputOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsMeasureComparisonFilterInputOperationFromValue(v string) (*ExperimentsMeasureComparisonFilterInputOperation, error) {
	ev := ExperimentsMeasureComparisonFilterInputOperation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsMeasureComparisonFilterInputOperation: valid values are %v", v, allowedExperimentsMeasureComparisonFilterInputOperationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsMeasureComparisonFilterInputOperation) IsValid() bool {
	for _, existing := range allowedExperimentsMeasureComparisonFilterInputOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsMeasureComparisonFilterInputOperation value.
func (v ExperimentsMeasureComparisonFilterInputOperation) Ptr() *ExperimentsMeasureComparisonFilterInputOperation {
	return &v
}
