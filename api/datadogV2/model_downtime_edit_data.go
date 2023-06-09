// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeEditData Object to update a downtime.
type DowntimeEditData struct {
	// Attributes of the downtime to update.
	Attributes DowntimeAttributeEditRequest `json:"attributes"`
	// ID of this downtime.
	Id string `json:"id"`
	// Downtime resource type.
	Type DowntimeResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewDowntimeEditData instantiates a new DowntimeEditData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeEditData(attributes DowntimeAttributeEditRequest, id string, typeVar DowntimeResourceType) *DowntimeEditData {
	this := DowntimeEditData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewDowntimeEditDataWithDefaults instantiates a new DowntimeEditData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeEditDataWithDefaults() *DowntimeEditData {
	this := DowntimeEditData{}
	var typeVar DowntimeResourceType = DOWNTIMERESOURCETYPE_DOWNTIME
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *DowntimeEditData) GetAttributes() DowntimeAttributeEditRequest {
	if o == nil {
		var ret DowntimeAttributeEditRequest
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *DowntimeEditData) GetAttributesOk() (*DowntimeAttributeEditRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *DowntimeEditData) SetAttributes(v DowntimeAttributeEditRequest) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *DowntimeEditData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DowntimeEditData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *DowntimeEditData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *DowntimeEditData) GetType() DowntimeResourceType {
	if o == nil {
		var ret DowntimeResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DowntimeEditData) GetTypeOk() (*DowntimeResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DowntimeEditData) SetType(v DowntimeResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeEditData) MarshalJSON() ([]byte, error) {
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
func (o *DowntimeEditData) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Attributes *DowntimeAttributeEditRequest `json:"attributes"`
		Id         *string                       `json:"id"`
		Type       *DowntimeResourceType         `json:"type"`
	}{}
	all := struct {
		Attributes DowntimeAttributeEditRequest `json:"attributes"`
		Id         string                       `json:"id"`
		Type       DowntimeResourceType         `json:"type"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if required.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if required.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}
	if v := all.Type; !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	o.Type = all.Type
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
