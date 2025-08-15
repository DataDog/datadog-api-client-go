// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateAppsDatastoreResponseData The definition of `CreateAppsDatastoreResponseData` object.
type CreateAppsDatastoreResponseData struct {
	// The `CreateAppsDatastoreResponseData` `id`.
	Id *string `json:"id,omitempty"`
	// Datastores resource type.
	Type CreateAppsDatastoreResponseDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateAppsDatastoreResponseData instantiates a new CreateAppsDatastoreResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateAppsDatastoreResponseData(typeVar CreateAppsDatastoreResponseDataType) *CreateAppsDatastoreResponseData {
	this := CreateAppsDatastoreResponseData{}
	this.Type = typeVar
	return &this
}

// NewCreateAppsDatastoreResponseDataWithDefaults instantiates a new CreateAppsDatastoreResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateAppsDatastoreResponseDataWithDefaults() *CreateAppsDatastoreResponseData {
	this := CreateAppsDatastoreResponseData{}
	var typeVar CreateAppsDatastoreResponseDataType = CREATEAPPSDATASTORERESPONSEDATATYPE_DATASTORES
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateAppsDatastoreResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppsDatastoreResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateAppsDatastoreResponseData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateAppsDatastoreResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *CreateAppsDatastoreResponseData) GetType() CreateAppsDatastoreResponseDataType {
	if o == nil {
		var ret CreateAppsDatastoreResponseDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateAppsDatastoreResponseData) GetTypeOk() (*CreateAppsDatastoreResponseDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateAppsDatastoreResponseData) SetType(v CreateAppsDatastoreResponseDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateAppsDatastoreResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
func (o *CreateAppsDatastoreResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                              `json:"id,omitempty"`
		Type *CreateAppsDatastoreResponseDataType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
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
