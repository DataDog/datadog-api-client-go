// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserOverrideIdentityProviderAttributes The definition of `UserOverrideIdentityProviderAttributes` object.
type UserOverrideIdentityProviderAttributes struct {
	// The `UserOverrideIdentityProviderAttributes` `authentication_method`.
	AuthenticationMethod *string `json:"authentication_method,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserOverrideIdentityProviderAttributes instantiates a new UserOverrideIdentityProviderAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserOverrideIdentityProviderAttributes() *UserOverrideIdentityProviderAttributes {
	this := UserOverrideIdentityProviderAttributes{}
	return &this
}

// NewUserOverrideIdentityProviderAttributesWithDefaults instantiates a new UserOverrideIdentityProviderAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserOverrideIdentityProviderAttributesWithDefaults() *UserOverrideIdentityProviderAttributes {
	this := UserOverrideIdentityProviderAttributes{}
	return &this
}

// GetAuthenticationMethod returns the AuthenticationMethod field value if set, zero value otherwise.
func (o *UserOverrideIdentityProviderAttributes) GetAuthenticationMethod() string {
	if o == nil || o.AuthenticationMethod == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOverrideIdentityProviderAttributes) GetAuthenticationMethodOk() (*string, bool) {
	if o == nil || o.AuthenticationMethod == nil {
		return nil, false
	}
	return o.AuthenticationMethod, true
}

// HasAuthenticationMethod returns a boolean if a field has been set.
func (o *UserOverrideIdentityProviderAttributes) HasAuthenticationMethod() bool {
	return o != nil && o.AuthenticationMethod != nil
}

// SetAuthenticationMethod gets a reference to the given string and assigns it to the AuthenticationMethod field.
func (o *UserOverrideIdentityProviderAttributes) SetAuthenticationMethod(v string) {
	o.AuthenticationMethod = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserOverrideIdentityProviderAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AuthenticationMethod != nil {
		toSerialize["authentication_method"] = o.AuthenticationMethod
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserOverrideIdentityProviderAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthenticationMethod *string `json:"authentication_method,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"authentication_method"})
	} else {
		return err
	}
	o.AuthenticationMethod = all.AuthenticationMethod

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
