// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomCostsFileUpdateData Custom Costs file to update.
type CustomCostsFileUpdateData struct {
	// Attributes of a Custom Costs file that can be updated.
	Attributes CustomCostsFileUpdateAttributes `json:"attributes"`
	// ID of the Custom Costs file.
	Id string `json:"id"`
	// Type of the Custom Costs file resource.
	Type CustomCostsFileUpdateType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomCostsFileUpdateData instantiates a new CustomCostsFileUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomCostsFileUpdateData(attributes CustomCostsFileUpdateAttributes, id string, typeVar CustomCostsFileUpdateType) *CustomCostsFileUpdateData {
	this := CustomCostsFileUpdateData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewCustomCostsFileUpdateDataWithDefaults instantiates a new CustomCostsFileUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomCostsFileUpdateDataWithDefaults() *CustomCostsFileUpdateData {
	this := CustomCostsFileUpdateData{}
	var typeVar CustomCostsFileUpdateType = CUSTOMCOSTSFILEUPDATETYPE_COST_METADATA
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *CustomCostsFileUpdateData) GetAttributes() CustomCostsFileUpdateAttributes {
	if o == nil {
		var ret CustomCostsFileUpdateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomCostsFileUpdateData) GetAttributesOk() (*CustomCostsFileUpdateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *CustomCostsFileUpdateData) SetAttributes(v CustomCostsFileUpdateAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *CustomCostsFileUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomCostsFileUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CustomCostsFileUpdateData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *CustomCostsFileUpdateData) GetType() CustomCostsFileUpdateType {
	if o == nil {
		var ret CustomCostsFileUpdateType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomCostsFileUpdateData) GetTypeOk() (*CustomCostsFileUpdateType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomCostsFileUpdateData) SetType(v CustomCostsFileUpdateType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomCostsFileUpdateData) MarshalJSON() ([]byte, error) {
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
func (o *CustomCostsFileUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CustomCostsFileUpdateAttributes `json:"attributes"`
		Id         *string                          `json:"id"`
		Type       *CustomCostsFileUpdateType       `json:"type"`
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
