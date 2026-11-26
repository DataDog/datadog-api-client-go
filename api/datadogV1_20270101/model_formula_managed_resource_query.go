// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaManagedResourceQuery Test-only generated schema for the dashboard versioning proof of concept.
type FormulaManagedResourceQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute []DatasetCompute `json:"compute"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaManagedResourceQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []DatasetGroupBy `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryFilter *string `json:"query_filter,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TableName string `json:"table_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaManagedResourceQuery instantiates a new FormulaManagedResourceQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaManagedResourceQuery(compute []DatasetCompute, dataSource FormulaManagedResourceQueryDataSource, name string, tableName string) *FormulaManagedResourceQuery {
	this := FormulaManagedResourceQuery{}
	this.Compute = compute
	this.DataSource = dataSource
	this.Name = name
	this.TableName = tableName
	return &this
}

// NewFormulaManagedResourceQueryWithDefaults instantiates a new FormulaManagedResourceQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaManagedResourceQueryWithDefaults() *FormulaManagedResourceQuery {
	this := FormulaManagedResourceQuery{}
	return &this
}

// GetCompute returns the Compute field value.
func (o *FormulaManagedResourceQuery) GetCompute() []DatasetCompute {
	if o == nil {
		var ret []DatasetCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *FormulaManagedResourceQuery) GetComputeOk() (*[]DatasetCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *FormulaManagedResourceQuery) SetCompute(v []DatasetCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaManagedResourceQuery) GetDataSource() FormulaManagedResourceQueryDataSource {
	if o == nil {
		var ret FormulaManagedResourceQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaManagedResourceQuery) GetDataSourceOk() (*FormulaManagedResourceQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaManagedResourceQuery) SetDataSource(v FormulaManagedResourceQueryDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *FormulaManagedResourceQuery) GetGroupBy() []DatasetGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []DatasetGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaManagedResourceQuery) GetGroupByOk() (*[]DatasetGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *FormulaManagedResourceQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []DatasetGroupBy and assigns it to the GroupBy field.
func (o *FormulaManagedResourceQuery) SetGroupBy(v []DatasetGroupBy) {
	o.GroupBy = v
}

// GetName returns the Name field value.
func (o *FormulaManagedResourceQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaManagedResourceQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaManagedResourceQuery) SetName(v string) {
	o.Name = v
}

// GetQueryFilter returns the QueryFilter field value if set, zero value otherwise.
func (o *FormulaManagedResourceQuery) GetQueryFilter() string {
	if o == nil || o.QueryFilter == nil {
		var ret string
		return ret
	}
	return *o.QueryFilter
}

// GetQueryFilterOk returns a tuple with the QueryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaManagedResourceQuery) GetQueryFilterOk() (*string, bool) {
	if o == nil || o.QueryFilter == nil {
		return nil, false
	}
	return o.QueryFilter, true
}

// HasQueryFilter returns a boolean if a field has been set.
func (o *FormulaManagedResourceQuery) HasQueryFilter() bool {
	return o != nil && o.QueryFilter != nil
}

// SetQueryFilter gets a reference to the given string and assigns it to the QueryFilter field.
func (o *FormulaManagedResourceQuery) SetQueryFilter(v string) {
	o.QueryFilter = &v
}

// GetTableName returns the TableName field value.
func (o *FormulaManagedResourceQuery) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *FormulaManagedResourceQuery) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value.
func (o *FormulaManagedResourceQuery) SetTableName(v string) {
	o.TableName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaManagedResourceQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["compute"] = o.Compute
	toSerialize["data_source"] = o.DataSource
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["name"] = o.Name
	if o.QueryFilter != nil {
		toSerialize["query_filter"] = o.QueryFilter
	}
	toSerialize["table_name"] = o.TableName
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaManagedResourceQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute     *[]DatasetCompute                      `json:"compute"`
		DataSource  *FormulaManagedResourceQueryDataSource `json:"data_source"`
		GroupBy     []DatasetGroupBy                       `json:"group_by,omitempty"`
		Name        *string                                `json:"name"`
		QueryFilter *string                                `json:"query_filter,omitempty"`
		TableName   *string                                `json:"table_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Compute == nil {
		return fmt.Errorf("required field compute missing")
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
	o.Compute = *all.Compute
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.GroupBy = all.GroupBy
	o.Name = *all.Name
	o.QueryFilter = all.QueryFilter
	o.TableName = *all.TableName

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
