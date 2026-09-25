// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeneralInvestigationAttributesWithoutTimeBounds Attributes for a general investigation without an explicit time window.
type GeneralInvestigationAttributesWithoutTimeBounds struct {
	// A free-form description of what to investigate, up to 4,096 characters.
	Description string `json:"description"`
	// Tags that scope the investigation.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGeneralInvestigationAttributesWithoutTimeBounds instantiates a new GeneralInvestigationAttributesWithoutTimeBounds object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeneralInvestigationAttributesWithoutTimeBounds(description string) *GeneralInvestigationAttributesWithoutTimeBounds {
	this := GeneralInvestigationAttributesWithoutTimeBounds{}
	this.Description = description
	return &this
}

// NewGeneralInvestigationAttributesWithoutTimeBoundsWithDefaults instantiates a new GeneralInvestigationAttributesWithoutTimeBounds object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeneralInvestigationAttributesWithoutTimeBoundsWithDefaults() *GeneralInvestigationAttributesWithoutTimeBounds {
	this := GeneralInvestigationAttributesWithoutTimeBounds{}
	return &this
}

// GetDescription returns the Description field value.
func (o *GeneralInvestigationAttributesWithoutTimeBounds) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GeneralInvestigationAttributesWithoutTimeBounds) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *GeneralInvestigationAttributesWithoutTimeBounds) SetDescription(v string) {
	o.Description = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GeneralInvestigationAttributesWithoutTimeBounds) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralInvestigationAttributesWithoutTimeBounds) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GeneralInvestigationAttributesWithoutTimeBounds) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GeneralInvestigationAttributesWithoutTimeBounds) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeneralInvestigationAttributesWithoutTimeBounds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeneralInvestigationAttributesWithoutTimeBounds) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string  `json:"description"`
		Tags        []string `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	o.Description = *all.Description
	o.Tags = all.Tags

	return nil
}
