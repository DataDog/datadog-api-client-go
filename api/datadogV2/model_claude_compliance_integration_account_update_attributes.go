// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ClaudeComplianceIntegrationAccountUpdateAttributes Writable attributes used to update a Claude Compliance integration account. Every field is optional. Only the provided fields are changed.
type ClaudeComplianceIntegrationAccountUpdateAttributes struct {
	// Authentication for updating the Claude Compliance integration account. Exactly one method is set.
	Authentication *ClaudeComplianceIntegrationAccountAuthenticationUpdate `json:"authentication,omitempty"`
	// Human-readable name of the Claude Compliance integration account.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClaudeComplianceIntegrationAccountUpdateAttributes instantiates a new ClaudeComplianceIntegrationAccountUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClaudeComplianceIntegrationAccountUpdateAttributes() *ClaudeComplianceIntegrationAccountUpdateAttributes {
	this := ClaudeComplianceIntegrationAccountUpdateAttributes{}
	return &this
}

// NewClaudeComplianceIntegrationAccountUpdateAttributesWithDefaults instantiates a new ClaudeComplianceIntegrationAccountUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClaudeComplianceIntegrationAccountUpdateAttributesWithDefaults() *ClaudeComplianceIntegrationAccountUpdateAttributes {
	this := ClaudeComplianceIntegrationAccountUpdateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *ClaudeComplianceIntegrationAccountUpdateAttributes) GetAuthentication() ClaudeComplianceIntegrationAccountAuthenticationUpdate {
	if o == nil || o.Authentication == nil {
		var ret ClaudeComplianceIntegrationAccountAuthenticationUpdate
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaudeComplianceIntegrationAccountUpdateAttributes) GetAuthenticationOk() (*ClaudeComplianceIntegrationAccountAuthenticationUpdate, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *ClaudeComplianceIntegrationAccountUpdateAttributes) HasAuthentication() bool {
	return o != nil && o.Authentication != nil
}

// SetAuthentication gets a reference to the given ClaudeComplianceIntegrationAccountAuthenticationUpdate and assigns it to the Authentication field.
func (o *ClaudeComplianceIntegrationAccountUpdateAttributes) SetAuthentication(v ClaudeComplianceIntegrationAccountAuthenticationUpdate) {
	o.Authentication = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClaudeComplianceIntegrationAccountUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaudeComplianceIntegrationAccountUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClaudeComplianceIntegrationAccountUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClaudeComplianceIntegrationAccountUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClaudeComplianceIntegrationAccountUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Authentication != nil {
		toSerialize["authentication"] = o.Authentication
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
func (o *ClaudeComplianceIntegrationAccountUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *ClaudeComplianceIntegrationAccountAuthenticationUpdate `json:"authentication,omitempty"`
		Name           *string                                                 `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"authentication", "name"})
	} else {
		return err
	}
	o.Authentication = all.Authentication
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
