// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyScalarResponseData The single JSON:API resource holding journey scalar results. Its attributes contain one value
// per group, suitable for a query value or top list widget.
type ProductAnalyticsJourneyScalarResponseData struct {
	// Attributes of a scalar analytics response, containing the result columns.
	Attributes ProductAnalyticsScalarResponseAttributes `json:"attributes"`
	// Identifier of this result.
	Id string `json:"id"`
	// The resource type identifier for a journey scalar response.
	Type ProductAnalyticsJourneyScalarResponseType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneyScalarResponseData instantiates a new ProductAnalyticsJourneyScalarResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneyScalarResponseData(attributes ProductAnalyticsScalarResponseAttributes, id string, typeVar ProductAnalyticsJourneyScalarResponseType) *ProductAnalyticsJourneyScalarResponseData {
	this := ProductAnalyticsJourneyScalarResponseData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewProductAnalyticsJourneyScalarResponseDataWithDefaults instantiates a new ProductAnalyticsJourneyScalarResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneyScalarResponseDataWithDefaults() *ProductAnalyticsJourneyScalarResponseData {
	this := ProductAnalyticsJourneyScalarResponseData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ProductAnalyticsJourneyScalarResponseData) GetAttributes() ProductAnalyticsScalarResponseAttributes {
	if o == nil {
		var ret ProductAnalyticsScalarResponseAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyScalarResponseData) GetAttributesOk() (*ProductAnalyticsScalarResponseAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ProductAnalyticsJourneyScalarResponseData) SetAttributes(v ProductAnalyticsScalarResponseAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *ProductAnalyticsJourneyScalarResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyScalarResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ProductAnalyticsJourneyScalarResponseData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *ProductAnalyticsJourneyScalarResponseData) GetType() ProductAnalyticsJourneyScalarResponseType {
	if o == nil {
		var ret ProductAnalyticsJourneyScalarResponseType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyScalarResponseData) GetTypeOk() (*ProductAnalyticsJourneyScalarResponseType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ProductAnalyticsJourneyScalarResponseData) SetType(v ProductAnalyticsJourneyScalarResponseType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneyScalarResponseData) MarshalJSON() ([]byte, error) {
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
func (o *ProductAnalyticsJourneyScalarResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ProductAnalyticsScalarResponseAttributes  `json:"attributes"`
		Id         *string                                    `json:"id"`
		Type       *ProductAnalyticsJourneyScalarResponseType `json:"type"`
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
