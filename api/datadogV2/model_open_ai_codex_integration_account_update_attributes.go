// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAICodexIntegrationAccountUpdateAttributes Writable attributes used to update an OpenAI Codex integration account. Every field is optional; only the fields provided are changed. When `dataflows` is provided, only the dataflow ids included in the request are modified; dataflows omitted from the map keep their current configuration.
type OpenAICodexIntegrationAccountUpdateAttributes struct {
	// Authentication for updating the OpenAI Codex integration account. Exactly one method is set.
	Authentication *OpenAICodexIntegrationAccountAuthenticationUpdate `json:"authentication,omitempty"`
	// Data Datadog collects from OpenAI Codex, keyed by dataflow id. Daily Codex analytics metrics are always collected; each dataflow here adds a further kind of collection on top of that baseline, so set `enabled` to start or stop it. Defaults listed on each dataflow apply when the account is created; on update, omitted fields keep their current values. Every dataflow requires an OpenAI enterprise subscription with Codex enabled; a dataflow enabled without it is stored but collects no data.
	Dataflows *OpenAICodexIntegrationDataflowsRequest `json:"dataflows,omitempty"`
	// Human-readable name of the OpenAI Codex integration account.
	Name *string `json:"name,omitempty"`
	// Settings for updating the OpenAI Codex integration account. Only the fields provided are changed.
	Settings *OpenAICodexIntegrationAccountSettingsUpdate `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpenAICodexIntegrationAccountUpdateAttributes instantiates a new OpenAICodexIntegrationAccountUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAICodexIntegrationAccountUpdateAttributes() *OpenAICodexIntegrationAccountUpdateAttributes {
	this := OpenAICodexIntegrationAccountUpdateAttributes{}
	return &this
}

// NewOpenAICodexIntegrationAccountUpdateAttributesWithDefaults instantiates a new OpenAICodexIntegrationAccountUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAICodexIntegrationAccountUpdateAttributesWithDefaults() *OpenAICodexIntegrationAccountUpdateAttributes {
	this := OpenAICodexIntegrationAccountUpdateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *OpenAICodexIntegrationAccountUpdateAttributes) GetAuthentication() OpenAICodexIntegrationAccountAuthenticationUpdate {
	if o == nil || o.Authentication == nil {
		var ret OpenAICodexIntegrationAccountAuthenticationUpdate
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAICodexIntegrationAccountUpdateAttributes) GetAuthenticationOk() (*OpenAICodexIntegrationAccountAuthenticationUpdate, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *OpenAICodexIntegrationAccountUpdateAttributes) HasAuthentication() bool {
	return o != nil && o.Authentication != nil
}

// SetAuthentication gets a reference to the given OpenAICodexIntegrationAccountAuthenticationUpdate and assigns it to the Authentication field.
func (o *OpenAICodexIntegrationAccountUpdateAttributes) SetAuthentication(v OpenAICodexIntegrationAccountAuthenticationUpdate) {
	o.Authentication = &v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *OpenAICodexIntegrationAccountUpdateAttributes) GetDataflows() OpenAICodexIntegrationDataflowsRequest {
	if o == nil || o.Dataflows == nil {
		var ret OpenAICodexIntegrationDataflowsRequest
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAICodexIntegrationAccountUpdateAttributes) GetDataflowsOk() (*OpenAICodexIntegrationDataflowsRequest, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *OpenAICodexIntegrationAccountUpdateAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given OpenAICodexIntegrationDataflowsRequest and assigns it to the Dataflows field.
func (o *OpenAICodexIntegrationAccountUpdateAttributes) SetDataflows(v OpenAICodexIntegrationDataflowsRequest) {
	o.Dataflows = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OpenAICodexIntegrationAccountUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAICodexIntegrationAccountUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OpenAICodexIntegrationAccountUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OpenAICodexIntegrationAccountUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *OpenAICodexIntegrationAccountUpdateAttributes) GetSettings() OpenAICodexIntegrationAccountSettingsUpdate {
	if o == nil || o.Settings == nil {
		var ret OpenAICodexIntegrationAccountSettingsUpdate
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAICodexIntegrationAccountUpdateAttributes) GetSettingsOk() (*OpenAICodexIntegrationAccountSettingsUpdate, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *OpenAICodexIntegrationAccountUpdateAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given OpenAICodexIntegrationAccountSettingsUpdate and assigns it to the Settings field.
func (o *OpenAICodexIntegrationAccountUpdateAttributes) SetSettings(v OpenAICodexIntegrationAccountSettingsUpdate) {
	o.Settings = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAICodexIntegrationAccountUpdateAttributes) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpenAICodexIntegrationAccountUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *OpenAICodexIntegrationAccountAuthenticationUpdate `json:"authentication,omitempty"`
		Dataflows      *OpenAICodexIntegrationDataflowsRequest            `json:"dataflows,omitempty"`
		Name           *string                                            `json:"name,omitempty"`
		Settings       *OpenAICodexIntegrationAccountSettingsUpdate       `json:"settings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"authentication", "dataflows", "name", "settings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Authentication = all.Authentication
	if all.Dataflows != nil && all.Dataflows.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Dataflows = all.Dataflows
	o.Name = all.Name
	if all.Settings != nil && all.Settings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Settings = all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
