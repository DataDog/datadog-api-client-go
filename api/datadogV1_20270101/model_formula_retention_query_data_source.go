// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaRetentionQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaRetentionQueryDataSource string

// List of FormulaRetentionQueryDataSource.
const (
	FORMULARETENTIONQUERYDATASOURCE_PRODUCT_ANALYTICS_RETENTION FormulaRetentionQueryDataSource = "product_analytics_retention"
)

var allowedFormulaRetentionQueryDataSourceEnumValues = []FormulaRetentionQueryDataSource{
	FORMULARETENTIONQUERYDATASOURCE_PRODUCT_ANALYTICS_RETENTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaRetentionQueryDataSource) GetAllowedValues() []FormulaRetentionQueryDataSource {
	return allowedFormulaRetentionQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaRetentionQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaRetentionQueryDataSource(value)
	return nil
}

// NewFormulaRetentionQueryDataSourceFromValue returns a pointer to a valid FormulaRetentionQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaRetentionQueryDataSourceFromValue(v string) (*FormulaRetentionQueryDataSource, error) {
	ev := FormulaRetentionQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaRetentionQueryDataSource: valid values are %v", v, allowedFormulaRetentionQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaRetentionQueryDataSource) IsValid() bool {
	for _, existing := range allowedFormulaRetentionQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaRetentionQueryDataSource value.
func (v FormulaRetentionQueryDataSource) Ptr() *FormulaRetentionQueryDataSource {
	return &v
}
