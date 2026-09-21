// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse Settings of the query history logs dataflow.
type SnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse struct {
	// Whether query logs are joined with Snowflake access history, which adds the objects each query read and wrote so you can follow how data is used and where it came from.
	JoinQueryHistoryWithAccessHistoryEnabled *bool `json:"join_query_history_with_access_history_enabled,omitempty"`
	// How often query history logs are collected, in minutes.
	QueryHistoryLogsIntervalMin *int64 `json:"query_history_logs_interval_min,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse instantiates a new SnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse() *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse {
	this := SnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse{}
	return &this
}

// NewSnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponseWithDefaults instantiates a new SnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponseWithDefaults() *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse {
	this := SnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse{}
	return &this
}

// GetJoinQueryHistoryWithAccessHistoryEnabled returns the JoinQueryHistoryWithAccessHistoryEnabled field value if set, zero value otherwise.
func (o *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse) GetJoinQueryHistoryWithAccessHistoryEnabled() bool {
	if o == nil || o.JoinQueryHistoryWithAccessHistoryEnabled == nil {
		var ret bool
		return ret
	}
	return *o.JoinQueryHistoryWithAccessHistoryEnabled
}

// GetJoinQueryHistoryWithAccessHistoryEnabledOk returns a tuple with the JoinQueryHistoryWithAccessHistoryEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse) GetJoinQueryHistoryWithAccessHistoryEnabledOk() (*bool, bool) {
	if o == nil || o.JoinQueryHistoryWithAccessHistoryEnabled == nil {
		return nil, false
	}
	return o.JoinQueryHistoryWithAccessHistoryEnabled, true
}

// HasJoinQueryHistoryWithAccessHistoryEnabled returns a boolean if a field has been set.
func (o *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse) HasJoinQueryHistoryWithAccessHistoryEnabled() bool {
	return o != nil && o.JoinQueryHistoryWithAccessHistoryEnabled != nil
}

// SetJoinQueryHistoryWithAccessHistoryEnabled gets a reference to the given bool and assigns it to the JoinQueryHistoryWithAccessHistoryEnabled field.
func (o *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse) SetJoinQueryHistoryWithAccessHistoryEnabled(v bool) {
	o.JoinQueryHistoryWithAccessHistoryEnabled = &v
}

// GetQueryHistoryLogsIntervalMin returns the QueryHistoryLogsIntervalMin field value if set, zero value otherwise.
func (o *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse) GetQueryHistoryLogsIntervalMin() int64 {
	if o == nil || o.QueryHistoryLogsIntervalMin == nil {
		var ret int64
		return ret
	}
	return *o.QueryHistoryLogsIntervalMin
}

// GetQueryHistoryLogsIntervalMinOk returns a tuple with the QueryHistoryLogsIntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse) GetQueryHistoryLogsIntervalMinOk() (*int64, bool) {
	if o == nil || o.QueryHistoryLogsIntervalMin == nil {
		return nil, false
	}
	return o.QueryHistoryLogsIntervalMin, true
}

// HasQueryHistoryLogsIntervalMin returns a boolean if a field has been set.
func (o *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse) HasQueryHistoryLogsIntervalMin() bool {
	return o != nil && o.QueryHistoryLogsIntervalMin != nil
}

// SetQueryHistoryLogsIntervalMin gets a reference to the given int64 and assigns it to the QueryHistoryLogsIntervalMin field.
func (o *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse) SetQueryHistoryLogsIntervalMin(v int64) {
	o.QueryHistoryLogsIntervalMin = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeQueryHistoryLogsIntegrationDataflowSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		JoinQueryHistoryWithAccessHistoryEnabled *bool  `json:"join_query_history_with_access_history_enabled,omitempty"`
		QueryHistoryLogsIntervalMin              *int64 `json:"query_history_logs_interval_min,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"join_query_history_with_access_history_enabled", "query_history_logs_interval_min"})
	} else {
		return err
	}
	o.JoinQueryHistoryWithAccessHistoryEnabled = all.JoinQueryHistoryWithAccessHistoryEnabled
	o.QueryHistoryLogsIntervalMin = all.QueryHistoryLogsIntervalMin

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
