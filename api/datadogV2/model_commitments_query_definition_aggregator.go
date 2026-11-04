// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsQueryDefinitionAggregator The reducer a scalar widget applies to the query.
type CommitmentsQueryDefinitionAggregator string

// List of CommitmentsQueryDefinitionAggregator.
const (
	COMMITMENTSQUERYDEFINITIONAGGREGATOR_SUM CommitmentsQueryDefinitionAggregator = "sum"
	COMMITMENTSQUERYDEFINITIONAGGREGATOR_AVG CommitmentsQueryDefinitionAggregator = "avg"
	COMMITMENTSQUERYDEFINITIONAGGREGATOR_MIN CommitmentsQueryDefinitionAggregator = "min"
	COMMITMENTSQUERYDEFINITIONAGGREGATOR_MAX CommitmentsQueryDefinitionAggregator = "max"
)

var allowedCommitmentsQueryDefinitionAggregatorEnumValues = []CommitmentsQueryDefinitionAggregator{
	COMMITMENTSQUERYDEFINITIONAGGREGATOR_SUM,
	COMMITMENTSQUERYDEFINITIONAGGREGATOR_AVG,
	COMMITMENTSQUERYDEFINITIONAGGREGATOR_MIN,
	COMMITMENTSQUERYDEFINITIONAGGREGATOR_MAX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CommitmentsQueryDefinitionAggregator) GetAllowedValues() []CommitmentsQueryDefinitionAggregator {
	return allowedCommitmentsQueryDefinitionAggregatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CommitmentsQueryDefinitionAggregator) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CommitmentsQueryDefinitionAggregator(value)
	return nil
}

// NewCommitmentsQueryDefinitionAggregatorFromValue returns a pointer to a valid CommitmentsQueryDefinitionAggregator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCommitmentsQueryDefinitionAggregatorFromValue(v string) (*CommitmentsQueryDefinitionAggregator, error) {
	ev := CommitmentsQueryDefinitionAggregator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CommitmentsQueryDefinitionAggregator: valid values are %v", v, allowedCommitmentsQueryDefinitionAggregatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CommitmentsQueryDefinitionAggregator) IsValid() bool {
	for _, existing := range allowedCommitmentsQueryDefinitionAggregatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommitmentsQueryDefinitionAggregator value.
func (v CommitmentsQueryDefinitionAggregator) Ptr() *CommitmentsQueryDefinitionAggregator {
	return &v
}
