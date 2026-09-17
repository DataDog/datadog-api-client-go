// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyFunnelRequestData The single JSON:API resource carrying a funnel query. Its attributes hold the time window to
// query and the journey whose step-to-step conversion should be measured.
type ProductAnalyticsJourneyFunnelRequestData struct {
	// Attributes of a journey funnel request.
	Attributes ProductAnalyticsJourneyFunnelRequestAttributes `json:"attributes"`
	// The resource type identifier for a journey funnel request.
	Type ProductAnalyticsJourneyRequestType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneyFunnelRequestData instantiates a new ProductAnalyticsJourneyFunnelRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneyFunnelRequestData(attributes ProductAnalyticsJourneyFunnelRequestAttributes, typeVar ProductAnalyticsJourneyRequestType) *ProductAnalyticsJourneyFunnelRequestData {
	this := ProductAnalyticsJourneyFunnelRequestData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewProductAnalyticsJourneyFunnelRequestDataWithDefaults instantiates a new ProductAnalyticsJourneyFunnelRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneyFunnelRequestDataWithDefaults() *ProductAnalyticsJourneyFunnelRequestData {
	this := ProductAnalyticsJourneyFunnelRequestData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ProductAnalyticsJourneyFunnelRequestData) GetAttributes() ProductAnalyticsJourneyFunnelRequestAttributes {
	if o == nil {
		var ret ProductAnalyticsJourneyFunnelRequestAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelRequestData) GetAttributesOk() (*ProductAnalyticsJourneyFunnelRequestAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ProductAnalyticsJourneyFunnelRequestData) SetAttributes(v ProductAnalyticsJourneyFunnelRequestAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *ProductAnalyticsJourneyFunnelRequestData) GetType() ProductAnalyticsJourneyRequestType {
	if o == nil {
		var ret ProductAnalyticsJourneyRequestType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelRequestData) GetTypeOk() (*ProductAnalyticsJourneyRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ProductAnalyticsJourneyFunnelRequestData) SetType(v ProductAnalyticsJourneyRequestType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneyFunnelRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJourneyFunnelRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ProductAnalyticsJourneyFunnelRequestAttributes `json:"attributes"`
		Type       *ProductAnalyticsJourneyRequestType             `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
