// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregateAugmentedQuery Test-only generated schema for the dashboard versioning proof of concept.
type AggregateAugmentedQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	AugmentQuery CompositeSourceSubQuery `json:"augment_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	BaseQuery CompositeSourceBaseQuery `json:"base_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute []CompositeCompute `json:"compute"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource AggregateAugmentedQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []CompositeGroupBy `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	JoinCondition AggregateAugmentedQueryJoinCondition `json:"join_condition"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAggregateAugmentedQuery instantiates a new AggregateAugmentedQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregateAugmentedQuery(augmentQuery CompositeSourceSubQuery, baseQuery CompositeSourceBaseQuery, compute []CompositeCompute, dataSource AggregateAugmentedQueryDataSource, joinCondition AggregateAugmentedQueryJoinCondition, name string) *AggregateAugmentedQuery {
	this := AggregateAugmentedQuery{}
	this.AugmentQuery = augmentQuery
	this.BaseQuery = baseQuery
	this.Compute = compute
	this.DataSource = dataSource
	this.JoinCondition = joinCondition
	this.Name = name
	return &this
}

// NewAggregateAugmentedQueryWithDefaults instantiates a new AggregateAugmentedQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregateAugmentedQueryWithDefaults() *AggregateAugmentedQuery {
	this := AggregateAugmentedQuery{}
	return &this
}

// GetAugmentQuery returns the AugmentQuery field value.
func (o *AggregateAugmentedQuery) GetAugmentQuery() CompositeSourceSubQuery {
	if o == nil {
		var ret CompositeSourceSubQuery
		return ret
	}
	return o.AugmentQuery
}

// GetAugmentQueryOk returns a tuple with the AugmentQuery field value
// and a boolean to check if the value has been set.
func (o *AggregateAugmentedQuery) GetAugmentQueryOk() (*CompositeSourceSubQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AugmentQuery, true
}

// SetAugmentQuery sets field value.
func (o *AggregateAugmentedQuery) SetAugmentQuery(v CompositeSourceSubQuery) {
	o.AugmentQuery = v
}

// GetBaseQuery returns the BaseQuery field value.
func (o *AggregateAugmentedQuery) GetBaseQuery() CompositeSourceBaseQuery {
	if o == nil {
		var ret CompositeSourceBaseQuery
		return ret
	}
	return o.BaseQuery
}

// GetBaseQueryOk returns a tuple with the BaseQuery field value
// and a boolean to check if the value has been set.
func (o *AggregateAugmentedQuery) GetBaseQueryOk() (*CompositeSourceBaseQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseQuery, true
}

// SetBaseQuery sets field value.
func (o *AggregateAugmentedQuery) SetBaseQuery(v CompositeSourceBaseQuery) {
	o.BaseQuery = v
}

// GetCompute returns the Compute field value.
func (o *AggregateAugmentedQuery) GetCompute() []CompositeCompute {
	if o == nil {
		var ret []CompositeCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *AggregateAugmentedQuery) GetComputeOk() (*[]CompositeCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *AggregateAugmentedQuery) SetCompute(v []CompositeCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *AggregateAugmentedQuery) GetDataSource() AggregateAugmentedQueryDataSource {
	if o == nil {
		var ret AggregateAugmentedQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *AggregateAugmentedQuery) GetDataSourceOk() (*AggregateAugmentedQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *AggregateAugmentedQuery) SetDataSource(v AggregateAugmentedQueryDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *AggregateAugmentedQuery) GetGroupBy() []CompositeGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []CompositeGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateAugmentedQuery) GetGroupByOk() (*[]CompositeGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *AggregateAugmentedQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []CompositeGroupBy and assigns it to the GroupBy field.
func (o *AggregateAugmentedQuery) SetGroupBy(v []CompositeGroupBy) {
	o.GroupBy = v
}

// GetJoinCondition returns the JoinCondition field value.
func (o *AggregateAugmentedQuery) GetJoinCondition() AggregateAugmentedQueryJoinCondition {
	if o == nil {
		var ret AggregateAugmentedQueryJoinCondition
		return ret
	}
	return o.JoinCondition
}

// GetJoinConditionOk returns a tuple with the JoinCondition field value
// and a boolean to check if the value has been set.
func (o *AggregateAugmentedQuery) GetJoinConditionOk() (*AggregateAugmentedQueryJoinCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinCondition, true
}

// SetJoinCondition sets field value.
func (o *AggregateAugmentedQuery) SetJoinCondition(v AggregateAugmentedQueryJoinCondition) {
	o.JoinCondition = v
}

// GetName returns the Name field value.
func (o *AggregateAugmentedQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AggregateAugmentedQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AggregateAugmentedQuery) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregateAugmentedQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["augment_query"] = o.AugmentQuery
	toSerialize["base_query"] = o.BaseQuery
	toSerialize["compute"] = o.Compute
	toSerialize["data_source"] = o.DataSource
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["join_condition"] = o.JoinCondition
	toSerialize["name"] = o.Name
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregateAugmentedQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AugmentQuery  *CompositeSourceSubQuery              `json:"augment_query"`
		BaseQuery     *CompositeSourceBaseQuery             `json:"base_query"`
		Compute       *[]CompositeCompute                   `json:"compute"`
		DataSource    *AggregateAugmentedQueryDataSource    `json:"data_source"`
		GroupBy       []CompositeGroupBy                    `json:"group_by,omitempty"`
		JoinCondition *AggregateAugmentedQueryJoinCondition `json:"join_condition"`
		Name          *string                               `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AugmentQuery == nil {
		return fmt.Errorf("required field augment_query missing")
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
	if all.JoinCondition == nil {
		return fmt.Errorf("required field join_condition missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}

	hasInvalidField := false
	o.AugmentQuery = *all.AugmentQuery
	o.BaseQuery = *all.BaseQuery
	o.Compute = *all.Compute
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.GroupBy = all.GroupBy
	if all.JoinCondition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JoinCondition = *all.JoinCondition
	o.Name = *all.Name

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
