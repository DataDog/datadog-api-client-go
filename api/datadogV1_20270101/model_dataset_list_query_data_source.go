// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetListQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type DatasetListQueryDataSource string

// List of DatasetListQueryDataSource.
const (
	DATASETLISTQUERYDATASOURCE_DATASET DatasetListQueryDataSource = "dataset"
)

var allowedDatasetListQueryDataSourceEnumValues = []DatasetListQueryDataSource{
	DATASETLISTQUERYDATASOURCE_DATASET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatasetListQueryDataSource) GetAllowedValues() []DatasetListQueryDataSource {
	return allowedDatasetListQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatasetListQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatasetListQueryDataSource(value)
	return nil
}

// NewDatasetListQueryDataSourceFromValue returns a pointer to a valid DatasetListQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatasetListQueryDataSourceFromValue(v string) (*DatasetListQueryDataSource, error) {
	ev := DatasetListQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatasetListQueryDataSource: valid values are %v", v, allowedDatasetListQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatasetListQueryDataSource) IsValid() bool {
	for _, existing := range allowedDatasetListQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatasetListQueryDataSource value.
func (v DatasetListQueryDataSource) Ptr() *DatasetListQueryDataSource {
	return &v
}
