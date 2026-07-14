// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PrintReportType JSON:API resource type for a print-only report.
type PrintReportType string

// List of PrintReportType.
const (
	PRINTREPORTTYPE_REPORT PrintReportType = "report"
)

var allowedPrintReportTypeEnumValues = []PrintReportType{
	PRINTREPORTTYPE_REPORT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PrintReportType) GetAllowedValues() []PrintReportType {
	return allowedPrintReportTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PrintReportType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PrintReportType(value)
	return nil
}

// NewPrintReportTypeFromValue returns a pointer to a valid PrintReportType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPrintReportTypeFromValue(v string) (*PrintReportType, error) {
	ev := PrintReportType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PrintReportType: valid values are %v", v, allowedPrintReportTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PrintReportType) IsValid() bool {
	for _, existing := range allowedPrintReportTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PrintReportType value.
func (v PrintReportType) Ptr() *PrintReportType {
	return &v
}
