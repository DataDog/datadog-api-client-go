// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeTaskHistoryLogsIntegrationDataflowSettingsRequest Settings of the task history logs dataflow. Only the fields provided are changed.
type SnowflakeTaskHistoryLogsIntegrationDataflowSettingsRequest struct {
	// How often task history logs are collected, in minutes. One of `5`, `15`, `30`, `60`, or `1440`. Defaults to `5`.
	TaskHistoryLogsIntervalMin *int64 `json:"task_history_logs_interval_min,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSnowflakeTaskHistoryLogsIntegrationDataflowSettingsRequest instantiates a new SnowflakeTaskHistoryLogsIntegrationDataflowSettingsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeTaskHistoryLogsIntegrationDataflowSettingsRequest() *SnowflakeTaskHistoryLogsIntegrationDataflowSettingsRequest {
	this := SnowflakeTaskHistoryLogsIntegrationDataflowSettingsRequest{}
	return &this
}

// NewSnowflakeTaskHistoryLogsIntegrationDataflowSettingsRequestWithDefaults instantiates a new SnowflakeTaskHistoryLogsIntegrationDataflowSettingsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeTaskHistoryLogsIntegrationDataflowSettingsRequestWithDefaults() *SnowflakeTaskHistoryLogsIntegrationDataflowSettingsRequest {
	this := SnowflakeTaskHistoryLogsIntegrationDataflowSettingsRequest{}
	return &this
}

// GetTaskHistoryLogsIntervalMin returns the TaskHistoryLogsIntervalMin field value if set, zero value otherwise.
func (o *SnowflakeTaskHistoryLogsIntegrationDataflowSettingsRequest) GetTaskHistoryLogsIntervalMin() int64 {
	if o == nil || o.TaskHistoryLogsIntervalMin == nil {
		var ret int64
		return ret
	}
	return *o.TaskHistoryLogsIntervalMin
}

// GetTaskHistoryLogsIntervalMinOk returns a tuple with the TaskHistoryLogsIntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeTaskHistoryLogsIntegrationDataflowSettingsRequest) GetTaskHistoryLogsIntervalMinOk() (*int64, bool) {
	if o == nil || o.TaskHistoryLogsIntervalMin == nil {
		return nil, false
	}
	return o.TaskHistoryLogsIntervalMin, true
}

// HasTaskHistoryLogsIntervalMin returns a boolean if a field has been set.
func (o *SnowflakeTaskHistoryLogsIntegrationDataflowSettingsRequest) HasTaskHistoryLogsIntervalMin() bool {
	return o != nil && o.TaskHistoryLogsIntervalMin != nil
}

// SetTaskHistoryLogsIntervalMin gets a reference to the given int64 and assigns it to the TaskHistoryLogsIntervalMin field.
func (o *SnowflakeTaskHistoryLogsIntegrationDataflowSettingsRequest) SetTaskHistoryLogsIntervalMin(v int64) {
	o.TaskHistoryLogsIntervalMin = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeTaskHistoryLogsIntegrationDataflowSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TaskHistoryLogsIntervalMin != nil {
		toSerialize["task_history_logs_interval_min"] = o.TaskHistoryLogsIntervalMin
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeTaskHistoryLogsIntegrationDataflowSettingsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TaskHistoryLogsIntervalMin *int64 `json:"task_history_logs_interval_min,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.TaskHistoryLogsIntervalMin = all.TaskHistoryLogsIntervalMin

	return nil
}
