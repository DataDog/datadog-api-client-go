// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceNowBasicAuth The definition of the `ServiceNowBasicAuth` object.
type ServiceNowBasicAuth struct {
	// The `ServiceNowBasicAuth` `instance`.
	Instance ServiceNowBasicAuthInstance `json:"instance"`
	// The `ServiceNowBasicAuth` `password`.
	Password ServiceNowBasicAuthPassword `json:"password"`
	// The definition of the `ServiceNowBasicAuth` object.
	Type ServiceNowBasicAuthType `json:"type"`
	// The `ServiceNowBasicAuth` `username`.
	Username ServiceNowBasicAuthUsername `json:"username"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceNowBasicAuth instantiates a new ServiceNowBasicAuth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceNowBasicAuth(instance ServiceNowBasicAuthInstance, password ServiceNowBasicAuthPassword, typeVar ServiceNowBasicAuthType, username ServiceNowBasicAuthUsername) *ServiceNowBasicAuth {
	this := ServiceNowBasicAuth{}
	this.Instance = instance
	this.Password = password
	this.Type = typeVar
	this.Username = username
	return &this
}

// NewServiceNowBasicAuthWithDefaults instantiates a new ServiceNowBasicAuth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceNowBasicAuthWithDefaults() *ServiceNowBasicAuth {
	this := ServiceNowBasicAuth{}
	return &this
}

// GetInstance returns the Instance field value.
func (o *ServiceNowBasicAuth) GetInstance() ServiceNowBasicAuthInstance {
	if o == nil {
		var ret ServiceNowBasicAuthInstance
		return ret
	}
	return o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value
// and a boolean to check if the value has been set.
func (o *ServiceNowBasicAuth) GetInstanceOk() (*ServiceNowBasicAuthInstance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instance, true
}

// SetInstance sets field value.
func (o *ServiceNowBasicAuth) SetInstance(v ServiceNowBasicAuthInstance) {
	o.Instance = v
}

// GetPassword returns the Password field value.
func (o *ServiceNowBasicAuth) GetPassword() ServiceNowBasicAuthPassword {
	if o == nil {
		var ret ServiceNowBasicAuthPassword
		return ret
	}
	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ServiceNowBasicAuth) GetPasswordOk() (*ServiceNowBasicAuthPassword, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value.
func (o *ServiceNowBasicAuth) SetPassword(v ServiceNowBasicAuthPassword) {
	o.Password = v
}

// GetType returns the Type field value.
func (o *ServiceNowBasicAuth) GetType() ServiceNowBasicAuthType {
	if o == nil {
		var ret ServiceNowBasicAuthType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceNowBasicAuth) GetTypeOk() (*ServiceNowBasicAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ServiceNowBasicAuth) SetType(v ServiceNowBasicAuthType) {
	o.Type = v
}

// GetUsername returns the Username field value.
func (o *ServiceNowBasicAuth) GetUsername() ServiceNowBasicAuthUsername {
	if o == nil {
		var ret ServiceNowBasicAuthUsername
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ServiceNowBasicAuth) GetUsernameOk() (*ServiceNowBasicAuthUsername, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value.
func (o *ServiceNowBasicAuth) SetUsername(v ServiceNowBasicAuthUsername) {
	o.Username = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceNowBasicAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["instance"] = o.Instance
	toSerialize["password"] = o.Password
	toSerialize["type"] = o.Type
	toSerialize["username"] = o.Username

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceNowBasicAuth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Instance *ServiceNowBasicAuthInstance `json:"instance"`
		Password *ServiceNowBasicAuthPassword `json:"password"`
		Type     *ServiceNowBasicAuthType     `json:"type"`
		Username *ServiceNowBasicAuthUsername `json:"username"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Instance == nil {
		return fmt.Errorf("required field instance missing")
	}
	if all.Password == nil {
		return fmt.Errorf("required field password missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Username == nil {
		return fmt.Errorf("required field username missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"instance", "password", "type", "username"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Instance = *all.Instance
	o.Password = *all.Password
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Username = *all.Username

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
