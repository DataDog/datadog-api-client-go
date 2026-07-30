// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioSettings Twilio interface settings.
type TwilioSettings struct {
	// Twilio Account SID that uniquely identifies your Twilio account.
	AccountSid string `json:"account_sid"`
	// When enabled, phone numbers in the `to` field and SMS message bodies are censored for privacy.
	CensorLogs *bool `json:"censor_logs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTwilioSettings instantiates a new TwilioSettings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTwilioSettings(accountSid string) *TwilioSettings {
	this := TwilioSettings{}
	this.AccountSid = accountSid
	return &this
}

// NewTwilioSettingsWithDefaults instantiates a new TwilioSettings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTwilioSettingsWithDefaults() *TwilioSettings {
	this := TwilioSettings{}
	return &this
}

// GetAccountSid returns the AccountSid field value.
func (o *TwilioSettings) GetAccountSid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountSid
}

// GetAccountSidOk returns a tuple with the AccountSid field value
// and a boolean to check if the value has been set.
func (o *TwilioSettings) GetAccountSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountSid, true
}

// SetAccountSid sets field value.
func (o *TwilioSettings) SetAccountSid(v string) {
	o.AccountSid = v
}

// GetCensorLogs returns the CensorLogs field value if set, zero value otherwise.
func (o *TwilioSettings) GetCensorLogs() bool {
	if o == nil || o.CensorLogs == nil {
		var ret bool
		return ret
	}
	return *o.CensorLogs
}

// GetCensorLogsOk returns a tuple with the CensorLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioSettings) GetCensorLogsOk() (*bool, bool) {
	if o == nil || o.CensorLogs == nil {
		return nil, false
	}
	return o.CensorLogs, true
}

// HasCensorLogs returns a boolean if a field has been set.
func (o *TwilioSettings) HasCensorLogs() bool {
	return o != nil && o.CensorLogs != nil
}

// SetCensorLogs gets a reference to the given bool and assigns it to the CensorLogs field.
func (o *TwilioSettings) SetCensorLogs(v bool) {
	o.CensorLogs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TwilioSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_sid"] = o.AccountSid
	if o.CensorLogs != nil {
		toSerialize["censor_logs"] = o.CensorLogs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TwilioSettings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountSid *string `json:"account_sid"`
		CensorLogs *bool   `json:"censor_logs,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountSid == nil {
		return fmt.Errorf("required field account_sid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_sid", "censor_logs"})
	} else {
		return err
	}
	o.AccountSid = *all.AccountSid
	o.CensorLogs = all.CensorLogs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
