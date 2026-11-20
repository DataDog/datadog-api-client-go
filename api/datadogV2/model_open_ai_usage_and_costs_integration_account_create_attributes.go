// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAIUsageAndCostsIntegrationAccountCreateAttributes Writable attributes used to create an OpenAI Usage and Costs integration account.
type OpenAIUsageAndCostsIntegrationAccountCreateAttributes struct {
	// Authentication for creating the OpenAI Usage and Costs integration account. Exactly one method is set.
	Authentication OpenAIUsageAndCostsIntegrationAccountAuthenticationRequest `json:"authentication"`
	// Data Datadog collects from OpenAI Usage and Costs, keyed by dataflow ID. Datadog always collects organization usage metrics. Each dataflow in this object adds another type of collection. Set `enabled` to start or stop a dataflow. The default for each dataflow applies when the account is created. On update, omitted fields keep their current values.
	Dataflows *OpenAIUsageAndCostsIntegrationDataflowsRequest `json:"dataflows,omitempty"`
	// Human-readable name of the OpenAI Usage and Costs integration account.
	Name string `json:"name"`
	// Settings for creating the OpenAI Usage and Costs integration account.
	Settings OpenAIUsageAndCostsIntegrationAccountSettingsRequest `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpenAIUsageAndCostsIntegrationAccountCreateAttributes instantiates a new OpenAIUsageAndCostsIntegrationAccountCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAIUsageAndCostsIntegrationAccountCreateAttributes(authentication OpenAIUsageAndCostsIntegrationAccountAuthenticationRequest, name string, settings OpenAIUsageAndCostsIntegrationAccountSettingsRequest) *OpenAIUsageAndCostsIntegrationAccountCreateAttributes {
	this := OpenAIUsageAndCostsIntegrationAccountCreateAttributes{}
	this.Authentication = authentication
	this.Name = name
	this.Settings = settings
	return &this
}

// NewOpenAIUsageAndCostsIntegrationAccountCreateAttributesWithDefaults instantiates a new OpenAIUsageAndCostsIntegrationAccountCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAIUsageAndCostsIntegrationAccountCreateAttributesWithDefaults() *OpenAIUsageAndCostsIntegrationAccountCreateAttributes {
	this := OpenAIUsageAndCostsIntegrationAccountCreateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value.
func (o *OpenAIUsageAndCostsIntegrationAccountCreateAttributes) GetAuthentication() OpenAIUsageAndCostsIntegrationAccountAuthenticationRequest {
	if o == nil {
		var ret OpenAIUsageAndCostsIntegrationAccountAuthenticationRequest
		return ret
	}
	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *OpenAIUsageAndCostsIntegrationAccountCreateAttributes) GetAuthenticationOk() (*OpenAIUsageAndCostsIntegrationAccountAuthenticationRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value.
func (o *OpenAIUsageAndCostsIntegrationAccountCreateAttributes) SetAuthentication(v OpenAIUsageAndCostsIntegrationAccountAuthenticationRequest) {
	o.Authentication = v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *OpenAIUsageAndCostsIntegrationAccountCreateAttributes) GetDataflows() OpenAIUsageAndCostsIntegrationDataflowsRequest {
	if o == nil || o.Dataflows == nil {
		var ret OpenAIUsageAndCostsIntegrationDataflowsRequest
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAIUsageAndCostsIntegrationAccountCreateAttributes) GetDataflowsOk() (*OpenAIUsageAndCostsIntegrationDataflowsRequest, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *OpenAIUsageAndCostsIntegrationAccountCreateAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given OpenAIUsageAndCostsIntegrationDataflowsRequest and assigns it to the Dataflows field.
func (o *OpenAIUsageAndCostsIntegrationAccountCreateAttributes) SetDataflows(v OpenAIUsageAndCostsIntegrationDataflowsRequest) {
	o.Dataflows = &v
}

// GetName returns the Name field value.
func (o *OpenAIUsageAndCostsIntegrationAccountCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpenAIUsageAndCostsIntegrationAccountCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OpenAIUsageAndCostsIntegrationAccountCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value.
func (o *OpenAIUsageAndCostsIntegrationAccountCreateAttributes) GetSettings() OpenAIUsageAndCostsIntegrationAccountSettingsRequest {
	if o == nil {
		var ret OpenAIUsageAndCostsIntegrationAccountSettingsRequest
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *OpenAIUsageAndCostsIntegrationAccountCreateAttributes) GetSettingsOk() (*OpenAIUsageAndCostsIntegrationAccountSettingsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *OpenAIUsageAndCostsIntegrationAccountCreateAttributes) SetSettings(v OpenAIUsageAndCostsIntegrationAccountSettingsRequest) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAIUsageAndCostsIntegrationAccountCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["authentication"] = o.Authentication
	if o.Dataflows != nil {
		toSerialize["dataflows"] = o.Dataflows
	}
	toSerialize["name"] = o.Name
	toSerialize["settings"] = o.Settings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpenAIUsageAndCostsIntegrationAccountCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *OpenAIUsageAndCostsIntegrationAccountAuthenticationRequest `json:"authentication"`
		Dataflows      *OpenAIUsageAndCostsIntegrationDataflowsRequest             `json:"dataflows,omitempty"`
		Name           *string                                                     `json:"name"`
		Settings       *OpenAIUsageAndCostsIntegrationAccountSettingsRequest       `json:"settings"`
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
	if all.Settings == nil {
		return fmt.Errorf("required field settings missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"authentication", "dataflows", "name", "settings"})
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
	if all.Settings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Settings = *all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
