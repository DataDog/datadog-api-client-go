// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest Settings of the query history logs dataflow. Only the fields provided are changed.
type SnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest struct {
	// Whether query logs are joined with Snowflake access history, which adds the objects each query read and wrote so you can follow how data is used and where it came from. Defaults to `false`.
	JoinQueryHistoryWithAccessHistoryEnabled *bool `json:"join_query_history_with_access_history_enabled,omitempty"`
	// How often query history logs are collected, in minutes. One of `5`, `15`, `30`, `60`, `720`, or `1440`. Defaults to `5`.
	QueryHistoryLogsIntervalMin *int64 `json:"query_history_logs_interval_min,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest instantiates a new SnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest() *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest {
	this := SnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest{}
	return &this
}

// NewSnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequestWithDefaults instantiates a new SnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequestWithDefaults() *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest {
	this := SnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest{}
	return &this
}

// GetJoinQueryHistoryWithAccessHistoryEnabled returns the JoinQueryHistoryWithAccessHistoryEnabled field value if set, zero value otherwise.
func (o *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest) GetJoinQueryHistoryWithAccessHistoryEnabled() bool {
	if o == nil || o.JoinQueryHistoryWithAccessHistoryEnabled == nil {
		var ret bool
		return ret
	}
	return *o.JoinQueryHistoryWithAccessHistoryEnabled
}

// GetJoinQueryHistoryWithAccessHistoryEnabledOk returns a tuple with the JoinQueryHistoryWithAccessHistoryEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest) GetJoinQueryHistoryWithAccessHistoryEnabledOk() (*bool, bool) {
	if o == nil || o.JoinQueryHistoryWithAccessHistoryEnabled == nil {
		return nil, false
	}
	return o.JoinQueryHistoryWithAccessHistoryEnabled, true
}

// HasJoinQueryHistoryWithAccessHistoryEnabled returns a boolean if a field has been set.
func (o *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest) HasJoinQueryHistoryWithAccessHistoryEnabled() bool {
	return o != nil && o.JoinQueryHistoryWithAccessHistoryEnabled != nil
}

// SetJoinQueryHistoryWithAccessHistoryEnabled gets a reference to the given bool and assigns it to the JoinQueryHistoryWithAccessHistoryEnabled field.
func (o *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest) SetJoinQueryHistoryWithAccessHistoryEnabled(v bool) {
	o.JoinQueryHistoryWithAccessHistoryEnabled = &v
}

// GetQueryHistoryLogsIntervalMin returns the QueryHistoryLogsIntervalMin field value if set, zero value otherwise.
func (o *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest) GetQueryHistoryLogsIntervalMin() int64 {
	if o == nil || o.QueryHistoryLogsIntervalMin == nil {
		var ret int64
		return ret
	}
	return *o.QueryHistoryLogsIntervalMin
}

// GetQueryHistoryLogsIntervalMinOk returns a tuple with the QueryHistoryLogsIntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest) GetQueryHistoryLogsIntervalMinOk() (*int64, bool) {
	if o == nil || o.QueryHistoryLogsIntervalMin == nil {
		return nil, false
	}
	return o.QueryHistoryLogsIntervalMin, true
}

// HasQueryHistoryLogsIntervalMin returns a boolean if a field has been set.
func (o *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest) HasQueryHistoryLogsIntervalMin() bool {
	return o != nil && o.QueryHistoryLogsIntervalMin != nil
}

// SetQueryHistoryLogsIntervalMin gets a reference to the given int64 and assigns it to the QueryHistoryLogsIntervalMin field.
func (o *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest) SetQueryHistoryLogsIntervalMin(v int64) {
	o.QueryHistoryLogsIntervalMin = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.JoinQueryHistoryWithAccessHistoryEnabled != nil {
		toSerialize["join_query_history_with_access_history_enabled"] = o.JoinQueryHistoryWithAccessHistoryEnabled
	}
	if o.QueryHistoryLogsIntervalMin != nil {
		toSerialize["query_history_logs_interval_min"] = o.QueryHistoryLogsIntervalMin
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		JoinQueryHistoryWithAccessHistoryEnabled *bool  `json:"join_query_history_with_access_history_enabled,omitempty"`
		QueryHistoryLogsIntervalMin              *int64 `json:"query_history_logs_interval_min,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.JoinQueryHistoryWithAccessHistoryEnabled = all.JoinQueryHistoryWithAccessHistoryEnabled
	o.QueryHistoryLogsIntervalMin = all.QueryHistoryLogsIntervalMin

	return nil
}
