// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DDSQLTableRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type DDSQLTableRequestResponseFormat string

// List of DDSQLTableRequestResponseFormat.
const (
	DDSQLTABLEREQUESTRESPONSEFORMAT_SCALAR DDSQLTableRequestResponseFormat = "scalar"
)

var allowedDDSQLTableRequestResponseFormatEnumValues = []DDSQLTableRequestResponseFormat{
	DDSQLTABLEREQUESTRESPONSEFORMAT_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DDSQLTableRequestResponseFormat) GetAllowedValues() []DDSQLTableRequestResponseFormat {
	return allowedDDSQLTableRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DDSQLTableRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DDSQLTableRequestResponseFormat(value)
	return nil
}

// NewDDSQLTableRequestResponseFormatFromValue returns a pointer to a valid DDSQLTableRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDDSQLTableRequestResponseFormatFromValue(v string) (*DDSQLTableRequestResponseFormat, error) {
	ev := DDSQLTableRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DDSQLTableRequestResponseFormat: valid values are %v", v, allowedDDSQLTableRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DDSQLTableRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedDDSQLTableRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DDSQLTableRequestResponseFormat value.
func (v DDSQLTableRequestResponseFormat) Ptr() *DDSQLTableRequestResponseFormat {
	return &v
}
