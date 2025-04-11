// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Layer Encapsulates a layer resource, holding attributes like rotation details, plus relationships to the members covering that layer.
type Layer struct {
	// Describes key properties of a Layer, including rotation details, name, start/end times, and any restrictions.
	Attributes *LayerAttributes `json:"attributes,omitempty"`
	// A unique identifier for this layer.
	Id *string `json:"id,omitempty"`
	// Holds references to objects related to the Layer entity, such as its members.
	Relationships *LayerRelationships `json:"relationships,omitempty"`
	// Layers resource type.
	Type *LayerType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLayer instantiates a new Layer object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLayer() *Layer {
	this := Layer{}
	var typeVar LayerType = LAYERTYPE_LAYERS
	this.Type = &typeVar
	return &this
}

// NewLayerWithDefaults instantiates a new Layer object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLayerWithDefaults() *Layer {
	this := Layer{}
	var typeVar LayerType = LAYERTYPE_LAYERS
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Layer) GetAttributes() LayerAttributes {
	if o == nil || o.Attributes == nil {
		var ret LayerAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Layer) GetAttributesOk() (*LayerAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Layer) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given LayerAttributes and assigns it to the Attributes field.
func (o *Layer) SetAttributes(v LayerAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Layer) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Layer) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Layer) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Layer) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *Layer) GetRelationships() LayerRelationships {
	if o == nil || o.Relationships == nil {
		var ret LayerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Layer) GetRelationshipsOk() (*LayerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *Layer) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given LayerRelationships and assigns it to the Relationships field.
func (o *Layer) SetRelationships(v LayerRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Layer) GetType() LayerType {
	if o == nil || o.Type == nil {
		var ret LayerType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Layer) GetTypeOk() (*LayerType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Layer) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given LayerType and assigns it to the Type field.
func (o *Layer) SetType(v LayerType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Layer) MarshalJSON() ([]byte, error) {
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
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
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
func (o *Layer) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *LayerAttributes    `json:"attributes,omitempty"`
		Id            *string             `json:"id,omitempty"`
		Relationships *LayerRelationships `json:"relationships,omitempty"`
		Type          *LayerType          `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
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
