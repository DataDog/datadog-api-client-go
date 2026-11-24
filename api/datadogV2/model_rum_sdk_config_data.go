// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMSdkConfigData The RUM SDK configuration data object.
type RUMSdkConfigData struct {
	// Attributes of the RUM SDK configuration.
	Attributes RUMSdkConfigAttributes `json:"attributes"`
	// The unique identifier of the RUM SDK configuration.
	Id string `json:"id"`
	// Metadata associated with a RUM SDK configuration.
	Meta *RUMSdkConfigMeta `json:"meta,omitempty"`
	// The type of the resource. The value should always be `rum_sdk_config`.
	Type RUMSdkConfigType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMSdkConfigData instantiates a new RUMSdkConfigData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMSdkConfigData(attributes RUMSdkConfigAttributes, id string, typeVar RUMSdkConfigType) *RUMSdkConfigData {
	this := RUMSdkConfigData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewRUMSdkConfigDataWithDefaults instantiates a new RUMSdkConfigData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMSdkConfigDataWithDefaults() *RUMSdkConfigData {
	this := RUMSdkConfigData{}
	var typeVar RUMSdkConfigType = RUMSDKCONFIGTYPE_RUM_SDK_CONFIG
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *RUMSdkConfigData) GetAttributes() RUMSdkConfigAttributes {
	if o == nil {
		var ret RUMSdkConfigAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigData) GetAttributesOk() (*RUMSdkConfigAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *RUMSdkConfigData) SetAttributes(v RUMSdkConfigAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *RUMSdkConfigData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *RUMSdkConfigData) SetId(v string) {
	o.Id = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *RUMSdkConfigData) GetMeta() RUMSdkConfigMeta {
	if o == nil || o.Meta == nil {
		var ret RUMSdkConfigMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigData) GetMetaOk() (*RUMSdkConfigMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *RUMSdkConfigData) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given RUMSdkConfigMeta and assigns it to the Meta field.
func (o *RUMSdkConfigData) SetMeta(v RUMSdkConfigMeta) {
	o.Meta = &v
}

// GetType returns the Type field value.
func (o *RUMSdkConfigData) GetType() RUMSdkConfigType {
	if o == nil {
		var ret RUMSdkConfigType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigData) GetTypeOk() (*RUMSdkConfigType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RUMSdkConfigData) SetType(v RUMSdkConfigType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMSdkConfigData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMSdkConfigData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *RUMSdkConfigAttributes `json:"attributes"`
		Id         *string                 `json:"id"`
		Meta       *RUMSdkConfigMeta       `json:"meta,omitempty"`
		Type       *RUMSdkConfigType       `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "meta", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta
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
