// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonthlyResourceType Type of a monthly resource usage record.
type MonthlyResourceType string

// List of MonthlyResourceType.
const (
	MONTHLYRESOURCETYPE_MONTHLY_RESOURCE MonthlyResourceType = "monthly_resource"
)

var allowedMonthlyResourceTypeEnumValues = []MonthlyResourceType{
	MONTHLYRESOURCETYPE_MONTHLY_RESOURCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonthlyResourceType) GetAllowedValues() []MonthlyResourceType {
	return allowedMonthlyResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonthlyResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonthlyResourceType(value)
	return nil
}

// NewMonthlyResourceTypeFromValue returns a pointer to a valid MonthlyResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonthlyResourceTypeFromValue(v string) (*MonthlyResourceType, error) {
	ev := MonthlyResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonthlyResourceType: valid values are %v", v, allowedMonthlyResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonthlyResourceType) IsValid() bool {
	for _, existing := range allowedMonthlyResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonthlyResourceType value.
func (v MonthlyResourceType) Ptr() *MonthlyResourceType {
	return &v
}
