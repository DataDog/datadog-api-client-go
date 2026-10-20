// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SavedFilterFlagReference A feature flag that references the saved filter.
type SavedFilterFlagReference struct {
	// The ID of the feature flag.
	Id uuid.UUID `json:"id"`
	// The name of the feature flag.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSavedFilterFlagReference instantiates a new SavedFilterFlagReference object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSavedFilterFlagReference(id uuid.UUID, name string) *SavedFilterFlagReference {
	this := SavedFilterFlagReference{}
	this.Id = id
	this.Name = name
	return &this
}

// NewSavedFilterFlagReferenceWithDefaults instantiates a new SavedFilterFlagReference object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSavedFilterFlagReferenceWithDefaults() *SavedFilterFlagReference {
	this := SavedFilterFlagReference{}
	return &this
}

// GetId returns the Id field value.
func (o *SavedFilterFlagReference) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SavedFilterFlagReference) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *SavedFilterFlagReference) SetId(v uuid.UUID) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *SavedFilterFlagReference) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SavedFilterFlagReference) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SavedFilterFlagReference) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SavedFilterFlagReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SavedFilterFlagReference) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *uuid.UUID `json:"id"`
		Name *string    `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "name"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
