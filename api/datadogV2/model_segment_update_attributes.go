// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentUpdateAttributes Attributes to update on a segment. Fields that are not set are left unchanged.
type SegmentUpdateAttributes struct {
	// Defines who belongs to the segment. It contains one or more blocks, up to 20 in total.
	// Use `combination` to combine named blocks with a boolean formula such as `A AND NOT B`.
	DataQuery *SegmentDataQuery `json:"data_query,omitempty"`
	// Description of the segment.
	Description *string `json:"description,omitempty"`
	// Name of the segment. It must be unique in the organization.
	Name *string `json:"name,omitempty"`
	// Tags of the segment. Replaces the stored tags.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSegmentUpdateAttributes instantiates a new SegmentUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSegmentUpdateAttributes() *SegmentUpdateAttributes {
	this := SegmentUpdateAttributes{}
	return &this
}

// NewSegmentUpdateAttributesWithDefaults instantiates a new SegmentUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSegmentUpdateAttributesWithDefaults() *SegmentUpdateAttributes {
	this := SegmentUpdateAttributes{}
	return &this
}

// GetDataQuery returns the DataQuery field value if set, zero value otherwise.
func (o *SegmentUpdateAttributes) GetDataQuery() SegmentDataQuery {
	if o == nil || o.DataQuery == nil {
		var ret SegmentDataQuery
		return ret
	}
	return *o.DataQuery
}

// GetDataQueryOk returns a tuple with the DataQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentUpdateAttributes) GetDataQueryOk() (*SegmentDataQuery, bool) {
	if o == nil || o.DataQuery == nil {
		return nil, false
	}
	return o.DataQuery, true
}

// HasDataQuery returns a boolean if a field has been set.
func (o *SegmentUpdateAttributes) HasDataQuery() bool {
	return o != nil && o.DataQuery != nil
}

// SetDataQuery gets a reference to the given SegmentDataQuery and assigns it to the DataQuery field.
func (o *SegmentUpdateAttributes) SetDataQuery(v SegmentDataQuery) {
	o.DataQuery = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SegmentUpdateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentUpdateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SegmentUpdateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SegmentUpdateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SegmentUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SegmentUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SegmentUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SegmentUpdateAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentUpdateAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SegmentUpdateAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SegmentUpdateAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SegmentUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DataQuery != nil {
		toSerialize["data_query"] = o.DataQuery
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SegmentUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataQuery   *SegmentDataQuery `json:"data_query,omitempty"`
		Description *string           `json:"description,omitempty"`
		Name        *string           `json:"name,omitempty"`
		Tags        []string          `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_query", "description", "name", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DataQuery != nil && all.DataQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DataQuery = all.DataQuery
	o.Description = all.Description
	o.Name = all.Name
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
