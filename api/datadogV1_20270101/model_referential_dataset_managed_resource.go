// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReferentialDatasetManagedResource Test-only generated schema for the dashboard versioning proof of concept.
type ReferentialDatasetManagedResource struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource ReferentialDatasetManagedResourceDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryFilter *string `json:"query_filter,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TableName string `json:"table_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewReferentialDatasetManagedResource instantiates a new ReferentialDatasetManagedResource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReferentialDatasetManagedResource(dataSource ReferentialDatasetManagedResourceDataSource, name string, tableName string) *ReferentialDatasetManagedResource {
	this := ReferentialDatasetManagedResource{}
	this.DataSource = dataSource
	this.Name = name
	this.TableName = tableName
	return &this
}

// NewReferentialDatasetManagedResourceWithDefaults instantiates a new ReferentialDatasetManagedResource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReferentialDatasetManagedResourceWithDefaults() *ReferentialDatasetManagedResource {
	this := ReferentialDatasetManagedResource{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *ReferentialDatasetManagedResource) GetDataSource() ReferentialDatasetManagedResourceDataSource {
	if o == nil {
		var ret ReferentialDatasetManagedResourceDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *ReferentialDatasetManagedResource) GetDataSourceOk() (*ReferentialDatasetManagedResourceDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *ReferentialDatasetManagedResource) SetDataSource(v ReferentialDatasetManagedResourceDataSource) {
	o.DataSource = v
}

// GetName returns the Name field value.
func (o *ReferentialDatasetManagedResource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReferentialDatasetManagedResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ReferentialDatasetManagedResource) SetName(v string) {
	o.Name = v
}

// GetQueryFilter returns the QueryFilter field value if set, zero value otherwise.
func (o *ReferentialDatasetManagedResource) GetQueryFilter() string {
	if o == nil || o.QueryFilter == nil {
		var ret string
		return ret
	}
	return *o.QueryFilter
}

// GetQueryFilterOk returns a tuple with the QueryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferentialDatasetManagedResource) GetQueryFilterOk() (*string, bool) {
	if o == nil || o.QueryFilter == nil {
		return nil, false
	}
	return o.QueryFilter, true
}

// HasQueryFilter returns a boolean if a field has been set.
func (o *ReferentialDatasetManagedResource) HasQueryFilter() bool {
	return o != nil && o.QueryFilter != nil
}

// SetQueryFilter gets a reference to the given string and assigns it to the QueryFilter field.
func (o *ReferentialDatasetManagedResource) SetQueryFilter(v string) {
	o.QueryFilter = &v
}

// GetTableName returns the TableName field value.
func (o *ReferentialDatasetManagedResource) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *ReferentialDatasetManagedResource) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value.
func (o *ReferentialDatasetManagedResource) SetTableName(v string) {
	o.TableName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReferentialDatasetManagedResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["name"] = o.Name
	if o.QueryFilter != nil {
		toSerialize["query_filter"] = o.QueryFilter
	}
	toSerialize["table_name"] = o.TableName
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReferentialDatasetManagedResource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource  *ReferentialDatasetManagedResourceDataSource `json:"data_source"`
		Name        *string                                      `json:"name"`
		QueryFilter *string                                      `json:"query_filter,omitempty"`
		TableName   *string                                      `json:"table_name"`
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
	if all.TableName == nil {
		return fmt.Errorf("required field table_name missing")
	}

	hasInvalidField := false
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Name = *all.Name
	o.QueryFilter = all.QueryFilter
	o.TableName = *all.TableName

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
