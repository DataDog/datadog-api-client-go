// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionsDatastoreDataset Test-only generated schema for the dashboard versioning proof of concept.
type ActionsDatastoreDataset struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Columns []DatasetColumn `json:"columns"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource ActionsDatastoreDatasetDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Id string `json:"id"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewActionsDatastoreDataset instantiates a new ActionsDatastoreDataset object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewActionsDatastoreDataset(columns []DatasetColumn, dataSource ActionsDatastoreDatasetDataSource, id string, name string) *ActionsDatastoreDataset {
	this := ActionsDatastoreDataset{}
	this.Columns = columns
	this.DataSource = dataSource
	this.Id = id
	this.Name = name
	return &this
}

// NewActionsDatastoreDatasetWithDefaults instantiates a new ActionsDatastoreDataset object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewActionsDatastoreDatasetWithDefaults() *ActionsDatastoreDataset {
	this := ActionsDatastoreDataset{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *ActionsDatastoreDataset) GetColumns() []DatasetColumn {
	if o == nil {
		var ret []DatasetColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *ActionsDatastoreDataset) GetColumnsOk() (*[]DatasetColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *ActionsDatastoreDataset) SetColumns(v []DatasetColumn) {
	o.Columns = v
}

// GetDataSource returns the DataSource field value.
func (o *ActionsDatastoreDataset) GetDataSource() ActionsDatastoreDatasetDataSource {
	if o == nil {
		var ret ActionsDatastoreDatasetDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *ActionsDatastoreDataset) GetDataSourceOk() (*ActionsDatastoreDatasetDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *ActionsDatastoreDataset) SetDataSource(v ActionsDatastoreDatasetDataSource) {
	o.DataSource = v
}

// GetId returns the Id field value.
func (o *ActionsDatastoreDataset) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ActionsDatastoreDataset) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ActionsDatastoreDataset) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *ActionsDatastoreDataset) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ActionsDatastoreDataset) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ActionsDatastoreDataset) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ActionsDatastoreDataset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["columns"] = o.Columns
	toSerialize["data_source"] = o.DataSource
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ActionsDatastoreDataset) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns    *[]DatasetColumn                   `json:"columns"`
		DataSource *ActionsDatastoreDatasetDataSource `json:"data_source"`
		Id         *string                            `json:"id"`
		Name       *string                            `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Columns == nil {
		return fmt.Errorf("required field columns missing")
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}

	hasInvalidField := false
	o.Columns = *all.Columns
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Id = *all.Id
	o.Name = *all.Name

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
