// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CompositeReferentialDataQueryComputeItemsAggregation Test-only generated schema for the dashboard versioning proof of concept.
type CompositeReferentialDataQueryComputeItemsAggregation string

// List of CompositeReferentialDataQueryComputeItemsAggregation.
const (
	COMPOSITEREFERENTIALDATAQUERYCOMPUTEITEMSAGGREGATION_SUM   CompositeReferentialDataQueryComputeItemsAggregation = "sum"
	COMPOSITEREFERENTIALDATAQUERYCOMPUTEITEMSAGGREGATION_MIN   CompositeReferentialDataQueryComputeItemsAggregation = "min"
	COMPOSITEREFERENTIALDATAQUERYCOMPUTEITEMSAGGREGATION_MAX   CompositeReferentialDataQueryComputeItemsAggregation = "max"
	COMPOSITEREFERENTIALDATAQUERYCOMPUTEITEMSAGGREGATION_AVG   CompositeReferentialDataQueryComputeItemsAggregation = "avg"
	COMPOSITEREFERENTIALDATAQUERYCOMPUTEITEMSAGGREGATION_COUNT CompositeReferentialDataQueryComputeItemsAggregation = "count"
)

var allowedCompositeReferentialDataQueryComputeItemsAggregationEnumValues = []CompositeReferentialDataQueryComputeItemsAggregation{
	COMPOSITEREFERENTIALDATAQUERYCOMPUTEITEMSAGGREGATION_SUM,
	COMPOSITEREFERENTIALDATAQUERYCOMPUTEITEMSAGGREGATION_MIN,
	COMPOSITEREFERENTIALDATAQUERYCOMPUTEITEMSAGGREGATION_MAX,
	COMPOSITEREFERENTIALDATAQUERYCOMPUTEITEMSAGGREGATION_AVG,
	COMPOSITEREFERENTIALDATAQUERYCOMPUTEITEMSAGGREGATION_COUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CompositeReferentialDataQueryComputeItemsAggregation) GetAllowedValues() []CompositeReferentialDataQueryComputeItemsAggregation {
	return allowedCompositeReferentialDataQueryComputeItemsAggregationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CompositeReferentialDataQueryComputeItemsAggregation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CompositeReferentialDataQueryComputeItemsAggregation(value)
	return nil
}

// NewCompositeReferentialDataQueryComputeItemsAggregationFromValue returns a pointer to a valid CompositeReferentialDataQueryComputeItemsAggregation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCompositeReferentialDataQueryComputeItemsAggregationFromValue(v string) (*CompositeReferentialDataQueryComputeItemsAggregation, error) {
	ev := CompositeReferentialDataQueryComputeItemsAggregation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CompositeReferentialDataQueryComputeItemsAggregation: valid values are %v", v, allowedCompositeReferentialDataQueryComputeItemsAggregationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CompositeReferentialDataQueryComputeItemsAggregation) IsValid() bool {
	for _, existing := range allowedCompositeReferentialDataQueryComputeItemsAggregationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CompositeReferentialDataQueryComputeItemsAggregation value.
func (v CompositeReferentialDataQueryComputeItemsAggregation) Ptr() *CompositeReferentialDataQueryComputeItemsAggregation {
	return &v
}
