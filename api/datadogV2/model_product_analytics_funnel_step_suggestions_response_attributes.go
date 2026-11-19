// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFunnelStepSuggestionsResponseAttributes Attributes of a step suggestions response.
type ProductAnalyticsFunnelStepSuggestionsResponseAttributes struct {
	// Event queries of the suggested next steps, most relevant first.
	Steps []string `json:"steps,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsFunnelStepSuggestionsResponseAttributes instantiates a new ProductAnalyticsFunnelStepSuggestionsResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsFunnelStepSuggestionsResponseAttributes() *ProductAnalyticsFunnelStepSuggestionsResponseAttributes {
	this := ProductAnalyticsFunnelStepSuggestionsResponseAttributes{}
	return &this
}

// NewProductAnalyticsFunnelStepSuggestionsResponseAttributesWithDefaults instantiates a new ProductAnalyticsFunnelStepSuggestionsResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsFunnelStepSuggestionsResponseAttributesWithDefaults() *ProductAnalyticsFunnelStepSuggestionsResponseAttributes {
	this := ProductAnalyticsFunnelStepSuggestionsResponseAttributes{}
	return &this
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelStepSuggestionsResponseAttributes) GetSteps() []string {
	if o == nil || o.Steps == nil {
		var ret []string
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelStepSuggestionsResponseAttributes) GetStepsOk() (*[]string, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return &o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelStepSuggestionsResponseAttributes) HasSteps() bool {
	return o != nil && o.Steps != nil
}

// SetSteps gets a reference to the given []string and assigns it to the Steps field.
func (o *ProductAnalyticsFunnelStepSuggestionsResponseAttributes) SetSteps(v []string) {
	o.Steps = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsFunnelStepSuggestionsResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsFunnelStepSuggestionsResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Steps []string `json:"steps,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"steps"})
	} else {
		return err
	}
	o.Steps = all.Steps

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
