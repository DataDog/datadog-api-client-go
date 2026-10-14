// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsDatadogEntryPointFilterOperation Comparison applied by the Datadog entry-point filter.
type ExperimentsDatadogEntryPointFilterOperation string

// List of ExperimentsDatadogEntryPointFilterOperation.
const (
	EXPERIMENTSDATADOGENTRYPOINTFILTEROPERATION_EQ     ExperimentsDatadogEntryPointFilterOperation = "eq"
	EXPERIMENTSDATADOGENTRYPOINTFILTEROPERATION_IN     ExperimentsDatadogEntryPointFilterOperation = "in"
	EXPERIMENTSDATADOGENTRYPOINTFILTEROPERATION_NEQ    ExperimentsDatadogEntryPointFilterOperation = "neq"
	EXPERIMENTSDATADOGENTRYPOINTFILTEROPERATION_NOT_IN ExperimentsDatadogEntryPointFilterOperation = "not_in"
	EXPERIMENTSDATADOGENTRYPOINTFILTEROPERATION_GTE    ExperimentsDatadogEntryPointFilterOperation = "gte"
	EXPERIMENTSDATADOGENTRYPOINTFILTEROPERATION_LTE    ExperimentsDatadogEntryPointFilterOperation = "lte"
	EXPERIMENTSDATADOGENTRYPOINTFILTEROPERATION_GT     ExperimentsDatadogEntryPointFilterOperation = "gt"
	EXPERIMENTSDATADOGENTRYPOINTFILTEROPERATION_LT     ExperimentsDatadogEntryPointFilterOperation = "lt"
)

var allowedExperimentsDatadogEntryPointFilterOperationEnumValues = []ExperimentsDatadogEntryPointFilterOperation{
	EXPERIMENTSDATADOGENTRYPOINTFILTEROPERATION_EQ,
	EXPERIMENTSDATADOGENTRYPOINTFILTEROPERATION_IN,
	EXPERIMENTSDATADOGENTRYPOINTFILTEROPERATION_NEQ,
	EXPERIMENTSDATADOGENTRYPOINTFILTEROPERATION_NOT_IN,
	EXPERIMENTSDATADOGENTRYPOINTFILTEROPERATION_GTE,
	EXPERIMENTSDATADOGENTRYPOINTFILTEROPERATION_LTE,
	EXPERIMENTSDATADOGENTRYPOINTFILTEROPERATION_GT,
	EXPERIMENTSDATADOGENTRYPOINTFILTEROPERATION_LT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsDatadogEntryPointFilterOperation) GetAllowedValues() []ExperimentsDatadogEntryPointFilterOperation {
	return allowedExperimentsDatadogEntryPointFilterOperationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsDatadogEntryPointFilterOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsDatadogEntryPointFilterOperation(value)
	return nil
}

// NewExperimentsDatadogEntryPointFilterOperationFromValue returns a pointer to a valid ExperimentsDatadogEntryPointFilterOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsDatadogEntryPointFilterOperationFromValue(v string) (*ExperimentsDatadogEntryPointFilterOperation, error) {
	ev := ExperimentsDatadogEntryPointFilterOperation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsDatadogEntryPointFilterOperation: valid values are %v", v, allowedExperimentsDatadogEntryPointFilterOperationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsDatadogEntryPointFilterOperation) IsValid() bool {
	for _, existing := range allowedExperimentsDatadogEntryPointFilterOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsDatadogEntryPointFilterOperation value.
func (v ExperimentsDatadogEntryPointFilterOperation) Ptr() *ExperimentsDatadogEntryPointFilterOperation {
	return &v
}
