// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonthlyResourcesView Whether to return billable resources or the full monthly resource breakdown.
type MonthlyResourcesView string

// List of MonthlyResourcesView.
const (
	MONTHLYRESOURCESVIEW_BILLABLE MonthlyResourcesView = "billable"
	MONTHLYRESOURCESVIEW_ALL      MonthlyResourcesView = "all"
)

var allowedMonthlyResourcesViewEnumValues = []MonthlyResourcesView{
	MONTHLYRESOURCESVIEW_BILLABLE,
	MONTHLYRESOURCESVIEW_ALL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonthlyResourcesView) GetAllowedValues() []MonthlyResourcesView {
	return allowedMonthlyResourcesViewEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonthlyResourcesView) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonthlyResourcesView(value)
	return nil
}

// NewMonthlyResourcesViewFromValue returns a pointer to a valid MonthlyResourcesView
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonthlyResourcesViewFromValue(v string) (*MonthlyResourcesView, error) {
	ev := MonthlyResourcesView(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonthlyResourcesView: valid values are %v", v, allowedMonthlyResourcesViewEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonthlyResourcesView) IsValid() bool {
	for _, existing := range allowedMonthlyResourcesViewEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonthlyResourcesView value.
func (v MonthlyResourcesView) Ptr() *MonthlyResourcesView {
	return &v
}
