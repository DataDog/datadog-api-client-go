// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LegacyEventOverlayQuery Test-only generated schema for the dashboard versioning proof of concept.
type LegacyEventOverlayQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Q string `json:"q"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TagsExecution *TagsExecution `json:"tags_execution,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewLegacyEventOverlayQuery instantiates a new LegacyEventOverlayQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLegacyEventOverlayQuery(q string) *LegacyEventOverlayQuery {
	this := LegacyEventOverlayQuery{}
	this.Q = q
	return &this
}

// NewLegacyEventOverlayQueryWithDefaults instantiates a new LegacyEventOverlayQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLegacyEventOverlayQueryWithDefaults() *LegacyEventOverlayQuery {
	this := LegacyEventOverlayQuery{}
	return &this
}

// GetQ returns the Q field value.
func (o *LegacyEventOverlayQuery) GetQ() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Q
}

// GetQOk returns a tuple with the Q field value
// and a boolean to check if the value has been set.
func (o *LegacyEventOverlayQuery) GetQOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Q, true
}

// SetQ sets field value.
func (o *LegacyEventOverlayQuery) SetQ(v string) {
	o.Q = v
}

// GetTagsExecution returns the TagsExecution field value if set, zero value otherwise.
func (o *LegacyEventOverlayQuery) GetTagsExecution() TagsExecution {
	if o == nil || o.TagsExecution == nil {
		var ret TagsExecution
		return ret
	}
	return *o.TagsExecution
}

// GetTagsExecutionOk returns a tuple with the TagsExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyEventOverlayQuery) GetTagsExecutionOk() (*TagsExecution, bool) {
	if o == nil || o.TagsExecution == nil {
		return nil, false
	}
	return o.TagsExecution, true
}

// HasTagsExecution returns a boolean if a field has been set.
func (o *LegacyEventOverlayQuery) HasTagsExecution() bool {
	return o != nil && o.TagsExecution != nil
}

// SetTagsExecution gets a reference to the given TagsExecution and assigns it to the TagsExecution field.
func (o *LegacyEventOverlayQuery) SetTagsExecution(v TagsExecution) {
	o.TagsExecution = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LegacyEventOverlayQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["q"] = o.Q
	if o.TagsExecution != nil {
		toSerialize["tags_execution"] = o.TagsExecution
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LegacyEventOverlayQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Q             *string        `json:"q"`
		TagsExecution *TagsExecution `json:"tags_execution,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Q == nil {
		return fmt.Errorf("required field q missing")
	}

	hasInvalidField := false
	o.Q = *all.Q
	if all.TagsExecution != nil && !all.TagsExecution.IsValid() {
		hasInvalidField = true
	} else {
		o.TagsExecution = all.TagsExecution
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
