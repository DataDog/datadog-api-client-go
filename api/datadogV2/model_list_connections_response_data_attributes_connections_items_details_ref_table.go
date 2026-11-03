// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable Configuration of a reference table connection.
type ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable struct {
	// Identifier of the reference table.
	TableId *string `json:"table_id,omitempty"`
	// Name of the reference table.
	TableName *string `json:"table_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable instantiates a new ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable() *ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable {
	this := ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable{}
	return &this
}

// NewListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTableWithDefaults instantiates a new ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTableWithDefaults() *ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable {
	this := ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable{}
	return &this
}

// GetTableId returns the TableId field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable) GetTableId() string {
	if o == nil || o.TableId == nil {
		var ret string
		return ret
	}
	return *o.TableId
}

// GetTableIdOk returns a tuple with the TableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable) GetTableIdOk() (*string, bool) {
	if o == nil || o.TableId == nil {
		return nil, false
	}
	return o.TableId, true
}

// HasTableId returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable) HasTableId() bool {
	return o != nil && o.TableId != nil
}

// SetTableId gets a reference to the given string and assigns it to the TableId field.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable) SetTableId(v string) {
	o.TableId = &v
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable) GetTableName() string {
	if o == nil || o.TableName == nil {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable) GetTableNameOk() (*string, bool) {
	if o == nil || o.TableName == nil {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable) HasTableName() bool {
	return o != nil && o.TableName != nil
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable) SetTableName(v string) {
	o.TableName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TableId != nil {
		toSerialize["table_id"] = o.TableId
	}
	if o.TableName != nil {
		toSerialize["table_name"] = o.TableName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TableId   *string `json:"table_id,omitempty"`
		TableName *string `json:"table_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"table_id", "table_name"})
	} else {
		return err
	}
	o.TableId = all.TableId
	o.TableName = all.TableName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
