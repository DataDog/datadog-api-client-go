// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventsQueryFragmentDefinitionDataSource Test-only generated schema for the dashboard versioning proof of concept.
type EventsQueryFragmentDefinitionDataSource string

// List of EventsQueryFragmentDefinitionDataSource.
const (
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_EVENTS               EventsQueryFragmentDefinitionDataSource = "events"
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_LOGS                 EventsQueryFragmentDefinitionDataSource = "logs"
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_RUM                  EventsQueryFragmentDefinitionDataSource = "rum"
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_SYNTHETICS_TEST_RUNS EventsQueryFragmentDefinitionDataSource = "synthetics_test_runs"
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_PRODUCT_ANALYTICS    EventsQueryFragmentDefinitionDataSource = "product_analytics"
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_NETWORK_DEVICE_FLOWS EventsQueryFragmentDefinitionDataSource = "network_device_flows"
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_CI_PIPELINES         EventsQueryFragmentDefinitionDataSource = "ci_pipelines"
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_DATABASE_QUERIES     EventsQueryFragmentDefinitionDataSource = "database_queries"
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_LLM_OBSERVABILITY    EventsQueryFragmentDefinitionDataSource = "llm_observability"
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_AUDIT                EventsQueryFragmentDefinitionDataSource = "audit"
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_CI_TESTS             EventsQueryFragmentDefinitionDataSource = "ci_tests"
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_SPANS                EventsQueryFragmentDefinitionDataSource = "spans"
)

var allowedEventsQueryFragmentDefinitionDataSourceEnumValues = []EventsQueryFragmentDefinitionDataSource{
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_EVENTS,
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_LOGS,
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_RUM,
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_SYNTHETICS_TEST_RUNS,
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_PRODUCT_ANALYTICS,
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_NETWORK_DEVICE_FLOWS,
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_CI_PIPELINES,
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_DATABASE_QUERIES,
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_LLM_OBSERVABILITY,
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_AUDIT,
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_CI_TESTS,
	EVENTSQUERYFRAGMENTDEFINITIONDATASOURCE_SPANS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventsQueryFragmentDefinitionDataSource) GetAllowedValues() []EventsQueryFragmentDefinitionDataSource {
	return allowedEventsQueryFragmentDefinitionDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventsQueryFragmentDefinitionDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventsQueryFragmentDefinitionDataSource(value)
	return nil
}

// NewEventsQueryFragmentDefinitionDataSourceFromValue returns a pointer to a valid EventsQueryFragmentDefinitionDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventsQueryFragmentDefinitionDataSourceFromValue(v string) (*EventsQueryFragmentDefinitionDataSource, error) {
	ev := EventsQueryFragmentDefinitionDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventsQueryFragmentDefinitionDataSource: valid values are %v", v, allowedEventsQueryFragmentDefinitionDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventsQueryFragmentDefinitionDataSource) IsValid() bool {
	for _, existing := range allowedEventsQueryFragmentDefinitionDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventsQueryFragmentDefinitionDataSource value.
func (v EventsQueryFragmentDefinitionDataSource) Ptr() *EventsQueryFragmentDefinitionDataSource {
	return &v
}
