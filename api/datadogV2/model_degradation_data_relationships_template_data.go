// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DegradationDataRelationshipsTemplateData The data object identifying the template the degradation was created from.
type DegradationDataRelationshipsTemplateData struct {
	// The ID of the degradation template.
	Id string `json:"id"`
	// Degradation templates resource type.
	Type PatchDegradationTemplateRequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDegradationDataRelationshipsTemplateData instantiates a new DegradationDataRelationshipsTemplateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDegradationDataRelationshipsTemplateData(id string, typeVar PatchDegradationTemplateRequestDataType) *DegradationDataRelationshipsTemplateData {
	this := DegradationDataRelationshipsTemplateData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewDegradationDataRelationshipsTemplateDataWithDefaults instantiates a new DegradationDataRelationshipsTemplateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDegradationDataRelationshipsTemplateDataWithDefaults() *DegradationDataRelationshipsTemplateData {
	this := DegradationDataRelationshipsTemplateData{}
	var typeVar PatchDegradationTemplateRequestDataType = PATCHDEGRADATIONTEMPLATEREQUESTDATATYPE_DEGRADATION_TEMPLATES
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *DegradationDataRelationshipsTemplateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DegradationDataRelationshipsTemplateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *DegradationDataRelationshipsTemplateData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *DegradationDataRelationshipsTemplateData) GetType() PatchDegradationTemplateRequestDataType {
	if o == nil {
		var ret PatchDegradationTemplateRequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DegradationDataRelationshipsTemplateData) GetTypeOk() (*PatchDegradationTemplateRequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DegradationDataRelationshipsTemplateData) SetType(v PatchDegradationTemplateRequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DegradationDataRelationshipsTemplateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DegradationDataRelationshipsTemplateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                                  `json:"id"`
		Type *PatchDegradationTemplateRequestDataType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
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
