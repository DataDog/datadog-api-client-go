// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArchiveSearchRehydrateRequestData Archive Search object to rehydrate.
type ArchiveSearchRehydrateRequestData struct {
	// Attributes accepted when rehydrating an Archive Search.
	Attributes ArchiveSearchRehydrateRequestAttributes `json:"attributes"`
	// Archive Search resource type.
	Type ArchiveSearchType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewArchiveSearchRehydrateRequestData instantiates a new ArchiveSearchRehydrateRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewArchiveSearchRehydrateRequestData(attributes ArchiveSearchRehydrateRequestAttributes, typeVar ArchiveSearchType) *ArchiveSearchRehydrateRequestData {
	this := ArchiveSearchRehydrateRequestData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewArchiveSearchRehydrateRequestDataWithDefaults instantiates a new ArchiveSearchRehydrateRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewArchiveSearchRehydrateRequestDataWithDefaults() *ArchiveSearchRehydrateRequestData {
	this := ArchiveSearchRehydrateRequestData{}
	var typeVar ArchiveSearchType = ARCHIVESEARCHTYPE_ARCHIVE_SEARCH
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ArchiveSearchRehydrateRequestData) GetAttributes() ArchiveSearchRehydrateRequestAttributes {
	if o == nil {
		var ret ArchiveSearchRehydrateRequestAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchRehydrateRequestData) GetAttributesOk() (*ArchiveSearchRehydrateRequestAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ArchiveSearchRehydrateRequestData) SetAttributes(v ArchiveSearchRehydrateRequestAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *ArchiveSearchRehydrateRequestData) GetType() ArchiveSearchType {
	if o == nil {
		var ret ArchiveSearchType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArchiveSearchRehydrateRequestData) GetTypeOk() (*ArchiveSearchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ArchiveSearchRehydrateRequestData) SetType(v ArchiveSearchType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ArchiveSearchRehydrateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ArchiveSearchRehydrateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ArchiveSearchRehydrateRequestAttributes `json:"attributes"`
		Type       *ArchiveSearchType                       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
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
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
