// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeEventTableLogsIntegrationDataflowSettingsRequest Settings of the event table dataflow. Each record type is collected independently so that you can control ingestion costs, and every record type is ingested into Datadog as logs tagged with its `record_type`. Only the fields provided are changed.
type SnowflakeEventTableLogsIntegrationDataflowSettingsRequest struct {
	// Whether records with a `record_type` of `event` are collected. Defaults to `false`.
	EventTableEventsEnabled *bool `json:"event_table_events_enabled,omitempty"`
	// Whether records with a `record_type` of `log` are collected. Defaults to `false`.
	EventTableLogsEnabled *bool `json:"event_table_logs_enabled,omitempty"`
	// How often event table records are collected, in minutes. One of `5`, `15`, `30`, `60`, or `1440`. Defaults to `5`.
	EventTableLogsIntervalMin *int64 `json:"event_table_logs_interval_min,omitempty"`
	// Whether records with a `record_type` of `span_event` are collected. Defaults to `false`.
	EventTableSpanEventsEnabled *bool `json:"event_table_span_events_enabled,omitempty"`
	// Whether records with a `record_type` of `span` are collected. Defaults to `false`.
	EventTableSpansEnabled *bool `json:"event_table_spans_enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSnowflakeEventTableLogsIntegrationDataflowSettingsRequest instantiates a new SnowflakeEventTableLogsIntegrationDataflowSettingsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeEventTableLogsIntegrationDataflowSettingsRequest() *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest {
	this := SnowflakeEventTableLogsIntegrationDataflowSettingsRequest{}
	return &this
}

// NewSnowflakeEventTableLogsIntegrationDataflowSettingsRequestWithDefaults instantiates a new SnowflakeEventTableLogsIntegrationDataflowSettingsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeEventTableLogsIntegrationDataflowSettingsRequestWithDefaults() *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest {
	this := SnowflakeEventTableLogsIntegrationDataflowSettingsRequest{}
	return &this
}

// GetEventTableEventsEnabled returns the EventTableEventsEnabled field value if set, zero value otherwise.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) GetEventTableEventsEnabled() bool {
	if o == nil || o.EventTableEventsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EventTableEventsEnabled
}

// GetEventTableEventsEnabledOk returns a tuple with the EventTableEventsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) GetEventTableEventsEnabledOk() (*bool, bool) {
	if o == nil || o.EventTableEventsEnabled == nil {
		return nil, false
	}
	return o.EventTableEventsEnabled, true
}

// HasEventTableEventsEnabled returns a boolean if a field has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) HasEventTableEventsEnabled() bool {
	return o != nil && o.EventTableEventsEnabled != nil
}

// SetEventTableEventsEnabled gets a reference to the given bool and assigns it to the EventTableEventsEnabled field.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) SetEventTableEventsEnabled(v bool) {
	o.EventTableEventsEnabled = &v
}

// GetEventTableLogsEnabled returns the EventTableLogsEnabled field value if set, zero value otherwise.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) GetEventTableLogsEnabled() bool {
	if o == nil || o.EventTableLogsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EventTableLogsEnabled
}

// GetEventTableLogsEnabledOk returns a tuple with the EventTableLogsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) GetEventTableLogsEnabledOk() (*bool, bool) {
	if o == nil || o.EventTableLogsEnabled == nil {
		return nil, false
	}
	return o.EventTableLogsEnabled, true
}

// HasEventTableLogsEnabled returns a boolean if a field has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) HasEventTableLogsEnabled() bool {
	return o != nil && o.EventTableLogsEnabled != nil
}

// SetEventTableLogsEnabled gets a reference to the given bool and assigns it to the EventTableLogsEnabled field.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) SetEventTableLogsEnabled(v bool) {
	o.EventTableLogsEnabled = &v
}

// GetEventTableLogsIntervalMin returns the EventTableLogsIntervalMin field value if set, zero value otherwise.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) GetEventTableLogsIntervalMin() int64 {
	if o == nil || o.EventTableLogsIntervalMin == nil {
		var ret int64
		return ret
	}
	return *o.EventTableLogsIntervalMin
}

// GetEventTableLogsIntervalMinOk returns a tuple with the EventTableLogsIntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) GetEventTableLogsIntervalMinOk() (*int64, bool) {
	if o == nil || o.EventTableLogsIntervalMin == nil {
		return nil, false
	}
	return o.EventTableLogsIntervalMin, true
}

// HasEventTableLogsIntervalMin returns a boolean if a field has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) HasEventTableLogsIntervalMin() bool {
	return o != nil && o.EventTableLogsIntervalMin != nil
}

// SetEventTableLogsIntervalMin gets a reference to the given int64 and assigns it to the EventTableLogsIntervalMin field.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) SetEventTableLogsIntervalMin(v int64) {
	o.EventTableLogsIntervalMin = &v
}

// GetEventTableSpanEventsEnabled returns the EventTableSpanEventsEnabled field value if set, zero value otherwise.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) GetEventTableSpanEventsEnabled() bool {
	if o == nil || o.EventTableSpanEventsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EventTableSpanEventsEnabled
}

// GetEventTableSpanEventsEnabledOk returns a tuple with the EventTableSpanEventsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) GetEventTableSpanEventsEnabledOk() (*bool, bool) {
	if o == nil || o.EventTableSpanEventsEnabled == nil {
		return nil, false
	}
	return o.EventTableSpanEventsEnabled, true
}

// HasEventTableSpanEventsEnabled returns a boolean if a field has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) HasEventTableSpanEventsEnabled() bool {
	return o != nil && o.EventTableSpanEventsEnabled != nil
}

// SetEventTableSpanEventsEnabled gets a reference to the given bool and assigns it to the EventTableSpanEventsEnabled field.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) SetEventTableSpanEventsEnabled(v bool) {
	o.EventTableSpanEventsEnabled = &v
}

// GetEventTableSpansEnabled returns the EventTableSpansEnabled field value if set, zero value otherwise.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) GetEventTableSpansEnabled() bool {
	if o == nil || o.EventTableSpansEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EventTableSpansEnabled
}

// GetEventTableSpansEnabledOk returns a tuple with the EventTableSpansEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) GetEventTableSpansEnabledOk() (*bool, bool) {
	if o == nil || o.EventTableSpansEnabled == nil {
		return nil, false
	}
	return o.EventTableSpansEnabled, true
}

// HasEventTableSpansEnabled returns a boolean if a field has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) HasEventTableSpansEnabled() bool {
	return o != nil && o.EventTableSpansEnabled != nil
}

// SetEventTableSpansEnabled gets a reference to the given bool and assigns it to the EventTableSpansEnabled field.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) SetEventTableSpansEnabled(v bool) {
	o.EventTableSpansEnabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EventTableEventsEnabled != nil {
		toSerialize["event_table_events_enabled"] = o.EventTableEventsEnabled
	}
	if o.EventTableLogsEnabled != nil {
		toSerialize["event_table_logs_enabled"] = o.EventTableLogsEnabled
	}
	if o.EventTableLogsIntervalMin != nil {
		toSerialize["event_table_logs_interval_min"] = o.EventTableLogsIntervalMin
	}
	if o.EventTableSpanEventsEnabled != nil {
		toSerialize["event_table_span_events_enabled"] = o.EventTableSpanEventsEnabled
	}
	if o.EventTableSpansEnabled != nil {
		toSerialize["event_table_spans_enabled"] = o.EventTableSpansEnabled
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EventTableEventsEnabled     *bool  `json:"event_table_events_enabled,omitempty"`
		EventTableLogsEnabled       *bool  `json:"event_table_logs_enabled,omitempty"`
		EventTableLogsIntervalMin   *int64 `json:"event_table_logs_interval_min,omitempty"`
		EventTableSpanEventsEnabled *bool  `json:"event_table_span_events_enabled,omitempty"`
		EventTableSpansEnabled      *bool  `json:"event_table_spans_enabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.EventTableEventsEnabled = all.EventTableEventsEnabled
	o.EventTableLogsEnabled = all.EventTableLogsEnabled
	o.EventTableLogsIntervalMin = all.EventTableLogsIntervalMin
	o.EventTableSpanEventsEnabled = all.EventTableSpanEventsEnabled
	o.EventTableSpansEnabled = all.EventTableSpansEnabled

	return nil
}
