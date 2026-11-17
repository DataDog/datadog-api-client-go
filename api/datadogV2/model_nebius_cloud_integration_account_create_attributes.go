// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudIntegrationAccountCreateAttributes Writable attributes used to create a Nebius Cloud integration account.
type NebiusCloudIntegrationAccountCreateAttributes struct {
	// Authentication for creating the Nebius Cloud integration account. Exactly one method is set.
	Authentication NebiusCloudIntegrationAccountAuthenticationRequest `json:"authentication"`
	// Data Datadog collects from Nebius Cloud, keyed by dataflow id. Each dataflow turns on a distinct kind of collection: set `enabled` to stop or start it, and use `settings` to configure what it collects. Defaults listed on each dataflow apply when the account is created; on update, omitted fields keep their current values. Every dataflow except `nebius-cloud-cost-collection` runs as the service account in `authentication`, using its authorized key pair and the static API key; enabling one of them without those keys is rejected with a `422` response, and one enabled with keys that later expire is stored but collects no data.
	Dataflows *NebiusCloudIntegrationDataflowsRequest `json:"dataflows,omitempty"`
	// Human-readable name of the Nebius Cloud integration account.
	Name string `json:"name"`
	// Settings for creating the Nebius Cloud integration account.
	Settings NebiusCloudIntegrationAccountSettingsRequest `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNebiusCloudIntegrationAccountCreateAttributes instantiates a new NebiusCloudIntegrationAccountCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNebiusCloudIntegrationAccountCreateAttributes(authentication NebiusCloudIntegrationAccountAuthenticationRequest, name string, settings NebiusCloudIntegrationAccountSettingsRequest) *NebiusCloudIntegrationAccountCreateAttributes {
	this := NebiusCloudIntegrationAccountCreateAttributes{}
	this.Authentication = authentication
	this.Name = name
	this.Settings = settings
	return &this
}

// NewNebiusCloudIntegrationAccountCreateAttributesWithDefaults instantiates a new NebiusCloudIntegrationAccountCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNebiusCloudIntegrationAccountCreateAttributesWithDefaults() *NebiusCloudIntegrationAccountCreateAttributes {
	this := NebiusCloudIntegrationAccountCreateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value.
func (o *NebiusCloudIntegrationAccountCreateAttributes) GetAuthentication() NebiusCloudIntegrationAccountAuthenticationRequest {
	if o == nil {
		var ret NebiusCloudIntegrationAccountAuthenticationRequest
		return ret
	}
	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountCreateAttributes) GetAuthenticationOk() (*NebiusCloudIntegrationAccountAuthenticationRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value.
func (o *NebiusCloudIntegrationAccountCreateAttributes) SetAuthentication(v NebiusCloudIntegrationAccountAuthenticationRequest) {
	o.Authentication = v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountCreateAttributes) GetDataflows() NebiusCloudIntegrationDataflowsRequest {
	if o == nil || o.Dataflows == nil {
		var ret NebiusCloudIntegrationDataflowsRequest
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountCreateAttributes) GetDataflowsOk() (*NebiusCloudIntegrationDataflowsRequest, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountCreateAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given NebiusCloudIntegrationDataflowsRequest and assigns it to the Dataflows field.
func (o *NebiusCloudIntegrationAccountCreateAttributes) SetDataflows(v NebiusCloudIntegrationDataflowsRequest) {
	o.Dataflows = &v
}

// GetName returns the Name field value.
func (o *NebiusCloudIntegrationAccountCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *NebiusCloudIntegrationAccountCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value.
func (o *NebiusCloudIntegrationAccountCreateAttributes) GetSettings() NebiusCloudIntegrationAccountSettingsRequest {
	if o == nil {
		var ret NebiusCloudIntegrationAccountSettingsRequest
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountCreateAttributes) GetSettingsOk() (*NebiusCloudIntegrationAccountSettingsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *NebiusCloudIntegrationAccountCreateAttributes) SetSettings(v NebiusCloudIntegrationAccountSettingsRequest) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NebiusCloudIntegrationAccountCreateAttributes) MarshalJSON() ([]byte, error) {
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
func (o *NebiusCloudIntegrationAccountCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *NebiusCloudIntegrationAccountAuthenticationRequest `json:"authentication"`
		Dataflows      *NebiusCloudIntegrationDataflowsRequest             `json:"dataflows,omitempty"`
		Name           *string                                             `json:"name"`
		Settings       *NebiusCloudIntegrationAccountSettingsRequest       `json:"settings"`
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
