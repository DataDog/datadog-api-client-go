// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DDSQLTimeseriesRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type DDSQLTimeseriesRequestResponseFormat string

// List of DDSQLTimeseriesRequestResponseFormat.
const (
	DDSQLTIMESERIESREQUESTRESPONSEFORMAT_TIMESERIES DDSQLTimeseriesRequestResponseFormat = "timeseries"
)

var allowedDDSQLTimeseriesRequestResponseFormatEnumValues = []DDSQLTimeseriesRequestResponseFormat{
	DDSQLTIMESERIESREQUESTRESPONSEFORMAT_TIMESERIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DDSQLTimeseriesRequestResponseFormat) GetAllowedValues() []DDSQLTimeseriesRequestResponseFormat {
	return allowedDDSQLTimeseriesRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DDSQLTimeseriesRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DDSQLTimeseriesRequestResponseFormat(value)
	return nil
}

// NewDDSQLTimeseriesRequestResponseFormatFromValue returns a pointer to a valid DDSQLTimeseriesRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDDSQLTimeseriesRequestResponseFormatFromValue(v string) (*DDSQLTimeseriesRequestResponseFormat, error) {
	ev := DDSQLTimeseriesRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DDSQLTimeseriesRequestResponseFormat: valid values are %v", v, allowedDDSQLTimeseriesRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DDSQLTimeseriesRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedDDSQLTimeseriesRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DDSQLTimeseriesRequestResponseFormat value.
func (v DDSQLTimeseriesRequestResponseFormat) Ptr() *DDSQLTimeseriesRequestResponseFormat {
	return &v
}
