// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LegacyEventManagementQuery Test-only generated schema for the dashboard versioning proof of concept.
type LegacyEventManagementQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Search string `json:"search"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TagsExecution TagsExecution `json:"tags_execution"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewLegacyEventManagementQuery instantiates a new LegacyEventManagementQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLegacyEventManagementQuery(search string, tagsExecution TagsExecution) *LegacyEventManagementQuery {
	this := LegacyEventManagementQuery{}
	this.Search = search
	this.TagsExecution = tagsExecution
	return &this
}

// NewLegacyEventManagementQueryWithDefaults instantiates a new LegacyEventManagementQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLegacyEventManagementQueryWithDefaults() *LegacyEventManagementQuery {
	this := LegacyEventManagementQuery{}
	return &this
}

// GetSearch returns the Search field value.
func (o *LegacyEventManagementQuery) GetSearch() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *LegacyEventManagementQuery) GetSearchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *LegacyEventManagementQuery) SetSearch(v string) {
	o.Search = v
}

// GetTagsExecution returns the TagsExecution field value.
func (o *LegacyEventManagementQuery) GetTagsExecution() TagsExecution {
	if o == nil {
		var ret TagsExecution
		return ret
	}
	return o.TagsExecution
}

// GetTagsExecutionOk returns a tuple with the TagsExecution field value
// and a boolean to check if the value has been set.
func (o *LegacyEventManagementQuery) GetTagsExecutionOk() (*TagsExecution, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagsExecution, true
}

// SetTagsExecution sets field value.
func (o *LegacyEventManagementQuery) SetTagsExecution(v TagsExecution) {
	o.TagsExecution = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LegacyEventManagementQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["search"] = o.Search
	toSerialize["tags_execution"] = o.TagsExecution
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LegacyEventManagementQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Search        *string        `json:"search"`
		TagsExecution *TagsExecution `json:"tags_execution"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Search == nil {
		return fmt.Errorf("required field search missing")
	}
	if all.TagsExecution == nil {
		return fmt.Errorf("required field tags_execution missing")
	}

	hasInvalidField := false
	o.Search = *all.Search
	if !all.TagsExecution.IsValid() {
		hasInvalidField = true
	} else {
		o.TagsExecution = *all.TagsExecution
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
