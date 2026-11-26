// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaUserJourneyQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaUserJourneyQueryDataSource string

// List of FormulaUserJourneyQueryDataSource.
const (
	FORMULAUSERJOURNEYQUERYDATASOURCE_PRODUCT_ANALYTICS_JOURNEY FormulaUserJourneyQueryDataSource = "product_analytics_journey"
)

var allowedFormulaUserJourneyQueryDataSourceEnumValues = []FormulaUserJourneyQueryDataSource{
	FORMULAUSERJOURNEYQUERYDATASOURCE_PRODUCT_ANALYTICS_JOURNEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaUserJourneyQueryDataSource) GetAllowedValues() []FormulaUserJourneyQueryDataSource {
	return allowedFormulaUserJourneyQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaUserJourneyQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaUserJourneyQueryDataSource(value)
	return nil
}

// NewFormulaUserJourneyQueryDataSourceFromValue returns a pointer to a valid FormulaUserJourneyQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaUserJourneyQueryDataSourceFromValue(v string) (*FormulaUserJourneyQueryDataSource, error) {
	ev := FormulaUserJourneyQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaUserJourneyQueryDataSource: valid values are %v", v, allowedFormulaUserJourneyQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaUserJourneyQueryDataSource) IsValid() bool {
	for _, existing := range allowedFormulaUserJourneyQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaUserJourneyQueryDataSource value.
func (v FormulaUserJourneyQueryDataSource) Ptr() *FormulaUserJourneyQueryDataSource {
	return &v
}
