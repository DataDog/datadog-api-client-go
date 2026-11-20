// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMMetricResponseAttributes The object describing a Datadog RUM-based metric.
type RUMMetricResponseAttributes struct {
	// The compute rule of the RUM-based metric.
	Compute *RUMMetricResponseCompute `json:"compute,omitempty"`
	// The type of RUM events the metric is generated from. It is set at creation and cannot be changed.
	EventType *RUMMetricEventType `json:"event_type,omitempty"`
	// The RUM-based metric filter. RUM events matching this filter are aggregated in this metric. Omitted if the metric only filters on the event type.
	Filter *RUMMetricResponseFilter `json:"filter,omitempty"`
	// The rules the metric is grouped by. An empty array is returned if the metric is not grouped.
	GroupBy []RUMMetricResponseGroupBy `json:"group_by,omitempty"`
	// The rule to count updatable events. Only present if `event_type` is `session` or `view`, and absent on older metrics created before this setting existed.
	Uniqueness *RUMMetricResponseUniqueness `json:"uniqueness,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMMetricResponseAttributes instantiates a new RUMMetricResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMMetricResponseAttributes() *RUMMetricResponseAttributes {
	this := RUMMetricResponseAttributes{}
	return &this
}

// NewRUMMetricResponseAttributesWithDefaults instantiates a new RUMMetricResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMMetricResponseAttributesWithDefaults() *RUMMetricResponseAttributes {
	this := RUMMetricResponseAttributes{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *RUMMetricResponseAttributes) GetCompute() RUMMetricResponseCompute {
	if o == nil || o.Compute == nil {
		var ret RUMMetricResponseCompute
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMMetricResponseAttributes) GetComputeOk() (*RUMMetricResponseCompute, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *RUMMetricResponseAttributes) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given RUMMetricResponseCompute and assigns it to the Compute field.
func (o *RUMMetricResponseAttributes) SetCompute(v RUMMetricResponseCompute) {
	o.Compute = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *RUMMetricResponseAttributes) GetEventType() RUMMetricEventType {
	if o == nil || o.EventType == nil {
		var ret RUMMetricEventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMMetricResponseAttributes) GetEventTypeOk() (*RUMMetricEventType, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *RUMMetricResponseAttributes) HasEventType() bool {
	return o != nil && o.EventType != nil
}

// SetEventType gets a reference to the given RUMMetricEventType and assigns it to the EventType field.
func (o *RUMMetricResponseAttributes) SetEventType(v RUMMetricEventType) {
	o.EventType = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *RUMMetricResponseAttributes) GetFilter() RUMMetricResponseFilter {
	if o == nil || o.Filter == nil {
		var ret RUMMetricResponseFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMMetricResponseAttributes) GetFilterOk() (*RUMMetricResponseFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *RUMMetricResponseAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given RUMMetricResponseFilter and assigns it to the Filter field.
func (o *RUMMetricResponseAttributes) SetFilter(v RUMMetricResponseFilter) {
	o.Filter = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *RUMMetricResponseAttributes) GetGroupBy() []RUMMetricResponseGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []RUMMetricResponseGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMMetricResponseAttributes) GetGroupByOk() (*[]RUMMetricResponseGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *RUMMetricResponseAttributes) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []RUMMetricResponseGroupBy and assigns it to the GroupBy field.
func (o *RUMMetricResponseAttributes) SetGroupBy(v []RUMMetricResponseGroupBy) {
	o.GroupBy = v
}

// GetUniqueness returns the Uniqueness field value if set, zero value otherwise.
func (o *RUMMetricResponseAttributes) GetUniqueness() RUMMetricResponseUniqueness {
	if o == nil || o.Uniqueness == nil {
		var ret RUMMetricResponseUniqueness
		return ret
	}
	return *o.Uniqueness
}

// GetUniquenessOk returns a tuple with the Uniqueness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMMetricResponseAttributes) GetUniquenessOk() (*RUMMetricResponseUniqueness, bool) {
	if o == nil || o.Uniqueness == nil {
		return nil, false
	}
	return o.Uniqueness, true
}

// HasUniqueness returns a boolean if a field has been set.
func (o *RUMMetricResponseAttributes) HasUniqueness() bool {
	return o != nil && o.Uniqueness != nil
}

// SetUniqueness gets a reference to the given RUMMetricResponseUniqueness and assigns it to the Uniqueness field.
func (o *RUMMetricResponseAttributes) SetUniqueness(v RUMMetricResponseUniqueness) {
	o.Uniqueness = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMMetricResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
	if o.EventType != nil {
		toSerialize["event_type"] = o.EventType
	}
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
func (o *RUMMetricResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute    *RUMMetricResponseCompute    `json:"compute,omitempty"`
		EventType  *RUMMetricEventType          `json:"event_type,omitempty"`
		Filter     *RUMMetricResponseFilter     `json:"filter,omitempty"`
		GroupBy    []RUMMetricResponseGroupBy   `json:"group_by,omitempty"`
		Uniqueness *RUMMetricResponseUniqueness `json:"uniqueness,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compute", "event_type", "filter", "group_by", "uniqueness"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Compute != nil && all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = all.Compute
	if all.EventType != nil && !all.EventType.IsValid() {
		hasInvalidField = true
	} else {
		o.EventType = all.EventType
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
