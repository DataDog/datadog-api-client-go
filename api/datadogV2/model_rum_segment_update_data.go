// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSegmentUpdateData Data object for a segment update request.
type RumSegmentUpdateData struct {
	// Attributes for updating a segment. All fields are optional.
	Attributes RumSegmentUpdateAttributes `json:"attributes"`
	// The identifier of the segment to update.
	Id string `json:"id"`
	// Type of the segment resource.
	Type RumSegmentResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumSegmentUpdateData instantiates a new RumSegmentUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumSegmentUpdateData(attributes RumSegmentUpdateAttributes, id string, typeVar RumSegmentResourceType) *RumSegmentUpdateData {
	this := RumSegmentUpdateData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewRumSegmentUpdateDataWithDefaults instantiates a new RumSegmentUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumSegmentUpdateDataWithDefaults() *RumSegmentUpdateData {
	this := RumSegmentUpdateData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *RumSegmentUpdateData) GetAttributes() RumSegmentUpdateAttributes {
	if o == nil {
		var ret RumSegmentUpdateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *RumSegmentUpdateData) GetAttributesOk() (*RumSegmentUpdateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *RumSegmentUpdateData) SetAttributes(v RumSegmentUpdateAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *RumSegmentUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RumSegmentUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *RumSegmentUpdateData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *RumSegmentUpdateData) GetType() RumSegmentResourceType {
	if o == nil {
		var ret RumSegmentResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RumSegmentUpdateData) GetTypeOk() (*RumSegmentResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RumSegmentUpdateData) SetType(v RumSegmentResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumSegmentUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumSegmentUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *RumSegmentUpdateAttributes `json:"attributes"`
		Id         *string                     `json:"id"`
		Type       *RumSegmentResourceType     `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
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
