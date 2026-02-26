// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RecommendedEntityWithSchema A recommended entity with its schema definition.
type RecommendedEntityWithSchema struct {
	// Unique identifier for the recommended entity.
	Id string `json:"id"`
	// Entity schema v3.
	Schema EntityV3 `json:"schema"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRecommendedEntityWithSchema instantiates a new RecommendedEntityWithSchema object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRecommendedEntityWithSchema(id string, schema EntityV3) *RecommendedEntityWithSchema {
	this := RecommendedEntityWithSchema{}
	this.Id = id
	this.Schema = schema
	return &this
}

// NewRecommendedEntityWithSchemaWithDefaults instantiates a new RecommendedEntityWithSchema object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRecommendedEntityWithSchemaWithDefaults() *RecommendedEntityWithSchema {
	this := RecommendedEntityWithSchema{}
	return &this
}

// GetId returns the Id field value.
func (o *RecommendedEntityWithSchema) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RecommendedEntityWithSchema) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *RecommendedEntityWithSchema) SetId(v string) {
	o.Id = v
}

// GetSchema returns the Schema field value.
func (o *RecommendedEntityWithSchema) GetSchema() EntityV3 {
	if o == nil {
		var ret EntityV3
		return ret
	}
	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *RecommendedEntityWithSchema) GetSchemaOk() (*EntityV3, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schema, true
}

// SetSchema sets field value.
func (o *RecommendedEntityWithSchema) SetSchema(v EntityV3) {
	o.Schema = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RecommendedEntityWithSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["schema"] = o.Schema

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RecommendedEntityWithSchema) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id     *string   `json:"id"`
		Schema *EntityV3 `json:"schema"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Schema == nil {
		return fmt.Errorf("required field schema missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "schema"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.Schema = *all.Schema

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
