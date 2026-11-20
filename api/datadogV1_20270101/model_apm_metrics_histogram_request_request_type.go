// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmMetricsHistogramRequestRequestType Test-only generated schema for the dashboard versioning proof of concept.
type ApmMetricsHistogramRequestRequestType string

// List of ApmMetricsHistogramRequestRequestType.
const (
	APMMETRICSHISTOGRAMREQUESTREQUESTTYPE_HISTOGRAM ApmMetricsHistogramRequestRequestType = "histogram"
)

var allowedApmMetricsHistogramRequestRequestTypeEnumValues = []ApmMetricsHistogramRequestRequestType{
	APMMETRICSHISTOGRAMREQUESTREQUESTTYPE_HISTOGRAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApmMetricsHistogramRequestRequestType) GetAllowedValues() []ApmMetricsHistogramRequestRequestType {
	return allowedApmMetricsHistogramRequestRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApmMetricsHistogramRequestRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApmMetricsHistogramRequestRequestType(value)
	return nil
}

// NewApmMetricsHistogramRequestRequestTypeFromValue returns a pointer to a valid ApmMetricsHistogramRequestRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApmMetricsHistogramRequestRequestTypeFromValue(v string) (*ApmMetricsHistogramRequestRequestType, error) {
	ev := ApmMetricsHistogramRequestRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApmMetricsHistogramRequestRequestType: valid values are %v", v, allowedApmMetricsHistogramRequestRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApmMetricsHistogramRequestRequestType) IsValid() bool {
	for _, existing := range allowedApmMetricsHistogramRequestRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApmMetricsHistogramRequestRequestType value.
func (v ApmMetricsHistogramRequestRequestType) Ptr() *ApmMetricsHistogramRequestRequestType {
	return &v
}
