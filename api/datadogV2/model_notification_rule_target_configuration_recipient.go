// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationRuleTargetConfigurationRecipient Recipient for an `EMAIL` target.
type NotificationRuleTargetConfigurationRecipient struct {
	// Email address to notify.
	Email *string `json:"email,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNotificationRuleTargetConfigurationRecipient instantiates a new NotificationRuleTargetConfigurationRecipient object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotificationRuleTargetConfigurationRecipient() *NotificationRuleTargetConfigurationRecipient {
	this := NotificationRuleTargetConfigurationRecipient{}
	return &this
}

// NewNotificationRuleTargetConfigurationRecipientWithDefaults instantiates a new NotificationRuleTargetConfigurationRecipient object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotificationRuleTargetConfigurationRecipientWithDefaults() *NotificationRuleTargetConfigurationRecipient {
	this := NotificationRuleTargetConfigurationRecipient{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *NotificationRuleTargetConfigurationRecipient) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRuleTargetConfigurationRecipient) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *NotificationRuleTargetConfigurationRecipient) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *NotificationRuleTargetConfigurationRecipient) SetEmail(v string) {
	o.Email = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotificationRuleTargetConfigurationRecipient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotificationRuleTargetConfigurationRecipient) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Email *string `json:"email,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"email"})
	} else {
		return err
	}
	o.Email = all.Email

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
