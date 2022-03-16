/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */


package datadog

import (
	"encoding/json"
	"fmt"

)


// LogsQueryFilter The search and filter query settings
type LogsQueryFilter struct {
	// The minimum time for the requested logs, supports date math and regular timestamps (milliseconds).
	From *string `json:"from,omitempty"`
	// For customers with multiple indexes, the indexes to search. Defaults to ['*'] which means all indexes.
	Indexes *[]string `json:"indexes,omitempty"`
	// The search query - following the log search syntax.
	Query *string `json:"query,omitempty"`
	// The maximum time for the requested logs, supports date math and regular timestamps (milliseconds).
	To *string `json:"to,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}



// NewLogsQueryFilter instantiates a new LogsQueryFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsQueryFilter() *LogsQueryFilter {
	this := LogsQueryFilter{}
	var from string = "now-15m"
	this.From = &from
	var query string = "*"
	this.Query = &query
	var to string = "now"
	this.To = &to
	return &this
}

// NewLogsQueryFilterWithDefaults instantiates a new LogsQueryFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsQueryFilterWithDefaults() *LogsQueryFilter {
	this := LogsQueryFilter{}
	var from string = "now-15m"
	this.From = &from
	var query string = "*"
	this.Query = &query
	var to string = "now"
	this.To = &to
	return &this
}
// GetFrom returns the From field value if set, zero value otherwise.
func (o *LogsQueryFilter) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsQueryFilter) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *LogsQueryFilter) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *LogsQueryFilter) SetFrom(v string) {
	o.From = &v
}


// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *LogsQueryFilter) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return *o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsQueryFilter) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *LogsQueryFilter) HasIndexes() bool {
	if o != nil && o.Indexes != nil {
		return true
	}

	return false
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *LogsQueryFilter) SetIndexes(v []string) {
	o.Indexes = &v
}


// GetQuery returns the Query field value if set, zero value otherwise.
func (o *LogsQueryFilter) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsQueryFilter) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *LogsQueryFilter) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *LogsQueryFilter) SetQuery(v string) {
	o.Query = &v
}


// GetTo returns the To field value if set, zero value otherwise.
func (o *LogsQueryFilter) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsQueryFilter) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *LogsQueryFilter) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *LogsQueryFilter) SetTo(v string) {
	o.To = &v
}



func (o LogsQueryFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}


func (o *LogsQueryFilter) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		From *string `json:"from,omitempty"`
		Indexes *[]string `json:"indexes,omitempty"`
		Query *string `json:"query,omitempty"`
		To *string `json:"to,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.From = all.From
	o.Indexes = all.Indexes
	o.Query = all.Query
	o.To = all.To
	return nil
}
