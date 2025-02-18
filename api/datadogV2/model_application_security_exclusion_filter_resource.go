// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityExclusionFilterResource A JSON:API resource for an Application Security exclusion filter.
type ApplicationSecurityExclusionFilterResource struct {
	// Attributes describing an Application Security exclusion filter.
	Attributes ApplicationSecurityExclusionFilterAttributes `json:"attributes"`
	// The identifier of the exclusion filter.
	Id *string `json:"id,omitempty"`
	// Extra information about the exclusion filter.
	Meta *ApplicationSecurityExclusionFilterMetadata `json:"meta,omitempty"`
	// Type of the resource. The value should always be `exclusion_filter`.
	Type ApplicationSecurityExclusionFilterType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityExclusionFilterResource instantiates a new ApplicationSecurityExclusionFilterResource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityExclusionFilterResource(attributes ApplicationSecurityExclusionFilterAttributes, typeVar ApplicationSecurityExclusionFilterType) *ApplicationSecurityExclusionFilterResource {
	this := ApplicationSecurityExclusionFilterResource{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewApplicationSecurityExclusionFilterResourceWithDefaults instantiates a new ApplicationSecurityExclusionFilterResource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityExclusionFilterResourceWithDefaults() *ApplicationSecurityExclusionFilterResource {
	this := ApplicationSecurityExclusionFilterResource{}
	var typeVar ApplicationSecurityExclusionFilterType = APPLICATIONSECURITYEXCLUSIONFILTERTYPE_EXCLUSION_FILTER
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ApplicationSecurityExclusionFilterResource) GetAttributes() ApplicationSecurityExclusionFilterAttributes {
	if o == nil {
		var ret ApplicationSecurityExclusionFilterAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterResource) GetAttributesOk() (*ApplicationSecurityExclusionFilterAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ApplicationSecurityExclusionFilterResource) SetAttributes(v ApplicationSecurityExclusionFilterAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterResource) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterResource) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterResource) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationSecurityExclusionFilterResource) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterResource) GetMeta() ApplicationSecurityExclusionFilterMetadata {
	if o == nil || o.Meta == nil {
		var ret ApplicationSecurityExclusionFilterMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterResource) GetMetaOk() (*ApplicationSecurityExclusionFilterMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterResource) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given ApplicationSecurityExclusionFilterMetadata and assigns it to the Meta field.
func (o *ApplicationSecurityExclusionFilterResource) SetMeta(v ApplicationSecurityExclusionFilterMetadata) {
	o.Meta = &v
}

// GetType returns the Type field value.
func (o *ApplicationSecurityExclusionFilterResource) GetType() ApplicationSecurityExclusionFilterType {
	if o == nil {
		var ret ApplicationSecurityExclusionFilterType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterResource) GetTypeOk() (*ApplicationSecurityExclusionFilterType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ApplicationSecurityExclusionFilterResource) SetType(v ApplicationSecurityExclusionFilterType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityExclusionFilterResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityExclusionFilterResource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ApplicationSecurityExclusionFilterAttributes `json:"attributes"`
		Id         *string                                       `json:"id,omitempty"`
		Meta       *ApplicationSecurityExclusionFilterMetadata   `json:"meta,omitempty"`
		Type       *ApplicationSecurityExclusionFilterType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "meta", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = all.Id
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
