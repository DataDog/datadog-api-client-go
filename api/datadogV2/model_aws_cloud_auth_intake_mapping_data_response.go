// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSCloudAuthIntakeMappingDataResponse Data for AWS cloud authentication intake mapping response
type AWSCloudAuthIntakeMappingDataResponse struct {
	// Unique identifier for the intake mapping
	Id string `json:"id"`
	// Type identifier for AWS cloud authentication intake mapping
	Type AWSCloudAuthIntakeMappingType `json:"type"`
	// Attributes for AWS cloud authentication intake mapping response
	Attributes AWSCloudAuthIntakeMappingAttributesResponse `json:"attributes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSCloudAuthIntakeMappingDataResponse instantiates a new AWSCloudAuthIntakeMappingDataResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSCloudAuthIntakeMappingDataResponse(id string, typeVar AWSCloudAuthIntakeMappingType, attributes AWSCloudAuthIntakeMappingAttributesResponse) *AWSCloudAuthIntakeMappingDataResponse {
	this := AWSCloudAuthIntakeMappingDataResponse{}
	this.Id = id
	this.Type = typeVar
	this.Attributes = attributes
	return &this
}

// NewAWSCloudAuthIntakeMappingDataResponseWithDefaults instantiates a new AWSCloudAuthIntakeMappingDataResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSCloudAuthIntakeMappingDataResponseWithDefaults() *AWSCloudAuthIntakeMappingDataResponse {
	this := AWSCloudAuthIntakeMappingDataResponse{}
	return &this
}

// GetId returns the Id field value.
func (o *AWSCloudAuthIntakeMappingDataResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AWSCloudAuthIntakeMappingDataResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *AWSCloudAuthIntakeMappingDataResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *AWSCloudAuthIntakeMappingDataResponse) GetType() AWSCloudAuthIntakeMappingType {
	if o == nil {
		var ret AWSCloudAuthIntakeMappingType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AWSCloudAuthIntakeMappingDataResponse) GetTypeOk() (*AWSCloudAuthIntakeMappingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AWSCloudAuthIntakeMappingDataResponse) SetType(v AWSCloudAuthIntakeMappingType) {
	o.Type = v
}

// GetAttributes returns the Attributes field value.
func (o *AWSCloudAuthIntakeMappingDataResponse) GetAttributes() AWSCloudAuthIntakeMappingAttributesResponse {
	if o == nil {
		var ret AWSCloudAuthIntakeMappingAttributesResponse
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AWSCloudAuthIntakeMappingDataResponse) GetAttributesOk() (*AWSCloudAuthIntakeMappingAttributesResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *AWSCloudAuthIntakeMappingDataResponse) SetAttributes(v AWSCloudAuthIntakeMappingAttributesResponse) {
	o.Attributes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSCloudAuthIntakeMappingDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSCloudAuthIntakeMappingDataResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id         *string                                      `json:"id"`
		Type       *AWSCloudAuthIntakeMappingType               `json:"type"`
		Attributes *AWSCloudAuthIntakeMappingAttributesResponse `json:"attributes"`
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
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type", "attributes"})
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
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
