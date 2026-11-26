// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricsHistogramRequestRequestType Test-only generated schema for the dashboard versioning proof of concept.
type MetricsHistogramRequestRequestType string

// List of MetricsHistogramRequestRequestType.
const (
	METRICSHISTOGRAMREQUESTREQUESTTYPE_HISTOGRAM MetricsHistogramRequestRequestType = "histogram"
)

var allowedMetricsHistogramRequestRequestTypeEnumValues = []MetricsHistogramRequestRequestType{
	METRICSHISTOGRAMREQUESTREQUESTTYPE_HISTOGRAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MetricsHistogramRequestRequestType) GetAllowedValues() []MetricsHistogramRequestRequestType {
	return allowedMetricsHistogramRequestRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricsHistogramRequestRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricsHistogramRequestRequestType(value)
	return nil
}

// NewMetricsHistogramRequestRequestTypeFromValue returns a pointer to a valid MetricsHistogramRequestRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricsHistogramRequestRequestTypeFromValue(v string) (*MetricsHistogramRequestRequestType, error) {
	ev := MetricsHistogramRequestRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricsHistogramRequestRequestType: valid values are %v", v, allowedMetricsHistogramRequestRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricsHistogramRequestRequestType) IsValid() bool {
	for _, existing := range allowedMetricsHistogramRequestRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricsHistogramRequestRequestType value.
func (v MetricsHistogramRequestRequestType) Ptr() *MetricsHistogramRequestRequestType {
	return &v
}
