// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPropertyNullFilterInputOperation Comparison applied by this filter.
type ExperimentsPropertyNullFilterInputOperation string

// List of ExperimentsPropertyNullFilterInputOperation.
const (
	EXPERIMENTSPROPERTYNULLFILTERINPUTOPERATION_IS_NULL     ExperimentsPropertyNullFilterInputOperation = "IS_NULL"
	EXPERIMENTSPROPERTYNULLFILTERINPUTOPERATION_IS_NOT_NULL ExperimentsPropertyNullFilterInputOperation = "IS_NOT_NULL"
)

var allowedExperimentsPropertyNullFilterInputOperationEnumValues = []ExperimentsPropertyNullFilterInputOperation{
	EXPERIMENTSPROPERTYNULLFILTERINPUTOPERATION_IS_NULL,
	EXPERIMENTSPROPERTYNULLFILTERINPUTOPERATION_IS_NOT_NULL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsPropertyNullFilterInputOperation) GetAllowedValues() []ExperimentsPropertyNullFilterInputOperation {
	return allowedExperimentsPropertyNullFilterInputOperationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsPropertyNullFilterInputOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsPropertyNullFilterInputOperation(value)
	return nil
}

// NewExperimentsPropertyNullFilterInputOperationFromValue returns a pointer to a valid ExperimentsPropertyNullFilterInputOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsPropertyNullFilterInputOperationFromValue(v string) (*ExperimentsPropertyNullFilterInputOperation, error) {
	ev := ExperimentsPropertyNullFilterInputOperation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsPropertyNullFilterInputOperation: valid values are %v", v, allowedExperimentsPropertyNullFilterInputOperationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsPropertyNullFilterInputOperation) IsValid() bool {
	for _, existing := range allowedExperimentsPropertyNullFilterInputOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsPropertyNullFilterInputOperation value.
func (v ExperimentsPropertyNullFilterInputOperation) Ptr() *ExperimentsPropertyNullFilterInputOperation {
	return &v
}
