// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisAggregation Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisAggregation struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute []StructuredAnalysisCompute `json:"compute,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []StructuredAnalysisGroupBy `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *float64 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *AnalysisAggregationSort `json:"sort,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TimeAggregation *StructuredAnalysisTimeAggregation `json:"time_aggregation,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type StructuredAnalysisAggregationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewStructuredAnalysisAggregation instantiates a new StructuredAnalysisAggregation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStructuredAnalysisAggregation(typeVar StructuredAnalysisAggregationType) *StructuredAnalysisAggregation {
	this := StructuredAnalysisAggregation{}
	this.Type = typeVar
	return &this
}

// NewStructuredAnalysisAggregationWithDefaults instantiates a new StructuredAnalysisAggregation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStructuredAnalysisAggregationWithDefaults() *StructuredAnalysisAggregation {
	this := StructuredAnalysisAggregation{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *StructuredAnalysisAggregation) GetCompute() []StructuredAnalysisCompute {
	if o == nil || o.Compute == nil {
		var ret []StructuredAnalysisCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisAggregation) GetComputeOk() (*[]StructuredAnalysisCompute, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return &o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *StructuredAnalysisAggregation) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given []StructuredAnalysisCompute and assigns it to the Compute field.
func (o *StructuredAnalysisAggregation) SetCompute(v []StructuredAnalysisCompute) {
	o.Compute = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *StructuredAnalysisAggregation) GetGroupBy() []StructuredAnalysisGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []StructuredAnalysisGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisAggregation) GetGroupByOk() (*[]StructuredAnalysisGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *StructuredAnalysisAggregation) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []StructuredAnalysisGroupBy and assigns it to the GroupBy field.
func (o *StructuredAnalysisAggregation) SetGroupBy(v []StructuredAnalysisGroupBy) {
	o.GroupBy = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *StructuredAnalysisAggregation) GetLimit() float64 {
	if o == nil || o.Limit == nil {
		var ret float64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisAggregation) GetLimitOk() (*float64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *StructuredAnalysisAggregation) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given float64 and assigns it to the Limit field.
func (o *StructuredAnalysisAggregation) SetLimit(v float64) {
	o.Limit = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *StructuredAnalysisAggregation) GetSort() AnalysisAggregationSort {
	if o == nil || o.Sort == nil {
		var ret AnalysisAggregationSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisAggregation) GetSortOk() (*AnalysisAggregationSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *StructuredAnalysisAggregation) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given AnalysisAggregationSort and assigns it to the Sort field.
func (o *StructuredAnalysisAggregation) SetSort(v AnalysisAggregationSort) {
	o.Sort = &v
}

// GetTimeAggregation returns the TimeAggregation field value if set, zero value otherwise.
func (o *StructuredAnalysisAggregation) GetTimeAggregation() StructuredAnalysisTimeAggregation {
	if o == nil || o.TimeAggregation == nil {
		var ret StructuredAnalysisTimeAggregation
		return ret
	}
	return *o.TimeAggregation
}

// GetTimeAggregationOk returns a tuple with the TimeAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisAggregation) GetTimeAggregationOk() (*StructuredAnalysisTimeAggregation, bool) {
	if o == nil || o.TimeAggregation == nil {
		return nil, false
	}
	return o.TimeAggregation, true
}

// HasTimeAggregation returns a boolean if a field has been set.
func (o *StructuredAnalysisAggregation) HasTimeAggregation() bool {
	return o != nil && o.TimeAggregation != nil
}

// SetTimeAggregation gets a reference to the given StructuredAnalysisTimeAggregation and assigns it to the TimeAggregation field.
func (o *StructuredAnalysisAggregation) SetTimeAggregation(v StructuredAnalysisTimeAggregation) {
	o.TimeAggregation = &v
}

// GetType returns the Type field value.
func (o *StructuredAnalysisAggregation) GetType() StructuredAnalysisAggregationType {
	if o == nil {
		var ret StructuredAnalysisAggregationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisAggregation) GetTypeOk() (*StructuredAnalysisAggregationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StructuredAnalysisAggregation) SetType(v StructuredAnalysisAggregationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StructuredAnalysisAggregation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.TimeAggregation != nil {
		toSerialize["time_aggregation"] = o.TimeAggregation
	}
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StructuredAnalysisAggregation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute         []StructuredAnalysisCompute        `json:"compute,omitempty"`
		GroupBy         []StructuredAnalysisGroupBy        `json:"group_by,omitempty"`
		Limit           *float64                           `json:"limit,omitempty"`
		Sort            *AnalysisAggregationSort           `json:"sort,omitempty"`
		TimeAggregation *StructuredAnalysisTimeAggregation `json:"time_aggregation,omitempty"`
		Type            *StructuredAnalysisAggregationType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Compute = all.Compute
	o.GroupBy = all.GroupBy
	o.Limit = all.Limit
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort
	if all.TimeAggregation != nil && all.TimeAggregation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeAggregation = all.TimeAggregation
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
