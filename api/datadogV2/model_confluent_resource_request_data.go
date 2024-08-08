// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConfluentResourceRequestData JSON:API request for updating a Confluent resource.
type ConfluentResourceRequestData struct {
	// Attributes object for updating a Confluent resource.
	Attributes ConfluentResourceRequestAttributes `json:"attributes"`
	// The ConfluentResourceRequestData id.
	Id datadog.NullableString `json:"id"`
	// The JSON:API type for this request.
	Type ConfluentResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConfluentResourceRequestData instantiates a new ConfluentResourceRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConfluentResourceRequestData(attributes ConfluentResourceRequestAttributes, id datadog.NullableString, typeVar ConfluentResourceType) *ConfluentResourceRequestData {
	this := ConfluentResourceRequestData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewConfluentResourceRequestDataWithDefaults instantiates a new ConfluentResourceRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConfluentResourceRequestDataWithDefaults() *ConfluentResourceRequestData {
	this := ConfluentResourceRequestData{}
	var typeVar ConfluentResourceType = CONFLUENTRESOURCETYPE_CONFLUENT_CLOUD_RESOURCES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ConfluentResourceRequestData) GetAttributes() ConfluentResourceRequestAttributes {
	if o == nil {
		var ret ConfluentResourceRequestAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ConfluentResourceRequestData) GetAttributesOk() (*ConfluentResourceRequestAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ConfluentResourceRequestData) SetAttributes(v ConfluentResourceRequestAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *ConfluentResourceRequestData) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ConfluentResourceRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value.
func (o *ConfluentResourceRequestData) SetId(v string) {
	o.Id.Set(&v)
}

// GetType returns the Type field value.
func (o *ConfluentResourceRequestData) GetType() ConfluentResourceType {
	if o == nil {
		var ret ConfluentResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConfluentResourceRequestData) GetTypeOk() (*ConfluentResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ConfluentResourceRequestData) SetType(v ConfluentResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConfluentResourceRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id.Get()
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConfluentResourceRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ConfluentResourceRequestAttributes `json:"attributes"`
		Id         datadog.NullableString              `json:"id"`
		Type       *ConfluentResourceType              `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if !all.Id.IsSet() {
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
	o.Id = all.Id
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
