// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmStatsHistogramRequestRequestType Test-only generated schema for the dashboard versioning proof of concept.
type ApmStatsHistogramRequestRequestType string

// List of ApmStatsHistogramRequestRequestType.
const (
	APMSTATSHISTOGRAMREQUESTREQUESTTYPE_HISTOGRAM ApmStatsHistogramRequestRequestType = "histogram"
)

var allowedApmStatsHistogramRequestRequestTypeEnumValues = []ApmStatsHistogramRequestRequestType{
	APMSTATSHISTOGRAMREQUESTREQUESTTYPE_HISTOGRAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApmStatsHistogramRequestRequestType) GetAllowedValues() []ApmStatsHistogramRequestRequestType {
	return allowedApmStatsHistogramRequestRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApmStatsHistogramRequestRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApmStatsHistogramRequestRequestType(value)
	return nil
}

// NewApmStatsHistogramRequestRequestTypeFromValue returns a pointer to a valid ApmStatsHistogramRequestRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApmStatsHistogramRequestRequestTypeFromValue(v string) (*ApmStatsHistogramRequestRequestType, error) {
	ev := ApmStatsHistogramRequestRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApmStatsHistogramRequestRequestType: valid values are %v", v, allowedApmStatsHistogramRequestRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApmStatsHistogramRequestRequestType) IsValid() bool {
	for _, existing := range allowedApmStatsHistogramRequestRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApmStatsHistogramRequestRequestType value.
func (v ApmStatsHistogramRequestRequestType) Ptr() *ApmStatsHistogramRequestRequestType {
	return &v
}
