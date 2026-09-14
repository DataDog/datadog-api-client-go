// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPromptIncludeItem An explicitly versioned whole-chat or selective-message include.
type LLMObsPromptIncludeItem struct {
	// An explicitly versioned prompt included as chat items. Omitting `items` includes every child message in its original order. When `items` is present, its zero-based indexes are inserted in the order provided; duplicate indexes are preserved.
	Include LLMObsPromptInclude `json:"include"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewLLMObsPromptIncludeItem instantiates a new LLMObsPromptIncludeItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPromptIncludeItem(include LLMObsPromptInclude) *LLMObsPromptIncludeItem {
	this := LLMObsPromptIncludeItem{}
	this.Include = include
	return &this
}

// NewLLMObsPromptIncludeItemWithDefaults instantiates a new LLMObsPromptIncludeItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPromptIncludeItemWithDefaults() *LLMObsPromptIncludeItem {
	this := LLMObsPromptIncludeItem{}
	return &this
}

// GetInclude returns the Include field value.
func (o *LLMObsPromptIncludeItem) GetInclude() LLMObsPromptInclude {
	if o == nil {
		var ret LLMObsPromptInclude
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptIncludeItem) GetIncludeOk() (*LLMObsPromptInclude, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *LLMObsPromptIncludeItem) SetInclude(v LLMObsPromptInclude) {
	o.Include = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPromptIncludeItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["include"] = o.Include
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPromptIncludeItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Include *LLMObsPromptInclude `json:"include"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Include == nil {
		return fmt.Errorf("required field include missing")
	}

	hasInvalidField := false
	if all.Include.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Include = *all.Include

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
