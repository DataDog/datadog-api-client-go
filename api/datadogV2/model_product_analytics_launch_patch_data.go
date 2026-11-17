// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchPatchData Data of the launch update.
type ProductAnalyticsLaunchPatchData struct {
	// Attributes to update on a launch. Fields that are not set are left unchanged, and `null` clears a nullable
	// field.
	// Setting a non-empty `feature_flag_id` clears the stored `feature_flag_key` when `feature_flag_key` is not set.
	Attributes ProductAnalyticsLaunchPatchAttributes `json:"attributes"`
	// UUID of the launch. It must match the ID in the path.
	Id string `json:"id"`
	// Launch resource type.
	Type ProductAnalyticsLaunchType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsLaunchPatchData instantiates a new ProductAnalyticsLaunchPatchData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsLaunchPatchData(attributes ProductAnalyticsLaunchPatchAttributes, id string, typeVar ProductAnalyticsLaunchType) *ProductAnalyticsLaunchPatchData {
	this := ProductAnalyticsLaunchPatchData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewProductAnalyticsLaunchPatchDataWithDefaults instantiates a new ProductAnalyticsLaunchPatchData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsLaunchPatchDataWithDefaults() *ProductAnalyticsLaunchPatchData {
	this := ProductAnalyticsLaunchPatchData{}
	var typeVar ProductAnalyticsLaunchType = PRODUCTANALYTICSLAUNCHTYPE_LIFTOFF
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ProductAnalyticsLaunchPatchData) GetAttributes() ProductAnalyticsLaunchPatchAttributes {
	if o == nil {
		var ret ProductAnalyticsLaunchPatchAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchPatchData) GetAttributesOk() (*ProductAnalyticsLaunchPatchAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ProductAnalyticsLaunchPatchData) SetAttributes(v ProductAnalyticsLaunchPatchAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *ProductAnalyticsLaunchPatchData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchPatchData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ProductAnalyticsLaunchPatchData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *ProductAnalyticsLaunchPatchData) GetType() ProductAnalyticsLaunchType {
	if o == nil {
		var ret ProductAnalyticsLaunchType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchPatchData) GetTypeOk() (*ProductAnalyticsLaunchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ProductAnalyticsLaunchPatchData) SetType(v ProductAnalyticsLaunchType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsLaunchPatchData) MarshalJSON() ([]byte, error) {
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
func (o *ProductAnalyticsLaunchPatchData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ProductAnalyticsLaunchPatchAttributes `json:"attributes"`
		Id         *string                                `json:"id"`
		Type       *ProductAnalyticsLaunchType            `json:"type"`
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
