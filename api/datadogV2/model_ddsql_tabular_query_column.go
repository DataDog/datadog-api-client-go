// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DdsqlTabularQueryColumn A single column of a DDSQL tabular query result.
type DdsqlTabularQueryColumn struct {
	// Name of the column as projected by the SQL statement.
	Name string `json:"name"`
	// DDSQL data type of the column's values, for example `VARCHAR`, `BIGINT`,
	// `DECIMAL`, `BOOLEAN`, `TIMESTAMP`, `JSON`, or an array variant such as
	// `VARCHAR[]`. See the
	// [DDSQL data-types reference](https://docs.datadoghq.com/ddsql_reference/#data-types)
	// for the full, up-to-date list.
	Type string `json:"type"`
	// Column values in row order. The element type matches the column's `type`;
	// for example a `VARCHAR` column carries strings, a `TIMESTAMP` column carries
	// Unix-millisecond integers. `null` is allowed for missing values.
	Values []interface{} `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDdsqlTabularQueryColumn instantiates a new DdsqlTabularQueryColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDdsqlTabularQueryColumn(name string, typeVar string, values []interface{}) *DdsqlTabularQueryColumn {
	this := DdsqlTabularQueryColumn{}
	this.Name = name
	this.Type = typeVar
	this.Values = values
	return &this
}

// NewDdsqlTabularQueryColumnWithDefaults instantiates a new DdsqlTabularQueryColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDdsqlTabularQueryColumnWithDefaults() *DdsqlTabularQueryColumn {
	this := DdsqlTabularQueryColumn{}
	return &this
}

// GetName returns the Name field value.
func (o *DdsqlTabularQueryColumn) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DdsqlTabularQueryColumn) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DdsqlTabularQueryColumn) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *DdsqlTabularQueryColumn) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DdsqlTabularQueryColumn) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DdsqlTabularQueryColumn) SetType(v string) {
	o.Type = v
}

// GetValues returns the Values field value.
func (o *DdsqlTabularQueryColumn) GetValues() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *DdsqlTabularQueryColumn) GetValuesOk() (*[]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *DdsqlTabularQueryColumn) SetValues(v []interface{}) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DdsqlTabularQueryColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DdsqlTabularQueryColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name   *string        `json:"name"`
		Type   *string        `json:"type"`
		Values *[]interface{} `json:"values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "type", "values"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Type = *all.Type
	o.Values = *all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
