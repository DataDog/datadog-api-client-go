// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserInfo
type UserInfo struct {
	// The organization name.
	OrgName string `json:"orgName"`
	// The user's email address.
	UserEmail string `json:"userEmail"`
	// The user's name.
	UserName *string `json:"userName,omitempty"`
	// The user's UUID.
	UserUuid string `json:"userUUID"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserInfo instantiates a new UserInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserInfo(orgName string, userEmail string, userUuid string) *UserInfo {
	this := UserInfo{}
	this.OrgName = orgName
	this.UserEmail = userEmail
	this.UserUuid = userUuid
	return &this
}

// NewUserInfoWithDefaults instantiates a new UserInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserInfoWithDefaults() *UserInfo {
	this := UserInfo{}
	return &this
}

// GetOrgName returns the OrgName field value.
func (o *UserInfo) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *UserInfo) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value.
func (o *UserInfo) SetOrgName(v string) {
	o.OrgName = v
}

// GetUserEmail returns the UserEmail field value.
func (o *UserInfo) GetUserEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value
// and a boolean to check if the value has been set.
func (o *UserInfo) GetUserEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserEmail, true
}

// SetUserEmail sets field value.
func (o *UserInfo) SetUserEmail(v string) {
	o.UserEmail = v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *UserInfo) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *UserInfo) HasUserName() bool {
	return o != nil && o.UserName != nil
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *UserInfo) SetUserName(v string) {
	o.UserName = &v
}

// GetUserUuid returns the UserUuid field value.
func (o *UserInfo) GetUserUuid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value
// and a boolean to check if the value has been set.
func (o *UserInfo) GetUserUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserUuid, true
}

// SetUserUuid sets field value.
func (o *UserInfo) SetUserUuid(v string) {
	o.UserUuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["orgName"] = o.OrgName
	toSerialize["userEmail"] = o.UserEmail
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}
	toSerialize["userUUID"] = o.UserUuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgName   *string `json:"orgName"`
		UserEmail *string `json:"userEmail"`
		UserName  *string `json:"userName,omitempty"`
		UserUuid  *string `json:"userUUID"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.OrgName == nil {
		return fmt.Errorf("required field orgName missing")
	}
	if all.UserEmail == nil {
		return fmt.Errorf("required field userEmail missing")
	}
	if all.UserUuid == nil {
		return fmt.Errorf("required field userUUID missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"orgName", "userEmail", "userName", "userUUID"})
	} else {
		return err
	}
	o.OrgName = *all.OrgName
	o.UserEmail = *all.UserEmail
	o.UserName = all.UserName
	o.UserUuid = *all.UserUuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
