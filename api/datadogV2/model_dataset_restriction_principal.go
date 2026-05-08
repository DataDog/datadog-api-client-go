// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetRestrictionPrincipal A user or role that is exempt from dataset restrictions and retains unrestricted
// access to all datasets for the product type.
type DatasetRestrictionPrincipal struct {
	// The unique identifier of the principal (a user UUID or role ID).
	Id string `json:"id"`
	// The human-readable display name of the principal as shown in the Datadog UI.
	Name string `json:"name"`
	// The kind of principal, such as `user` for an individual user account or `role`
	// for a Datadog role.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatasetRestrictionPrincipal instantiates a new DatasetRestrictionPrincipal object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatasetRestrictionPrincipal(id string, name string, typeVar string) *DatasetRestrictionPrincipal {
	this := DatasetRestrictionPrincipal{}
	this.Id = id
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewDatasetRestrictionPrincipalWithDefaults instantiates a new DatasetRestrictionPrincipal object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatasetRestrictionPrincipalWithDefaults() *DatasetRestrictionPrincipal {
	this := DatasetRestrictionPrincipal{}
	return &this
}

// GetId returns the Id field value.
func (o *DatasetRestrictionPrincipal) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DatasetRestrictionPrincipal) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *DatasetRestrictionPrincipal) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *DatasetRestrictionPrincipal) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatasetRestrictionPrincipal) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DatasetRestrictionPrincipal) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *DatasetRestrictionPrincipal) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DatasetRestrictionPrincipal) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DatasetRestrictionPrincipal) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatasetRestrictionPrincipal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatasetRestrictionPrincipal) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string `json:"id"`
		Name *string `json:"name"`
		Type *string `json:"type"`
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
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "name", "type"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.Name = *all.Name
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
