/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */


package datadog

import (
	"encoding/json"
	"fmt"

)


// SyntheticsPrivateLocationSecretsAuthentication Authentication part of the secrets.
type SyntheticsPrivateLocationSecretsAuthentication struct {
	// Access key for the private location.
	Id *string `json:"id,omitempty"`
	// Secret access key for the private location.
	Key *string `json:"key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}



// NewSyntheticsPrivateLocationSecretsAuthentication instantiates a new SyntheticsPrivateLocationSecretsAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsPrivateLocationSecretsAuthentication() *SyntheticsPrivateLocationSecretsAuthentication {
	this := SyntheticsPrivateLocationSecretsAuthentication{}
	return &this
}

// NewSyntheticsPrivateLocationSecretsAuthenticationWithDefaults instantiates a new SyntheticsPrivateLocationSecretsAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsPrivateLocationSecretsAuthenticationWithDefaults() *SyntheticsPrivateLocationSecretsAuthentication {
	this := SyntheticsPrivateLocationSecretsAuthentication{}
	return &this
}
// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsPrivateLocationSecretsAuthentication) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsPrivateLocationSecretsAuthentication) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsPrivateLocationSecretsAuthentication) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticsPrivateLocationSecretsAuthentication) SetId(v string) {
	o.Id = &v
}


// GetKey returns the Key field value if set, zero value otherwise.
func (o *SyntheticsPrivateLocationSecretsAuthentication) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsPrivateLocationSecretsAuthentication) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SyntheticsPrivateLocationSecretsAuthentication) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SyntheticsPrivateLocationSecretsAuthentication) SetKey(v string) {
	o.Key = &v
}



func (o SyntheticsPrivateLocationSecretsAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}


func (o *SyntheticsPrivateLocationSecretsAuthentication) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Id *string `json:"id,omitempty"`
		Key *string `json:"key,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Id = all.Id
	o.Key = all.Key
	return nil
}
