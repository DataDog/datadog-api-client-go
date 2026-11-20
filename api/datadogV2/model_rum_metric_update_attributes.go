// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMMetricUpdateAttributes The RUM-based metric properties to update. Omitted properties keep their current value.
type RUMMetricUpdateAttributes struct {
	// The compute rule properties that can be updated. Only valid for metrics whose `aggregation_type` is `distribution`.
	Compute *RUMMetricUpdateCompute `json:"compute,omitempty"`
	// The RUM-based metric filter. Events matching this filter are aggregated in this metric. The `event_type` filter is added automatically.
	Filter *RUMMetricFilter `json:"filter,omitempty"`
	// The rules to group the metric by. When provided, replaces all the existing group by rules.
	GroupBy []RUMMetricGroupBy `json:"group_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMMetricUpdateAttributes instantiates a new RUMMetricUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMMetricUpdateAttributes() *RUMMetricUpdateAttributes {
	this := RUMMetricUpdateAttributes{}
	return &this
}

// NewRUMMetricUpdateAttributesWithDefaults instantiates a new RUMMetricUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMMetricUpdateAttributesWithDefaults() *RUMMetricUpdateAttributes {
	this := RUMMetricUpdateAttributes{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *RUMMetricUpdateAttributes) GetCompute() RUMMetricUpdateCompute {
	if o == nil || o.Compute == nil {
		var ret RUMMetricUpdateCompute
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMMetricUpdateAttributes) GetComputeOk() (*RUMMetricUpdateCompute, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *RUMMetricUpdateAttributes) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given RUMMetricUpdateCompute and assigns it to the Compute field.
func (o *RUMMetricUpdateAttributes) SetCompute(v RUMMetricUpdateCompute) {
	o.Compute = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *RUMMetricUpdateAttributes) GetFilter() RUMMetricFilter {
	if o == nil || o.Filter == nil {
		var ret RUMMetricFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMMetricUpdateAttributes) GetFilterOk() (*RUMMetricFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *RUMMetricUpdateAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given RUMMetricFilter and assigns it to the Filter field.
func (o *RUMMetricUpdateAttributes) SetFilter(v RUMMetricFilter) {
	o.Filter = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *RUMMetricUpdateAttributes) GetGroupBy() []RUMMetricGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []RUMMetricGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMMetricUpdateAttributes) GetGroupByOk() (*[]RUMMetricGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *RUMMetricUpdateAttributes) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []RUMMetricGroupBy and assigns it to the GroupBy field.
func (o *RUMMetricUpdateAttributes) SetGroupBy(v []RUMMetricGroupBy) {
	o.GroupBy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMMetricUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMMetricUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute *RUMMetricUpdateCompute `json:"compute,omitempty"`
		Filter  *RUMMetricFilter        `json:"filter,omitempty"`
		GroupBy []RUMMetricGroupBy      `json:"group_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compute", "filter", "group_by"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Compute != nil && all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = all.Compute
	if all.Filter != nil && all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filter = all.Filter
	o.GroupBy = all.GroupBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
