// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeneralInvestigationAttributesWithTimeBounds Attributes for a general investigation with an explicit time window.
type GeneralInvestigationAttributesWithTimeBounds struct {
	// A free-form description of what to investigate, up to 4,096 characters.
	Description string `json:"description"`
	// The end of the investigation window, in Unix milliseconds.
	EndTime int64 `json:"end_time"`
	// The start of the investigation window, in Unix milliseconds.
	StartTime int64 `json:"start_time"`
	// Tags scoping the investigation.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGeneralInvestigationAttributesWithTimeBounds instantiates a new GeneralInvestigationAttributesWithTimeBounds object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeneralInvestigationAttributesWithTimeBounds(description string, endTime int64, startTime int64) *GeneralInvestigationAttributesWithTimeBounds {
	this := GeneralInvestigationAttributesWithTimeBounds{}
	this.Description = description
	this.EndTime = endTime
	this.StartTime = startTime
	return &this
}

// NewGeneralInvestigationAttributesWithTimeBoundsWithDefaults instantiates a new GeneralInvestigationAttributesWithTimeBounds object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeneralInvestigationAttributesWithTimeBoundsWithDefaults() *GeneralInvestigationAttributesWithTimeBounds {
	this := GeneralInvestigationAttributesWithTimeBounds{}
	return &this
}

// GetDescription returns the Description field value.
func (o *GeneralInvestigationAttributesWithTimeBounds) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GeneralInvestigationAttributesWithTimeBounds) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *GeneralInvestigationAttributesWithTimeBounds) SetDescription(v string) {
	o.Description = v
}

// GetEndTime returns the EndTime field value.
func (o *GeneralInvestigationAttributesWithTimeBounds) GetEndTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *GeneralInvestigationAttributesWithTimeBounds) GetEndTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value.
func (o *GeneralInvestigationAttributesWithTimeBounds) SetEndTime(v int64) {
	o.EndTime = v
}

// GetStartTime returns the StartTime field value.
func (o *GeneralInvestigationAttributesWithTimeBounds) GetStartTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *GeneralInvestigationAttributesWithTimeBounds) GetStartTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value.
func (o *GeneralInvestigationAttributesWithTimeBounds) SetStartTime(v int64) {
	o.StartTime = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GeneralInvestigationAttributesWithTimeBounds) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralInvestigationAttributesWithTimeBounds) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GeneralInvestigationAttributesWithTimeBounds) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GeneralInvestigationAttributesWithTimeBounds) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeneralInvestigationAttributesWithTimeBounds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["end_time"] = o.EndTime
	toSerialize["start_time"] = o.StartTime
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeneralInvestigationAttributesWithTimeBounds) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string  `json:"description"`
		EndTime     *int64   `json:"end_time"`
		StartTime   *int64   `json:"start_time"`
		Tags        []string `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.EndTime == nil {
		return fmt.Errorf("required field end_time missing")
	}
	if all.StartTime == nil {
		return fmt.Errorf("required field start_time missing")
	}
	o.Description = *all.Description
	o.EndTime = *all.EndTime
	o.StartTime = *all.StartTime
	o.Tags = all.Tags

	return nil
}
