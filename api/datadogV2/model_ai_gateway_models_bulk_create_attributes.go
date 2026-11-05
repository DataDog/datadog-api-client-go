// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelsBulkCreateAttributes Attributes used to create several model catalog entries.
type AiGatewayModelsBulkCreateAttributes struct {
	//
	Models []AiGatewayModelCreateAttributes `json:"models"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelsBulkCreateAttributes instantiates a new AiGatewayModelsBulkCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelsBulkCreateAttributes(models []AiGatewayModelCreateAttributes) *AiGatewayModelsBulkCreateAttributes {
	this := AiGatewayModelsBulkCreateAttributes{}
	this.Models = models
	return &this
}

// NewAiGatewayModelsBulkCreateAttributesWithDefaults instantiates a new AiGatewayModelsBulkCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelsBulkCreateAttributesWithDefaults() *AiGatewayModelsBulkCreateAttributes {
	this := AiGatewayModelsBulkCreateAttributes{}
	return &this
}

// GetModels returns the Models field value.
func (o *AiGatewayModelsBulkCreateAttributes) GetModels() []AiGatewayModelCreateAttributes {
	if o == nil {
		var ret []AiGatewayModelCreateAttributes
		return ret
	}
	return o.Models
}

// GetModelsOk returns a tuple with the Models field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelsBulkCreateAttributes) GetModelsOk() (*[]AiGatewayModelCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Models, true
}

// SetModels sets field value.
func (o *AiGatewayModelsBulkCreateAttributes) SetModels(v []AiGatewayModelCreateAttributes) {
	o.Models = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelsBulkCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["models"] = o.Models

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelsBulkCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Models *[]AiGatewayModelCreateAttributes `json:"models"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Models == nil {
		return fmt.Errorf("required field models missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"models"})
	} else {
		return err
	}
	o.Models = *all.Models

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
