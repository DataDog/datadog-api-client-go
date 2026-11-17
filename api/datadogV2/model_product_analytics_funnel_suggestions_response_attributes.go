// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFunnelSuggestionsResponseAttributes Attributes of a feature suggestions response.
type ProductAnalyticsFunnelSuggestionsResponseAttributes struct {
	// Suggested features, grouped by application and feature area.
	Applications []ProductAnalyticsFunnelSuggestionsApplication `json:"applications,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsFunnelSuggestionsResponseAttributes instantiates a new ProductAnalyticsFunnelSuggestionsResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsFunnelSuggestionsResponseAttributes() *ProductAnalyticsFunnelSuggestionsResponseAttributes {
	this := ProductAnalyticsFunnelSuggestionsResponseAttributes{}
	return &this
}

// NewProductAnalyticsFunnelSuggestionsResponseAttributesWithDefaults instantiates a new ProductAnalyticsFunnelSuggestionsResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsFunnelSuggestionsResponseAttributesWithDefaults() *ProductAnalyticsFunnelSuggestionsResponseAttributes {
	this := ProductAnalyticsFunnelSuggestionsResponseAttributes{}
	return &this
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelSuggestionsResponseAttributes) GetApplications() []ProductAnalyticsFunnelSuggestionsApplication {
	if o == nil || o.Applications == nil {
		var ret []ProductAnalyticsFunnelSuggestionsApplication
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelSuggestionsResponseAttributes) GetApplicationsOk() (*[]ProductAnalyticsFunnelSuggestionsApplication, bool) {
	if o == nil || o.Applications == nil {
		return nil, false
	}
	return &o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelSuggestionsResponseAttributes) HasApplications() bool {
	return o != nil && o.Applications != nil
}

// SetApplications gets a reference to the given []ProductAnalyticsFunnelSuggestionsApplication and assigns it to the Applications field.
func (o *ProductAnalyticsFunnelSuggestionsResponseAttributes) SetApplications(v []ProductAnalyticsFunnelSuggestionsApplication) {
	o.Applications = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsFunnelSuggestionsResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Applications != nil {
		toSerialize["applications"] = o.Applications
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsFunnelSuggestionsResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Applications []ProductAnalyticsFunnelSuggestionsApplication `json:"applications,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"applications"})
	} else {
		return err
	}
	o.Applications = all.Applications

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
