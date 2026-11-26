// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaProductAnalyticsExtendedQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaProductAnalyticsExtendedQueryDataSource string

// List of FormulaProductAnalyticsExtendedQueryDataSource.
const (
	FORMULAPRODUCTANALYTICSEXTENDEDQUERYDATASOURCE_PRODUCT_ANALYTICS_EXTENDED FormulaProductAnalyticsExtendedQueryDataSource = "product_analytics_extended"
)

var allowedFormulaProductAnalyticsExtendedQueryDataSourceEnumValues = []FormulaProductAnalyticsExtendedQueryDataSource{
	FORMULAPRODUCTANALYTICSEXTENDEDQUERYDATASOURCE_PRODUCT_ANALYTICS_EXTENDED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaProductAnalyticsExtendedQueryDataSource) GetAllowedValues() []FormulaProductAnalyticsExtendedQueryDataSource {
	return allowedFormulaProductAnalyticsExtendedQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaProductAnalyticsExtendedQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaProductAnalyticsExtendedQueryDataSource(value)
	return nil
}

// NewFormulaProductAnalyticsExtendedQueryDataSourceFromValue returns a pointer to a valid FormulaProductAnalyticsExtendedQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaProductAnalyticsExtendedQueryDataSourceFromValue(v string) (*FormulaProductAnalyticsExtendedQueryDataSource, error) {
	ev := FormulaProductAnalyticsExtendedQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaProductAnalyticsExtendedQueryDataSource: valid values are %v", v, allowedFormulaProductAnalyticsExtendedQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaProductAnalyticsExtendedQueryDataSource) IsValid() bool {
	for _, existing := range allowedFormulaProductAnalyticsExtendedQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaProductAnalyticsExtendedQueryDataSource value.
func (v FormulaProductAnalyticsExtendedQueryDataSource) Ptr() *FormulaProductAnalyticsExtendedQueryDataSource {
	return &v
}
