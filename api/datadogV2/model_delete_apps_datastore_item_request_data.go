// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeleteAppsDatastoreItemRequestData The definition of `DeleteAppsDatastoreItemRequestData` object.
type DeleteAppsDatastoreItemRequestData struct {
	// The definition of `DeleteAppsDatastoreItemRequestDataAttributes` object.
	Attributes *DeleteAppsDatastoreItemRequestDataAttributes `json:"attributes,omitempty"`
	// The `DeleteAppsDatastoreItemRequestData` `id`.
	Id *string `json:"id,omitempty"`
	// Items resource type.
	Type DeleteAppsDatastoreItemRequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeleteAppsDatastoreItemRequestData instantiates a new DeleteAppsDatastoreItemRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeleteAppsDatastoreItemRequestData(typeVar DeleteAppsDatastoreItemRequestDataType) *DeleteAppsDatastoreItemRequestData {
	this := DeleteAppsDatastoreItemRequestData{}
	this.Type = typeVar
	return &this
}

// NewDeleteAppsDatastoreItemRequestDataWithDefaults instantiates a new DeleteAppsDatastoreItemRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeleteAppsDatastoreItemRequestDataWithDefaults() *DeleteAppsDatastoreItemRequestData {
	this := DeleteAppsDatastoreItemRequestData{}
	var typeVar DeleteAppsDatastoreItemRequestDataType = DELETEAPPSDATASTOREITEMREQUESTDATATYPE_ITEMS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DeleteAppsDatastoreItemRequestData) GetAttributes() DeleteAppsDatastoreItemRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret DeleteAppsDatastoreItemRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAppsDatastoreItemRequestData) GetAttributesOk() (*DeleteAppsDatastoreItemRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DeleteAppsDatastoreItemRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given DeleteAppsDatastoreItemRequestDataAttributes and assigns it to the Attributes field.
func (o *DeleteAppsDatastoreItemRequestData) SetAttributes(v DeleteAppsDatastoreItemRequestDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeleteAppsDatastoreItemRequestData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAppsDatastoreItemRequestData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeleteAppsDatastoreItemRequestData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeleteAppsDatastoreItemRequestData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *DeleteAppsDatastoreItemRequestData) GetType() DeleteAppsDatastoreItemRequestDataType {
	if o == nil {
		var ret DeleteAppsDatastoreItemRequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeleteAppsDatastoreItemRequestData) GetTypeOk() (*DeleteAppsDatastoreItemRequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DeleteAppsDatastoreItemRequestData) SetType(v DeleteAppsDatastoreItemRequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeleteAppsDatastoreItemRequestData) MarshalJSON() ([]byte, error) {
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
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeleteAppsDatastoreItemRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *DeleteAppsDatastoreItemRequestDataAttributes `json:"attributes,omitempty"`
		Id         *string                                       `json:"id,omitempty"`
		Type       *DeleteAppsDatastoreItemRequestDataType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
