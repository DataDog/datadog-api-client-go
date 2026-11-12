// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AgentHealthIssuesResponseMeta Pagination metadata for the matching Agent Health issues.
type AgentHealthIssuesResponseMeta struct {
	// Pagination counts for the Agent Health issues.
	Page AgentHealthIssuesResponsePage `json:"page"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAgentHealthIssuesResponseMeta instantiates a new AgentHealthIssuesResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAgentHealthIssuesResponseMeta(page AgentHealthIssuesResponsePage) *AgentHealthIssuesResponseMeta {
	this := AgentHealthIssuesResponseMeta{}
	this.Page = page
	return &this
}

// NewAgentHealthIssuesResponseMetaWithDefaults instantiates a new AgentHealthIssuesResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAgentHealthIssuesResponseMetaWithDefaults() *AgentHealthIssuesResponseMeta {
	this := AgentHealthIssuesResponseMeta{}
	return &this
}

// GetPage returns the Page field value.
func (o *AgentHealthIssuesResponseMeta) GetPage() AgentHealthIssuesResponsePage {
	if o == nil {
		var ret AgentHealthIssuesResponsePage
		return ret
	}
	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *AgentHealthIssuesResponseMeta) GetPageOk() (*AgentHealthIssuesResponsePage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value.
func (o *AgentHealthIssuesResponseMeta) SetPage(v AgentHealthIssuesResponsePage) {
	o.Page = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AgentHealthIssuesResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["page"] = o.Page

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AgentHealthIssuesResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Page *AgentHealthIssuesResponsePage `json:"page"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Page == nil {
		return fmt.Errorf("required field page missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"page"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Page = *all.Page

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
