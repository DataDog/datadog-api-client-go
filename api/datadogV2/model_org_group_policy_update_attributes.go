// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicyUpdateAttributes Attributes for updating an org group policy.
type OrgGroupPolicyUpdateAttributes struct {
	// The policy content as key-value pairs.
	Content map[string]interface{} `json:"content,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupPolicyUpdateAttributes instantiates a new OrgGroupPolicyUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupPolicyUpdateAttributes() *OrgGroupPolicyUpdateAttributes {
	this := OrgGroupPolicyUpdateAttributes{}
	return &this
}

// NewOrgGroupPolicyUpdateAttributesWithDefaults instantiates a new OrgGroupPolicyUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupPolicyUpdateAttributesWithDefaults() *OrgGroupPolicyUpdateAttributes {
	this := OrgGroupPolicyUpdateAttributes{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *OrgGroupPolicyUpdateAttributes) GetContent() map[string]interface{} {
	if o == nil || o.Content == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyUpdateAttributes) GetContentOk() (*map[string]interface{}, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return &o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *OrgGroupPolicyUpdateAttributes) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given map[string]interface{} and assigns it to the Content field.
func (o *OrgGroupPolicyUpdateAttributes) SetContent(v map[string]interface{}) {
	o.Content = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupPolicyUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupPolicyUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Content map[string]interface{} `json:"content,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content"})
	} else {
		return err
	}
	o.Content = all.Content

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
