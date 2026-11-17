// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentTemplateBlock Selects users or accounts with a segment template.
type SegmentTemplateBlock struct {
	// Start of the time range, in epoch milliseconds.
	From *string `json:"from,omitempty"`
	// Name of the block, used in `combination`.
	Name *string `json:"name,omitempty"`
	// Template parameter values, keyed by parameter name.
	Parameters map[string]interface{} `json:"parameters"`
	// ID of the segment template, such as `time-spent-v1`.
	TemplateId string `json:"template_id"`
	// End of the time range, in epoch milliseconds.
	To *string `json:"to,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSegmentTemplateBlock instantiates a new SegmentTemplateBlock object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSegmentTemplateBlock(parameters map[string]interface{}, templateId string) *SegmentTemplateBlock {
	this := SegmentTemplateBlock{}
	this.Parameters = parameters
	this.TemplateId = templateId
	return &this
}

// NewSegmentTemplateBlockWithDefaults instantiates a new SegmentTemplateBlock object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSegmentTemplateBlockWithDefaults() *SegmentTemplateBlock {
	this := SegmentTemplateBlock{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *SegmentTemplateBlock) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentTemplateBlock) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *SegmentTemplateBlock) HasFrom() bool {
	return o != nil && o.From != nil
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *SegmentTemplateBlock) SetFrom(v string) {
	o.From = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SegmentTemplateBlock) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentTemplateBlock) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SegmentTemplateBlock) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SegmentTemplateBlock) SetName(v string) {
	o.Name = &v
}

// GetParameters returns the Parameters field value.
func (o *SegmentTemplateBlock) GetParameters() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *SegmentTemplateBlock) GetParametersOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value.
func (o *SegmentTemplateBlock) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

// GetTemplateId returns the TemplateId field value.
func (o *SegmentTemplateBlock) GetTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *SegmentTemplateBlock) GetTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value.
func (o *SegmentTemplateBlock) SetTemplateId(v string) {
	o.TemplateId = v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *SegmentTemplateBlock) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentTemplateBlock) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *SegmentTemplateBlock) HasTo() bool {
	return o != nil && o.To != nil
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *SegmentTemplateBlock) SetTo(v string) {
	o.To = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SegmentTemplateBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["parameters"] = o.Parameters
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
func (o *SegmentTemplateBlock) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From       *string                 `json:"from,omitempty"`
		Name       *string                 `json:"name,omitempty"`
		Parameters *map[string]interface{} `json:"parameters"`
		TemplateId *string                 `json:"template_id"`
		To         *string                 `json:"to,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Parameters == nil {
		return fmt.Errorf("required field parameters missing")
	}
	if all.TemplateId == nil {
		return fmt.Errorf("required field template_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"from", "name", "parameters", "template_id", "to"})
	} else {
		return err
	}
	o.From = all.From
	o.Name = all.Name
	o.Parameters = *all.Parameters
	o.TemplateId = *all.TemplateId
	o.To = all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
