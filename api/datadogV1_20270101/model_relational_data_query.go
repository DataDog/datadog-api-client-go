// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RelationalDataQuery Test-only generated schema for the dashboard versioning proof of concept.
type RelationalDataQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource RelationalDataQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Filter *string `json:"filter,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *float64 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *RelationalDataQuerySort `json:"sort,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TableName string `json:"table_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRelationalDataQuery instantiates a new RelationalDataQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRelationalDataQuery(dataSource RelationalDataQueryDataSource, tableName string) *RelationalDataQuery {
	this := RelationalDataQuery{}
	this.DataSource = dataSource
	this.TableName = tableName
	return &this
}

// NewRelationalDataQueryWithDefaults instantiates a new RelationalDataQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRelationalDataQueryWithDefaults() *RelationalDataQuery {
	this := RelationalDataQuery{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *RelationalDataQuery) GetDataSource() RelationalDataQueryDataSource {
	if o == nil {
		var ret RelationalDataQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *RelationalDataQuery) GetDataSourceOk() (*RelationalDataQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *RelationalDataQuery) SetDataSource(v RelationalDataQueryDataSource) {
	o.DataSource = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *RelationalDataQuery) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationalDataQuery) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *RelationalDataQuery) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *RelationalDataQuery) SetFilter(v string) {
	o.Filter = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *RelationalDataQuery) GetLimit() float64 {
	if o == nil || o.Limit == nil {
		var ret float64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationalDataQuery) GetLimitOk() (*float64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *RelationalDataQuery) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given float64 and assigns it to the Limit field.
func (o *RelationalDataQuery) SetLimit(v float64) {
	o.Limit = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *RelationalDataQuery) GetSort() RelationalDataQuerySort {
	if o == nil || o.Sort == nil {
		var ret RelationalDataQuerySort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationalDataQuery) GetSortOk() (*RelationalDataQuerySort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *RelationalDataQuery) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given RelationalDataQuerySort and assigns it to the Sort field.
func (o *RelationalDataQuery) SetSort(v RelationalDataQuerySort) {
	o.Sort = &v
}

// GetTableName returns the TableName field value.
func (o *RelationalDataQuery) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *RelationalDataQuery) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value.
func (o *RelationalDataQuery) SetTableName(v string) {
	o.TableName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RelationalDataQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	toSerialize["table_name"] = o.TableName
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RelationalDataQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource *RelationalDataQueryDataSource `json:"data_source"`
		Filter     *string                        `json:"filter,omitempty"`
		Limit      *float64                       `json:"limit,omitempty"`
		Sort       *RelationalDataQuerySort       `json:"sort,omitempty"`
		TableName  *string                        `json:"table_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
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
	o.Filter = all.Filter
	o.Limit = all.Limit
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort
	o.TableName = *all.TableName

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
