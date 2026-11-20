// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TabularLocalDatasetRequestRequestType Test-only generated schema for the dashboard versioning proof of concept.
type TabularLocalDatasetRequestRequestType string

// List of TabularLocalDatasetRequestRequestType.
const (
	TABULARLOCALDATASETREQUESTREQUESTTYPE_LOCAL_DATASET TabularLocalDatasetRequestRequestType = "local_dataset"
)

var allowedTabularLocalDatasetRequestRequestTypeEnumValues = []TabularLocalDatasetRequestRequestType{
	TABULARLOCALDATASETREQUESTREQUESTTYPE_LOCAL_DATASET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TabularLocalDatasetRequestRequestType) GetAllowedValues() []TabularLocalDatasetRequestRequestType {
	return allowedTabularLocalDatasetRequestRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TabularLocalDatasetRequestRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TabularLocalDatasetRequestRequestType(value)
	return nil
}

// NewTabularLocalDatasetRequestRequestTypeFromValue returns a pointer to a valid TabularLocalDatasetRequestRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTabularLocalDatasetRequestRequestTypeFromValue(v string) (*TabularLocalDatasetRequestRequestType, error) {
	ev := TabularLocalDatasetRequestRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TabularLocalDatasetRequestRequestType: valid values are %v", v, allowedTabularLocalDatasetRequestRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TabularLocalDatasetRequestRequestType) IsValid() bool {
	for _, existing := range allowedTabularLocalDatasetRequestRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TabularLocalDatasetRequestRequestType value.
func (v TabularLocalDatasetRequestRequestType) Ptr() *TabularLocalDatasetRequestRequestType {
	return &v
}
