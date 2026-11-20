// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnthropicUsageAndCostsIntegrationAccountCreateAttributes Writable attributes used to create an Anthropic Usage and Costs integration account.
type AnthropicUsageAndCostsIntegrationAccountCreateAttributes struct {
	// Authentication for creating the Anthropic Usage and Costs integration account. Exactly one method is set.
	Authentication AnthropicUsageAndCostsIntegrationAccountAuthenticationRequest `json:"authentication"`
	// Data Datadog collects from Anthropic Usage and Costs, keyed by dataflow ID. Datadog always collects usage metrics. Each dataflow in this object adds another type of collection. Set `enabled` to start or stop a dataflow. The default for each dataflow applies when the account is created. On update, omitted fields keep their current values. If a dataflow depends on a Datadog product, its description names that product. If you enable a dataflow without that product, Datadog stores the setting but collects no data.
	Dataflows *AnthropicUsageAndCostsIntegrationDataflowsRequest `json:"dataflows,omitempty"`
	// Human-readable name of the Anthropic Usage and Costs integration account.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnthropicUsageAndCostsIntegrationAccountCreateAttributes instantiates a new AnthropicUsageAndCostsIntegrationAccountCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnthropicUsageAndCostsIntegrationAccountCreateAttributes(authentication AnthropicUsageAndCostsIntegrationAccountAuthenticationRequest, name string) *AnthropicUsageAndCostsIntegrationAccountCreateAttributes {
	this := AnthropicUsageAndCostsIntegrationAccountCreateAttributes{}
	this.Authentication = authentication
	this.Name = name
	return &this
}

// NewAnthropicUsageAndCostsIntegrationAccountCreateAttributesWithDefaults instantiates a new AnthropicUsageAndCostsIntegrationAccountCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnthropicUsageAndCostsIntegrationAccountCreateAttributesWithDefaults() *AnthropicUsageAndCostsIntegrationAccountCreateAttributes {
	this := AnthropicUsageAndCostsIntegrationAccountCreateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value.
func (o *AnthropicUsageAndCostsIntegrationAccountCreateAttributes) GetAuthentication() AnthropicUsageAndCostsIntegrationAccountAuthenticationRequest {
	if o == nil {
		var ret AnthropicUsageAndCostsIntegrationAccountAuthenticationRequest
		return ret
	}
	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *AnthropicUsageAndCostsIntegrationAccountCreateAttributes) GetAuthenticationOk() (*AnthropicUsageAndCostsIntegrationAccountAuthenticationRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value.
func (o *AnthropicUsageAndCostsIntegrationAccountCreateAttributes) SetAuthentication(v AnthropicUsageAndCostsIntegrationAccountAuthenticationRequest) {
	o.Authentication = v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *AnthropicUsageAndCostsIntegrationAccountCreateAttributes) GetDataflows() AnthropicUsageAndCostsIntegrationDataflowsRequest {
	if o == nil || o.Dataflows == nil {
		var ret AnthropicUsageAndCostsIntegrationDataflowsRequest
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnthropicUsageAndCostsIntegrationAccountCreateAttributes) GetDataflowsOk() (*AnthropicUsageAndCostsIntegrationDataflowsRequest, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *AnthropicUsageAndCostsIntegrationAccountCreateAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given AnthropicUsageAndCostsIntegrationDataflowsRequest and assigns it to the Dataflows field.
func (o *AnthropicUsageAndCostsIntegrationAccountCreateAttributes) SetDataflows(v AnthropicUsageAndCostsIntegrationDataflowsRequest) {
	o.Dataflows = &v
}

// GetName returns the Name field value.
func (o *AnthropicUsageAndCostsIntegrationAccountCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AnthropicUsageAndCostsIntegrationAccountCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AnthropicUsageAndCostsIntegrationAccountCreateAttributes) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnthropicUsageAndCostsIntegrationAccountCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["authentication"] = o.Authentication
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
func (o *AnthropicUsageAndCostsIntegrationAccountCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *AnthropicUsageAndCostsIntegrationAccountAuthenticationRequest `json:"authentication"`
		Dataflows      *AnthropicUsageAndCostsIntegrationDataflowsRequest             `json:"dataflows,omitempty"`
		Name           *string                                                        `json:"name"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"authentication", "dataflows", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Authentication = *all.Authentication
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
