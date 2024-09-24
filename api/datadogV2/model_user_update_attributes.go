// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserUpdateAttributes Attributes of the edited user.
type UserUpdateAttributes struct {
	// The `UserUpdateAttributes` `created_at`.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// If the user is enabled or disabled.
	Disabled *bool `json:"disabled,omitempty"`
	// The email of the user.
	Email *string `json:"email,omitempty"`
	// The `UserUpdateAttributes` `handle`.
	Handle *string `json:"handle,omitempty"`
	// The `UserUpdateAttributes` `modified_at`.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The name of the user.
	Name *string `json:"name,omitempty"`
	// The `UserUpdateAttributes` `service_account`.
	ServiceAccount *bool `json:"service_account,omitempty"`
	// The `UserUpdateAttributes` `title`.
	Title datadog.NullableString `json:"title,omitempty"`
	// The `UserUpdateAttributes` `verified`.
	Verified *bool `json:"verified,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserUpdateAttributes instantiates a new UserUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserUpdateAttributes() *UserUpdateAttributes {
	this := UserUpdateAttributes{}
	return &this
}

// NewUserUpdateAttributesWithDefaults instantiates a new UserUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserUpdateAttributesWithDefaults() *UserUpdateAttributes {
	this := UserUpdateAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UserUpdateAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserUpdateAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *UserUpdateAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *UserUpdateAttributes) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateAttributes) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *UserUpdateAttributes) HasDisabled() bool {
	return o != nil && o.Disabled != nil
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *UserUpdateAttributes) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserUpdateAttributes) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateAttributes) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserUpdateAttributes) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserUpdateAttributes) SetEmail(v string) {
	o.Email = &v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *UserUpdateAttributes) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateAttributes) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *UserUpdateAttributes) HasHandle() bool {
	return o != nil && o.Handle != nil
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *UserUpdateAttributes) SetHandle(v string) {
	o.Handle = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *UserUpdateAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *UserUpdateAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *UserUpdateAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetServiceAccount returns the ServiceAccount field value if set, zero value otherwise.
func (o *UserUpdateAttributes) GetServiceAccount() bool {
	if o == nil || o.ServiceAccount == nil {
		var ret bool
		return ret
	}
	return *o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateAttributes) GetServiceAccountOk() (*bool, bool) {
	if o == nil || o.ServiceAccount == nil {
		return nil, false
	}
	return o.ServiceAccount, true
}

// HasServiceAccount returns a boolean if a field has been set.
func (o *UserUpdateAttributes) HasServiceAccount() bool {
	return o != nil && o.ServiceAccount != nil
}

// SetServiceAccount gets a reference to the given bool and assigns it to the ServiceAccount field.
func (o *UserUpdateAttributes) SetServiceAccount(v bool) {
	o.ServiceAccount = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUpdateAttributes) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UserUpdateAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *UserUpdateAttributes) HasTitle() bool {
	return o != nil && o.Title.IsSet()
}

// SetTitle gets a reference to the given datadog.NullableString and assigns it to the Title field.
func (o *UserUpdateAttributes) SetTitle(v string) {
	o.Title.Set(&v)
}

// SetTitleNil sets the value for Title to be an explicit nil.
func (o *UserUpdateAttributes) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil.
func (o *UserUpdateAttributes) UnsetTitle() {
	o.Title.Unset()
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *UserUpdateAttributes) GetVerified() bool {
	if o == nil || o.Verified == nil {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateAttributes) GetVerifiedOk() (*bool, bool) {
	if o == nil || o.Verified == nil {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *UserUpdateAttributes) HasVerified() bool {
	return o != nil && o.Verified != nil
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *UserUpdateAttributes) SetVerified(v bool) {
	o.Verified = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
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
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ServiceAccount != nil {
		toSerialize["service_account"] = o.ServiceAccount
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Verified != nil {
		toSerialize["verified"] = o.Verified
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt      *time.Time             `json:"created_at,omitempty"`
		Disabled       *bool                  `json:"disabled,omitempty"`
		Email          *string                `json:"email,omitempty"`
		Handle         *string                `json:"handle,omitempty"`
		ModifiedAt     *time.Time             `json:"modified_at,omitempty"`
		Name           *string                `json:"name,omitempty"`
		ServiceAccount *bool                  `json:"service_account,omitempty"`
		Title          datadog.NullableString `json:"title,omitempty"`
		Verified       *bool                  `json:"verified,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "disabled", "email", "handle", "modified_at", "name", "service_account", "title", "verified"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.Disabled = all.Disabled
	o.Email = all.Email
	o.Handle = all.Handle
	o.ModifiedAt = all.ModifiedAt
	o.Name = all.Name
	o.ServiceAccount = all.ServiceAccount
	o.Title = all.Title
	o.Verified = all.Verified

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
