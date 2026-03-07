// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSegmentTemplateInstance A template-based query block within a segment data query.
type RumSegmentTemplateInstance struct {
	// The start of the time range in milliseconds since epoch.
	From *int64 `json:"from,omitempty"`
	// The template parameters as key-value pairs.
	Parameters map[string]string `json:"parameters,omitempty"`
	// The identifier of the template.
	TemplateId string `json:"template_id"`
	// The end of the time range in milliseconds since epoch.
	To *int64 `json:"to,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumSegmentTemplateInstance instantiates a new RumSegmentTemplateInstance object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumSegmentTemplateInstance(templateId string) *RumSegmentTemplateInstance {
	this := RumSegmentTemplateInstance{}
	this.TemplateId = templateId
	return &this
}

// NewRumSegmentTemplateInstanceWithDefaults instantiates a new RumSegmentTemplateInstance object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumSegmentTemplateInstanceWithDefaults() *RumSegmentTemplateInstance {
	this := RumSegmentTemplateInstance{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *RumSegmentTemplateInstance) GetFrom() int64 {
	if o == nil || o.From == nil {
		var ret int64
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSegmentTemplateInstance) GetFromOk() (*int64, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *RumSegmentTemplateInstance) HasFrom() bool {
	return o != nil && o.From != nil
}

// SetFrom gets a reference to the given int64 and assigns it to the From field.
func (o *RumSegmentTemplateInstance) SetFrom(v int64) {
	o.From = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *RumSegmentTemplateInstance) GetParameters() map[string]string {
	if o == nil || o.Parameters == nil {
		var ret map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSegmentTemplateInstance) GetParametersOk() (*map[string]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *RumSegmentTemplateInstance) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *RumSegmentTemplateInstance) SetParameters(v map[string]string) {
	o.Parameters = v
}

// GetTemplateId returns the TemplateId field value.
func (o *RumSegmentTemplateInstance) GetTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *RumSegmentTemplateInstance) GetTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value.
func (o *RumSegmentTemplateInstance) SetTemplateId(v string) {
	o.TemplateId = v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *RumSegmentTemplateInstance) GetTo() int64 {
	if o == nil || o.To == nil {
		var ret int64
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSegmentTemplateInstance) GetToOk() (*int64, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *RumSegmentTemplateInstance) HasTo() bool {
	return o != nil && o.To != nil
}

// SetTo gets a reference to the given int64 and assigns it to the To field.
func (o *RumSegmentTemplateInstance) SetTo(v int64) {
	o.To = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumSegmentTemplateInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	toSerialize["template_id"] = o.TemplateId
	if o.To != nil {
		toSerialize["to"] = o.To
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumSegmentTemplateInstance) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From       *int64            `json:"from,omitempty"`
		Parameters map[string]string `json:"parameters,omitempty"`
		TemplateId *string           `json:"template_id"`
		To         *int64            `json:"to,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TemplateId == nil {
		return fmt.Errorf("required field template_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"from", "parameters", "template_id", "to"})
	} else {
		return err
	}
	o.From = all.From
	o.Parameters = all.Parameters
	o.TemplateId = *all.TemplateId
	o.To = all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
