// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostAIPreferredTagsAttributes Attributes for the preferred tags response.
type CostAIPreferredTagsAttributes struct {
	// A contextual message about the preferred tags configuration.
	Message string `json:"message"`
	// The list of preferred cost allocation tags.
	PreferredTags []string `json:"preferred_tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostAIPreferredTagsAttributes instantiates a new CostAIPreferredTagsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostAIPreferredTagsAttributes(message string, preferredTags []string) *CostAIPreferredTagsAttributes {
	this := CostAIPreferredTagsAttributes{}
	this.Message = message
	this.PreferredTags = preferredTags
	return &this
}

// NewCostAIPreferredTagsAttributesWithDefaults instantiates a new CostAIPreferredTagsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostAIPreferredTagsAttributesWithDefaults() *CostAIPreferredTagsAttributes {
	this := CostAIPreferredTagsAttributes{}
	return &this
}

// GetMessage returns the Message field value.
func (o *CostAIPreferredTagsAttributes) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CostAIPreferredTagsAttributes) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *CostAIPreferredTagsAttributes) SetMessage(v string) {
	o.Message = v
}

// GetPreferredTags returns the PreferredTags field value.
func (o *CostAIPreferredTagsAttributes) GetPreferredTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PreferredTags
}

// GetPreferredTagsOk returns a tuple with the PreferredTags field value
// and a boolean to check if the value has been set.
func (o *CostAIPreferredTagsAttributes) GetPreferredTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreferredTags, true
}

// SetPreferredTags sets field value.
func (o *CostAIPreferredTagsAttributes) SetPreferredTags(v []string) {
	o.PreferredTags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostAIPreferredTagsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["message"] = o.Message
	toSerialize["preferred_tags"] = o.PreferredTags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostAIPreferredTagsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message       *string   `json:"message"`
		PreferredTags *[]string `json:"preferred_tags"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.PreferredTags == nil {
		return fmt.Errorf("required field preferred_tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"message", "preferred_tags"})
	} else {
		return err
	}
	o.Message = *all.Message
	o.PreferredTags = *all.PreferredTags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
