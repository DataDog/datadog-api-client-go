// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSegmentCreateAttributes Attributes for creating a new segment.
type RumSegmentCreateAttributes struct {
	// Query definition for the segment. Contains one or more query blocks and an optional combination formula.
	DataQuery RumSegmentDataQuery `json:"data_query"`
	// A description of the segment.
	Description *string `json:"description,omitempty"`
	// The name of the segment.
	Name string `json:"name"`
	// A list of tags for the segment.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumSegmentCreateAttributes instantiates a new RumSegmentCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumSegmentCreateAttributes(dataQuery RumSegmentDataQuery, name string) *RumSegmentCreateAttributes {
	this := RumSegmentCreateAttributes{}
	this.DataQuery = dataQuery
	this.Name = name
	return &this
}

// NewRumSegmentCreateAttributesWithDefaults instantiates a new RumSegmentCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumSegmentCreateAttributesWithDefaults() *RumSegmentCreateAttributes {
	this := RumSegmentCreateAttributes{}
	return &this
}

// GetDataQuery returns the DataQuery field value.
func (o *RumSegmentCreateAttributes) GetDataQuery() RumSegmentDataQuery {
	if o == nil {
		var ret RumSegmentDataQuery
		return ret
	}
	return o.DataQuery
}

// GetDataQueryOk returns a tuple with the DataQuery field value
// and a boolean to check if the value has been set.
func (o *RumSegmentCreateAttributes) GetDataQueryOk() (*RumSegmentDataQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataQuery, true
}

// SetDataQuery sets field value.
func (o *RumSegmentCreateAttributes) SetDataQuery(v RumSegmentDataQuery) {
	o.DataQuery = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RumSegmentCreateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSegmentCreateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RumSegmentCreateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RumSegmentCreateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value.
func (o *RumSegmentCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RumSegmentCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RumSegmentCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RumSegmentCreateAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSegmentCreateAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RumSegmentCreateAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *RumSegmentCreateAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumSegmentCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_query"] = o.DataQuery
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
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
func (o *RumSegmentCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataQuery   *RumSegmentDataQuery `json:"data_query"`
		Description *string              `json:"description,omitempty"`
		Name        *string              `json:"name"`
		Tags        []string             `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataQuery == nil {
		return fmt.Errorf("required field data_query missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_query", "description", "name", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DataQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DataQuery = *all.DataQuery
	o.Description = all.Description
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
