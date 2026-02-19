// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumStaticSegmentCreateAttributes Attributes for creating a new static segment.
type RumStaticSegmentCreateAttributes struct {
	// A description of the static segment.
	Description string `json:"description"`
	// The journey query object used to compute the static segment user list.
	JourneyQueryObject RumStaticSegmentJourneyQueryObject `json:"journey_query_object"`
	// The name of the static segment.
	Name string `json:"name"`
	// A list of tags for the static segment.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumStaticSegmentCreateAttributes instantiates a new RumStaticSegmentCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumStaticSegmentCreateAttributes(description string, journeyQueryObject RumStaticSegmentJourneyQueryObject, name string) *RumStaticSegmentCreateAttributes {
	this := RumStaticSegmentCreateAttributes{}
	this.Description = description
	this.JourneyQueryObject = journeyQueryObject
	this.Name = name
	return &this
}

// NewRumStaticSegmentCreateAttributesWithDefaults instantiates a new RumStaticSegmentCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumStaticSegmentCreateAttributesWithDefaults() *RumStaticSegmentCreateAttributes {
	this := RumStaticSegmentCreateAttributes{}
	return &this
}

// GetDescription returns the Description field value.
func (o *RumStaticSegmentCreateAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RumStaticSegmentCreateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *RumStaticSegmentCreateAttributes) SetDescription(v string) {
	o.Description = v
}

// GetJourneyQueryObject returns the JourneyQueryObject field value.
func (o *RumStaticSegmentCreateAttributes) GetJourneyQueryObject() RumStaticSegmentJourneyQueryObject {
	if o == nil {
		var ret RumStaticSegmentJourneyQueryObject
		return ret
	}
	return o.JourneyQueryObject
}

// GetJourneyQueryObjectOk returns a tuple with the JourneyQueryObject field value
// and a boolean to check if the value has been set.
func (o *RumStaticSegmentCreateAttributes) GetJourneyQueryObjectOk() (*RumStaticSegmentJourneyQueryObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JourneyQueryObject, true
}

// SetJourneyQueryObject sets field value.
func (o *RumStaticSegmentCreateAttributes) SetJourneyQueryObject(v RumStaticSegmentJourneyQueryObject) {
	o.JourneyQueryObject = v
}

// GetName returns the Name field value.
func (o *RumStaticSegmentCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RumStaticSegmentCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RumStaticSegmentCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RumStaticSegmentCreateAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumStaticSegmentCreateAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RumStaticSegmentCreateAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *RumStaticSegmentCreateAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumStaticSegmentCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["journey_query_object"] = o.JourneyQueryObject
	toSerialize["name"] = o.Name
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumStaticSegmentCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description        *string                             `json:"description"`
		JourneyQueryObject *RumStaticSegmentJourneyQueryObject `json:"journey_query_object"`
		Name               *string                             `json:"name"`
		Tags               []string                            `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.JourneyQueryObject == nil {
		return fmt.Errorf("required field journey_query_object missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "journey_query_object", "name", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = *all.Description
	if all.JourneyQueryObject.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JourneyQueryObject = *all.JourneyQueryObject
	o.Name = *all.Name
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
