// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScalarFormulaRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type ScalarFormulaRequestResponseFormat string

// List of ScalarFormulaRequestResponseFormat.
const (
	SCALARFORMULAREQUESTRESPONSEFORMAT_SCALAR ScalarFormulaRequestResponseFormat = "scalar"
)

var allowedScalarFormulaRequestResponseFormatEnumValues = []ScalarFormulaRequestResponseFormat{
	SCALARFORMULAREQUESTRESPONSEFORMAT_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScalarFormulaRequestResponseFormat) GetAllowedValues() []ScalarFormulaRequestResponseFormat {
	return allowedScalarFormulaRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScalarFormulaRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScalarFormulaRequestResponseFormat(value)
	return nil
}

// NewScalarFormulaRequestResponseFormatFromValue returns a pointer to a valid ScalarFormulaRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScalarFormulaRequestResponseFormatFromValue(v string) (*ScalarFormulaRequestResponseFormat, error) {
	ev := ScalarFormulaRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScalarFormulaRequestResponseFormat: valid values are %v", v, allowedScalarFormulaRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScalarFormulaRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedScalarFormulaRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScalarFormulaRequestResponseFormat value.
func (v ScalarFormulaRequestResponseFormat) Ptr() *ScalarFormulaRequestResponseFormat {
	return &v
}
