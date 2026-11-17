// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFunnelSuggestionsRequestAttributes Attributes of a feature suggestions request.
type ProductAnalyticsFunnelSuggestionsRequestAttributes struct {
	// Only return suggestions for this RUM application.
	ApplicationId *string `json:"application_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsFunnelSuggestionsRequestAttributes instantiates a new ProductAnalyticsFunnelSuggestionsRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsFunnelSuggestionsRequestAttributes() *ProductAnalyticsFunnelSuggestionsRequestAttributes {
	this := ProductAnalyticsFunnelSuggestionsRequestAttributes{}
	return &this
}

// NewProductAnalyticsFunnelSuggestionsRequestAttributesWithDefaults instantiates a new ProductAnalyticsFunnelSuggestionsRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsFunnelSuggestionsRequestAttributesWithDefaults() *ProductAnalyticsFunnelSuggestionsRequestAttributes {
	this := ProductAnalyticsFunnelSuggestionsRequestAttributes{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelSuggestionsRequestAttributes) GetApplicationId() string {
	if o == nil || o.ApplicationId == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelSuggestionsRequestAttributes) GetApplicationIdOk() (*string, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelSuggestionsRequestAttributes) HasApplicationId() bool {
	return o != nil && o.ApplicationId != nil
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *ProductAnalyticsFunnelSuggestionsRequestAttributes) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsFunnelSuggestionsRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApplicationId != nil {
		toSerialize["application_id"] = o.ApplicationId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsFunnelSuggestionsRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId *string `json:"application_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id"})
	} else {
		return err
	}
	o.ApplicationId = all.ApplicationId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
