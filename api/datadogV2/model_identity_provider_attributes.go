// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IdentityProviderAttributes Attributes of an organization identity provider.
type IdentityProviderAttributes struct {
	// The authentication method used by this identity provider.
	AuthenticationMethod string `json:"authentication_method"`
	// Whether this identity provider is enabled for the organization.
	Enabled bool `json:"enabled"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIdentityProviderAttributes instantiates a new IdentityProviderAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIdentityProviderAttributes(authenticationMethod string, enabled bool) *IdentityProviderAttributes {
	this := IdentityProviderAttributes{}
	this.AuthenticationMethod = authenticationMethod
	this.Enabled = enabled
	return &this
}

// NewIdentityProviderAttributesWithDefaults instantiates a new IdentityProviderAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIdentityProviderAttributesWithDefaults() *IdentityProviderAttributes {
	this := IdentityProviderAttributes{}
	return &this
}

// GetAuthenticationMethod returns the AuthenticationMethod field value.
func (o *IdentityProviderAttributes) GetAuthenticationMethod() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderAttributes) GetAuthenticationMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationMethod, true
}

// SetAuthenticationMethod sets field value.
func (o *IdentityProviderAttributes) SetAuthenticationMethod(v string) {
	o.AuthenticationMethod = v
}

// GetEnabled returns the Enabled field value.
func (o *IdentityProviderAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *IdentityProviderAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IdentityProviderAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["authentication_method"] = o.AuthenticationMethod
	toSerialize["enabled"] = o.Enabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IdentityProviderAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthenticationMethod *string `json:"authentication_method"`
		Enabled              *bool   `json:"enabled"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthenticationMethod == nil {
		return fmt.Errorf("required field authentication_method missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"authentication_method", "enabled"})
	} else {
		return err
	}
	o.AuthenticationMethod = *all.AuthenticationMethod
	o.Enabled = *all.Enabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
