// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentCreatePageAttributes Attributes for creating a page from an incident.
type IncidentCreatePageAttributes struct {
	// Description of the page.
	Description datadog.NullableString `json:"description,omitempty"`
	// List of services associated with the page.
	Services datadog.NullableList[string] `json:"services,omitempty"`
	// List of tags for the page.
	Tags datadog.NullableList[string] `json:"tags,omitempty"`
	// Target for creating a page from an incident.
	Target IncidentPageTarget `json:"target"`
	// Title of the page.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentCreatePageAttributes instantiates a new IncidentCreatePageAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentCreatePageAttributes(target IncidentPageTarget, title string) *IncidentCreatePageAttributes {
	this := IncidentCreatePageAttributes{}
	this.Target = target
	this.Title = title
	return &this
}

// NewIncidentCreatePageAttributesWithDefaults instantiates a new IncidentCreatePageAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentCreatePageAttributesWithDefaults() *IncidentCreatePageAttributes {
	this := IncidentCreatePageAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentCreatePageAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentCreatePageAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *IncidentCreatePageAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *IncidentCreatePageAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *IncidentCreatePageAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *IncidentCreatePageAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetServices returns the Services field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentCreatePageAttributes) GetServices() []string {
	if o == nil || o.Services.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Services.Get()
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentCreatePageAttributes) GetServicesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Services.Get(), o.Services.IsSet()
}

// HasServices returns a boolean if a field has been set.
func (o *IncidentCreatePageAttributes) HasServices() bool {
	return o != nil && o.Services.IsSet()
}

// SetServices gets a reference to the given datadog.NullableList[string] and assigns it to the Services field.
func (o *IncidentCreatePageAttributes) SetServices(v []string) {
	o.Services.Set(&v)
}

// SetServicesNil sets the value for Services to be an explicit nil.
func (o *IncidentCreatePageAttributes) SetServicesNil() {
	o.Services.Set(nil)
}

// UnsetServices ensures that no value is present for Services, not even an explicit nil.
func (o *IncidentCreatePageAttributes) UnsetServices() {
	o.Services.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentCreatePageAttributes) GetTags() []string {
	if o == nil || o.Tags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentCreatePageAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *IncidentCreatePageAttributes) HasTags() bool {
	return o != nil && o.Tags.IsSet()
}

// SetTags gets a reference to the given datadog.NullableList[string] and assigns it to the Tags field.
func (o *IncidentCreatePageAttributes) SetTags(v []string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil.
func (o *IncidentCreatePageAttributes) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil.
func (o *IncidentCreatePageAttributes) UnsetTags() {
	o.Tags.Unset()
}

// GetTarget returns the Target field value.
func (o *IncidentCreatePageAttributes) GetTarget() IncidentPageTarget {
	if o == nil {
		var ret IncidentPageTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *IncidentCreatePageAttributes) GetTargetOk() (*IncidentPageTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *IncidentCreatePageAttributes) SetTarget(v IncidentPageTarget) {
	o.Target = v
}

// GetTitle returns the Title field value.
func (o *IncidentCreatePageAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *IncidentCreatePageAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *IncidentCreatePageAttributes) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentCreatePageAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Services.IsSet() {
		toSerialize["services"] = o.Services.Get()
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	toSerialize["target"] = o.Target
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentCreatePageAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description datadog.NullableString       `json:"description,omitempty"`
		Services    datadog.NullableList[string] `json:"services,omitempty"`
		Tags        datadog.NullableList[string] `json:"tags,omitempty"`
		Target      *IncidentPageTarget          `json:"target"`
		Title       *string                      `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "services", "tags", "target", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.Services = all.Services
	o.Tags = all.Tags
	if all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = *all.Target
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
