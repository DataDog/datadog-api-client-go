// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgAuthorizedClientUpdateAttributes Attributes for updating an org authorized client.
type OrgAuthorizedClientUpdateAttributes struct {
	// Whether to disable or enable this client for the organization.
	Disabled *bool `json:"disabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgAuthorizedClientUpdateAttributes instantiates a new OrgAuthorizedClientUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgAuthorizedClientUpdateAttributes() *OrgAuthorizedClientUpdateAttributes {
	this := OrgAuthorizedClientUpdateAttributes{}
	return &this
}

// NewOrgAuthorizedClientUpdateAttributesWithDefaults instantiates a new OrgAuthorizedClientUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgAuthorizedClientUpdateAttributesWithDefaults() *OrgAuthorizedClientUpdateAttributes {
	this := OrgAuthorizedClientUpdateAttributes{}
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *OrgAuthorizedClientUpdateAttributes) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgAuthorizedClientUpdateAttributes) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *OrgAuthorizedClientUpdateAttributes) HasDisabled() bool {
	return o != nil && o.Disabled != nil
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *OrgAuthorizedClientUpdateAttributes) SetDisabled(v bool) {
	o.Disabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgAuthorizedClientUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgAuthorizedClientUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Disabled *bool `json:"disabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"disabled"})
	} else {
		return err
	}
	o.Disabled = all.Disabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
