// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipUntaggedFindingsData The data wrapper for an ownership untagged findings response.
type OwnershipUntaggedFindingsData struct {
	// The counts of findings without a team tag by ownership confidence.
	Attributes OwnershipUntaggedFindingsAttributes `json:"attributes"`
	// The identifier of the ownership untagged findings resource.
	Id string `json:"id"`
	// The type of the ownership untagged findings resource. The value should always be `ownership_untagged_findings`.
	Type OwnershipUntaggedFindingsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOwnershipUntaggedFindingsData instantiates a new OwnershipUntaggedFindingsData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOwnershipUntaggedFindingsData(attributes OwnershipUntaggedFindingsAttributes, id string, typeVar OwnershipUntaggedFindingsType) *OwnershipUntaggedFindingsData {
	this := OwnershipUntaggedFindingsData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewOwnershipUntaggedFindingsDataWithDefaults instantiates a new OwnershipUntaggedFindingsData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOwnershipUntaggedFindingsDataWithDefaults() *OwnershipUntaggedFindingsData {
	this := OwnershipUntaggedFindingsData{}
	var typeVar OwnershipUntaggedFindingsType = OWNERSHIPUNTAGGEDFINDINGSTYPE_OWNERSHIP_UNTAGGED_FINDINGS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *OwnershipUntaggedFindingsData) GetAttributes() OwnershipUntaggedFindingsAttributes {
	if o == nil {
		var ret OwnershipUntaggedFindingsAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OwnershipUntaggedFindingsData) GetAttributesOk() (*OwnershipUntaggedFindingsAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *OwnershipUntaggedFindingsData) SetAttributes(v OwnershipUntaggedFindingsAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *OwnershipUntaggedFindingsData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OwnershipUntaggedFindingsData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *OwnershipUntaggedFindingsData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *OwnershipUntaggedFindingsData) GetType() OwnershipUntaggedFindingsType {
	if o == nil {
		var ret OwnershipUntaggedFindingsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OwnershipUntaggedFindingsData) GetTypeOk() (*OwnershipUntaggedFindingsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OwnershipUntaggedFindingsData) SetType(v OwnershipUntaggedFindingsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OwnershipUntaggedFindingsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OwnershipUntaggedFindingsData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *OwnershipUntaggedFindingsAttributes `json:"attributes"`
		Id         *string                              `json:"id"`
		Type       *OwnershipUntaggedFindingsType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
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
