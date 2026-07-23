// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateDegradationTemplateRequestData The data object for creating a degradation template.
type CreateDegradationTemplateRequestData struct {
	// The attributes for creating a degradation template.
	Attributes *CreateDegradationTemplateRequestDataAttributes `json:"attributes,omitempty"`
	// Degradation templates resource type.
	Type PatchDegradationTemplateRequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateDegradationTemplateRequestData instantiates a new CreateDegradationTemplateRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateDegradationTemplateRequestData(typeVar PatchDegradationTemplateRequestDataType) *CreateDegradationTemplateRequestData {
	this := CreateDegradationTemplateRequestData{}
	this.Type = typeVar
	return &this
}

// NewCreateDegradationTemplateRequestDataWithDefaults instantiates a new CreateDegradationTemplateRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateDegradationTemplateRequestDataWithDefaults() *CreateDegradationTemplateRequestData {
	this := CreateDegradationTemplateRequestData{}
	var typeVar PatchDegradationTemplateRequestDataType = PATCHDEGRADATIONTEMPLATEREQUESTDATATYPE_DEGRADATION_TEMPLATES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CreateDegradationTemplateRequestData) GetAttributes() CreateDegradationTemplateRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret CreateDegradationTemplateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDegradationTemplateRequestData) GetAttributesOk() (*CreateDegradationTemplateRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CreateDegradationTemplateRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given CreateDegradationTemplateRequestDataAttributes and assigns it to the Attributes field.
func (o *CreateDegradationTemplateRequestData) SetAttributes(v CreateDegradationTemplateRequestDataAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value.
func (o *CreateDegradationTemplateRequestData) GetType() PatchDegradationTemplateRequestDataType {
	if o == nil {
		var ret PatchDegradationTemplateRequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateDegradationTemplateRequestData) GetTypeOk() (*PatchDegradationTemplateRequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateDegradationTemplateRequestData) SetType(v PatchDegradationTemplateRequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateDegradationTemplateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateDegradationTemplateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CreateDegradationTemplateRequestDataAttributes `json:"attributes,omitempty"`
		Type       *PatchDegradationTemplateRequestDataType        `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
