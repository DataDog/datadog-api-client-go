// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GraphQuerySearchFilterOperator Test-only generated schema for the dashboard versioning proof of concept.
type GraphQuerySearchFilterOperator string

// List of GraphQuerySearchFilterOperator.
const (
	GRAPHQUERYSEARCHFILTEROPERATOR_EQ    GraphQuerySearchFilterOperator = "="
	GRAPHQUERYSEARCHFILTEROPERATOR_GT    GraphQuerySearchFilterOperator = ">"
	GRAPHQUERYSEARCHFILTEROPERATOR_GT_EQ GraphQuerySearchFilterOperator = ">="
	GRAPHQUERYSEARCHFILTEROPERATOR_LT    GraphQuerySearchFilterOperator = "<"
	GRAPHQUERYSEARCHFILTEROPERATOR_LT_EQ GraphQuerySearchFilterOperator = "<="
)

var allowedGraphQuerySearchFilterOperatorEnumValues = []GraphQuerySearchFilterOperator{
	GRAPHQUERYSEARCHFILTEROPERATOR_EQ,
	GRAPHQUERYSEARCHFILTEROPERATOR_GT,
	GRAPHQUERYSEARCHFILTEROPERATOR_GT_EQ,
	GRAPHQUERYSEARCHFILTEROPERATOR_LT,
	GRAPHQUERYSEARCHFILTEROPERATOR_LT_EQ,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GraphQuerySearchFilterOperator) GetAllowedValues() []GraphQuerySearchFilterOperator {
	return allowedGraphQuerySearchFilterOperatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GraphQuerySearchFilterOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GraphQuerySearchFilterOperator(value)
	return nil
}

// NewGraphQuerySearchFilterOperatorFromValue returns a pointer to a valid GraphQuerySearchFilterOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGraphQuerySearchFilterOperatorFromValue(v string) (*GraphQuerySearchFilterOperator, error) {
	ev := GraphQuerySearchFilterOperator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GraphQuerySearchFilterOperator: valid values are %v", v, allowedGraphQuerySearchFilterOperatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GraphQuerySearchFilterOperator) IsValid() bool {
	for _, existing := range allowedGraphQuerySearchFilterOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GraphQuerySearchFilterOperator value.
func (v GraphQuerySearchFilterOperator) Ptr() *GraphQuerySearchFilterOperator {
	return &v
}
