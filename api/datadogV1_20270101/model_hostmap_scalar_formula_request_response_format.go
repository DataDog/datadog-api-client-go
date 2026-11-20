// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostmapScalarFormulaRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type HostmapScalarFormulaRequestResponseFormat string

// List of HostmapScalarFormulaRequestResponseFormat.
const (
	HOSTMAPSCALARFORMULAREQUESTRESPONSEFORMAT_SCALAR HostmapScalarFormulaRequestResponseFormat = "scalar"
)

var allowedHostmapScalarFormulaRequestResponseFormatEnumValues = []HostmapScalarFormulaRequestResponseFormat{
	HOSTMAPSCALARFORMULAREQUESTRESPONSEFORMAT_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HostmapScalarFormulaRequestResponseFormat) GetAllowedValues() []HostmapScalarFormulaRequestResponseFormat {
	return allowedHostmapScalarFormulaRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HostmapScalarFormulaRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HostmapScalarFormulaRequestResponseFormat(value)
	return nil
}

// NewHostmapScalarFormulaRequestResponseFormatFromValue returns a pointer to a valid HostmapScalarFormulaRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHostmapScalarFormulaRequestResponseFormatFromValue(v string) (*HostmapScalarFormulaRequestResponseFormat, error) {
	ev := HostmapScalarFormulaRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HostmapScalarFormulaRequestResponseFormat: valid values are %v", v, allowedHostmapScalarFormulaRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HostmapScalarFormulaRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedHostmapScalarFormulaRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostmapScalarFormulaRequestResponseFormat value.
func (v HostmapScalarFormulaRequestResponseFormat) Ptr() *HostmapScalarFormulaRequestResponseFormat {
	return &v
}
