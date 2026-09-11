// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeSecurityLogsIntegrationDataflowSettingsResponse Settings of the security logs dataflow.
type SnowflakeSecurityLogsIntegrationDataflowSettingsResponse struct {
	// How often security logs are collected, in minutes.
	SecurityLogsIntervalMin *int64 `json:"security_logs_interval_min,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSnowflakeSecurityLogsIntegrationDataflowSettingsResponse instantiates a new SnowflakeSecurityLogsIntegrationDataflowSettingsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeSecurityLogsIntegrationDataflowSettingsResponse() *SnowflakeSecurityLogsIntegrationDataflowSettingsResponse {
	this := SnowflakeSecurityLogsIntegrationDataflowSettingsResponse{}
	return &this
}

// NewSnowflakeSecurityLogsIntegrationDataflowSettingsResponseWithDefaults instantiates a new SnowflakeSecurityLogsIntegrationDataflowSettingsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeSecurityLogsIntegrationDataflowSettingsResponseWithDefaults() *SnowflakeSecurityLogsIntegrationDataflowSettingsResponse {
	this := SnowflakeSecurityLogsIntegrationDataflowSettingsResponse{}
	return &this
}

// GetSecurityLogsIntervalMin returns the SecurityLogsIntervalMin field value if set, zero value otherwise.
func (o *SnowflakeSecurityLogsIntegrationDataflowSettingsResponse) GetSecurityLogsIntervalMin() int64 {
	if o == nil || o.SecurityLogsIntervalMin == nil {
		var ret int64
		return ret
	}
	return *o.SecurityLogsIntervalMin
}

// GetSecurityLogsIntervalMinOk returns a tuple with the SecurityLogsIntervalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeSecurityLogsIntegrationDataflowSettingsResponse) GetSecurityLogsIntervalMinOk() (*int64, bool) {
	if o == nil || o.SecurityLogsIntervalMin == nil {
		return nil, false
	}
	return o.SecurityLogsIntervalMin, true
}

// HasSecurityLogsIntervalMin returns a boolean if a field has been set.
func (o *SnowflakeSecurityLogsIntegrationDataflowSettingsResponse) HasSecurityLogsIntervalMin() bool {
	return o != nil && o.SecurityLogsIntervalMin != nil
}

// SetSecurityLogsIntervalMin gets a reference to the given int64 and assigns it to the SecurityLogsIntervalMin field.
func (o *SnowflakeSecurityLogsIntegrationDataflowSettingsResponse) SetSecurityLogsIntervalMin(v int64) {
	o.SecurityLogsIntervalMin = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeSecurityLogsIntegrationDataflowSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.SecurityLogsIntervalMin != nil {
		toSerialize["security_logs_interval_min"] = o.SecurityLogsIntervalMin
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeSecurityLogsIntegrationDataflowSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SecurityLogsIntervalMin *int64 `json:"security_logs_interval_min,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"security_logs_interval_min"})
	} else {
		return err
	}
	o.SecurityLogsIntervalMin = all.SecurityLogsIntervalMin

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
