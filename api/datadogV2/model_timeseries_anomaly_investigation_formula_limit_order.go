// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationFormulaLimitOrder Sort order used when applying a formula series limit.
type TimeseriesAnomalyInvestigationFormulaLimitOrder string

// List of TimeseriesAnomalyInvestigationFormulaLimitOrder.
const (
	TIMESERIESANOMALYINVESTIGATIONFORMULALIMITORDER_ASC  TimeseriesAnomalyInvestigationFormulaLimitOrder = "asc"
	TIMESERIESANOMALYINVESTIGATIONFORMULALIMITORDER_DESC TimeseriesAnomalyInvestigationFormulaLimitOrder = "desc"
)

var allowedTimeseriesAnomalyInvestigationFormulaLimitOrderEnumValues = []TimeseriesAnomalyInvestigationFormulaLimitOrder{
	TIMESERIESANOMALYINVESTIGATIONFORMULALIMITORDER_ASC,
	TIMESERIESANOMALYINVESTIGATIONFORMULALIMITORDER_DESC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesAnomalyInvestigationFormulaLimitOrder) GetAllowedValues() []TimeseriesAnomalyInvestigationFormulaLimitOrder {
	return allowedTimeseriesAnomalyInvestigationFormulaLimitOrderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesAnomalyInvestigationFormulaLimitOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesAnomalyInvestigationFormulaLimitOrder(value)
	return nil
}

// NewTimeseriesAnomalyInvestigationFormulaLimitOrderFromValue returns a pointer to a valid TimeseriesAnomalyInvestigationFormulaLimitOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesAnomalyInvestigationFormulaLimitOrderFromValue(v string) (*TimeseriesAnomalyInvestigationFormulaLimitOrder, error) {
	ev := TimeseriesAnomalyInvestigationFormulaLimitOrder(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesAnomalyInvestigationFormulaLimitOrder: valid values are %v", v, allowedTimeseriesAnomalyInvestigationFormulaLimitOrderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesAnomalyInvestigationFormulaLimitOrder) IsValid() bool {
	for _, existing := range allowedTimeseriesAnomalyInvestigationFormulaLimitOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesAnomalyInvestigationFormulaLimitOrder value.
func (v TimeseriesAnomalyInvestigationFormulaLimitOrder) Ptr() *TimeseriesAnomalyInvestigationFormulaLimitOrder {
	return &v
}
