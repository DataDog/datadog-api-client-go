// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchDegradationTemplateRequestDataAttributes The supported attributes for updating a degradation template.
type PatchDegradationTemplateRequestDataAttributes struct {
	// The components affected by a degradation created from this template.
	ComponentsAffected []PatchDegradationTemplateRequestDataAttributesComponentsAffectedItems `json:"components_affected,omitempty"`
	// The title used for a degradation created from this template.
	DegradationTitle *string `json:"degradation_title,omitempty"`
	// The name of the degradation template.
	Name *string `json:"name,omitempty"`
	// The pre-filled updates for a degradation created from this template.
	Updates []PatchDegradationTemplateRequestDataAttributesUpdatesItems `json:"updates,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPatchDegradationTemplateRequestDataAttributes instantiates a new PatchDegradationTemplateRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPatchDegradationTemplateRequestDataAttributes() *PatchDegradationTemplateRequestDataAttributes {
	this := PatchDegradationTemplateRequestDataAttributes{}
	return &this
}

// NewPatchDegradationTemplateRequestDataAttributesWithDefaults instantiates a new PatchDegradationTemplateRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPatchDegradationTemplateRequestDataAttributesWithDefaults() *PatchDegradationTemplateRequestDataAttributes {
	this := PatchDegradationTemplateRequestDataAttributes{}
	return &this
}

// GetComponentsAffected returns the ComponentsAffected field value if set, zero value otherwise.
func (o *PatchDegradationTemplateRequestDataAttributes) GetComponentsAffected() []PatchDegradationTemplateRequestDataAttributesComponentsAffectedItems {
	if o == nil || o.ComponentsAffected == nil {
		var ret []PatchDegradationTemplateRequestDataAttributesComponentsAffectedItems
		return ret
	}
	return o.ComponentsAffected
}

// GetComponentsAffectedOk returns a tuple with the ComponentsAffected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchDegradationTemplateRequestDataAttributes) GetComponentsAffectedOk() (*[]PatchDegradationTemplateRequestDataAttributesComponentsAffectedItems, bool) {
	if o == nil || o.ComponentsAffected == nil {
		return nil, false
	}
	return &o.ComponentsAffected, true
}

// HasComponentsAffected returns a boolean if a field has been set.
func (o *PatchDegradationTemplateRequestDataAttributes) HasComponentsAffected() bool {
	return o != nil && o.ComponentsAffected != nil
}

// SetComponentsAffected gets a reference to the given []PatchDegradationTemplateRequestDataAttributesComponentsAffectedItems and assigns it to the ComponentsAffected field.
func (o *PatchDegradationTemplateRequestDataAttributes) SetComponentsAffected(v []PatchDegradationTemplateRequestDataAttributesComponentsAffectedItems) {
	o.ComponentsAffected = v
}

// GetDegradationTitle returns the DegradationTitle field value if set, zero value otherwise.
func (o *PatchDegradationTemplateRequestDataAttributes) GetDegradationTitle() string {
	if o == nil || o.DegradationTitle == nil {
		var ret string
		return ret
	}
	return *o.DegradationTitle
}

// GetDegradationTitleOk returns a tuple with the DegradationTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchDegradationTemplateRequestDataAttributes) GetDegradationTitleOk() (*string, bool) {
	if o == nil || o.DegradationTitle == nil {
		return nil, false
	}
	return o.DegradationTitle, true
}

// HasDegradationTitle returns a boolean if a field has been set.
func (o *PatchDegradationTemplateRequestDataAttributes) HasDegradationTitle() bool {
	return o != nil && o.DegradationTitle != nil
}

// SetDegradationTitle gets a reference to the given string and assigns it to the DegradationTitle field.
func (o *PatchDegradationTemplateRequestDataAttributes) SetDegradationTitle(v string) {
	o.DegradationTitle = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchDegradationTemplateRequestDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchDegradationTemplateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchDegradationTemplateRequestDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchDegradationTemplateRequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetUpdates returns the Updates field value if set, zero value otherwise.
func (o *PatchDegradationTemplateRequestDataAttributes) GetUpdates() []PatchDegradationTemplateRequestDataAttributesUpdatesItems {
	if o == nil || o.Updates == nil {
		var ret []PatchDegradationTemplateRequestDataAttributesUpdatesItems
		return ret
	}
	return o.Updates
}

// GetUpdatesOk returns a tuple with the Updates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchDegradationTemplateRequestDataAttributes) GetUpdatesOk() (*[]PatchDegradationTemplateRequestDataAttributesUpdatesItems, bool) {
	if o == nil || o.Updates == nil {
		return nil, false
	}
	return &o.Updates, true
}

// HasUpdates returns a boolean if a field has been set.
func (o *PatchDegradationTemplateRequestDataAttributes) HasUpdates() bool {
	return o != nil && o.Updates != nil
}

// SetUpdates gets a reference to the given []PatchDegradationTemplateRequestDataAttributesUpdatesItems and assigns it to the Updates field.
func (o *PatchDegradationTemplateRequestDataAttributes) SetUpdates(v []PatchDegradationTemplateRequestDataAttributesUpdatesItems) {
	o.Updates = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PatchDegradationTemplateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ComponentsAffected != nil {
		toSerialize["components_affected"] = o.ComponentsAffected
	}
	if o.DegradationTitle != nil {
		toSerialize["degradation_title"] = o.DegradationTitle
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Updates != nil {
		toSerialize["updates"] = o.Updates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PatchDegradationTemplateRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentsAffected []PatchDegradationTemplateRequestDataAttributesComponentsAffectedItems `json:"components_affected,omitempty"`
		DegradationTitle   *string                                                                `json:"degradation_title,omitempty"`
		Name               *string                                                                `json:"name,omitempty"`
		Updates            []PatchDegradationTemplateRequestDataAttributesUpdatesItems            `json:"updates,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"components_affected", "degradation_title", "name", "updates"})
	} else {
		return err
	}
	o.ComponentsAffected = all.ComponentsAffected
	o.DegradationTitle = all.DegradationTitle
	o.Name = all.Name
	o.Updates = all.Updates

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
