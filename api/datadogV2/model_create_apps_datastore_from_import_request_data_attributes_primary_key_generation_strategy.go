// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy The `attributes` `primary_key_generation_strategy`.
type CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy string

// List of CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy.
const (
	CREATEAPPSDATASTOREFROMIMPORTREQUESTDATAATTRIBUTESPRIMARYKEYGENERATIONSTRATEGY_NONE CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy = "none"
	CREATEAPPSDATASTOREFROMIMPORTREQUESTDATAATTRIBUTESPRIMARYKEYGENERATIONSTRATEGY_UUID CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy = "uuid"
)

var allowedCreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategyEnumValues = []CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy{
	CREATEAPPSDATASTOREFROMIMPORTREQUESTDATAATTRIBUTESPRIMARYKEYGENERATIONSTRATEGY_NONE,
	CREATEAPPSDATASTOREFROMIMPORTREQUESTDATAATTRIBUTESPRIMARYKEYGENERATIONSTRATEGY_UUID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy) GetAllowedValues() []CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy {
	return allowedCreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy(value)
	return nil
}

// NewCreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategyFromValue returns a pointer to a valid CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategyFromValue(v string) (*CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy, error) {
	ev := CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy: valid values are %v", v, allowedCreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy) IsValid() bool {
	for _, existing := range allowedCreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy value.
func (v CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy) Ptr() *CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy {
	return &v
}
