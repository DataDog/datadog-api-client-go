// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetListQueryDataSourceType Identifies this as a published-dataset list query.
type DatasetListQueryDataSourceType string

// List of DatasetListQueryDataSourceType.
const (
	DATASETLISTQUERYDATASOURCETYPE_DATASET DatasetListQueryDataSourceType = "dataset"
)

var allowedDatasetListQueryDataSourceTypeEnumValues = []DatasetListQueryDataSourceType{
	DATASETLISTQUERYDATASOURCETYPE_DATASET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatasetListQueryDataSourceType) GetAllowedValues() []DatasetListQueryDataSourceType {
	return allowedDatasetListQueryDataSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatasetListQueryDataSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatasetListQueryDataSourceType(value)
	return nil
}

// NewDatasetListQueryDataSourceTypeFromValue returns a pointer to a valid DatasetListQueryDataSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatasetListQueryDataSourceTypeFromValue(v string) (*DatasetListQueryDataSourceType, error) {
	ev := DatasetListQueryDataSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatasetListQueryDataSourceType: valid values are %v", v, allowedDatasetListQueryDataSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatasetListQueryDataSourceType) IsValid() bool {
	for _, existing := range allowedDatasetListQueryDataSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatasetListQueryDataSourceType value.
func (v DatasetListQueryDataSourceType) Ptr() *DatasetListQueryDataSourceType {
	return &v
}
