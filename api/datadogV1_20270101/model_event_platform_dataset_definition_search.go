// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventPlatformDatasetDefinitionSearch Test-only generated schema for the dashboard versioning proof of concept.
type EventPlatformDatasetDefinitionSearch struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEventPlatformDatasetDefinitionSearch instantiates a new EventPlatformDatasetDefinitionSearch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventPlatformDatasetDefinitionSearch(query string) *EventPlatformDatasetDefinitionSearch {
	this := EventPlatformDatasetDefinitionSearch{}
	this.Query = query
	return &this
}

// NewEventPlatformDatasetDefinitionSearchWithDefaults instantiates a new EventPlatformDatasetDefinitionSearch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventPlatformDatasetDefinitionSearchWithDefaults() *EventPlatformDatasetDefinitionSearch {
	this := EventPlatformDatasetDefinitionSearch{}
	return &this
}

// GetQuery returns the Query field value.
func (o *EventPlatformDatasetDefinitionSearch) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *EventPlatformDatasetDefinitionSearch) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *EventPlatformDatasetDefinitionSearch) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventPlatformDatasetDefinitionSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["query"] = o.Query
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventPlatformDatasetDefinitionSearch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Query *string `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	o.Query = *all.Query

	return nil
}
