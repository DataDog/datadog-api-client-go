// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFunnelSuggestionsDomain A feature area of an application, such as checkout or settings.
type ProductAnalyticsFunnelSuggestionsDomain struct {
	// Features in this area.
	Features []ProductAnalyticsFunnelSuggestionsFeature `json:"features,omitempty"`
	// Name of the feature area.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsFunnelSuggestionsDomain instantiates a new ProductAnalyticsFunnelSuggestionsDomain object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsFunnelSuggestionsDomain() *ProductAnalyticsFunnelSuggestionsDomain {
	this := ProductAnalyticsFunnelSuggestionsDomain{}
	return &this
}

// NewProductAnalyticsFunnelSuggestionsDomainWithDefaults instantiates a new ProductAnalyticsFunnelSuggestionsDomain object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsFunnelSuggestionsDomainWithDefaults() *ProductAnalyticsFunnelSuggestionsDomain {
	this := ProductAnalyticsFunnelSuggestionsDomain{}
	return &this
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelSuggestionsDomain) GetFeatures() []ProductAnalyticsFunnelSuggestionsFeature {
	if o == nil || o.Features == nil {
		var ret []ProductAnalyticsFunnelSuggestionsFeature
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelSuggestionsDomain) GetFeaturesOk() (*[]ProductAnalyticsFunnelSuggestionsFeature, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return &o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelSuggestionsDomain) HasFeatures() bool {
	return o != nil && o.Features != nil
}

// SetFeatures gets a reference to the given []ProductAnalyticsFunnelSuggestionsFeature and assigns it to the Features field.
func (o *ProductAnalyticsFunnelSuggestionsDomain) SetFeatures(v []ProductAnalyticsFunnelSuggestionsFeature) {
	o.Features = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelSuggestionsDomain) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelSuggestionsDomain) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelSuggestionsDomain) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductAnalyticsFunnelSuggestionsDomain) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsFunnelSuggestionsDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsFunnelSuggestionsDomain) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Features []ProductAnalyticsFunnelSuggestionsFeature `json:"features,omitempty"`
		Name     *string                                    `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"features", "name"})
	} else {
		return err
	}
	o.Features = all.Features
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
