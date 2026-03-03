// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SamlConfigurationsUpdateAttributes Attributes for updating SAML preferences.
type SamlConfigurationsUpdateAttributes struct {
	// List of role UUIDs to assign to JIT-provisioned users. Exactly 1 role is required.
	DefaultRoleUuids []uuid.UUID `json:"default_role_uuids"`
	// List of domains for Just-In-Time user provisioning. Maximum 50 domains.
	// Each domain must be between 1 and 256 characters.
	JitDomains []string `json:"jit_domains"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSamlConfigurationsUpdateAttributes instantiates a new SamlConfigurationsUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSamlConfigurationsUpdateAttributes(defaultRoleUuids []uuid.UUID, jitDomains []string) *SamlConfigurationsUpdateAttributes {
	this := SamlConfigurationsUpdateAttributes{}
	this.DefaultRoleUuids = defaultRoleUuids
	this.JitDomains = jitDomains
	return &this
}

// NewSamlConfigurationsUpdateAttributesWithDefaults instantiates a new SamlConfigurationsUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSamlConfigurationsUpdateAttributesWithDefaults() *SamlConfigurationsUpdateAttributes {
	this := SamlConfigurationsUpdateAttributes{}
	return &this
}

// GetDefaultRoleUuids returns the DefaultRoleUuids field value.
func (o *SamlConfigurationsUpdateAttributes) GetDefaultRoleUuids() []uuid.UUID {
	if o == nil {
		var ret []uuid.UUID
		return ret
	}
	return o.DefaultRoleUuids
}

// GetDefaultRoleUuidsOk returns a tuple with the DefaultRoleUuids field value
// and a boolean to check if the value has been set.
func (o *SamlConfigurationsUpdateAttributes) GetDefaultRoleUuidsOk() (*[]uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultRoleUuids, true
}

// SetDefaultRoleUuids sets field value.
func (o *SamlConfigurationsUpdateAttributes) SetDefaultRoleUuids(v []uuid.UUID) {
	o.DefaultRoleUuids = v
}

// GetJitDomains returns the JitDomains field value.
func (o *SamlConfigurationsUpdateAttributes) GetJitDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.JitDomains
}

// GetJitDomainsOk returns a tuple with the JitDomains field value
// and a boolean to check if the value has been set.
func (o *SamlConfigurationsUpdateAttributes) GetJitDomainsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JitDomains, true
}

// SetJitDomains sets field value.
func (o *SamlConfigurationsUpdateAttributes) SetJitDomains(v []string) {
	o.JitDomains = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SamlConfigurationsUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["default_role_uuids"] = o.DefaultRoleUuids
	toSerialize["jit_domains"] = o.JitDomains

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SamlConfigurationsUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DefaultRoleUuids *[]uuid.UUID `json:"default_role_uuids"`
		JitDomains       *[]string    `json:"jit_domains"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DefaultRoleUuids == nil {
		return fmt.Errorf("required field default_role_uuids missing")
	}
	if all.JitDomains == nil {
		return fmt.Errorf("required field jit_domains missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"default_role_uuids", "jit_domains"})
	} else {
		return err
	}
	o.DefaultRoleUuids = *all.DefaultRoleUuids
	o.JitDomains = *all.JitDomains

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
