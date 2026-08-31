// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageQuotaResponseData A usage quota resource.
type UsageQuotaResponseData struct {
	// Attributes of a usage quota.
	Attributes UsageQuotaResponseAttributes `json:"attributes"`
	// An opaque usage quota identifier. Clients must pass this value back verbatim in update and delete requests and must not infer any structure from it.
	Id string `json:"id"`
	// The JSON:API resource type for a usage quota.
	Type UsageQuotaType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageQuotaResponseData instantiates a new UsageQuotaResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageQuotaResponseData(attributes UsageQuotaResponseAttributes, id string, typeVar UsageQuotaType) *UsageQuotaResponseData {
	this := UsageQuotaResponseData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewUsageQuotaResponseDataWithDefaults instantiates a new UsageQuotaResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageQuotaResponseDataWithDefaults() *UsageQuotaResponseData {
	this := UsageQuotaResponseData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *UsageQuotaResponseData) GetAttributes() UsageQuotaResponseAttributes {
	if o == nil {
		var ret UsageQuotaResponseAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *UsageQuotaResponseData) GetAttributesOk() (*UsageQuotaResponseAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *UsageQuotaResponseData) SetAttributes(v UsageQuotaResponseAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *UsageQuotaResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UsageQuotaResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *UsageQuotaResponseData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *UsageQuotaResponseData) GetType() UsageQuotaType {
	if o == nil {
		var ret UsageQuotaType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UsageQuotaResponseData) GetTypeOk() (*UsageQuotaType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *UsageQuotaResponseData) SetType(v UsageQuotaType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageQuotaResponseData) MarshalJSON() ([]byte, error) {
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
func (o *UsageQuotaResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *UsageQuotaResponseAttributes `json:"attributes"`
		Id         *string                       `json:"id"`
		Type       *UsageQuotaType               `json:"type"`
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
