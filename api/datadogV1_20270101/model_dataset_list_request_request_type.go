// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetListRequestRequestType Test-only generated schema for the dashboard versioning proof of concept.
type DatasetListRequestRequestType string

// List of DatasetListRequestRequestType.
const (
	DATASETLISTREQUESTREQUESTTYPE_DATASET_LIST_ROWS DatasetListRequestRequestType = "dataset_list_rows"
)

var allowedDatasetListRequestRequestTypeEnumValues = []DatasetListRequestRequestType{
	DATASETLISTREQUESTREQUESTTYPE_DATASET_LIST_ROWS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DatasetListRequestRequestType) GetAllowedValues() []DatasetListRequestRequestType {
	return allowedDatasetListRequestRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DatasetListRequestRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DatasetListRequestRequestType(value)
	return nil
}

// NewDatasetListRequestRequestTypeFromValue returns a pointer to a valid DatasetListRequestRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDatasetListRequestRequestTypeFromValue(v string) (*DatasetListRequestRequestType, error) {
	ev := DatasetListRequestRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DatasetListRequestRequestType: valid values are %v", v, allowedDatasetListRequestRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DatasetListRequestRequestType) IsValid() bool {
	for _, existing := range allowedDatasetListRequestRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatasetListRequestRequestType value.
func (v DatasetListRequestRequestType) Ptr() *DatasetListRequestRequestType {
	return &v
}
