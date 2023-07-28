// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsAggregateSort A sort rule
type LogsAggregateSort struct {
	// An aggregation function
	Aggregation *LogsAggregationFunction `json:"aggregation,omitempty"`
	// The metric to sort by (only used for `type=measure`)
	Metric *string `json:"metric,omitempty"`
	// The order to use, ascending or descending
	Order *LogsSortOrder `json:"order,omitempty"`
	// The type of sorting algorithm
	Type *LogsAggregateSortType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewLogsAggregateSort instantiates a new LogsAggregateSort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsAggregateSort() *LogsAggregateSort {
	this := LogsAggregateSort{}
	var typeVar LogsAggregateSortType = LOGSAGGREGATESORTTYPE_ALPHABETICAL
	this.Type = &typeVar
	return &this
}

// NewLogsAggregateSortWithDefaults instantiates a new LogsAggregateSort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsAggregateSortWithDefaults() *LogsAggregateSort {
	this := LogsAggregateSort{}
	var typeVar LogsAggregateSortType = LOGSAGGREGATESORTTYPE_ALPHABETICAL
	this.Type = &typeVar
	return &this
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *LogsAggregateSort) GetAggregation() LogsAggregationFunction {
	if o == nil || o.Aggregation == nil {
		var ret LogsAggregationFunction
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsAggregateSort) GetAggregationOk() (*LogsAggregationFunction, bool) {
	if o == nil || o.Aggregation == nil {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *LogsAggregateSort) HasAggregation() bool {
	return o != nil && o.Aggregation != nil
}

// SetAggregation gets a reference to the given LogsAggregationFunction and assigns it to the Aggregation field.
func (o *LogsAggregateSort) SetAggregation(v LogsAggregationFunction) {
	o.Aggregation = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *LogsAggregateSort) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsAggregateSort) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *LogsAggregateSort) HasMetric() bool {
	return o != nil && o.Metric != nil
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *LogsAggregateSort) SetMetric(v string) {
	o.Metric = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *LogsAggregateSort) GetOrder() LogsSortOrder {
	if o == nil || o.Order == nil {
		var ret LogsSortOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsAggregateSort) GetOrderOk() (*LogsSortOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *LogsAggregateSort) HasOrder() bool {
	return o != nil && o.Order != nil
}

// SetOrder gets a reference to the given LogsSortOrder and assigns it to the Order field.
func (o *LogsAggregateSort) SetOrder(v LogsSortOrder) {
	o.Order = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LogsAggregateSort) GetType() LogsAggregateSortType {
	if o == nil || o.Type == nil {
		var ret LogsAggregateSortType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsAggregateSort) GetTypeOk() (*LogsAggregateSortType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LogsAggregateSort) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given LogsAggregateSortType and assigns it to the Type field.
func (o *LogsAggregateSort) SetType(v LogsAggregateSortType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsAggregateSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Aggregation != nil {
		toSerialize["aggregation"] = o.Aggregation
	}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsAggregateSort) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Aggregation *LogsAggregationFunction `json:"aggregation,omitempty"`
		Metric      *string                  `json:"metric,omitempty"`
		Order       *LogsSortOrder           `json:"order,omitempty"`
		Type        *LogsAggregateSortType   `json:"type,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation", "metric", "order", "type"})
	} else {
		return err
	}

	var hasInvalidField bool
	if v := all.Aggregation; v != nil && !v.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregation = all.Aggregation
	}
	o.Metric = all.Metric
	if v := all.Order; v != nil && !v.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = all.Order
	}
	if v := all.Type; v != nil && !v.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}

	return nil
}
