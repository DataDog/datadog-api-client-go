// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorNotificationRuleBundleConfig Use bundle config to enable alert bundling to reduce monitor signal noises. **Note**: This feature is in preview and is subject to change.
// If you have any feedback, contact [Datadog support](https://docs.datadoghq.com/help/).
type MonitorNotificationRuleBundleConfig struct {
	// Duration of the bundling period.
	Duration int32 `json:"duration"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorNotificationRuleBundleConfig instantiates a new MonitorNotificationRuleBundleConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorNotificationRuleBundleConfig(duration int32) *MonitorNotificationRuleBundleConfig {
	this := MonitorNotificationRuleBundleConfig{}
	this.Duration = duration
	return &this
}

// NewMonitorNotificationRuleBundleConfigWithDefaults instantiates a new MonitorNotificationRuleBundleConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorNotificationRuleBundleConfigWithDefaults() *MonitorNotificationRuleBundleConfig {
	this := MonitorNotificationRuleBundleConfig{}
	return &this
}

// GetDuration returns the Duration field value.
func (o *MonitorNotificationRuleBundleConfig) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleBundleConfig) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value.
func (o *MonitorNotificationRuleBundleConfig) SetDuration(v int32) {
	o.Duration = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorNotificationRuleBundleConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["duration"] = o.Duration

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorNotificationRuleBundleConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Duration *int32 `json:"duration"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Duration == nil {
		return fmt.Errorf("required field duration missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"duration"})
	} else {
		return err
	}
	o.Duration = *all.Duration

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
