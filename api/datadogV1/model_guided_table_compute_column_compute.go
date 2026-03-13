// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableComputeColumnCompute Aggregation configuration.
type GuidedTableComputeColumnCompute struct {
	// Aggregation type (e.g. `avg`, `sum`, `count`).
	Aggregation string `json:"aggregation"`
	// Additional filter criteria.
	Filter *string `json:"filter,omitempty"`
	// Attribute to aggregate on. Depends on the data source (e.g. a log attribute).
	Property *string `json:"property,omitempty"`
	// Name of the source query.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGuidedTableComputeColumnCompute instantiates a new GuidedTableComputeColumnCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGuidedTableComputeColumnCompute(aggregation string, query string) *GuidedTableComputeColumnCompute {
	this := GuidedTableComputeColumnCompute{}
	this.Aggregation = aggregation
	this.Query = query
	return &this
}

// NewGuidedTableComputeColumnComputeWithDefaults instantiates a new GuidedTableComputeColumnCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGuidedTableComputeColumnComputeWithDefaults() *GuidedTableComputeColumnCompute {
	this := GuidedTableComputeColumnCompute{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *GuidedTableComputeColumnCompute) GetAggregation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *GuidedTableComputeColumnCompute) GetAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *GuidedTableComputeColumnCompute) SetAggregation(v string) {
	o.Aggregation = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *GuidedTableComputeColumnCompute) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableComputeColumnCompute) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *GuidedTableComputeColumnCompute) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *GuidedTableComputeColumnCompute) SetFilter(v string) {
	o.Filter = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *GuidedTableComputeColumnCompute) GetProperty() string {
	if o == nil || o.Property == nil {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableComputeColumnCompute) GetPropertyOk() (*string, bool) {
	if o == nil || o.Property == nil {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *GuidedTableComputeColumnCompute) HasProperty() bool {
	return o != nil && o.Property != nil
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *GuidedTableComputeColumnCompute) SetProperty(v string) {
	o.Property = &v
}

// GetQuery returns the Query field value.
func (o *GuidedTableComputeColumnCompute) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *GuidedTableComputeColumnCompute) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *GuidedTableComputeColumnCompute) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GuidedTableComputeColumnCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Property != nil {
		toSerialize["property"] = o.Property
	}
	toSerialize["query"] = o.Query

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GuidedTableComputeColumnCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *string `json:"aggregation"`
		Filter      *string `json:"filter,omitempty"`
		Property    *string `json:"property,omitempty"`
		Query       *string `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregation == nil {
		return fmt.Errorf("required field aggregation missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation", "filter", "property", "query"})
	} else {
		return err
	}
	o.Aggregation = *all.Aggregation
	o.Filter = all.Filter
	o.Property = all.Property
	o.Query = *all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
