// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityOverrideRequestData Data of the severity override request.
type SeverityOverrideRequestData struct {
	// Attributes of the severity override request.
	Attributes SeverityOverrideRequestDataAttributes `json:"attributes"`
	// Unique identifier of the severity override request. If not provided, an identifier is generated.
	Id *string `json:"id,omitempty"`
	// Relationships of the severity override request.
	Relationships SeverityOverrideRequestDataRelationships `json:"relationships"`
	// Severity override resource type.
	Type SeverityOverrideDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSeverityOverrideRequestData instantiates a new SeverityOverrideRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSeverityOverrideRequestData(attributes SeverityOverrideRequestDataAttributes, relationships SeverityOverrideRequestDataRelationships, typeVar SeverityOverrideDataType) *SeverityOverrideRequestData {
	this := SeverityOverrideRequestData{}
	this.Attributes = attributes
	this.Relationships = relationships
	this.Type = typeVar
	return &this
}

// NewSeverityOverrideRequestDataWithDefaults instantiates a new SeverityOverrideRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSeverityOverrideRequestDataWithDefaults() *SeverityOverrideRequestData {
	this := SeverityOverrideRequestData{}
	var typeVar SeverityOverrideDataType = SEVERITYOVERRIDEDATATYPE_SEVERITY_OVERRIDE
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *SeverityOverrideRequestData) GetAttributes() SeverityOverrideRequestDataAttributes {
	if o == nil {
		var ret SeverityOverrideRequestDataAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SeverityOverrideRequestData) GetAttributesOk() (*SeverityOverrideRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *SeverityOverrideRequestData) SetAttributes(v SeverityOverrideRequestDataAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SeverityOverrideRequestData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeverityOverrideRequestData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SeverityOverrideRequestData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SeverityOverrideRequestData) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value.
func (o *SeverityOverrideRequestData) GetRelationships() SeverityOverrideRequestDataRelationships {
	if o == nil {
		var ret SeverityOverrideRequestDataRelationships
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *SeverityOverrideRequestData) GetRelationshipsOk() (*SeverityOverrideRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value.
func (o *SeverityOverrideRequestData) SetRelationships(v SeverityOverrideRequestDataRelationships) {
	o.Relationships = v
}

// GetType returns the Type field value.
func (o *SeverityOverrideRequestData) GetType() SeverityOverrideDataType {
	if o == nil {
		var ret SeverityOverrideDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SeverityOverrideRequestData) GetTypeOk() (*SeverityOverrideDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SeverityOverrideRequestData) SetType(v SeverityOverrideDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SeverityOverrideRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["relationships"] = o.Relationships
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SeverityOverrideRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *SeverityOverrideRequestDataAttributes    `json:"attributes"`
		Id            *string                                   `json:"id,omitempty"`
		Relationships *SeverityOverrideRequestDataRelationships `json:"relationships"`
		Type          *SeverityOverrideDataType                 `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Relationships == nil {
		return fmt.Errorf("required field relationships missing")
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
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = all.Id
	if all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = *all.Relationships
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
