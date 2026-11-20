// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregateSegmentQuery Test-only generated schema for the dashboard versioning proof of concept.
type AggregateSegmentQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	BaseQuery CompositeEventsPlatformQuery `json:"base_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute []CompositeCompute `json:"compute"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource AggregateSegmentQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []CompositeGroupBy `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SegmentQuery AggregateSegmentQuerySegmentQuery `json:"segment_query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAggregateSegmentQuery instantiates a new AggregateSegmentQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregateSegmentQuery(baseQuery CompositeEventsPlatformQuery, compute []CompositeCompute, dataSource AggregateSegmentQueryDataSource, name string, segmentQuery AggregateSegmentQuerySegmentQuery) *AggregateSegmentQuery {
	this := AggregateSegmentQuery{}
	this.BaseQuery = baseQuery
	this.Compute = compute
	this.DataSource = dataSource
	this.Name = name
	this.SegmentQuery = segmentQuery
	return &this
}

// NewAggregateSegmentQueryWithDefaults instantiates a new AggregateSegmentQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregateSegmentQueryWithDefaults() *AggregateSegmentQuery {
	this := AggregateSegmentQuery{}
	return &this
}

// GetBaseQuery returns the BaseQuery field value.
func (o *AggregateSegmentQuery) GetBaseQuery() CompositeEventsPlatformQuery {
	if o == nil {
		var ret CompositeEventsPlatformQuery
		return ret
	}
	return o.BaseQuery
}

// GetBaseQueryOk returns a tuple with the BaseQuery field value
// and a boolean to check if the value has been set.
func (o *AggregateSegmentQuery) GetBaseQueryOk() (*CompositeEventsPlatformQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseQuery, true
}

// SetBaseQuery sets field value.
func (o *AggregateSegmentQuery) SetBaseQuery(v CompositeEventsPlatformQuery) {
	o.BaseQuery = v
}

// GetCompute returns the Compute field value.
func (o *AggregateSegmentQuery) GetCompute() []CompositeCompute {
	if o == nil {
		var ret []CompositeCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *AggregateSegmentQuery) GetComputeOk() (*[]CompositeCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *AggregateSegmentQuery) SetCompute(v []CompositeCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *AggregateSegmentQuery) GetDataSource() AggregateSegmentQueryDataSource {
	if o == nil {
		var ret AggregateSegmentQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *AggregateSegmentQuery) GetDataSourceOk() (*AggregateSegmentQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *AggregateSegmentQuery) SetDataSource(v AggregateSegmentQueryDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *AggregateSegmentQuery) GetGroupBy() []CompositeGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []CompositeGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateSegmentQuery) GetGroupByOk() (*[]CompositeGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *AggregateSegmentQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []CompositeGroupBy and assigns it to the GroupBy field.
func (o *AggregateSegmentQuery) SetGroupBy(v []CompositeGroupBy) {
	o.GroupBy = v
}

// GetName returns the Name field value.
func (o *AggregateSegmentQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AggregateSegmentQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AggregateSegmentQuery) SetName(v string) {
	o.Name = v
}

// GetSegmentQuery returns the SegmentQuery field value.
func (o *AggregateSegmentQuery) GetSegmentQuery() AggregateSegmentQuerySegmentQuery {
	if o == nil {
		var ret AggregateSegmentQuerySegmentQuery
		return ret
	}
	return o.SegmentQuery
}

// GetSegmentQueryOk returns a tuple with the SegmentQuery field value
// and a boolean to check if the value has been set.
func (o *AggregateSegmentQuery) GetSegmentQueryOk() (*AggregateSegmentQuerySegmentQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SegmentQuery, true
}

// SetSegmentQuery sets field value.
func (o *AggregateSegmentQuery) SetSegmentQuery(v AggregateSegmentQuerySegmentQuery) {
	o.SegmentQuery = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregateSegmentQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["base_query"] = o.BaseQuery
	toSerialize["compute"] = o.Compute
	toSerialize["data_source"] = o.DataSource
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["name"] = o.Name
	toSerialize["segment_query"] = o.SegmentQuery
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregateSegmentQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BaseQuery    *CompositeEventsPlatformQuery      `json:"base_query"`
		Compute      *[]CompositeCompute                `json:"compute"`
		DataSource   *AggregateSegmentQueryDataSource   `json:"data_source"`
		GroupBy      []CompositeGroupBy                 `json:"group_by,omitempty"`
		Name         *string                            `json:"name"`
		SegmentQuery *AggregateSegmentQuerySegmentQuery `json:"segment_query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BaseQuery == nil {
		return fmt.Errorf("required field base_query missing")
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
	if all.SegmentQuery == nil {
		return fmt.Errorf("required field segment_query missing")
	}

	hasInvalidField := false
	if all.BaseQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BaseQuery = *all.BaseQuery
	o.Compute = *all.Compute
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.GroupBy = all.GroupBy
	o.Name = *all.Name
	if all.SegmentQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SegmentQuery = *all.SegmentQuery

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
