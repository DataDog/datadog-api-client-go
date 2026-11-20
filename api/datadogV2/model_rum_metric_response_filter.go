// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMMetricResponseFilter The RUM-based metric filter. RUM events matching this filter are aggregated in this metric. Omitted if the metric only filters on the event type.
type RUMMetricResponseFilter struct {
	// The search query, following the [RUM search syntax](https://docs.datadoghq.com/real_user_monitoring/explorer/search_syntax/). It does not include the event type filter.
	Query *string `json:"query,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMMetricResponseFilter instantiates a new RUMMetricResponseFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMMetricResponseFilter() *RUMMetricResponseFilter {
	this := RUMMetricResponseFilter{}
	return &this
}

// NewRUMMetricResponseFilterWithDefaults instantiates a new RUMMetricResponseFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMMetricResponseFilterWithDefaults() *RUMMetricResponseFilter {
	this := RUMMetricResponseFilter{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *RUMMetricResponseFilter) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMMetricResponseFilter) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *RUMMetricResponseFilter) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *RUMMetricResponseFilter) SetQuery(v string) {
	o.Query = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMMetricResponseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMMetricResponseFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query *string `json:"query,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"query"})
	} else {
		return err
	}
	o.Query = all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
