// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationWithId The custom destination object with an already assigned ID.
type CustomDestinationWithId struct {
	// The custom destination attributes.
	Attributes CustomDestinationAttributes `json:"attributes"`
	// The ID of the custom destination.
	Id string `json:"id"`
	// The custom destination type.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewCustomDestinationWithId instantiates a new CustomDestinationWithId object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomDestinationWithId(attributes CustomDestinationAttributes, id string, typeVar string) *CustomDestinationWithId {
	this := CustomDestinationWithId{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewCustomDestinationWithIdWithDefaults instantiates a new CustomDestinationWithId object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomDestinationWithIdWithDefaults() *CustomDestinationWithId {
	this := CustomDestinationWithId{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *CustomDestinationWithId) GetAttributes() CustomDestinationAttributes {
	if o == nil {
		var ret CustomDestinationAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationWithId) GetAttributesOk() (*CustomDestinationAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *CustomDestinationWithId) SetAttributes(v CustomDestinationAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *CustomDestinationWithId) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationWithId) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CustomDestinationWithId) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *CustomDestinationWithId) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationWithId) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomDestinationWithId) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomDestinationWithId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomDestinationWithId) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CustomDestinationAttributes `json:"attributes"`
		Id         *string                      `json:"id"`
		Type       *string                      `json:"type"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
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
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
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
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
