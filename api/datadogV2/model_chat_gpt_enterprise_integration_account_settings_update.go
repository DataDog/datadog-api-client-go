// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChatGPTEnterpriseIntegrationAccountSettingsUpdate Settings for updating the ChatGPT Enterprise integration account. Only the fields provided are changed.
type ChatGPTEnterpriseIntegrationAccountSettingsUpdate struct {
	// OpenAI API Platform organization ID (`org-...`) that the Costs API uses for your Enterprise usage. This is not the ChatGPT workspace UUID or your Datadog organization ID. A value that does not match `^org-[A-Za-z0-9]+$` is rejected with a `422` response.
	OpenaiOrgId *string `json:"openai_org_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewChatGPTEnterpriseIntegrationAccountSettingsUpdate instantiates a new ChatGPTEnterpriseIntegrationAccountSettingsUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChatGPTEnterpriseIntegrationAccountSettingsUpdate() *ChatGPTEnterpriseIntegrationAccountSettingsUpdate {
	this := ChatGPTEnterpriseIntegrationAccountSettingsUpdate{}
	return &this
}

// NewChatGPTEnterpriseIntegrationAccountSettingsUpdateWithDefaults instantiates a new ChatGPTEnterpriseIntegrationAccountSettingsUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChatGPTEnterpriseIntegrationAccountSettingsUpdateWithDefaults() *ChatGPTEnterpriseIntegrationAccountSettingsUpdate {
	this := ChatGPTEnterpriseIntegrationAccountSettingsUpdate{}
	return &this
}

// GetOpenaiOrgId returns the OpenaiOrgId field value if set, zero value otherwise.
func (o *ChatGPTEnterpriseIntegrationAccountSettingsUpdate) GetOpenaiOrgId() string {
	if o == nil || o.OpenaiOrgId == nil {
		var ret string
		return ret
	}
	return *o.OpenaiOrgId
}

// GetOpenaiOrgIdOk returns a tuple with the OpenaiOrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatGPTEnterpriseIntegrationAccountSettingsUpdate) GetOpenaiOrgIdOk() (*string, bool) {
	if o == nil || o.OpenaiOrgId == nil {
		return nil, false
	}
	return o.OpenaiOrgId, true
}

// HasOpenaiOrgId returns a boolean if a field has been set.
func (o *ChatGPTEnterpriseIntegrationAccountSettingsUpdate) HasOpenaiOrgId() bool {
	return o != nil && o.OpenaiOrgId != nil
}

// SetOpenaiOrgId gets a reference to the given string and assigns it to the OpenaiOrgId field.
func (o *ChatGPTEnterpriseIntegrationAccountSettingsUpdate) SetOpenaiOrgId(v string) {
	o.OpenaiOrgId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChatGPTEnterpriseIntegrationAccountSettingsUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.OpenaiOrgId != nil {
		toSerialize["openai_org_id"] = o.OpenaiOrgId
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChatGPTEnterpriseIntegrationAccountSettingsUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OpenaiOrgId *string `json:"openai_org_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.OpenaiOrgId = all.OpenaiOrgId

	return nil
}
