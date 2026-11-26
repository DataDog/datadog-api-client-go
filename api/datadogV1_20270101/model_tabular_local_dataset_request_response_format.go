// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TabularLocalDatasetRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type TabularLocalDatasetRequestResponseFormat string

// List of TabularLocalDatasetRequestResponseFormat.
const (
	TABULARLOCALDATASETREQUESTRESPONSEFORMAT_TABULAR TabularLocalDatasetRequestResponseFormat = "tabular"
)

var allowedTabularLocalDatasetRequestResponseFormatEnumValues = []TabularLocalDatasetRequestResponseFormat{
	TABULARLOCALDATASETREQUESTRESPONSEFORMAT_TABULAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TabularLocalDatasetRequestResponseFormat) GetAllowedValues() []TabularLocalDatasetRequestResponseFormat {
	return allowedTabularLocalDatasetRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TabularLocalDatasetRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TabularLocalDatasetRequestResponseFormat(value)
	return nil
}

// NewTabularLocalDatasetRequestResponseFormatFromValue returns a pointer to a valid TabularLocalDatasetRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTabularLocalDatasetRequestResponseFormatFromValue(v string) (*TabularLocalDatasetRequestResponseFormat, error) {
	ev := TabularLocalDatasetRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TabularLocalDatasetRequestResponseFormat: valid values are %v", v, allowedTabularLocalDatasetRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TabularLocalDatasetRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedTabularLocalDatasetRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TabularLocalDatasetRequestResponseFormat value.
func (v TabularLocalDatasetRequestResponseFormat) Ptr() *TabularLocalDatasetRequestResponseFormat {
	return &v
}
