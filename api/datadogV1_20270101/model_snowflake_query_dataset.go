// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeQueryDataset Test-only generated schema for the dashboard versioning proof of concept.
type SnowflakeQueryDataset struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource SnowflakeQueryDatasetDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SqlQuery string `json:"sql_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Warehouse *string `json:"warehouse,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSnowflakeQueryDataset instantiates a new SnowflakeQueryDataset object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeQueryDataset(dataSource SnowflakeQueryDatasetDataSource, name string, sqlQuery string) *SnowflakeQueryDataset {
	this := SnowflakeQueryDataset{}
	this.DataSource = dataSource
	this.Name = name
	this.SqlQuery = sqlQuery
	return &this
}

// NewSnowflakeQueryDatasetWithDefaults instantiates a new SnowflakeQueryDataset object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeQueryDatasetWithDefaults() *SnowflakeQueryDataset {
	this := SnowflakeQueryDataset{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *SnowflakeQueryDataset) GetDataSource() SnowflakeQueryDatasetDataSource {
	if o == nil {
		var ret SnowflakeQueryDatasetDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *SnowflakeQueryDataset) GetDataSourceOk() (*SnowflakeQueryDatasetDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *SnowflakeQueryDataset) SetDataSource(v SnowflakeQueryDatasetDataSource) {
	o.DataSource = v
}

// GetName returns the Name field value.
func (o *SnowflakeQueryDataset) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SnowflakeQueryDataset) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SnowflakeQueryDataset) SetName(v string) {
	o.Name = v
}

// GetSqlQuery returns the SqlQuery field value.
func (o *SnowflakeQueryDataset) GetSqlQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SqlQuery
}

// GetSqlQueryOk returns a tuple with the SqlQuery field value
// and a boolean to check if the value has been set.
func (o *SnowflakeQueryDataset) GetSqlQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SqlQuery, true
}

// SetSqlQuery sets field value.
func (o *SnowflakeQueryDataset) SetSqlQuery(v string) {
	o.SqlQuery = v
}

// GetWarehouse returns the Warehouse field value if set, zero value otherwise.
func (o *SnowflakeQueryDataset) GetWarehouse() string {
	if o == nil || o.Warehouse == nil {
		var ret string
		return ret
	}
	return *o.Warehouse
}

// GetWarehouseOk returns a tuple with the Warehouse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeQueryDataset) GetWarehouseOk() (*string, bool) {
	if o == nil || o.Warehouse == nil {
		return nil, false
	}
	return o.Warehouse, true
}

// HasWarehouse returns a boolean if a field has been set.
func (o *SnowflakeQueryDataset) HasWarehouse() bool {
	return o != nil && o.Warehouse != nil
}

// SetWarehouse gets a reference to the given string and assigns it to the Warehouse field.
func (o *SnowflakeQueryDataset) SetWarehouse(v string) {
	o.Warehouse = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeQueryDataset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["name"] = o.Name
	toSerialize["sql_query"] = o.SqlQuery
	if o.Warehouse != nil {
		toSerialize["warehouse"] = o.Warehouse
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeQueryDataset) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource *SnowflakeQueryDatasetDataSource `json:"data_source"`
		Name       *string                          `json:"name"`
		SqlQuery   *string                          `json:"sql_query"`
		Warehouse  *string                          `json:"warehouse,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.SqlQuery == nil {
		return fmt.Errorf("required field sql_query missing")
	}

	hasInvalidField := false
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Name = *all.Name
	o.SqlQuery = *all.SqlQuery
	o.Warehouse = all.Warehouse

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
