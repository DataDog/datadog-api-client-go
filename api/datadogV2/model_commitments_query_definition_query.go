// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsQueryDefinitionQuery A metric query of a query definition.
type CommitmentsQueryDefinitionQuery struct {
	// The reducer a scalar widget applies to the query.
	Aggregator CommitmentsQueryDefinitionAggregator `json:"aggregator"`
	// The data source of the query.
	DataSource CommitmentsQueryDefinitionDataSource `json:"data_source"`
	// The name formulas use to reference the query.
	Name string `json:"name"`
	// The metric query.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsQueryDefinitionQuery instantiates a new CommitmentsQueryDefinitionQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsQueryDefinitionQuery(aggregator CommitmentsQueryDefinitionAggregator, dataSource CommitmentsQueryDefinitionDataSource, name string, query string) *CommitmentsQueryDefinitionQuery {
	this := CommitmentsQueryDefinitionQuery{}
	this.Aggregator = aggregator
	this.DataSource = dataSource
	this.Name = name
	this.Query = query
	return &this
}

// NewCommitmentsQueryDefinitionQueryWithDefaults instantiates a new CommitmentsQueryDefinitionQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsQueryDefinitionQueryWithDefaults() *CommitmentsQueryDefinitionQuery {
	this := CommitmentsQueryDefinitionQuery{}
	return &this
}

// GetAggregator returns the Aggregator field value.
func (o *CommitmentsQueryDefinitionQuery) GetAggregator() CommitmentsQueryDefinitionAggregator {
	if o == nil {
		var ret CommitmentsQueryDefinitionAggregator
		return ret
	}
	return o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value
// and a boolean to check if the value has been set.
func (o *CommitmentsQueryDefinitionQuery) GetAggregatorOk() (*CommitmentsQueryDefinitionAggregator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregator, true
}

// SetAggregator sets field value.
func (o *CommitmentsQueryDefinitionQuery) SetAggregator(v CommitmentsQueryDefinitionAggregator) {
	o.Aggregator = v
}

// GetDataSource returns the DataSource field value.
func (o *CommitmentsQueryDefinitionQuery) GetDataSource() CommitmentsQueryDefinitionDataSource {
	if o == nil {
		var ret CommitmentsQueryDefinitionDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *CommitmentsQueryDefinitionQuery) GetDataSourceOk() (*CommitmentsQueryDefinitionDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *CommitmentsQueryDefinitionQuery) SetDataSource(v CommitmentsQueryDefinitionDataSource) {
	o.DataSource = v
}

// GetName returns the Name field value.
func (o *CommitmentsQueryDefinitionQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CommitmentsQueryDefinitionQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CommitmentsQueryDefinitionQuery) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value.
func (o *CommitmentsQueryDefinitionQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *CommitmentsQueryDefinitionQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *CommitmentsQueryDefinitionQuery) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsQueryDefinitionQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregator"] = o.Aggregator
	toSerialize["data_source"] = o.DataSource
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsQueryDefinitionQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregator *CommitmentsQueryDefinitionAggregator `json:"aggregator"`
		DataSource *CommitmentsQueryDefinitionDataSource `json:"data_source"`
		Name       *string                               `json:"name"`
		Query      *string                               `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregator == nil {
		return fmt.Errorf("required field aggregator missing")
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregator", "data_source", "name", "query"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Aggregator.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregator = *all.Aggregator
	}
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Name = *all.Name
	o.Query = *all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
