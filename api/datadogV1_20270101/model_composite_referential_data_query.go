// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CompositeReferentialDataQuery Test-only generated schema for the dashboard versioning proof of concept.
type CompositeReferentialDataQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Columns []ListColumn `json:"columns,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute []CompositeReferentialDataQueryComputeItems `json:"compute,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource CompositeReferentialDataQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []string `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryFilter *string `json:"query_filter,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TableName string `json:"table_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCompositeReferentialDataQuery instantiates a new CompositeReferentialDataQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCompositeReferentialDataQuery(dataSource CompositeReferentialDataQueryDataSource, name string, tableName string) *CompositeReferentialDataQuery {
	this := CompositeReferentialDataQuery{}
	this.DataSource = dataSource
	this.Name = name
	this.TableName = tableName
	return &this
}

// NewCompositeReferentialDataQueryWithDefaults instantiates a new CompositeReferentialDataQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCompositeReferentialDataQueryWithDefaults() *CompositeReferentialDataQuery {
	this := CompositeReferentialDataQuery{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *CompositeReferentialDataQuery) GetColumns() []ListColumn {
	if o == nil || o.Columns == nil {
		var ret []ListColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeReferentialDataQuery) GetColumnsOk() (*[]ListColumn, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *CompositeReferentialDataQuery) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []ListColumn and assigns it to the Columns field.
func (o *CompositeReferentialDataQuery) SetColumns(v []ListColumn) {
	o.Columns = v
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *CompositeReferentialDataQuery) GetCompute() []CompositeReferentialDataQueryComputeItems {
	if o == nil || o.Compute == nil {
		var ret []CompositeReferentialDataQueryComputeItems
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeReferentialDataQuery) GetComputeOk() (*[]CompositeReferentialDataQueryComputeItems, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return &o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *CompositeReferentialDataQuery) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given []CompositeReferentialDataQueryComputeItems and assigns it to the Compute field.
func (o *CompositeReferentialDataQuery) SetCompute(v []CompositeReferentialDataQueryComputeItems) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *CompositeReferentialDataQuery) GetDataSource() CompositeReferentialDataQueryDataSource {
	if o == nil {
		var ret CompositeReferentialDataQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *CompositeReferentialDataQuery) GetDataSourceOk() (*CompositeReferentialDataQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *CompositeReferentialDataQuery) SetDataSource(v CompositeReferentialDataQueryDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *CompositeReferentialDataQuery) GetGroupBy() []string {
	if o == nil || o.GroupBy == nil {
		var ret []string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeReferentialDataQuery) GetGroupByOk() (*[]string, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *CompositeReferentialDataQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []string and assigns it to the GroupBy field.
func (o *CompositeReferentialDataQuery) SetGroupBy(v []string) {
	o.GroupBy = v
}

// GetName returns the Name field value.
func (o *CompositeReferentialDataQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CompositeReferentialDataQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CompositeReferentialDataQuery) SetName(v string) {
	o.Name = v
}

// GetQueryFilter returns the QueryFilter field value if set, zero value otherwise.
func (o *CompositeReferentialDataQuery) GetQueryFilter() string {
	if o == nil || o.QueryFilter == nil {
		var ret string
		return ret
	}
	return *o.QueryFilter
}

// GetQueryFilterOk returns a tuple with the QueryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeReferentialDataQuery) GetQueryFilterOk() (*string, bool) {
	if o == nil || o.QueryFilter == nil {
		return nil, false
	}
	return o.QueryFilter, true
}

// HasQueryFilter returns a boolean if a field has been set.
func (o *CompositeReferentialDataQuery) HasQueryFilter() bool {
	return o != nil && o.QueryFilter != nil
}

// SetQueryFilter gets a reference to the given string and assigns it to the QueryFilter field.
func (o *CompositeReferentialDataQuery) SetQueryFilter(v string) {
	o.QueryFilter = &v
}

// GetTableName returns the TableName field value.
func (o *CompositeReferentialDataQuery) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *CompositeReferentialDataQuery) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value.
func (o *CompositeReferentialDataQuery) SetTableName(v string) {
	o.TableName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CompositeReferentialDataQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
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
func (o *CompositeReferentialDataQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns     []ListColumn                                `json:"columns,omitempty"`
		Compute     []CompositeReferentialDataQueryComputeItems `json:"compute,omitempty"`
		DataSource  *CompositeReferentialDataQueryDataSource    `json:"data_source"`
		GroupBy     []string                                    `json:"group_by,omitempty"`
		Name        *string                                     `json:"name"`
		QueryFilter *string                                     `json:"query_filter,omitempty"`
		TableName   *string                                     `json:"table_name"`
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
	o.Columns = all.Columns
	o.Compute = all.Compute
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
