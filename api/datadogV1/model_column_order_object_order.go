// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ColumnOrderObjectOrder Order object
type ColumnOrderObjectOrder string

// List of ColumnOrderObjectOrder.
const (
	COLUMNORDEROBJECTORDER_ASC  ColumnOrderObjectOrder = "asc"
	COLUMNORDEROBJECTORDER_DESC ColumnOrderObjectOrder = "desc"
)

var allowedColumnOrderObjectOrderEnumValues = []ColumnOrderObjectOrder{
	COLUMNORDEROBJECTORDER_ASC,
	COLUMNORDEROBJECTORDER_DESC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ColumnOrderObjectOrder) GetAllowedValues() []ColumnOrderObjectOrder {
	return allowedColumnOrderObjectOrderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ColumnOrderObjectOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ColumnOrderObjectOrder(value)
	return nil
}

// NewColumnOrderObjectOrderFromValue returns a pointer to a valid ColumnOrderObjectOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewColumnOrderObjectOrderFromValue(v string) (*ColumnOrderObjectOrder, error) {
	ev := ColumnOrderObjectOrder(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ColumnOrderObjectOrder: valid values are %v", v, allowedColumnOrderObjectOrderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ColumnOrderObjectOrder) IsValid() bool {
	for _, existing := range allowedColumnOrderObjectOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ColumnOrderObjectOrder value.
func (v ColumnOrderObjectOrder) Ptr() *ColumnOrderObjectOrder {
	return &v
}
