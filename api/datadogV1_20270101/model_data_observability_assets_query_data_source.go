// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataObservabilityAssetsQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type DataObservabilityAssetsQueryDataSource string

// List of DataObservabilityAssetsQueryDataSource.
const (
	DATAOBSERVABILITYASSETSQUERYDATASOURCE_DATA_OBSERVABILITY_ASSETS DataObservabilityAssetsQueryDataSource = "data_observability_assets"
)

var allowedDataObservabilityAssetsQueryDataSourceEnumValues = []DataObservabilityAssetsQueryDataSource{
	DATAOBSERVABILITYASSETSQUERYDATASOURCE_DATA_OBSERVABILITY_ASSETS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DataObservabilityAssetsQueryDataSource) GetAllowedValues() []DataObservabilityAssetsQueryDataSource {
	return allowedDataObservabilityAssetsQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DataObservabilityAssetsQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DataObservabilityAssetsQueryDataSource(value)
	return nil
}

// NewDataObservabilityAssetsQueryDataSourceFromValue returns a pointer to a valid DataObservabilityAssetsQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDataObservabilityAssetsQueryDataSourceFromValue(v string) (*DataObservabilityAssetsQueryDataSource, error) {
	ev := DataObservabilityAssetsQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DataObservabilityAssetsQueryDataSource: valid values are %v", v, allowedDataObservabilityAssetsQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DataObservabilityAssetsQueryDataSource) IsValid() bool {
	for _, existing := range allowedDataObservabilityAssetsQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataObservabilityAssetsQueryDataSource value.
func (v DataObservabilityAssetsQueryDataSource) Ptr() *DataObservabilityAssetsQueryDataSource {
	return &v
}
