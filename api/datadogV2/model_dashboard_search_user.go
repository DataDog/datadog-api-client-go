// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardSearchUser User information.
type DashboardSearchUser struct {
	// User handle or email.
	Handle string `json:"handle"`
	// User display name.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDashboardSearchUser instantiates a new DashboardSearchUser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardSearchUser(handle string, name string) *DashboardSearchUser {
	this := DashboardSearchUser{}
	this.Handle = handle
	this.Name = name
	return &this
}

// NewDashboardSearchUserWithDefaults instantiates a new DashboardSearchUser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardSearchUserWithDefaults() *DashboardSearchUser {
	this := DashboardSearchUser{}
	return &this
}

// GetHandle returns the Handle field value.
func (o *DashboardSearchUser) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchUser) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *DashboardSearchUser) SetHandle(v string) {
	o.Handle = v
}

// GetName returns the Name field value.
func (o *DashboardSearchUser) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchUser) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DashboardSearchUser) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardSearchUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["handle"] = o.Handle
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardSearchUser) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Handle *string `json:"handle"`
		Name   *string `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Handle == nil {
		return fmt.Errorf("required field handle missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"handle", "name"})
	} else {
		return err
	}
	o.Handle = *all.Handle
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
