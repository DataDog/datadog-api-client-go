// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrderByOptions Test-only generated schema for the dashboard versioning proof of concept.
type OrderByOptions string

// List of OrderByOptions.
const (
	ORDERBYOPTIONS_CHANGE  OrderByOptions = "change"
	ORDERBYOPTIONS_NAME    OrderByOptions = "name"
	ORDERBYOPTIONS_PRESENT OrderByOptions = "present"
	ORDERBYOPTIONS_PAST    OrderByOptions = "past"
)

var allowedOrderByOptionsEnumValues = []OrderByOptions{
	ORDERBYOPTIONS_CHANGE,
	ORDERBYOPTIONS_NAME,
	ORDERBYOPTIONS_PRESENT,
	ORDERBYOPTIONS_PAST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrderByOptions) GetAllowedValues() []OrderByOptions {
	return allowedOrderByOptionsEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrderByOptions) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrderByOptions(value)
	return nil
}

// NewOrderByOptionsFromValue returns a pointer to a valid OrderByOptions
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrderByOptionsFromValue(v string) (*OrderByOptions, error) {
	ev := OrderByOptions(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrderByOptions: valid values are %v", v, allowedOrderByOptionsEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrderByOptions) IsValid() bool {
	for _, existing := range allowedOrderByOptionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderByOptions value.
func (v OrderByOptions) Ptr() *OrderByOptions {
	return &v
}
