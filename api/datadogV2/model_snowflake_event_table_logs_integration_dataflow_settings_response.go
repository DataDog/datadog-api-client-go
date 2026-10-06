// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeEventTableLogsIntegrationDataflowSettingsResponse Settings of the event table dataflow. Each record type is collected independently so that you can control ingestion costs, and every record type is ingested into Datadog as logs tagged with its `record_type`.
type SnowflakeEventTableLogsIntegrationDataflowSettingsResponse struct {
	// Whether records with a `record_type` of `event` are collected.
	EventTableEventsEnabled *bool `json:"event_table_events_enabled,omitempty"`
	// Whether records with a `record_type` of `log` are collected.
	EventTableLogsEnabled *bool `json:"event_table_logs_enabled,omitempty"`
	// How often event table records are collected, in minutes.
	EventTableLogsIntervalMin *int64 `json:"event_table_logs_interval_min,omitempty"`
	// Whether records with a `record_type` of `span_event` are collected.
	EventTableSpanEventsEnabled *bool `json:"event_table_span_events_enabled,omitempty"`
	// Whether records with a `record_type` of `span` are collected.
	EventTableSpansEnabled *bool `json:"event_table_spans_enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSnowflakeEventTableLogsIntegrationDataflowSettingsResponse instantiates a new SnowflakeEventTableLogsIntegrationDataflowSettingsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeEventTableLogsIntegrationDataflowSettingsResponse() *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse {
	this := SnowflakeEventTableLogsIntegrationDataflowSettingsResponse{}
	return &this
}

// NewSnowflakeEventTableLogsIntegrationDataflowSettingsResponseWithDefaults instantiates a new SnowflakeEventTableLogsIntegrationDataflowSettingsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeEventTableLogsIntegrationDataflowSettingsResponseWithDefaults() *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse {
	this := SnowflakeEventTableLogsIntegrationDataflowSettingsResponse{}
	return &this
}

// GetEventTableEventsEnabled returns the EventTableEventsEnabled field value if set, zero value otherwise.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) GetEventTableEventsEnabled() bool {
	if o == nil || o.EventTableEventsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EventTableEventsEnabled
}

// GetEventTableEventsEnabledOk returns a tuple with the EventTableEventsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) GetEventTableEventsEnabledOk() (*bool, bool) {
	if o == nil || o.EventTableEventsEnabled == nil {
		return nil, false
	}
	return o.EventTableEventsEnabled, true
}

// HasEventTableEventsEnabled returns a boolean if a field has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) HasEventTableEventsEnabled() bool {
	return o != nil && o.EventTableEventsEnabled != nil
}

// SetEventTableEventsEnabled gets a reference to the given bool and assigns it to the EventTableEventsEnabled field.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) SetEventTableEventsEnabled(v bool) {
	o.EventTableEventsEnabled = &v
}

// GetEventTableLogsEnabled returns the EventTableLogsEnabled field value if set, zero value otherwise.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) GetEventTableLogsEnabled() bool {
	if o == nil || o.EventTableLogsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EventTableLogsEnabled
}

// GetEventTableLogsEnabledOk returns a tuple with the EventTableLogsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) GetEventTableLogsEnabledOk() (*bool, bool) {
	if o == nil || o.EventTableLogsEnabled == nil {
		return nil, false
	}
	return o.EventTableLogsEnabled, true
}

// HasEventTableLogsEnabled returns a boolean if a field has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) HasEventTableLogsEnabled() bool {
	return o != nil && o.EventTableLogsEnabled != nil
}

// SetEventTableLogsEnabled gets a reference to the given bool and assigns it to the EventTableLogsEnabled field.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) SetEventTableLogsEnabled(v bool) {
	o.EventTableLogsEnabled = &v
}

// GetEventTableLogsIntervalMin returns the EventTableLogsIntervalMin field value if set, zero value otherwise.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) GetEventTableLogsIntervalMin() int64 {
	if o == nil || o.EventTableLogsIntervalMin == nil {
		var ret int64
		return ret
	}
	return *o.EventTableLogsIntervalMin
}

// GetEventTableLogsIntervalMinOk returns a tuple with the EventTableLogsIntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) GetEventTableLogsIntervalMinOk() (*int64, bool) {
	if o == nil || o.EventTableLogsIntervalMin == nil {
		return nil, false
	}
	return o.EventTableLogsIntervalMin, true
}

// HasEventTableLogsIntervalMin returns a boolean if a field has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) HasEventTableLogsIntervalMin() bool {
	return o != nil && o.EventTableLogsIntervalMin != nil
}

// SetEventTableLogsIntervalMin gets a reference to the given int64 and assigns it to the EventTableLogsIntervalMin field.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) SetEventTableLogsIntervalMin(v int64) {
	o.EventTableLogsIntervalMin = &v
}

// GetEventTableSpanEventsEnabled returns the EventTableSpanEventsEnabled field value if set, zero value otherwise.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) GetEventTableSpanEventsEnabled() bool {
	if o == nil || o.EventTableSpanEventsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EventTableSpanEventsEnabled
}

// GetEventTableSpanEventsEnabledOk returns a tuple with the EventTableSpanEventsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) GetEventTableSpanEventsEnabledOk() (*bool, bool) {
	if o == nil || o.EventTableSpanEventsEnabled == nil {
		return nil, false
	}
	return o.EventTableSpanEventsEnabled, true
}

// HasEventTableSpanEventsEnabled returns a boolean if a field has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) HasEventTableSpanEventsEnabled() bool {
	return o != nil && o.EventTableSpanEventsEnabled != nil
}

// SetEventTableSpanEventsEnabled gets a reference to the given bool and assigns it to the EventTableSpanEventsEnabled field.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) SetEventTableSpanEventsEnabled(v bool) {
	o.EventTableSpanEventsEnabled = &v
}

// GetEventTableSpansEnabled returns the EventTableSpansEnabled field value if set, zero value otherwise.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) GetEventTableSpansEnabled() bool {
	if o == nil || o.EventTableSpansEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EventTableSpansEnabled
}

// GetEventTableSpansEnabledOk returns a tuple with the EventTableSpansEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) GetEventTableSpansEnabledOk() (*bool, bool) {
	if o == nil || o.EventTableSpansEnabled == nil {
		return nil, false
	}
	return o.EventTableSpansEnabled, true
}

// HasEventTableSpansEnabled returns a boolean if a field has been set.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) HasEventTableSpansEnabled() bool {
	return o != nil && o.EventTableSpansEnabled != nil
}

// SetEventTableSpansEnabled gets a reference to the given bool and assigns it to the EventTableSpansEnabled field.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) SetEventTableSpansEnabled(v bool) {
	o.EventTableSpansEnabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeEventTableLogsIntegrationDataflowSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"event_table_events_enabled", "event_table_logs_enabled", "event_table_logs_interval_min", "event_table_span_events_enabled", "event_table_spans_enabled"})
	} else {
		return err
	}
	o.EventTableEventsEnabled = all.EventTableEventsEnabled
	o.EventTableLogsEnabled = all.EventTableLogsEnabled
	o.EventTableLogsIntervalMin = all.EventTableLogsIntervalMin
	o.EventTableSpanEventsEnabled = all.EventTableSpanEventsEnabled
	o.EventTableSpansEnabled = all.EventTableSpansEnabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
