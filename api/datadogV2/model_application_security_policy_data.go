// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityPolicyData Object for a single WAF policy.
type ApplicationSecurityPolicyData struct {
	// A WAF policy.
	Attributes *ApplicationSecurityPolicyAttributes `json:"attributes,omitempty"`
	// The ID of the policy.
	Id *string `json:"id,omitempty"`
	// Metadata associated with the WAF policy.
	Metadata *ApplicationSecurityPolicyMetadata `json:"metadata,omitempty"`
	// The type of the resource. The value should always be `policy`.
	Type *ApplicationSecurityPolicyType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityPolicyData instantiates a new ApplicationSecurityPolicyData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityPolicyData() *ApplicationSecurityPolicyData {
	this := ApplicationSecurityPolicyData{}
	var typeVar ApplicationSecurityPolicyType = APPLICATIONSECURITYPOLICYTYPE_POLICY
	this.Type = &typeVar
	return &this
}

// NewApplicationSecurityPolicyDataWithDefaults instantiates a new ApplicationSecurityPolicyData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityPolicyDataWithDefaults() *ApplicationSecurityPolicyData {
	this := ApplicationSecurityPolicyData{}
	var typeVar ApplicationSecurityPolicyType = APPLICATIONSECURITYPOLICYTYPE_POLICY
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ApplicationSecurityPolicyData) GetAttributes() ApplicationSecurityPolicyAttributes {
	if o == nil || o.Attributes == nil {
		var ret ApplicationSecurityPolicyAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyData) GetAttributesOk() (*ApplicationSecurityPolicyAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ApplicationSecurityPolicyData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given ApplicationSecurityPolicyAttributes and assigns it to the Attributes field.
func (o *ApplicationSecurityPolicyData) SetAttributes(v ApplicationSecurityPolicyAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationSecurityPolicyData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationSecurityPolicyData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationSecurityPolicyData) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ApplicationSecurityPolicyData) GetMetadata() ApplicationSecurityPolicyMetadata {
	if o == nil || o.Metadata == nil {
		var ret ApplicationSecurityPolicyMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyData) GetMetadataOk() (*ApplicationSecurityPolicyMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ApplicationSecurityPolicyData) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given ApplicationSecurityPolicyMetadata and assigns it to the Metadata field.
func (o *ApplicationSecurityPolicyData) SetMetadata(v ApplicationSecurityPolicyMetadata) {
	o.Metadata = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApplicationSecurityPolicyData) GetType() ApplicationSecurityPolicyType {
	if o == nil || o.Type == nil {
		var ret ApplicationSecurityPolicyType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityPolicyData) GetTypeOk() (*ApplicationSecurityPolicyType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApplicationSecurityPolicyData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ApplicationSecurityPolicyType and assigns it to the Type field.
func (o *ApplicationSecurityPolicyData) SetType(v ApplicationSecurityPolicyType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityPolicyData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityPolicyData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ApplicationSecurityPolicyAttributes `json:"attributes,omitempty"`
		Id         *string                              `json:"id,omitempty"`
		Metadata   *ApplicationSecurityPolicyMetadata   `json:"metadata,omitempty"`
		Type       *ApplicationSecurityPolicyType       `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "metadata", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Metadata != nil && all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = all.Metadata
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
