// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeSecurityLogsIntegrationDataflowSettingsRequest Settings of the security logs dataflow. Only the fields provided are changed.
type SnowflakeSecurityLogsIntegrationDataflowSettingsRequest struct {
	// How often security logs are collected, in minutes. One of `5`, `15`, `30`, `60`, `360`, `720`, or `1440`. Defaults to `5`.
	SecurityLogsIntervalMin *int64 `json:"security_logs_interval_min,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSnowflakeSecurityLogsIntegrationDataflowSettingsRequest instantiates a new SnowflakeSecurityLogsIntegrationDataflowSettingsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeSecurityLogsIntegrationDataflowSettingsRequest() *SnowflakeSecurityLogsIntegrationDataflowSettingsRequest {
	this := SnowflakeSecurityLogsIntegrationDataflowSettingsRequest{}
	return &this
}

// NewSnowflakeSecurityLogsIntegrationDataflowSettingsRequestWithDefaults instantiates a new SnowflakeSecurityLogsIntegrationDataflowSettingsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeSecurityLogsIntegrationDataflowSettingsRequestWithDefaults() *SnowflakeSecurityLogsIntegrationDataflowSettingsRequest {
	this := SnowflakeSecurityLogsIntegrationDataflowSettingsRequest{}
	return &this
}

// GetSecurityLogsIntervalMin returns the SecurityLogsIntervalMin field value if set, zero value otherwise.
func (o *SnowflakeSecurityLogsIntegrationDataflowSettingsRequest) GetSecurityLogsIntervalMin() int64 {
	if o == nil || o.SecurityLogsIntervalMin == nil {
		var ret int64
		return ret
	}
	return *o.SecurityLogsIntervalMin
}

// GetSecurityLogsIntervalMinOk returns a tuple with the SecurityLogsIntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeSecurityLogsIntegrationDataflowSettingsRequest) GetSecurityLogsIntervalMinOk() (*int64, bool) {
	if o == nil || o.SecurityLogsIntervalMin == nil {
		return nil, false
	}
	return o.SecurityLogsIntervalMin, true
}

// HasSecurityLogsIntervalMin returns a boolean if a field has been set.
func (o *SnowflakeSecurityLogsIntegrationDataflowSettingsRequest) HasSecurityLogsIntervalMin() bool {
	return o != nil && o.SecurityLogsIntervalMin != nil
}

// SetSecurityLogsIntervalMin gets a reference to the given int64 and assigns it to the SecurityLogsIntervalMin field.
func (o *SnowflakeSecurityLogsIntegrationDataflowSettingsRequest) SetSecurityLogsIntervalMin(v int64) {
	o.SecurityLogsIntervalMin = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeSecurityLogsIntegrationDataflowSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.SecurityLogsIntervalMin != nil {
		toSerialize["security_logs_interval_min"] = o.SecurityLogsIntervalMin
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeSecurityLogsIntegrationDataflowSettingsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SecurityLogsIntervalMin *int64 `json:"security_logs_interval_min,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.SecurityLogsIntervalMin = all.SecurityLogsIntervalMin

	return nil
}
