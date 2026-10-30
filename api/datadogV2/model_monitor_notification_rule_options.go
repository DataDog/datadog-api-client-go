// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorNotificationRuleOptions Additional options for the monitor notification rule.
type MonitorNotificationRuleOptions struct {
	// Whether notifications from this rule are threaded. When `true`, follow-up notifications for the same monitor alert are posted in Slack as replies to the first notification.
	// Cannot be `true` when `bundle_config` is set.
	IsThreaded bool `json:"is_threaded"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorNotificationRuleOptions instantiates a new MonitorNotificationRuleOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorNotificationRuleOptions(isThreaded bool) *MonitorNotificationRuleOptions {
	this := MonitorNotificationRuleOptions{}
	this.IsThreaded = isThreaded
	return &this
}

// NewMonitorNotificationRuleOptionsWithDefaults instantiates a new MonitorNotificationRuleOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorNotificationRuleOptionsWithDefaults() *MonitorNotificationRuleOptions {
	this := MonitorNotificationRuleOptions{}
	return &this
}

// GetIsThreaded returns the IsThreaded field value.
func (o *MonitorNotificationRuleOptions) GetIsThreaded() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsThreaded
}

// GetIsThreadedOk returns a tuple with the IsThreaded field value
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleOptions) GetIsThreadedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsThreaded, true
}

// SetIsThreaded sets field value.
func (o *MonitorNotificationRuleOptions) SetIsThreaded(v bool) {
	o.IsThreaded = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorNotificationRuleOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["is_threaded"] = o.IsThreaded

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorNotificationRuleOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsThreaded *bool `json:"is_threaded"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IsThreaded == nil {
		return fmt.Errorf("required field is_threaded missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_threaded"})
	} else {
		return err
	}
	o.IsThreaded = *all.IsThreaded

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
