// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserTeamUserAttributes The definition of `UserTeamUserAttributes` object.
type UserTeamUserAttributes struct {
	// The `UserTeamUserAttributes` `disabled`.
	Disabled *bool `json:"disabled,omitempty"`
	// The `UserTeamUserAttributes` `email`.
	Email *string `json:"email,omitempty"`
	// The `UserTeamUserAttributes` `handle`.
	Handle *string `json:"handle,omitempty"`
	// The `UserTeamUserAttributes` `icon`.
	Icon *string `json:"icon,omitempty"`
	// The `UserTeamUserAttributes` `name`.
	Name datadog.NullableString `json:"name,omitempty"`
	// The `UserTeamUserAttributes` `service_account`.
	ServiceAccount *bool `json:"service_account,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserTeamUserAttributes instantiates a new UserTeamUserAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserTeamUserAttributes() *UserTeamUserAttributes {
	this := UserTeamUserAttributes{}
	return &this
}

// NewUserTeamUserAttributesWithDefaults instantiates a new UserTeamUserAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserTeamUserAttributesWithDefaults() *UserTeamUserAttributes {
	this := UserTeamUserAttributes{}
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *UserTeamUserAttributes) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTeamUserAttributes) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *UserTeamUserAttributes) HasDisabled() bool {
	return o != nil && o.Disabled != nil
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *UserTeamUserAttributes) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserTeamUserAttributes) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTeamUserAttributes) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserTeamUserAttributes) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserTeamUserAttributes) SetEmail(v string) {
	o.Email = &v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *UserTeamUserAttributes) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTeamUserAttributes) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *UserTeamUserAttributes) HasHandle() bool {
	return o != nil && o.Handle != nil
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *UserTeamUserAttributes) SetHandle(v string) {
	o.Handle = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *UserTeamUserAttributes) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTeamUserAttributes) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *UserTeamUserAttributes) HasIcon() bool {
	return o != nil && o.Icon != nil
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *UserTeamUserAttributes) SetIcon(v string) {
	o.Icon = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserTeamUserAttributes) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UserTeamUserAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UserTeamUserAttributes) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given datadog.NullableString and assigns it to the Name field.
func (o *UserTeamUserAttributes) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *UserTeamUserAttributes) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *UserTeamUserAttributes) UnsetName() {
	o.Name.Unset()
}

// GetServiceAccount returns the ServiceAccount field value if set, zero value otherwise.
func (o *UserTeamUserAttributes) GetServiceAccount() bool {
	if o == nil || o.ServiceAccount == nil {
		var ret bool
		return ret
	}
	return *o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTeamUserAttributes) GetServiceAccountOk() (*bool, bool) {
	if o == nil || o.ServiceAccount == nil {
		return nil, false
	}
	return o.ServiceAccount, true
}

// HasServiceAccount returns a boolean if a field has been set.
func (o *UserTeamUserAttributes) HasServiceAccount() bool {
	return o != nil && o.ServiceAccount != nil
}

// SetServiceAccount gets a reference to the given bool and assigns it to the ServiceAccount field.
func (o *UserTeamUserAttributes) SetServiceAccount(v bool) {
	o.ServiceAccount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserTeamUserAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ServiceAccount != nil {
		toSerialize["service_account"] = o.ServiceAccount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserTeamUserAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Disabled       *bool                  `json:"disabled,omitempty"`
		Email          *string                `json:"email,omitempty"`
		Handle         *string                `json:"handle,omitempty"`
		Icon           *string                `json:"icon,omitempty"`
		Name           datadog.NullableString `json:"name,omitempty"`
		ServiceAccount *bool                  `json:"service_account,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"disabled", "email", "handle", "icon", "name", "service_account"})
	} else {
		return err
	}
	o.Disabled = all.Disabled
	o.Email = all.Email
	o.Handle = all.Handle
	o.Icon = all.Icon
	o.Name = all.Name
	o.ServiceAccount = all.ServiceAccount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
