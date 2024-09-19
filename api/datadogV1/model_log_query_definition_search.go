// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV1

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"

)


// LogQueryDefinitionSearch The query being made on the logs.
type LogQueryDefinitionSearch struct {
	// Search value to apply.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewLogQueryDefinitionSearch instantiates a new LogQueryDefinitionSearch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogQueryDefinitionSearch(query string) *LogQueryDefinitionSearch {
	this := LogQueryDefinitionSearch{}
	this.Query = query
	return &this
}

// NewLogQueryDefinitionSearchWithDefaults instantiates a new LogQueryDefinitionSearch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogQueryDefinitionSearchWithDefaults() *LogQueryDefinitionSearch {
	this := LogQueryDefinitionSearch{}
	return &this
}
// GetQuery returns the Query field value.
func (o *LogQueryDefinitionSearch) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *LogQueryDefinitionSearch) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *LogQueryDefinitionSearch) SetQuery(v string) {
	o.Query = v
}



// MarshalJSON serializes the struct using spec logic.
func (o LogQueryDefinitionSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["query"] = o.Query

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogQueryDefinitionSearch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query *string `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{ "query",  })
	} else {
		return err
	}
	o.Query = *all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
