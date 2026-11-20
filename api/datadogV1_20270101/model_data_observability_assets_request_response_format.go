// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataObservabilityAssetsRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type DataObservabilityAssetsRequestResponseFormat string

// List of DataObservabilityAssetsRequestResponseFormat.
const (
	DATAOBSERVABILITYASSETSREQUESTRESPONSEFORMAT_EVENT_LIST DataObservabilityAssetsRequestResponseFormat = "event_list"
)

var allowedDataObservabilityAssetsRequestResponseFormatEnumValues = []DataObservabilityAssetsRequestResponseFormat{
	DATAOBSERVABILITYASSETSREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DataObservabilityAssetsRequestResponseFormat) GetAllowedValues() []DataObservabilityAssetsRequestResponseFormat {
	return allowedDataObservabilityAssetsRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DataObservabilityAssetsRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DataObservabilityAssetsRequestResponseFormat(value)
	return nil
}

// NewDataObservabilityAssetsRequestResponseFormatFromValue returns a pointer to a valid DataObservabilityAssetsRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDataObservabilityAssetsRequestResponseFormatFromValue(v string) (*DataObservabilityAssetsRequestResponseFormat, error) {
	ev := DataObservabilityAssetsRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DataObservabilityAssetsRequestResponseFormat: valid values are %v", v, allowedDataObservabilityAssetsRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DataObservabilityAssetsRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedDataObservabilityAssetsRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataObservabilityAssetsRequestResponseFormat value.
func (v DataObservabilityAssetsRequestResponseFormat) Ptr() *DataObservabilityAssetsRequestResponseFormat {
	return &v
}
