// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetRestrictionsType JSON:API resource type for dataset restrictions.
type DatasetRestrictionsType string

// List of DatasetRestrictionsType.
const (
	DATASETRESTRICTIONSTYPE_DATASET_RESTRICTIONS DatasetRestrictionsType = "dataset_restrictions"
)

var allowedDatasetRestrictionsTypeEnumValues = []DatasetRestrictionsType{
	DATASETRESTRICTIONSTYPE_DATASET_RESTRICTIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatasetRestrictionsType) GetAllowedValues() []DatasetRestrictionsType {
	return allowedDatasetRestrictionsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatasetRestrictionsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatasetRestrictionsType(value)
	return nil
}

// NewDatasetRestrictionsTypeFromValue returns a pointer to a valid DatasetRestrictionsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatasetRestrictionsTypeFromValue(v string) (*DatasetRestrictionsType, error) {
	ev := DatasetRestrictionsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatasetRestrictionsType: valid values are %v", v, allowedDatasetRestrictionsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatasetRestrictionsType) IsValid() bool {
	for _, existing := range allowedDatasetRestrictionsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatasetRestrictionsType value.
func (v DatasetRestrictionsType) Ptr() *DatasetRestrictionsType {
	return &v
}
