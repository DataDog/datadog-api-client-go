// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableProjectionUnlistedColumns Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableProjectionUnlistedColumns string

// List of QueryTableProjectionUnlistedColumns.
const (
	QUERYTABLEPROJECTIONUNLISTEDCOLUMNS_SHOW QueryTableProjectionUnlistedColumns = "show"
	QUERYTABLEPROJECTIONUNLISTEDCOLUMNS_HIDE QueryTableProjectionUnlistedColumns = "hide"
)

var allowedQueryTableProjectionUnlistedColumnsEnumValues = []QueryTableProjectionUnlistedColumns{
	QUERYTABLEPROJECTIONUNLISTEDCOLUMNS_SHOW,
	QUERYTABLEPROJECTIONUNLISTEDCOLUMNS_HIDE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryTableProjectionUnlistedColumns) GetAllowedValues() []QueryTableProjectionUnlistedColumns {
	return allowedQueryTableProjectionUnlistedColumnsEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryTableProjectionUnlistedColumns) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryTableProjectionUnlistedColumns(value)
	return nil
}

// NewQueryTableProjectionUnlistedColumnsFromValue returns a pointer to a valid QueryTableProjectionUnlistedColumns
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryTableProjectionUnlistedColumnsFromValue(v string) (*QueryTableProjectionUnlistedColumns, error) {
	ev := QueryTableProjectionUnlistedColumns(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryTableProjectionUnlistedColumns: valid values are %v", v, allowedQueryTableProjectionUnlistedColumnsEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryTableProjectionUnlistedColumns) IsValid() bool {
	for _, existing := range allowedQueryTableProjectionUnlistedColumnsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryTableProjectionUnlistedColumns value.
func (v QueryTableProjectionUnlistedColumns) Ptr() *QueryTableProjectionUnlistedColumns {
	return &v
}
