// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaDataQualityMetricsQueryAnomalyOptionsDirection Test-only generated schema for the dashboard versioning proof of concept.
type FormulaDataQualityMetricsQueryAnomalyOptionsDirection string

// List of FormulaDataQualityMetricsQueryAnomalyOptionsDirection.
const (
	FORMULADATAQUALITYMETRICSQUERYANOMALYOPTIONSDIRECTION_ABOVE FormulaDataQualityMetricsQueryAnomalyOptionsDirection = "above"
	FORMULADATAQUALITYMETRICSQUERYANOMALYOPTIONSDIRECTION_BELOW FormulaDataQualityMetricsQueryAnomalyOptionsDirection = "below"
	FORMULADATAQUALITYMETRICSQUERYANOMALYOPTIONSDIRECTION_BOTH  FormulaDataQualityMetricsQueryAnomalyOptionsDirection = "both"
)

var allowedFormulaDataQualityMetricsQueryAnomalyOptionsDirectionEnumValues = []FormulaDataQualityMetricsQueryAnomalyOptionsDirection{
	FORMULADATAQUALITYMETRICSQUERYANOMALYOPTIONSDIRECTION_ABOVE,
	FORMULADATAQUALITYMETRICSQUERYANOMALYOPTIONSDIRECTION_BELOW,
	FORMULADATAQUALITYMETRICSQUERYANOMALYOPTIONSDIRECTION_BOTH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaDataQualityMetricsQueryAnomalyOptionsDirection) GetAllowedValues() []FormulaDataQualityMetricsQueryAnomalyOptionsDirection {
	return allowedFormulaDataQualityMetricsQueryAnomalyOptionsDirectionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaDataQualityMetricsQueryAnomalyOptionsDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaDataQualityMetricsQueryAnomalyOptionsDirection(value)
	return nil
}

// NewFormulaDataQualityMetricsQueryAnomalyOptionsDirectionFromValue returns a pointer to a valid FormulaDataQualityMetricsQueryAnomalyOptionsDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaDataQualityMetricsQueryAnomalyOptionsDirectionFromValue(v string) (*FormulaDataQualityMetricsQueryAnomalyOptionsDirection, error) {
	ev := FormulaDataQualityMetricsQueryAnomalyOptionsDirection(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaDataQualityMetricsQueryAnomalyOptionsDirection: valid values are %v", v, allowedFormulaDataQualityMetricsQueryAnomalyOptionsDirectionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaDataQualityMetricsQueryAnomalyOptionsDirection) IsValid() bool {
	for _, existing := range allowedFormulaDataQualityMetricsQueryAnomalyOptionsDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaDataQualityMetricsQueryAnomalyOptionsDirection value.
func (v FormulaDataQualityMetricsQueryAnomalyOptionsDirection) Ptr() *FormulaDataQualityMetricsQueryAnomalyOptionsDirection {
	return &v
}
