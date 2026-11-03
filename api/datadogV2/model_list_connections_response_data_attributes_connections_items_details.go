// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListConnectionsResponseDataAttributesConnectionsItemsDetails Type-specific configuration of the connection. Only the field matching the connection type is set.
type ListConnectionsResponseDataAttributesConnectionsItemsDetails struct {
	// Configuration of a reference table connection.
	RefTable *ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable `json:"ref_table,omitempty"`
	// Configuration of a Snowflake connection.
	Snowflake *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake `json:"snowflake,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListConnectionsResponseDataAttributesConnectionsItemsDetails instantiates a new ListConnectionsResponseDataAttributesConnectionsItemsDetails object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListConnectionsResponseDataAttributesConnectionsItemsDetails() *ListConnectionsResponseDataAttributesConnectionsItemsDetails {
	this := ListConnectionsResponseDataAttributesConnectionsItemsDetails{}
	return &this
}

// NewListConnectionsResponseDataAttributesConnectionsItemsDetailsWithDefaults instantiates a new ListConnectionsResponseDataAttributesConnectionsItemsDetails object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListConnectionsResponseDataAttributesConnectionsItemsDetailsWithDefaults() *ListConnectionsResponseDataAttributesConnectionsItemsDetails {
	this := ListConnectionsResponseDataAttributesConnectionsItemsDetails{}
	return &this
}

// GetRefTable returns the RefTable field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetails) GetRefTable() ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable {
	if o == nil || o.RefTable == nil {
		var ret ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable
		return ret
	}
	return *o.RefTable
}

// GetRefTableOk returns a tuple with the RefTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetails) GetRefTableOk() (*ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable, bool) {
	if o == nil || o.RefTable == nil {
		return nil, false
	}
	return o.RefTable, true
}

// HasRefTable returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetails) HasRefTable() bool {
	return o != nil && o.RefTable != nil
}

// SetRefTable gets a reference to the given ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable and assigns it to the RefTable field.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetails) SetRefTable(v ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable) {
	o.RefTable = &v
}

// GetSnowflake returns the Snowflake field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetails) GetSnowflake() ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake {
	if o == nil || o.Snowflake == nil {
		var ret ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake
		return ret
	}
	return *o.Snowflake
}

// GetSnowflakeOk returns a tuple with the Snowflake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetails) GetSnowflakeOk() (*ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake, bool) {
	if o == nil || o.Snowflake == nil {
		return nil, false
	}
	return o.Snowflake, true
}

// HasSnowflake returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetails) HasSnowflake() bool {
	return o != nil && o.Snowflake != nil
}

// SetSnowflake gets a reference to the given ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake and assigns it to the Snowflake field.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetails) SetSnowflake(v ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake) {
	o.Snowflake = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListConnectionsResponseDataAttributesConnectionsItemsDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.RefTable != nil {
		toSerialize["ref_table"] = o.RefTable
	}
	if o.Snowflake != nil {
		toSerialize["snowflake"] = o.Snowflake
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListConnectionsResponseDataAttributesConnectionsItemsDetails) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RefTable  *ListConnectionsResponseDataAttributesConnectionsItemsDetailsRefTable  `json:"ref_table,omitempty"`
		Snowflake *ListConnectionsResponseDataAttributesConnectionsItemsDetailsSnowflake `json:"snowflake,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ref_table", "snowflake"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.RefTable != nil && all.RefTable.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RefTable = all.RefTable
	if all.Snowflake != nil && all.Snowflake.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Snowflake = all.Snowflake

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
