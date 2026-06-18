// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// VercelConfigAttributes Attributes of the Datadog Vercel integration configuration.
type VercelConfigAttributes struct {
	// Datadog API key reference used by the Vercel integration to send telemetry.
	ApiKey VercelApiKey `json:"apiKey"`
	// Logs forwarding configuration for the Vercel integration.
	LogsConfig VercelLogsConfig `json:"logsConfig"`
	// Tracing configuration for the Vercel integration.
	TraceConfig VercelTraceConfig `json:"traceConfig"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVercelConfigAttributes instantiates a new VercelConfigAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVercelConfigAttributes(apiKey VercelApiKey, logsConfig VercelLogsConfig, traceConfig VercelTraceConfig) *VercelConfigAttributes {
	this := VercelConfigAttributes{}
	this.ApiKey = apiKey
	this.LogsConfig = logsConfig
	this.TraceConfig = traceConfig
	return &this
}

// NewVercelConfigAttributesWithDefaults instantiates a new VercelConfigAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVercelConfigAttributesWithDefaults() *VercelConfigAttributes {
	this := VercelConfigAttributes{}
	return &this
}

// GetApiKey returns the ApiKey field value.
func (o *VercelConfigAttributes) GetApiKey() VercelApiKey {
	if o == nil {
		var ret VercelApiKey
		return ret
	}
	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *VercelConfigAttributes) GetApiKeyOk() (*VercelApiKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value.
func (o *VercelConfigAttributes) SetApiKey(v VercelApiKey) {
	o.ApiKey = v
}

// GetLogsConfig returns the LogsConfig field value.
func (o *VercelConfigAttributes) GetLogsConfig() VercelLogsConfig {
	if o == nil {
		var ret VercelLogsConfig
		return ret
	}
	return o.LogsConfig
}

// GetLogsConfigOk returns a tuple with the LogsConfig field value
// and a boolean to check if the value has been set.
func (o *VercelConfigAttributes) GetLogsConfigOk() (*VercelLogsConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogsConfig, true
}

// SetLogsConfig sets field value.
func (o *VercelConfigAttributes) SetLogsConfig(v VercelLogsConfig) {
	o.LogsConfig = v
}

// GetTraceConfig returns the TraceConfig field value.
func (o *VercelConfigAttributes) GetTraceConfig() VercelTraceConfig {
	if o == nil {
		var ret VercelTraceConfig
		return ret
	}
	return o.TraceConfig
}

// GetTraceConfigOk returns a tuple with the TraceConfig field value
// and a boolean to check if the value has been set.
func (o *VercelConfigAttributes) GetTraceConfigOk() (*VercelTraceConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TraceConfig, true
}

// SetTraceConfig sets field value.
func (o *VercelConfigAttributes) SetTraceConfig(v VercelTraceConfig) {
	o.TraceConfig = v
}

// MarshalJSON serializes the struct using spec logic.
func (o VercelConfigAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["apiKey"] = o.ApiKey
	toSerialize["logsConfig"] = o.LogsConfig
	toSerialize["traceConfig"] = o.TraceConfig

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *VercelConfigAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiKey      *VercelApiKey      `json:"apiKey"`
		LogsConfig  *VercelLogsConfig  `json:"logsConfig"`
		TraceConfig *VercelTraceConfig `json:"traceConfig"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApiKey == nil {
		return fmt.Errorf("required field apiKey missing")
	}
	if all.LogsConfig == nil {
		return fmt.Errorf("required field logsConfig missing")
	}
	if all.TraceConfig == nil {
		return fmt.Errorf("required field traceConfig missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"apiKey", "logsConfig", "traceConfig"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ApiKey.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ApiKey = *all.ApiKey
	if all.LogsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LogsConfig = *all.LogsConfig
	if all.TraceConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TraceConfig = *all.TraceConfig

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
