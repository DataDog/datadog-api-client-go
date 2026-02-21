// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserIncludedAttributes Attributes of an included user.
type UserIncludedAttributes struct {
	// The email address of the user.
	Email string `json:"email"`
	// The handle of the user.
	Handle string `json:"handle"`
	// The icon URL for the user.
	Icon string `json:"icon"`
	// The name of the user.
	Name string `json:"name"`
	// The UUID of the user.
	Uuid uuid.UUID `json:"uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserIncludedAttributes instantiates a new UserIncludedAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserIncludedAttributes(email string, handle string, icon string, name string, uuid uuid.UUID) *UserIncludedAttributes {
	this := UserIncludedAttributes{}
	this.Email = email
	this.Handle = handle
	this.Icon = icon
	this.Name = name
	this.Uuid = uuid
	return &this
}

// NewUserIncludedAttributesWithDefaults instantiates a new UserIncludedAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserIncludedAttributesWithDefaults() *UserIncludedAttributes {
	this := UserIncludedAttributes{}
	return &this
}

// GetEmail returns the Email field value.
func (o *UserIncludedAttributes) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserIncludedAttributes) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value.
func (o *UserIncludedAttributes) SetEmail(v string) {
	o.Email = v
}

// GetHandle returns the Handle field value.
func (o *UserIncludedAttributes) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *UserIncludedAttributes) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *UserIncludedAttributes) SetHandle(v string) {
	o.Handle = v
}

// GetIcon returns the Icon field value.
func (o *UserIncludedAttributes) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *UserIncludedAttributes) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value.
func (o *UserIncludedAttributes) SetIcon(v string) {
	o.Icon = v
}

// GetName returns the Name field value.
func (o *UserIncludedAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserIncludedAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *UserIncludedAttributes) SetName(v string) {
	o.Name = v
}

// GetUuid returns the Uuid field value.
func (o *UserIncludedAttributes) GetUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *UserIncludedAttributes) GetUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value.
func (o *UserIncludedAttributes) SetUuid(v uuid.UUID) {
	o.Uuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserIncludedAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["email"] = o.Email
	toSerialize["handle"] = o.Handle
	toSerialize["icon"] = o.Icon
	toSerialize["name"] = o.Name
	toSerialize["uuid"] = o.Uuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserIncludedAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Email  *string    `json:"email"`
		Handle *string    `json:"handle"`
		Icon   *string    `json:"icon"`
		Name   *string    `json:"name"`
		Uuid   *uuid.UUID `json:"uuid"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Email == nil {
		return fmt.Errorf("required field email missing")
	}
	if all.Handle == nil {
		return fmt.Errorf("required field handle missing")
	}
	if all.Icon == nil {
		return fmt.Errorf("required field icon missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Uuid == nil {
		return fmt.Errorf("required field uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"email", "handle", "icon", "name", "uuid"})
	} else {
		return err
	}
	o.Email = *all.Email
	o.Handle = *all.Handle
	o.Icon = *all.Icon
	o.Name = *all.Name
	o.Uuid = *all.Uuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
