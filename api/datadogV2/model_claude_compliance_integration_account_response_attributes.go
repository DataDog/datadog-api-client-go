// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ClaudeComplianceIntegrationAccountResponseAttributes Attributes of a Claude Compliance integration account returned in responses.
type ClaudeComplianceIntegrationAccountResponseAttributes struct {
	// Authentication configured on the Claude Compliance integration account.
	Authentication *ClaudeComplianceIntegrationAccountAuthenticationResponse `json:"authentication,omitempty"`
	// Data Datadog collects from Claude Compliance, keyed by dataflow ID.
	Dataflows *ClaudeComplianceIntegrationDataflowsResponse `json:"dataflows,omitempty"`
	// Human-readable name of the Claude Compliance integration account.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClaudeComplianceIntegrationAccountResponseAttributes instantiates a new ClaudeComplianceIntegrationAccountResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClaudeComplianceIntegrationAccountResponseAttributes(name string) *ClaudeComplianceIntegrationAccountResponseAttributes {
	this := ClaudeComplianceIntegrationAccountResponseAttributes{}
	this.Name = name
	return &this
}

// NewClaudeComplianceIntegrationAccountResponseAttributesWithDefaults instantiates a new ClaudeComplianceIntegrationAccountResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClaudeComplianceIntegrationAccountResponseAttributesWithDefaults() *ClaudeComplianceIntegrationAccountResponseAttributes {
	this := ClaudeComplianceIntegrationAccountResponseAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *ClaudeComplianceIntegrationAccountResponseAttributes) GetAuthentication() ClaudeComplianceIntegrationAccountAuthenticationResponse {
	if o == nil || o.Authentication == nil {
		var ret ClaudeComplianceIntegrationAccountAuthenticationResponse
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaudeComplianceIntegrationAccountResponseAttributes) GetAuthenticationOk() (*ClaudeComplianceIntegrationAccountAuthenticationResponse, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *ClaudeComplianceIntegrationAccountResponseAttributes) HasAuthentication() bool {
	return o != nil && o.Authentication != nil
}

// SetAuthentication gets a reference to the given ClaudeComplianceIntegrationAccountAuthenticationResponse and assigns it to the Authentication field.
func (o *ClaudeComplianceIntegrationAccountResponseAttributes) SetAuthentication(v ClaudeComplianceIntegrationAccountAuthenticationResponse) {
	o.Authentication = &v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *ClaudeComplianceIntegrationAccountResponseAttributes) GetDataflows() ClaudeComplianceIntegrationDataflowsResponse {
	if o == nil || o.Dataflows == nil {
		var ret ClaudeComplianceIntegrationDataflowsResponse
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaudeComplianceIntegrationAccountResponseAttributes) GetDataflowsOk() (*ClaudeComplianceIntegrationDataflowsResponse, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *ClaudeComplianceIntegrationAccountResponseAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given ClaudeComplianceIntegrationDataflowsResponse and assigns it to the Dataflows field.
func (o *ClaudeComplianceIntegrationAccountResponseAttributes) SetDataflows(v ClaudeComplianceIntegrationDataflowsResponse) {
	o.Dataflows = &v
}

// GetName returns the Name field value.
func (o *ClaudeComplianceIntegrationAccountResponseAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClaudeComplianceIntegrationAccountResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ClaudeComplianceIntegrationAccountResponseAttributes) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClaudeComplianceIntegrationAccountResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Authentication != nil {
		toSerialize["authentication"] = o.Authentication
	}
	if o.Dataflows != nil {
		toSerialize["dataflows"] = o.Dataflows
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClaudeComplianceIntegrationAccountResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *ClaudeComplianceIntegrationAccountAuthenticationResponse `json:"authentication,omitempty"`
		Dataflows      *ClaudeComplianceIntegrationDataflowsResponse             `json:"dataflows,omitempty"`
		Name           *string                                                   `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"authentication", "dataflows", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Authentication = all.Authentication
	if all.Dataflows != nil && all.Dataflows.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Dataflows = all.Dataflows
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
