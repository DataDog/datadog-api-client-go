// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ClaudeComplianceIntegrationAccountCreateAttributes Writable attributes used to create a Claude Compliance integration account.
type ClaudeComplianceIntegrationAccountCreateAttributes struct {
	// Authentication for creating the Claude Compliance integration account. Exactly one method is set.
	Authentication ClaudeComplianceIntegrationAccountAuthenticationRequest `json:"authentication"`
	// Human-readable name of the Claude Compliance integration account.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClaudeComplianceIntegrationAccountCreateAttributes instantiates a new ClaudeComplianceIntegrationAccountCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClaudeComplianceIntegrationAccountCreateAttributes(authentication ClaudeComplianceIntegrationAccountAuthenticationRequest, name string) *ClaudeComplianceIntegrationAccountCreateAttributes {
	this := ClaudeComplianceIntegrationAccountCreateAttributes{}
	this.Authentication = authentication
	this.Name = name
	return &this
}

// NewClaudeComplianceIntegrationAccountCreateAttributesWithDefaults instantiates a new ClaudeComplianceIntegrationAccountCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClaudeComplianceIntegrationAccountCreateAttributesWithDefaults() *ClaudeComplianceIntegrationAccountCreateAttributes {
	this := ClaudeComplianceIntegrationAccountCreateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value.
func (o *ClaudeComplianceIntegrationAccountCreateAttributes) GetAuthentication() ClaudeComplianceIntegrationAccountAuthenticationRequest {
	if o == nil {
		var ret ClaudeComplianceIntegrationAccountAuthenticationRequest
		return ret
	}
	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *ClaudeComplianceIntegrationAccountCreateAttributes) GetAuthenticationOk() (*ClaudeComplianceIntegrationAccountAuthenticationRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value.
func (o *ClaudeComplianceIntegrationAccountCreateAttributes) SetAuthentication(v ClaudeComplianceIntegrationAccountAuthenticationRequest) {
	o.Authentication = v
}

// GetName returns the Name field value.
func (o *ClaudeComplianceIntegrationAccountCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClaudeComplianceIntegrationAccountCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ClaudeComplianceIntegrationAccountCreateAttributes) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClaudeComplianceIntegrationAccountCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["authentication"] = o.Authentication
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClaudeComplianceIntegrationAccountCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *ClaudeComplianceIntegrationAccountAuthenticationRequest `json:"authentication"`
		Name           *string                                                  `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Authentication == nil {
		return fmt.Errorf("required field authentication missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"authentication", "name"})
	} else {
		return err
	}
	o.Authentication = *all.Authentication
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
