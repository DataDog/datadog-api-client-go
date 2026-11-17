// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CoreWeaveCloudIntegrationAccountCreateAttributes Writable attributes used to create a CoreWeave integration account.
type CoreWeaveCloudIntegrationAccountCreateAttributes struct {
	// Authentication for creating the CoreWeave integration account. Exactly one method is set.
	Authentication CoreWeaveCloudIntegrationAccountAuthenticationRequest `json:"authentication"`
	// Data Datadog collects from CoreWeave, keyed by dataflow id. Each dataflow turns on a distinct kind of collection: set `enabled` to start or stop it. Defaults listed on each dataflow apply when the account is created; on update, omitted dataflow ids keep their current values. Collection runs with the API access token in `authentication`, whose user must have the Observability Viewer role; a dataflow enabled without that role is stored but collects no data.
	Dataflows *CoreWeaveCloudIntegrationDataflowsRequest `json:"dataflows,omitempty"`
	// Human-readable name of the CoreWeave integration account.
	Name string `json:"name"`
	// Settings for creating the CoreWeave integration account.
	Settings CoreWeaveCloudIntegrationAccountSettingsRequest `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCoreWeaveCloudIntegrationAccountCreateAttributes instantiates a new CoreWeaveCloudIntegrationAccountCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCoreWeaveCloudIntegrationAccountCreateAttributes(authentication CoreWeaveCloudIntegrationAccountAuthenticationRequest, name string, settings CoreWeaveCloudIntegrationAccountSettingsRequest) *CoreWeaveCloudIntegrationAccountCreateAttributes {
	this := CoreWeaveCloudIntegrationAccountCreateAttributes{}
	this.Authentication = authentication
	this.Name = name
	this.Settings = settings
	return &this
}

// NewCoreWeaveCloudIntegrationAccountCreateAttributesWithDefaults instantiates a new CoreWeaveCloudIntegrationAccountCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCoreWeaveCloudIntegrationAccountCreateAttributesWithDefaults() *CoreWeaveCloudIntegrationAccountCreateAttributes {
	this := CoreWeaveCloudIntegrationAccountCreateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value.
func (o *CoreWeaveCloudIntegrationAccountCreateAttributes) GetAuthentication() CoreWeaveCloudIntegrationAccountAuthenticationRequest {
	if o == nil {
		var ret CoreWeaveCloudIntegrationAccountAuthenticationRequest
		return ret
	}
	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *CoreWeaveCloudIntegrationAccountCreateAttributes) GetAuthenticationOk() (*CoreWeaveCloudIntegrationAccountAuthenticationRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value.
func (o *CoreWeaveCloudIntegrationAccountCreateAttributes) SetAuthentication(v CoreWeaveCloudIntegrationAccountAuthenticationRequest) {
	o.Authentication = v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *CoreWeaveCloudIntegrationAccountCreateAttributes) GetDataflows() CoreWeaveCloudIntegrationDataflowsRequest {
	if o == nil || o.Dataflows == nil {
		var ret CoreWeaveCloudIntegrationDataflowsRequest
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreWeaveCloudIntegrationAccountCreateAttributes) GetDataflowsOk() (*CoreWeaveCloudIntegrationDataflowsRequest, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *CoreWeaveCloudIntegrationAccountCreateAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given CoreWeaveCloudIntegrationDataflowsRequest and assigns it to the Dataflows field.
func (o *CoreWeaveCloudIntegrationAccountCreateAttributes) SetDataflows(v CoreWeaveCloudIntegrationDataflowsRequest) {
	o.Dataflows = &v
}

// GetName returns the Name field value.
func (o *CoreWeaveCloudIntegrationAccountCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CoreWeaveCloudIntegrationAccountCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CoreWeaveCloudIntegrationAccountCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value.
func (o *CoreWeaveCloudIntegrationAccountCreateAttributes) GetSettings() CoreWeaveCloudIntegrationAccountSettingsRequest {
	if o == nil {
		var ret CoreWeaveCloudIntegrationAccountSettingsRequest
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *CoreWeaveCloudIntegrationAccountCreateAttributes) GetSettingsOk() (*CoreWeaveCloudIntegrationAccountSettingsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *CoreWeaveCloudIntegrationAccountCreateAttributes) SetSettings(v CoreWeaveCloudIntegrationAccountSettingsRequest) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CoreWeaveCloudIntegrationAccountCreateAttributes) MarshalJSON() ([]byte, error) {
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
func (o *CoreWeaveCloudIntegrationAccountCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *CoreWeaveCloudIntegrationAccountAuthenticationRequest `json:"authentication"`
		Dataflows      *CoreWeaveCloudIntegrationDataflowsRequest             `json:"dataflows,omitempty"`
		Name           *string                                                `json:"name"`
		Settings       *CoreWeaveCloudIntegrationAccountSettingsRequest       `json:"settings"`
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
