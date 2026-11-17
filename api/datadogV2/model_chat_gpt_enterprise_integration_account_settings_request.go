// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChatGPTEnterpriseIntegrationAccountSettingsRequest Settings for creating the ChatGPT Enterprise integration account.
type ChatGPTEnterpriseIntegrationAccountSettingsRequest struct {
	// OpenAI API Platform organization ID (`org-...`) that the Costs API uses for your Enterprise usage. This is not the ChatGPT workspace UUID or your Datadog organization ID. A value that does not match `^org-[A-Za-z0-9]+$` is rejected with a `422` response.
	OpenaiOrgId string `json:"openai_org_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewChatGPTEnterpriseIntegrationAccountSettingsRequest instantiates a new ChatGPTEnterpriseIntegrationAccountSettingsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChatGPTEnterpriseIntegrationAccountSettingsRequest(openaiOrgId string) *ChatGPTEnterpriseIntegrationAccountSettingsRequest {
	this := ChatGPTEnterpriseIntegrationAccountSettingsRequest{}
	this.OpenaiOrgId = openaiOrgId
	return &this
}

// NewChatGPTEnterpriseIntegrationAccountSettingsRequestWithDefaults instantiates a new ChatGPTEnterpriseIntegrationAccountSettingsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChatGPTEnterpriseIntegrationAccountSettingsRequestWithDefaults() *ChatGPTEnterpriseIntegrationAccountSettingsRequest {
	this := ChatGPTEnterpriseIntegrationAccountSettingsRequest{}
	return &this
}

// GetOpenaiOrgId returns the OpenaiOrgId field value.
func (o *ChatGPTEnterpriseIntegrationAccountSettingsRequest) GetOpenaiOrgId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OpenaiOrgId
}

// GetOpenaiOrgIdOk returns a tuple with the OpenaiOrgId field value
// and a boolean to check if the value has been set.
func (o *ChatGPTEnterpriseIntegrationAccountSettingsRequest) GetOpenaiOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpenaiOrgId, true
}

// SetOpenaiOrgId sets field value.
func (o *ChatGPTEnterpriseIntegrationAccountSettingsRequest) SetOpenaiOrgId(v string) {
	o.OpenaiOrgId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChatGPTEnterpriseIntegrationAccountSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["openai_org_id"] = o.OpenaiOrgId
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChatGPTEnterpriseIntegrationAccountSettingsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OpenaiOrgId *string `json:"openai_org_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.OpenaiOrgId == nil {
		return fmt.Errorf("required field openai_org_id missing")
	}
	o.OpenaiOrgId = *all.OpenaiOrgId

	return nil
}
