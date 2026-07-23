// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DegradationTemplateData The data object for a degradation template.
type DegradationTemplateData struct {
	// The attributes of a degradation template.
	Attributes *DegradationTemplateDataAttributes `json:"attributes,omitempty"`
	// The ID of the degradation template.
	Id *string `json:"id,omitempty"`
	// The relationships of a degradation template.
	Relationships *DegradationTemplateDataRelationships `json:"relationships,omitempty"`
	// Degradation templates resource type.
	Type PatchDegradationTemplateRequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDegradationTemplateData instantiates a new DegradationTemplateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDegradationTemplateData(typeVar PatchDegradationTemplateRequestDataType) *DegradationTemplateData {
	this := DegradationTemplateData{}
	this.Type = typeVar
	return &this
}

// NewDegradationTemplateDataWithDefaults instantiates a new DegradationTemplateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDegradationTemplateDataWithDefaults() *DegradationTemplateData {
	this := DegradationTemplateData{}
	var typeVar PatchDegradationTemplateRequestDataType = PATCHDEGRADATIONTEMPLATEREQUESTDATATYPE_DEGRADATION_TEMPLATES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DegradationTemplateData) GetAttributes() DegradationTemplateDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret DegradationTemplateDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationTemplateData) GetAttributesOk() (*DegradationTemplateDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DegradationTemplateData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given DegradationTemplateDataAttributes and assigns it to the Attributes field.
func (o *DegradationTemplateData) SetAttributes(v DegradationTemplateDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DegradationTemplateData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationTemplateData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DegradationTemplateData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DegradationTemplateData) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *DegradationTemplateData) GetRelationships() DegradationTemplateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret DegradationTemplateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationTemplateData) GetRelationshipsOk() (*DegradationTemplateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *DegradationTemplateData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given DegradationTemplateDataRelationships and assigns it to the Relationships field.
func (o *DegradationTemplateData) SetRelationships(v DegradationTemplateDataRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *DegradationTemplateData) GetType() PatchDegradationTemplateRequestDataType {
	if o == nil {
		var ret PatchDegradationTemplateRequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DegradationTemplateData) GetTypeOk() (*PatchDegradationTemplateRequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DegradationTemplateData) SetType(v PatchDegradationTemplateRequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DegradationTemplateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DegradationTemplateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *DegradationTemplateDataAttributes       `json:"attributes,omitempty"`
		Id            *string                                  `json:"id,omitempty"`
		Relationships *DegradationTemplateDataRelationships    `json:"relationships,omitempty"`
		Type          *PatchDegradationTemplateRequestDataType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
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
