// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake Configuration of a Snowflake connection.
type ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake struct {
	// Identifier of the Snowflake account.
	AccountId *string `json:"account_id,omitempty"`
	// Name of the Snowflake database.
	DatabaseName *string `json:"database_name,omitempty"`
	// Name of the Snowflake schema.
	Schema *string `json:"schema,omitempty"`
	// Name of the Snowflake table.
	TableName *string `json:"table_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake instantiates a new ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake() *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake {
	this := ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake{}
	return &this
}

// NewListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflakeWithDefaults instantiates a new ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflakeWithDefaults() *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake {
	this := ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake) HasAccountId() bool {
	return o != nil && o.AccountId != nil
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake) SetAccountId(v string) {
	o.AccountId = &v
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake) GetDatabaseName() string {
	if o == nil || o.DatabaseName == nil {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake) GetDatabaseNameOk() (*string, bool) {
	if o == nil || o.DatabaseName == nil {
		return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake) HasDatabaseName() bool {
	return o != nil && o.DatabaseName != nil
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake) SetSchema(v string) {
	o.Schema = &v
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake) GetTableName() string {
	if o == nil || o.TableName == nil {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake) GetTableNameOk() (*string, bool) {
	if o == nil || o.TableName == nil {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake) HasTableName() bool {
	return o != nil && o.TableName != nil
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake) SetTableName(v string) {
	o.TableName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.DatabaseName != nil {
		toSerialize["database_name"] = o.DatabaseName
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
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
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId    *string `json:"account_id,omitempty"`
		DatabaseName *string `json:"database_name,omitempty"`
		Schema       *string `json:"schema,omitempty"`
		TableName    *string `json:"table_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "database_name", "schema", "table_name"})
	} else {
		return err
	}
	o.AccountId = all.AccountId
	o.DatabaseName = all.DatabaseName
	o.Schema = all.Schema
	o.TableName = all.TableName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
