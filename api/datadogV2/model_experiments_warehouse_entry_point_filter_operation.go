// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsWarehouseEntryPointFilterOperation Comparison applied by the warehouse entry-point filter.
type ExperimentsWarehouseEntryPointFilterOperation string

// List of ExperimentsWarehouseEntryPointFilterOperation.
const (
	EXPERIMENTSWAREHOUSEENTRYPOINTFILTEROPERATION_IS     ExperimentsWarehouseEntryPointFilterOperation = "IS"
	EXPERIMENTSWAREHOUSEENTRYPOINTFILTEROPERATION_IS_NOT ExperimentsWarehouseEntryPointFilterOperation = "IS_NOT"
)

var allowedExperimentsWarehouseEntryPointFilterOperationEnumValues = []ExperimentsWarehouseEntryPointFilterOperation{
	EXPERIMENTSWAREHOUSEENTRYPOINTFILTEROPERATION_IS,
	EXPERIMENTSWAREHOUSEENTRYPOINTFILTEROPERATION_IS_NOT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsWarehouseEntryPointFilterOperation) GetAllowedValues() []ExperimentsWarehouseEntryPointFilterOperation {
	return allowedExperimentsWarehouseEntryPointFilterOperationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsWarehouseEntryPointFilterOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsWarehouseEntryPointFilterOperation(value)
	return nil
}

// NewExperimentsWarehouseEntryPointFilterOperationFromValue returns a pointer to a valid ExperimentsWarehouseEntryPointFilterOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsWarehouseEntryPointFilterOperationFromValue(v string) (*ExperimentsWarehouseEntryPointFilterOperation, error) {
	ev := ExperimentsWarehouseEntryPointFilterOperation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsWarehouseEntryPointFilterOperation: valid values are %v", v, allowedExperimentsWarehouseEntryPointFilterOperationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsWarehouseEntryPointFilterOperation) IsValid() bool {
	for _, existing := range allowedExperimentsWarehouseEntryPointFilterOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsWarehouseEntryPointFilterOperation value.
func (v ExperimentsWarehouseEntryPointFilterOperation) Ptr() *ExperimentsWarehouseEntryPointFilterOperation {
	return &v
}
