// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BigQueryDatasetDataSource Test-only generated schema for the dashboard versioning proof of concept.
type BigQueryDatasetDataSource string

// List of BigQueryDatasetDataSource.
const (
	BIGQUERYDATASETDATASOURCE_BIGQUERY BigQueryDatasetDataSource = "bigquery"
)

var allowedBigQueryDatasetDataSourceEnumValues = []BigQueryDatasetDataSource{
	BIGQUERYDATASETDATASOURCE_BIGQUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BigQueryDatasetDataSource) GetAllowedValues() []BigQueryDatasetDataSource {
	return allowedBigQueryDatasetDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BigQueryDatasetDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BigQueryDatasetDataSource(value)
	return nil
}

// NewBigQueryDatasetDataSourceFromValue returns a pointer to a valid BigQueryDatasetDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBigQueryDatasetDataSourceFromValue(v string) (*BigQueryDatasetDataSource, error) {
	ev := BigQueryDatasetDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BigQueryDatasetDataSource: valid values are %v", v, allowedBigQueryDatasetDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BigQueryDatasetDataSource) IsValid() bool {
	for _, existing := range allowedBigQueryDatasetDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BigQueryDatasetDataSource value.
func (v BigQueryDatasetDataSource) Ptr() *BigQueryDatasetDataSource {
	return &v
}
