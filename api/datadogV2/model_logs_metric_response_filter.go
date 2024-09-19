// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV2

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"

)


// LogsMetricResponseFilter The log-based metric filter. Logs matching this filter will be aggregated in this metric.
type LogsMetricResponseFilter struct {
	// The search query - following the log search syntax.
	Query *string `json:"query,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewLogsMetricResponseFilter instantiates a new LogsMetricResponseFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsMetricResponseFilter() *LogsMetricResponseFilter {
	this := LogsMetricResponseFilter{}
	return &this
}

// NewLogsMetricResponseFilterWithDefaults instantiates a new LogsMetricResponseFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsMetricResponseFilterWithDefaults() *LogsMetricResponseFilter {
	this := LogsMetricResponseFilter{}
	return &this
}
// GetQuery returns the Query field value if set, zero value otherwise.
func (o *LogsMetricResponseFilter) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsMetricResponseFilter) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *LogsMetricResponseFilter) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *LogsMetricResponseFilter) SetQuery(v string) {
	o.Query = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o LogsMetricResponseFilter) MarshalJSON() ([]byte, error) {
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
func (o *LogsMetricResponseFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query *string `json:"query,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{ "query",  })
	} else {
		return err
	}
	o.Query = all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
