// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableEventsQueryDataSource Events data source.
type GuidedTableEventsQueryDataSource string

// List of GuidedTableEventsQueryDataSource.
const (
	GUIDEDTABLEEVENTSQUERYDATASOURCE_LOGS              GuidedTableEventsQueryDataSource = "logs"
	GUIDEDTABLEEVENTSQUERYDATASOURCE_RUM               GuidedTableEventsQueryDataSource = "rum"
	GUIDEDTABLEEVENTSQUERYDATASOURCE_SPANS             GuidedTableEventsQueryDataSource = "spans"
	GUIDEDTABLEEVENTSQUERYDATASOURCE_CI_PIPELINES      GuidedTableEventsQueryDataSource = "ci_pipelines"
	GUIDEDTABLEEVENTSQUERYDATASOURCE_EVENTS            GuidedTableEventsQueryDataSource = "events"
	GUIDEDTABLEEVENTSQUERYDATASOURCE_PRODUCT_ANALYTICS GuidedTableEventsQueryDataSource = "product_analytics"
)

var allowedGuidedTableEventsQueryDataSourceEnumValues = []GuidedTableEventsQueryDataSource{
	GUIDEDTABLEEVENTSQUERYDATASOURCE_LOGS,
	GUIDEDTABLEEVENTSQUERYDATASOURCE_RUM,
	GUIDEDTABLEEVENTSQUERYDATASOURCE_SPANS,
	GUIDEDTABLEEVENTSQUERYDATASOURCE_CI_PIPELINES,
	GUIDEDTABLEEVENTSQUERYDATASOURCE_EVENTS,
	GUIDEDTABLEEVENTSQUERYDATASOURCE_PRODUCT_ANALYTICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GuidedTableEventsQueryDataSource) GetAllowedValues() []GuidedTableEventsQueryDataSource {
	return allowedGuidedTableEventsQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GuidedTableEventsQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GuidedTableEventsQueryDataSource(value)
	return nil
}

// NewGuidedTableEventsQueryDataSourceFromValue returns a pointer to a valid GuidedTableEventsQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGuidedTableEventsQueryDataSourceFromValue(v string) (*GuidedTableEventsQueryDataSource, error) {
	ev := GuidedTableEventsQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GuidedTableEventsQueryDataSource: valid values are %v", v, allowedGuidedTableEventsQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GuidedTableEventsQueryDataSource) IsValid() bool {
	for _, existing := range allowedGuidedTableEventsQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GuidedTableEventsQueryDataSource value.
func (v GuidedTableEventsQueryDataSource) Ptr() *GuidedTableEventsQueryDataSource {
	return &v
}
