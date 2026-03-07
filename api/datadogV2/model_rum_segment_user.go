// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSegmentUser A user who performed an action on a segment.
type RumSegmentUser struct {
	// The email handle of the user.
	Handle string `json:"handle"`
	// The URL of the user icon.
	Icon string `json:"icon"`
	// The numeric identifier of the user.
	Id string `json:"id"`
	// The display name of the user.
	Name string `json:"name"`
	// The unique identifier of the user.
	Uuid string `json:"uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumSegmentUser instantiates a new RumSegmentUser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumSegmentUser(handle string, icon string, id string, name string, uuid string) *RumSegmentUser {
	this := RumSegmentUser{}
	this.Handle = handle
	this.Icon = icon
	this.Id = id
	this.Name = name
	this.Uuid = uuid
	return &this
}

// NewRumSegmentUserWithDefaults instantiates a new RumSegmentUser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumSegmentUserWithDefaults() *RumSegmentUser {
	this := RumSegmentUser{}
	return &this
}

// GetHandle returns the Handle field value.
func (o *RumSegmentUser) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *RumSegmentUser) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *RumSegmentUser) SetHandle(v string) {
	o.Handle = v
}

// GetIcon returns the Icon field value.
func (o *RumSegmentUser) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *RumSegmentUser) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value.
func (o *RumSegmentUser) SetIcon(v string) {
	o.Icon = v
}

// GetId returns the Id field value.
func (o *RumSegmentUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RumSegmentUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *RumSegmentUser) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *RumSegmentUser) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RumSegmentUser) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RumSegmentUser) SetName(v string) {
	o.Name = v
}

// GetUuid returns the Uuid field value.
func (o *RumSegmentUser) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *RumSegmentUser) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value.
func (o *RumSegmentUser) SetUuid(v string) {
	o.Uuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumSegmentUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["handle"] = o.Handle
	toSerialize["icon"] = o.Icon
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["uuid"] = o.Uuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumSegmentUser) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Handle *string `json:"handle"`
		Icon   *string `json:"icon"`
		Id     *string `json:"id"`
		Name   *string `json:"name"`
		Uuid   *string `json:"uuid"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Handle == nil {
		return fmt.Errorf("required field handle missing")
	}
	if all.Icon == nil {
		return fmt.Errorf("required field icon missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Uuid == nil {
		return fmt.Errorf("required field uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"handle", "icon", "id", "name", "uuid"})
	} else {
		return err
	}
	o.Handle = *all.Handle
	o.Icon = *all.Icon
	o.Id = *all.Id
	o.Name = *all.Name
	o.Uuid = *all.Uuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
