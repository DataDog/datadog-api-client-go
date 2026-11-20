// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMMetricCreateAttributes The object describing the Datadog RUM-based metric to create.
type RUMMetricCreateAttributes struct {
	// The compute rule to compute the RUM-based metric. It is set at creation and only `include_percentiles` can be updated afterward.
	Compute RUMMetricCompute `json:"compute"`
	// The type of RUM events the metric is generated from. It is set at creation and cannot be changed.
	EventType RUMMetricEventType `json:"event_type"`
	// The RUM-based metric filter. Events matching this filter are aggregated in this metric. The `event_type` filter is added automatically.
	Filter *RUMMetricFilter `json:"filter,omitempty"`
	// The rules to group the metric by. Each rule creates a tag on the generated metric. If omitted, the metric is not grouped.
	GroupBy []RUMMetricGroupBy `json:"group_by,omitempty"`
	// The rule to count updatable events.
	// Required when `event_type` is `session` or `view`, and must be omitted for other event types. It cannot be changed after creation.
	Uniqueness *RUMMetricUniqueness `json:"uniqueness,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMMetricCreateAttributes instantiates a new RUMMetricCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMMetricCreateAttributes(compute RUMMetricCompute, eventType RUMMetricEventType) *RUMMetricCreateAttributes {
	this := RUMMetricCreateAttributes{}
	this.Compute = compute
	this.EventType = eventType
	return &this
}

// NewRUMMetricCreateAttributesWithDefaults instantiates a new RUMMetricCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMMetricCreateAttributesWithDefaults() *RUMMetricCreateAttributes {
	this := RUMMetricCreateAttributes{}
	return &this
}

// GetCompute returns the Compute field value.
func (o *RUMMetricCreateAttributes) GetCompute() RUMMetricCompute {
	if o == nil {
		var ret RUMMetricCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *RUMMetricCreateAttributes) GetComputeOk() (*RUMMetricCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *RUMMetricCreateAttributes) SetCompute(v RUMMetricCompute) {
	o.Compute = v
}

// GetEventType returns the EventType field value.
func (o *RUMMetricCreateAttributes) GetEventType() RUMMetricEventType {
	if o == nil {
		var ret RUMMetricEventType
		return ret
	}
	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *RUMMetricCreateAttributes) GetEventTypeOk() (*RUMMetricEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value.
func (o *RUMMetricCreateAttributes) SetEventType(v RUMMetricEventType) {
	o.EventType = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *RUMMetricCreateAttributes) GetFilter() RUMMetricFilter {
	if o == nil || o.Filter == nil {
		var ret RUMMetricFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMMetricCreateAttributes) GetFilterOk() (*RUMMetricFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *RUMMetricCreateAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given RUMMetricFilter and assigns it to the Filter field.
func (o *RUMMetricCreateAttributes) SetFilter(v RUMMetricFilter) {
	o.Filter = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *RUMMetricCreateAttributes) GetGroupBy() []RUMMetricGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []RUMMetricGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMMetricCreateAttributes) GetGroupByOk() (*[]RUMMetricGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *RUMMetricCreateAttributes) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []RUMMetricGroupBy and assigns it to the GroupBy field.
func (o *RUMMetricCreateAttributes) SetGroupBy(v []RUMMetricGroupBy) {
	o.GroupBy = v
}

// GetUniqueness returns the Uniqueness field value if set, zero value otherwise.
func (o *RUMMetricCreateAttributes) GetUniqueness() RUMMetricUniqueness {
	if o == nil || o.Uniqueness == nil {
		var ret RUMMetricUniqueness
		return ret
	}
	return *o.Uniqueness
}

// GetUniquenessOk returns a tuple with the Uniqueness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMMetricCreateAttributes) GetUniquenessOk() (*RUMMetricUniqueness, bool) {
	if o == nil || o.Uniqueness == nil {
		return nil, false
	}
	return o.Uniqueness, true
}

// HasUniqueness returns a boolean if a field has been set.
func (o *RUMMetricCreateAttributes) HasUniqueness() bool {
	return o != nil && o.Uniqueness != nil
}

// SetUniqueness gets a reference to the given RUMMetricUniqueness and assigns it to the Uniqueness field.
func (o *RUMMetricCreateAttributes) SetUniqueness(v RUMMetricUniqueness) {
	o.Uniqueness = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMMetricCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["compute"] = o.Compute
	toSerialize["event_type"] = o.EventType
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Uniqueness != nil {
		toSerialize["uniqueness"] = o.Uniqueness
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMMetricCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute    *RUMMetricCompute    `json:"compute"`
		EventType  *RUMMetricEventType  `json:"event_type"`
		Filter     *RUMMetricFilter     `json:"filter,omitempty"`
		GroupBy    []RUMMetricGroupBy   `json:"group_by,omitempty"`
		Uniqueness *RUMMetricUniqueness `json:"uniqueness,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Compute == nil {
		return fmt.Errorf("required field compute missing")
	}
	if all.EventType == nil {
		return fmt.Errorf("required field event_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compute", "event_type", "filter", "group_by", "uniqueness"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = *all.Compute
	if !all.EventType.IsValid() {
		hasInvalidField = true
	} else {
		o.EventType = *all.EventType
	}
	if all.Filter != nil && all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filter = all.Filter
	o.GroupBy = all.GroupBy
	if all.Uniqueness != nil && all.Uniqueness.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Uniqueness = all.Uniqueness

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
