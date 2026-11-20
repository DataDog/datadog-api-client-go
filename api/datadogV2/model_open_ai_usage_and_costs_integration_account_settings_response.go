// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAIUsageAndCostsIntegrationAccountSettingsResponse Settings configured on the OpenAI Usage and Costs integration account.
type OpenAIUsageAndCostsIntegrationAccountSettingsResponse struct {
	// Comma-separated list of custom tags applied to the data collected for this account, for example `team:ai,env:prod`.
	CustomTags *string `json:"custom_tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpenAIUsageAndCostsIntegrationAccountSettingsResponse instantiates a new OpenAIUsageAndCostsIntegrationAccountSettingsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAIUsageAndCostsIntegrationAccountSettingsResponse() *OpenAIUsageAndCostsIntegrationAccountSettingsResponse {
	this := OpenAIUsageAndCostsIntegrationAccountSettingsResponse{}
	return &this
}

// NewOpenAIUsageAndCostsIntegrationAccountSettingsResponseWithDefaults instantiates a new OpenAIUsageAndCostsIntegrationAccountSettingsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAIUsageAndCostsIntegrationAccountSettingsResponseWithDefaults() *OpenAIUsageAndCostsIntegrationAccountSettingsResponse {
	this := OpenAIUsageAndCostsIntegrationAccountSettingsResponse{}
	return &this
}

// GetCustomTags returns the CustomTags field value if set, zero value otherwise.
func (o *OpenAIUsageAndCostsIntegrationAccountSettingsResponse) GetCustomTags() string {
	if o == nil || o.CustomTags == nil {
		var ret string
		return ret
	}
	return *o.CustomTags
}

// GetCustomTagsOk returns a tuple with the CustomTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAIUsageAndCostsIntegrationAccountSettingsResponse) GetCustomTagsOk() (*string, bool) {
	if o == nil || o.CustomTags == nil {
		return nil, false
	}
	return o.CustomTags, true
}

// HasCustomTags returns a boolean if a field has been set.
func (o *OpenAIUsageAndCostsIntegrationAccountSettingsResponse) HasCustomTags() bool {
	return o != nil && o.CustomTags != nil
}

// SetCustomTags gets a reference to the given string and assigns it to the CustomTags field.
func (o *OpenAIUsageAndCostsIntegrationAccountSettingsResponse) SetCustomTags(v string) {
	o.CustomTags = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAIUsageAndCostsIntegrationAccountSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CustomTags != nil {
		toSerialize["custom_tags"] = o.CustomTags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpenAIUsageAndCostsIntegrationAccountSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomTags *string `json:"custom_tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"custom_tags"})
	} else {
		return err
	}
	o.CustomTags = all.CustomTags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
