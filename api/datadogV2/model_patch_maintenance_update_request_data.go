// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchMaintenanceUpdateRequestData The data object for editing a maintenance update.
type PatchMaintenanceUpdateRequestData struct {
	// Attributes for editing a maintenance update.
	Attributes *PatchMaintenanceUpdateRequestDataAttributes `json:"attributes,omitempty"`
	// The ID of the maintenance update to edit. Must match the `update_id` path parameter.
	Id string `json:"id"`
	// Maintenance updates resource type.
	Type PatchMaintenanceUpdateRequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPatchMaintenanceUpdateRequestData instantiates a new PatchMaintenanceUpdateRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPatchMaintenanceUpdateRequestData(id string, typeVar PatchMaintenanceUpdateRequestDataType) *PatchMaintenanceUpdateRequestData {
	this := PatchMaintenanceUpdateRequestData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewPatchMaintenanceUpdateRequestDataWithDefaults instantiates a new PatchMaintenanceUpdateRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPatchMaintenanceUpdateRequestDataWithDefaults() *PatchMaintenanceUpdateRequestData {
	this := PatchMaintenanceUpdateRequestData{}
	var typeVar PatchMaintenanceUpdateRequestDataType = PATCHMAINTENANCEUPDATEREQUESTDATATYPE_MAINTENANCE_UPDATES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PatchMaintenanceUpdateRequestData) GetAttributes() PatchMaintenanceUpdateRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret PatchMaintenanceUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchMaintenanceUpdateRequestData) GetAttributesOk() (*PatchMaintenanceUpdateRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PatchMaintenanceUpdateRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given PatchMaintenanceUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *PatchMaintenanceUpdateRequestData) SetAttributes(v PatchMaintenanceUpdateRequestDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *PatchMaintenanceUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchMaintenanceUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *PatchMaintenanceUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *PatchMaintenanceUpdateRequestData) GetType() PatchMaintenanceUpdateRequestDataType {
	if o == nil {
		var ret PatchMaintenanceUpdateRequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PatchMaintenanceUpdateRequestData) GetTypeOk() (*PatchMaintenanceUpdateRequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *PatchMaintenanceUpdateRequestData) SetType(v PatchMaintenanceUpdateRequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PatchMaintenanceUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PatchMaintenanceUpdateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *PatchMaintenanceUpdateRequestDataAttributes `json:"attributes,omitempty"`
		Id         *string                                      `json:"id"`
		Type       *PatchMaintenanceUpdateRequestDataType       `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
