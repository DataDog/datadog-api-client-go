// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudIntegrationAccountSettingsResponse Settings configured on the Elastic Cloud integration account.
type ElasticCloudIntegrationAccountSettingsResponse struct {
	// Comma-separated list of custom tags for this Elastic Cloud deployment.
	Tags *string `json:"tags,omitempty"`
	// Elastic Cloud deployment URL.
	Url string `json:"url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticCloudIntegrationAccountSettingsResponse instantiates a new ElasticCloudIntegrationAccountSettingsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticCloudIntegrationAccountSettingsResponse(url string) *ElasticCloudIntegrationAccountSettingsResponse {
	this := ElasticCloudIntegrationAccountSettingsResponse{}
	this.Url = url
	return &this
}

// NewElasticCloudIntegrationAccountSettingsResponseWithDefaults instantiates a new ElasticCloudIntegrationAccountSettingsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticCloudIntegrationAccountSettingsResponseWithDefaults() *ElasticCloudIntegrationAccountSettingsResponse {
	this := ElasticCloudIntegrationAccountSettingsResponse{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ElasticCloudIntegrationAccountSettingsResponse) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationAccountSettingsResponse) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ElasticCloudIntegrationAccountSettingsResponse) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *ElasticCloudIntegrationAccountSettingsResponse) SetTags(v string) {
	o.Tags = &v
}

// GetUrl returns the Url field value.
func (o *ElasticCloudIntegrationAccountSettingsResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ElasticCloudIntegrationAccountSettingsResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *ElasticCloudIntegrationAccountSettingsResponse) SetUrl(v string) {
	o.Url = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticCloudIntegrationAccountSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticCloudIntegrationAccountSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Tags *string `json:"tags,omitempty"`
		Url  *string `json:"url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"tags", "url"})
	} else {
		return err
	}
	o.Tags = all.Tags
	o.Url = *all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
