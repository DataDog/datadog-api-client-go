// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListAugmentedQuery Test-only generated schema for the dashboard versioning proof of concept.
type ListAugmentedQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	AugmentQuery CompositeSourceSubQuery `json:"augment_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	BaseQuery CompositeSourceBaseQuery `json:"base_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Context *EventsContext `json:"context,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource ListAugmentedQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	JoinCondition ListAugmentedQueryJoinCondition `json:"join_condition"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewListAugmentedQuery instantiates a new ListAugmentedQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListAugmentedQuery(augmentQuery CompositeSourceSubQuery, baseQuery CompositeSourceBaseQuery, dataSource ListAugmentedQueryDataSource, joinCondition ListAugmentedQueryJoinCondition, name string) *ListAugmentedQuery {
	this := ListAugmentedQuery{}
	this.AugmentQuery = augmentQuery
	this.BaseQuery = baseQuery
	this.DataSource = dataSource
	this.JoinCondition = joinCondition
	this.Name = name
	return &this
}

// NewListAugmentedQueryWithDefaults instantiates a new ListAugmentedQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListAugmentedQueryWithDefaults() *ListAugmentedQuery {
	this := ListAugmentedQuery{}
	return &this
}

// GetAugmentQuery returns the AugmentQuery field value.
func (o *ListAugmentedQuery) GetAugmentQuery() CompositeSourceSubQuery {
	if o == nil {
		var ret CompositeSourceSubQuery
		return ret
	}
	return o.AugmentQuery
}

// GetAugmentQueryOk returns a tuple with the AugmentQuery field value
// and a boolean to check if the value has been set.
func (o *ListAugmentedQuery) GetAugmentQueryOk() (*CompositeSourceSubQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AugmentQuery, true
}

// SetAugmentQuery sets field value.
func (o *ListAugmentedQuery) SetAugmentQuery(v CompositeSourceSubQuery) {
	o.AugmentQuery = v
}

// GetBaseQuery returns the BaseQuery field value.
func (o *ListAugmentedQuery) GetBaseQuery() CompositeSourceBaseQuery {
	if o == nil {
		var ret CompositeSourceBaseQuery
		return ret
	}
	return o.BaseQuery
}

// GetBaseQueryOk returns a tuple with the BaseQuery field value
// and a boolean to check if the value has been set.
func (o *ListAugmentedQuery) GetBaseQueryOk() (*CompositeSourceBaseQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseQuery, true
}

// SetBaseQuery sets field value.
func (o *ListAugmentedQuery) SetBaseQuery(v CompositeSourceBaseQuery) {
	o.BaseQuery = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ListAugmentedQuery) GetContext() EventsContext {
	if o == nil || o.Context == nil {
		var ret EventsContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAugmentedQuery) GetContextOk() (*EventsContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ListAugmentedQuery) HasContext() bool {
	return o != nil && o.Context != nil
}

// SetContext gets a reference to the given EventsContext and assigns it to the Context field.
func (o *ListAugmentedQuery) SetContext(v EventsContext) {
	o.Context = &v
}

// GetDataSource returns the DataSource field value.
func (o *ListAugmentedQuery) GetDataSource() ListAugmentedQueryDataSource {
	if o == nil {
		var ret ListAugmentedQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *ListAugmentedQuery) GetDataSourceOk() (*ListAugmentedQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *ListAugmentedQuery) SetDataSource(v ListAugmentedQueryDataSource) {
	o.DataSource = v
}

// GetJoinCondition returns the JoinCondition field value.
func (o *ListAugmentedQuery) GetJoinCondition() ListAugmentedQueryJoinCondition {
	if o == nil {
		var ret ListAugmentedQueryJoinCondition
		return ret
	}
	return o.JoinCondition
}

// GetJoinConditionOk returns a tuple with the JoinCondition field value
// and a boolean to check if the value has been set.
func (o *ListAugmentedQuery) GetJoinConditionOk() (*ListAugmentedQueryJoinCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinCondition, true
}

// SetJoinCondition sets field value.
func (o *ListAugmentedQuery) SetJoinCondition(v ListAugmentedQueryJoinCondition) {
	o.JoinCondition = v
}

// GetName returns the Name field value.
func (o *ListAugmentedQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListAugmentedQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ListAugmentedQuery) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListAugmentedQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["augment_query"] = o.AugmentQuery
	toSerialize["base_query"] = o.BaseQuery
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["join_condition"] = o.JoinCondition
	toSerialize["name"] = o.Name
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListAugmentedQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AugmentQuery  *CompositeSourceSubQuery         `json:"augment_query"`
		BaseQuery     *CompositeSourceBaseQuery        `json:"base_query"`
		Context       *EventsContext                   `json:"context,omitempty"`
		DataSource    *ListAugmentedQueryDataSource    `json:"data_source"`
		JoinCondition *ListAugmentedQueryJoinCondition `json:"join_condition"`
		Name          *string                          `json:"name"`
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
	if all.Context != nil && all.Context.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Context = all.Context
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
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
