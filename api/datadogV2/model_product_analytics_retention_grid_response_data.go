// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionGridResponseData The single JSON:API resource holding a computed retention grid. Its attributes contain the
// return periods forming the columns and the cohorts forming the rows.
type ProductAnalyticsRetentionGridResponseData struct {
	// Attributes of a retention grid response, containing the cohort rows and the period columns.
	Attributes ProductAnalyticsRetentionGridResponseAttributes `json:"attributes"`
	// Unique identifier for this response data object.
	Id string `json:"id"`
	// The resource type identifier for a retention grid response.
	Type ProductAnalyticsRetentionGridResponseType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionGridResponseData instantiates a new ProductAnalyticsRetentionGridResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionGridResponseData(attributes ProductAnalyticsRetentionGridResponseAttributes, id string, typeVar ProductAnalyticsRetentionGridResponseType) *ProductAnalyticsRetentionGridResponseData {
	this := ProductAnalyticsRetentionGridResponseData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewProductAnalyticsRetentionGridResponseDataWithDefaults instantiates a new ProductAnalyticsRetentionGridResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionGridResponseDataWithDefaults() *ProductAnalyticsRetentionGridResponseData {
	this := ProductAnalyticsRetentionGridResponseData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ProductAnalyticsRetentionGridResponseData) GetAttributes() ProductAnalyticsRetentionGridResponseAttributes {
	if o == nil {
		var ret ProductAnalyticsRetentionGridResponseAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridResponseData) GetAttributesOk() (*ProductAnalyticsRetentionGridResponseAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ProductAnalyticsRetentionGridResponseData) SetAttributes(v ProductAnalyticsRetentionGridResponseAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *ProductAnalyticsRetentionGridResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ProductAnalyticsRetentionGridResponseData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *ProductAnalyticsRetentionGridResponseData) GetType() ProductAnalyticsRetentionGridResponseType {
	if o == nil {
		var ret ProductAnalyticsRetentionGridResponseType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridResponseData) GetTypeOk() (*ProductAnalyticsRetentionGridResponseType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ProductAnalyticsRetentionGridResponseData) SetType(v ProductAnalyticsRetentionGridResponseType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionGridResponseData) MarshalJSON() ([]byte, error) {
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
func (o *ProductAnalyticsRetentionGridResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ProductAnalyticsRetentionGridResponseAttributes `json:"attributes"`
		Id         *string                                          `json:"id"`
		Type       *ProductAnalyticsRetentionGridResponseType       `json:"type"`
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
