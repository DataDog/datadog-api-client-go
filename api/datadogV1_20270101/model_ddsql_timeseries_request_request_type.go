// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DDSQLTimeseriesRequestRequestType Test-only generated schema for the dashboard versioning proof of concept.
type DDSQLTimeseriesRequestRequestType string

// List of DDSQLTimeseriesRequestRequestType.
const (
	DDSQLTIMESERIESREQUESTREQUESTTYPE_DDSQL DDSQLTimeseriesRequestRequestType = "ddsql"
)

var allowedDDSQLTimeseriesRequestRequestTypeEnumValues = []DDSQLTimeseriesRequestRequestType{
	DDSQLTIMESERIESREQUESTREQUESTTYPE_DDSQL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DDSQLTimeseriesRequestRequestType) GetAllowedValues() []DDSQLTimeseriesRequestRequestType {
	return allowedDDSQLTimeseriesRequestRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DDSQLTimeseriesRequestRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DDSQLTimeseriesRequestRequestType(value)
	return nil
}

// NewDDSQLTimeseriesRequestRequestTypeFromValue returns a pointer to a valid DDSQLTimeseriesRequestRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDDSQLTimeseriesRequestRequestTypeFromValue(v string) (*DDSQLTimeseriesRequestRequestType, error) {
	ev := DDSQLTimeseriesRequestRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DDSQLTimeseriesRequestRequestType: valid values are %v", v, allowedDDSQLTimeseriesRequestRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DDSQLTimeseriesRequestRequestType) IsValid() bool {
	for _, existing := range allowedDDSQLTimeseriesRequestRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DDSQLTimeseriesRequestRequestType value.
func (v DDSQLTimeseriesRequestRequestType) Ptr() *DDSQLTimeseriesRequestRequestType {
	return &v
}
