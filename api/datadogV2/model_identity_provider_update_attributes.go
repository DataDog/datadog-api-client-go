// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IdentityProviderUpdateAttributes Attributes for updating an organization identity provider.
type IdentityProviderUpdateAttributes struct {
	// Whether to enable or disable this identity provider for the organization.
	Enabled bool `json:"enabled"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIdentityProviderUpdateAttributes instantiates a new IdentityProviderUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIdentityProviderUpdateAttributes(enabled bool) *IdentityProviderUpdateAttributes {
	this := IdentityProviderUpdateAttributes{}
	this.Enabled = enabled
	return &this
}

// NewIdentityProviderUpdateAttributesWithDefaults instantiates a new IdentityProviderUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIdentityProviderUpdateAttributesWithDefaults() *IdentityProviderUpdateAttributes {
	this := IdentityProviderUpdateAttributes{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *IdentityProviderUpdateAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderUpdateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *IdentityProviderUpdateAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IdentityProviderUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IdentityProviderUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled *bool `json:"enabled"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled"})
	} else {
		return err
	}
	o.Enabled = *all.Enabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
