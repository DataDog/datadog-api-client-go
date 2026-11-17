// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFunnelSuggestionsApplication Suggested features of one application.
type ProductAnalyticsFunnelSuggestionsApplication struct {
	// ID of the RUM application.
	ApplicationId *string `json:"application_id,omitempty"`
	// Feature areas of the application.
	Domains []ProductAnalyticsFunnelSuggestionsDomain `json:"domains,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsFunnelSuggestionsApplication instantiates a new ProductAnalyticsFunnelSuggestionsApplication object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsFunnelSuggestionsApplication() *ProductAnalyticsFunnelSuggestionsApplication {
	this := ProductAnalyticsFunnelSuggestionsApplication{}
	return &this
}

// NewProductAnalyticsFunnelSuggestionsApplicationWithDefaults instantiates a new ProductAnalyticsFunnelSuggestionsApplication object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsFunnelSuggestionsApplicationWithDefaults() *ProductAnalyticsFunnelSuggestionsApplication {
	this := ProductAnalyticsFunnelSuggestionsApplication{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelSuggestionsApplication) GetApplicationId() string {
	if o == nil || o.ApplicationId == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelSuggestionsApplication) GetApplicationIdOk() (*string, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelSuggestionsApplication) HasApplicationId() bool {
	return o != nil && o.ApplicationId != nil
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *ProductAnalyticsFunnelSuggestionsApplication) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelSuggestionsApplication) GetDomains() []ProductAnalyticsFunnelSuggestionsDomain {
	if o == nil || o.Domains == nil {
		var ret []ProductAnalyticsFunnelSuggestionsDomain
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelSuggestionsApplication) GetDomainsOk() (*[]ProductAnalyticsFunnelSuggestionsDomain, bool) {
	if o == nil || o.Domains == nil {
		return nil, false
	}
	return &o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelSuggestionsApplication) HasDomains() bool {
	return o != nil && o.Domains != nil
}

// SetDomains gets a reference to the given []ProductAnalyticsFunnelSuggestionsDomain and assigns it to the Domains field.
func (o *ProductAnalyticsFunnelSuggestionsApplication) SetDomains(v []ProductAnalyticsFunnelSuggestionsDomain) {
	o.Domains = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsFunnelSuggestionsApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApplicationId != nil {
		toSerialize["application_id"] = o.ApplicationId
	}
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsFunnelSuggestionsApplication) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId *string                                   `json:"application_id,omitempty"`
		Domains       []ProductAnalyticsFunnelSuggestionsDomain `json:"domains,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "domains"})
	} else {
		return err
	}
	o.ApplicationId = all.ApplicationId
	o.Domains = all.Domains

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
