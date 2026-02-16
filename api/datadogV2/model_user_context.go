// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserContext
type UserContext struct {
	//
	UserInfo UserInfo `json:"userInfo"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserContext instantiates a new UserContext object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserContext(userInfo UserInfo) *UserContext {
	this := UserContext{}
	this.UserInfo = userInfo
	return &this
}

// NewUserContextWithDefaults instantiates a new UserContext object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserContextWithDefaults() *UserContext {
	this := UserContext{}
	return &this
}

// GetUserInfo returns the UserInfo field value.
func (o *UserContext) GetUserInfo() UserInfo {
	if o == nil {
		var ret UserInfo
		return ret
	}
	return o.UserInfo
}

// GetUserInfoOk returns a tuple with the UserInfo field value
// and a boolean to check if the value has been set.
func (o *UserContext) GetUserInfoOk() (*UserInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserInfo, true
}

// SetUserInfo sets field value.
func (o *UserContext) SetUserInfo(v UserInfo) {
	o.UserInfo = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["userInfo"] = o.UserInfo

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserContext) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		UserInfo *UserInfo `json:"userInfo"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.UserInfo == nil {
		return fmt.Errorf("required field userInfo missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"userInfo"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.UserInfo.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UserInfo = *all.UserInfo

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
