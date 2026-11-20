// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableProjectionColumnOrder Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableProjectionColumnOrder string

// List of QueryTableProjectionColumnOrder.
const (
	QUERYTABLEPROJECTIONCOLUMNORDER_QUERY      QueryTableProjectionColumnOrder = "query"
	QUERYTABLEPROJECTIONCOLUMNORDER_DIMENSIONS QueryTableProjectionColumnOrder = "dimensions"
)

var allowedQueryTableProjectionColumnOrderEnumValues = []QueryTableProjectionColumnOrder{
	QUERYTABLEPROJECTIONCOLUMNORDER_QUERY,
	QUERYTABLEPROJECTIONCOLUMNORDER_DIMENSIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryTableProjectionColumnOrder) GetAllowedValues() []QueryTableProjectionColumnOrder {
	return allowedQueryTableProjectionColumnOrderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryTableProjectionColumnOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryTableProjectionColumnOrder(value)
	return nil
}

// NewQueryTableProjectionColumnOrderFromValue returns a pointer to a valid QueryTableProjectionColumnOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryTableProjectionColumnOrderFromValue(v string) (*QueryTableProjectionColumnOrder, error) {
	ev := QueryTableProjectionColumnOrder(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryTableProjectionColumnOrder: valid values are %v", v, allowedQueryTableProjectionColumnOrderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryTableProjectionColumnOrder) IsValid() bool {
	for _, existing := range allowedQueryTableProjectionColumnOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryTableProjectionColumnOrder value.
func (v QueryTableProjectionColumnOrder) Ptr() *QueryTableProjectionColumnOrder {
	return &v
}
