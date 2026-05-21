// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRenderedTemplateDataAttributes Attributes of a rendered template.
type IncidentRenderedTemplateDataAttributes struct {
	// The rendered template content.
	RenderedContent string `json:"rendered_content"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentRenderedTemplateDataAttributes instantiates a new IncidentRenderedTemplateDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentRenderedTemplateDataAttributes(renderedContent string) *IncidentRenderedTemplateDataAttributes {
	this := IncidentRenderedTemplateDataAttributes{}
	this.RenderedContent = renderedContent
	return &this
}

// NewIncidentRenderedTemplateDataAttributesWithDefaults instantiates a new IncidentRenderedTemplateDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentRenderedTemplateDataAttributesWithDefaults() *IncidentRenderedTemplateDataAttributes {
	this := IncidentRenderedTemplateDataAttributes{}
	return &this
}

// GetRenderedContent returns the RenderedContent field value.
func (o *IncidentRenderedTemplateDataAttributes) GetRenderedContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RenderedContent
}

// GetRenderedContentOk returns a tuple with the RenderedContent field value
// and a boolean to check if the value has been set.
func (o *IncidentRenderedTemplateDataAttributes) GetRenderedContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RenderedContent, true
}

// SetRenderedContent sets field value.
func (o *IncidentRenderedTemplateDataAttributes) SetRenderedContent(v string) {
	o.RenderedContent = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentRenderedTemplateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["rendered_content"] = o.RenderedContent

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentRenderedTemplateDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RenderedContent *string `json:"rendered_content"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RenderedContent == nil {
		return fmt.Errorf("required field rendered_content missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rendered_content"})
	} else {
		return err
	}
	o.RenderedContent = *all.RenderedContent

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
