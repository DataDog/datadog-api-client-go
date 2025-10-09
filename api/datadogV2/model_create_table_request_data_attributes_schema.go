// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateTableRequestDataAttributesSchema The definition of `CreateTableRequestDataAttributesSchema` object.
type CreateTableRequestDataAttributesSchema struct {
	// The `schema` `fields`.
	Fields []CreateTableRequestDataAttributesSchemaFieldsItems `json:"fields"`
	// List of field names that serve as primary keys for the table. Only one primary key is supported, and it is used as an ID to retrieve rows.
	PrimaryKeys []string `json:"primary_keys"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateTableRequestDataAttributesSchema instantiates a new CreateTableRequestDataAttributesSchema object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateTableRequestDataAttributesSchema(fields []CreateTableRequestDataAttributesSchemaFieldsItems, primaryKeys []string) *CreateTableRequestDataAttributesSchema {
	this := CreateTableRequestDataAttributesSchema{}
	this.Fields = fields
	this.PrimaryKeys = primaryKeys
	return &this
}

// NewCreateTableRequestDataAttributesSchemaWithDefaults instantiates a new CreateTableRequestDataAttributesSchema object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateTableRequestDataAttributesSchemaWithDefaults() *CreateTableRequestDataAttributesSchema {
	this := CreateTableRequestDataAttributesSchema{}
	return &this
}

// GetFields returns the Fields field value.
func (o *CreateTableRequestDataAttributesSchema) GetFields() []CreateTableRequestDataAttributesSchemaFieldsItems {
	if o == nil {
		var ret []CreateTableRequestDataAttributesSchemaFieldsItems
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributesSchema) GetFieldsOk() (*[]CreateTableRequestDataAttributesSchemaFieldsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value.
func (o *CreateTableRequestDataAttributesSchema) SetFields(v []CreateTableRequestDataAttributesSchemaFieldsItems) {
	o.Fields = v
}

// GetPrimaryKeys returns the PrimaryKeys field value.
func (o *CreateTableRequestDataAttributesSchema) GetPrimaryKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PrimaryKeys
}

// GetPrimaryKeysOk returns a tuple with the PrimaryKeys field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributesSchema) GetPrimaryKeysOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryKeys, true
}

// SetPrimaryKeys sets field value.
func (o *CreateTableRequestDataAttributesSchema) SetPrimaryKeys(v []string) {
	o.PrimaryKeys = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateTableRequestDataAttributesSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["fields"] = o.Fields
	toSerialize["primary_keys"] = o.PrimaryKeys

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateTableRequestDataAttributesSchema) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fields      *[]CreateTableRequestDataAttributesSchemaFieldsItems `json:"fields"`
		PrimaryKeys *[]string                                            `json:"primary_keys"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Fields == nil {
		return fmt.Errorf("required field fields missing")
	}
	if all.PrimaryKeys == nil {
		return fmt.Errorf("required field primary_keys missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fields", "primary_keys"})
	} else {
		return err
	}
	o.Fields = *all.Fields
	o.PrimaryKeys = *all.PrimaryKeys

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
