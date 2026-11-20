// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaDatasetResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type FormulaDatasetResponseFormat string

// List of FormulaDatasetResponseFormat.
const (
	FORMULADATASETRESPONSEFORMAT_TIMESERIES FormulaDatasetResponseFormat = "timeseries"
	FORMULADATASETRESPONSEFORMAT_SCALAR     FormulaDatasetResponseFormat = "scalar"
)

var allowedFormulaDatasetResponseFormatEnumValues = []FormulaDatasetResponseFormat{
	FORMULADATASETRESPONSEFORMAT_TIMESERIES,
	FORMULADATASETRESPONSEFORMAT_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaDatasetResponseFormat) GetAllowedValues() []FormulaDatasetResponseFormat {
	return allowedFormulaDatasetResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaDatasetResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaDatasetResponseFormat(value)
	return nil
}

// NewFormulaDatasetResponseFormatFromValue returns a pointer to a valid FormulaDatasetResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaDatasetResponseFormatFromValue(v string) (*FormulaDatasetResponseFormat, error) {
	ev := FormulaDatasetResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaDatasetResponseFormat: valid values are %v", v, allowedFormulaDatasetResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaDatasetResponseFormat) IsValid() bool {
	for _, existing := range allowedFormulaDatasetResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaDatasetResponseFormat value.
func (v FormulaDatasetResponseFormat) Ptr() *FormulaDatasetResponseFormat {
	return &v
}
