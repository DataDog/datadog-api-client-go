// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPromptInclude An explicitly versioned prompt included as chat items. Omitting `items` includes every child message in its original order. When `items` is present, its zero-based indexes are inserted in the order provided; duplicate indexes are preserved.
type LLMObsPromptInclude struct {
	// Optional ordered zero-based child-message indexes. Order and duplicate indexes are preserved.
	Items []int64 `json:"items,omitempty"`
	// Customer-provided identifier of the included prompt. It cannot contain spaces, tabs, line breaks, braces, an equals sign, a comma, or quotes.
	PromptId string `json:"prompt_id"`
	// Positive sequential version number of the included prompt.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewLLMObsPromptInclude instantiates a new LLMObsPromptInclude object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPromptInclude(promptId string, version int64) *LLMObsPromptInclude {
	this := LLMObsPromptInclude{}
	this.PromptId = promptId
	this.Version = version
	return &this
}

// NewLLMObsPromptIncludeWithDefaults instantiates a new LLMObsPromptInclude object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPromptIncludeWithDefaults() *LLMObsPromptInclude {
	this := LLMObsPromptInclude{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *LLMObsPromptInclude) GetItems() []int64 {
	if o == nil || o.Items == nil {
		var ret []int64
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptInclude) GetItemsOk() (*[]int64, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *LLMObsPromptInclude) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []int64 and assigns it to the Items field.
func (o *LLMObsPromptInclude) SetItems(v []int64) {
	o.Items = v
}

// GetPromptId returns the PromptId field value.
func (o *LLMObsPromptInclude) GetPromptId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PromptId
}

// GetPromptIdOk returns a tuple with the PromptId field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptInclude) GetPromptIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromptId, true
}

// SetPromptId sets field value.
func (o *LLMObsPromptInclude) SetPromptId(v string) {
	o.PromptId = v
}

// GetVersion returns the Version field value.
func (o *LLMObsPromptInclude) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptInclude) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *LLMObsPromptInclude) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPromptInclude) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	toSerialize["prompt_id"] = o.PromptId
	toSerialize["version"] = o.Version
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPromptInclude) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items    []int64 `json:"items,omitempty"`
		PromptId *string `json:"prompt_id"`
		Version  *int64  `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PromptId == nil {
		return fmt.Errorf("required field prompt_id missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	o.Items = all.Items
	o.PromptId = *all.PromptId
	o.Version = *all.Version

	return nil
}
