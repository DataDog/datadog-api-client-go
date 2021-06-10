/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// EventQueryDefinition The event query.
type EventQueryDefinition struct {
	// The query being made on the event.
	Search string `json:"search"`
	// The execution method for multi-value filters. Can be either and or or.
	TagsExecution string `json:"tags_execution"`
}

// NewEventQueryDefinition instantiates a new EventQueryDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventQueryDefinition(search string, tagsExecution string) *EventQueryDefinition {
	this := EventQueryDefinition{}
	this.Search = search
	this.TagsExecution = tagsExecution
	return &this
}

// NewEventQueryDefinitionWithDefaults instantiates a new EventQueryDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventQueryDefinitionWithDefaults() *EventQueryDefinition {
	this := EventQueryDefinition{}
	return &this
}

// GetSearch returns the Search field value
func (o *EventQueryDefinition) GetSearch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *EventQueryDefinition) GetSearchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value
func (o *EventQueryDefinition) SetSearch(v string) {
	o.Search = v
}

// GetTagsExecution returns the TagsExecution field value
func (o *EventQueryDefinition) GetTagsExecution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TagsExecution
}

// GetTagsExecutionOk returns a tuple with the TagsExecution field value
// and a boolean to check if the value has been set.
func (o *EventQueryDefinition) GetTagsExecutionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagsExecution, true
}

// SetTagsExecution sets field value
func (o *EventQueryDefinition) SetTagsExecution(v string) {
	o.TagsExecution = v
}

func (o EventQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["search"] = o.Search
	}
	if true {
		toSerialize["tags_execution"] = o.TagsExecution
	}
	return json.Marshal(toSerialize)
}

func (o *EventQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Search        *string `json:"search"`
		TagsExecution *string `json:"tags_execution"`
	}{}
	all := struct {
		Search        string `json:"search"`
		TagsExecution string `json:"tags_execution"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Search == nil {
		return fmt.Errorf("Required field search missing")
	}
	if required.TagsExecution == nil {
		return fmt.Errorf("Required field tags_execution missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Search = all.Search
	o.TagsExecution = all.TagsExecution
	return nil
}

type NullableEventQueryDefinition struct {
	value *EventQueryDefinition
	isSet bool
}

func (v NullableEventQueryDefinition) Get() *EventQueryDefinition {
	return v.value
}

func (v *NullableEventQueryDefinition) Set(val *EventQueryDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableEventQueryDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableEventQueryDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventQueryDefinition(val *EventQueryDefinition) *NullableEventQueryDefinition {
	return &NullableEventQueryDefinition{value: val, isSet: true}
}

func (v NullableEventQueryDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventQueryDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
