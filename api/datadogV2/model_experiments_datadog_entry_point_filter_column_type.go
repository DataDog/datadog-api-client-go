// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsDatadogEntryPointFilterColumnType Data type of the column evaluated by the entry-point filter.
type ExperimentsDatadogEntryPointFilterColumnType string

// List of ExperimentsDatadogEntryPointFilterColumnType.
const (
	EXPERIMENTSDATADOGENTRYPOINTFILTERCOLUMNTYPE_VARCHAR       ExperimentsDatadogEntryPointFilterColumnType = "varchar"
	EXPERIMENTSDATADOGENTRYPOINTFILTERCOLUMNTYPE_INT           ExperimentsDatadogEntryPointFilterColumnType = "int"
	EXPERIMENTSDATADOGENTRYPOINTFILTERCOLUMNTYPE_DOUBLE        ExperimentsDatadogEntryPointFilterColumnType = "double"
	EXPERIMENTSDATADOGENTRYPOINTFILTERCOLUMNTYPE_BOOLEAN       ExperimentsDatadogEntryPointFilterColumnType = "boolean"
	EXPERIMENTSDATADOGENTRYPOINTFILTERCOLUMNTYPE_VARCHAR_ARRAY ExperimentsDatadogEntryPointFilterColumnType = "varchar_array"
	EXPERIMENTSDATADOGENTRYPOINTFILTERCOLUMNTYPE_INT_ARRAY     ExperimentsDatadogEntryPointFilterColumnType = "int_array"
	EXPERIMENTSDATADOGENTRYPOINTFILTERCOLUMNTYPE_DOUBLE_ARRAY  ExperimentsDatadogEntryPointFilterColumnType = "double_array"
	EXPERIMENTSDATADOGENTRYPOINTFILTERCOLUMNTYPE_RAW           ExperimentsDatadogEntryPointFilterColumnType = "raw"
)

var allowedExperimentsDatadogEntryPointFilterColumnTypeEnumValues = []ExperimentsDatadogEntryPointFilterColumnType{
	EXPERIMENTSDATADOGENTRYPOINTFILTERCOLUMNTYPE_VARCHAR,
	EXPERIMENTSDATADOGENTRYPOINTFILTERCOLUMNTYPE_INT,
	EXPERIMENTSDATADOGENTRYPOINTFILTERCOLUMNTYPE_DOUBLE,
	EXPERIMENTSDATADOGENTRYPOINTFILTERCOLUMNTYPE_BOOLEAN,
	EXPERIMENTSDATADOGENTRYPOINTFILTERCOLUMNTYPE_VARCHAR_ARRAY,
	EXPERIMENTSDATADOGENTRYPOINTFILTERCOLUMNTYPE_INT_ARRAY,
	EXPERIMENTSDATADOGENTRYPOINTFILTERCOLUMNTYPE_DOUBLE_ARRAY,
	EXPERIMENTSDATADOGENTRYPOINTFILTERCOLUMNTYPE_RAW,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsDatadogEntryPointFilterColumnType) GetAllowedValues() []ExperimentsDatadogEntryPointFilterColumnType {
	return allowedExperimentsDatadogEntryPointFilterColumnTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsDatadogEntryPointFilterColumnType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsDatadogEntryPointFilterColumnType(value)
	return nil
}

// NewExperimentsDatadogEntryPointFilterColumnTypeFromValue returns a pointer to a valid ExperimentsDatadogEntryPointFilterColumnType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsDatadogEntryPointFilterColumnTypeFromValue(v string) (*ExperimentsDatadogEntryPointFilterColumnType, error) {
	ev := ExperimentsDatadogEntryPointFilterColumnType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsDatadogEntryPointFilterColumnType: valid values are %v", v, allowedExperimentsDatadogEntryPointFilterColumnTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsDatadogEntryPointFilterColumnType) IsValid() bool {
	for _, existing := range allowedExperimentsDatadogEntryPointFilterColumnTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsDatadogEntryPointFilterColumnType value.
func (v ExperimentsDatadogEntryPointFilterColumnType) Ptr() *ExperimentsDatadogEntryPointFilterColumnType {
	return &v
}
