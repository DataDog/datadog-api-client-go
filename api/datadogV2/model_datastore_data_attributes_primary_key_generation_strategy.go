// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatastoreDataAttributesPrimaryKeyGenerationStrategy The `attributes` `primary_key_generation_strategy`.
type DatastoreDataAttributesPrimaryKeyGenerationStrategy string

// List of DatastoreDataAttributesPrimaryKeyGenerationStrategy.
const (
	DATASTOREDATAATTRIBUTESPRIMARYKEYGENERATIONSTRATEGY_NONE DatastoreDataAttributesPrimaryKeyGenerationStrategy = "none"
	DATASTOREDATAATTRIBUTESPRIMARYKEYGENERATIONSTRATEGY_UUID DatastoreDataAttributesPrimaryKeyGenerationStrategy = "uuid"
)

var allowedDatastoreDataAttributesPrimaryKeyGenerationStrategyEnumValues = []DatastoreDataAttributesPrimaryKeyGenerationStrategy{
	DATASTOREDATAATTRIBUTESPRIMARYKEYGENERATIONSTRATEGY_NONE,
	DATASTOREDATAATTRIBUTESPRIMARYKEYGENERATIONSTRATEGY_UUID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatastoreDataAttributesPrimaryKeyGenerationStrategy) GetAllowedValues() []DatastoreDataAttributesPrimaryKeyGenerationStrategy {
	return allowedDatastoreDataAttributesPrimaryKeyGenerationStrategyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatastoreDataAttributesPrimaryKeyGenerationStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatastoreDataAttributesPrimaryKeyGenerationStrategy(value)
	return nil
}

// NewDatastoreDataAttributesPrimaryKeyGenerationStrategyFromValue returns a pointer to a valid DatastoreDataAttributesPrimaryKeyGenerationStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatastoreDataAttributesPrimaryKeyGenerationStrategyFromValue(v string) (*DatastoreDataAttributesPrimaryKeyGenerationStrategy, error) {
	ev := DatastoreDataAttributesPrimaryKeyGenerationStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatastoreDataAttributesPrimaryKeyGenerationStrategy: valid values are %v", v, allowedDatastoreDataAttributesPrimaryKeyGenerationStrategyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatastoreDataAttributesPrimaryKeyGenerationStrategy) IsValid() bool {
	for _, existing := range allowedDatastoreDataAttributesPrimaryKeyGenerationStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatastoreDataAttributesPrimaryKeyGenerationStrategy value.
func (v DatastoreDataAttributesPrimaryKeyGenerationStrategy) Ptr() *DatastoreDataAttributesPrimaryKeyGenerationStrategy {
	return &v
}
