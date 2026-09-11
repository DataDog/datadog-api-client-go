// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeIntegrationAccountCreateAttributes Writable attributes used to create a Snowflake integration account.
type SnowflakeIntegrationAccountCreateAttributes struct {
	// RSA key pair authentication, the only method Snowflake integration accounts support. Generate an RSA key pair and assign the public key to the Snowflake user named in `settings.username`. Because an update replaces this object as a whole, every required field must be sent again on each update, even when only one of them is changing.
	Authentication SnowflakeIntegrationAccountAuthenticationRequest `json:"authentication"`
	// Data Datadog collects from Snowflake, keyed by dataflow id. Each dataflow turns on a distinct kind of collection: set `enabled` to start or stop it, and use `settings` to configure what it collects. Defaults listed on each dataflow apply when the account is created; on update, omitted fields keep their current values. Every dataflow reads from Snowflake as the user in `settings.username`, so that user's role must be granted access to the underlying views; a dataflow enabled without those grants is stored but collects no data.
	Dataflows *SnowflakeIntegrationDataflowsRequest `json:"dataflows,omitempty"`
	// Human-readable name of the Snowflake integration account. Must be
	// unique within your Datadog organization.
	Name string `json:"name"`
	// Settings for creating the Snowflake integration account.
	Settings SnowflakeIntegrationAccountSettingsRequest `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSnowflakeIntegrationAccountCreateAttributes instantiates a new SnowflakeIntegrationAccountCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeIntegrationAccountCreateAttributes(authentication SnowflakeIntegrationAccountAuthenticationRequest, name string, settings SnowflakeIntegrationAccountSettingsRequest) *SnowflakeIntegrationAccountCreateAttributes {
	this := SnowflakeIntegrationAccountCreateAttributes{}
	this.Authentication = authentication
	this.Name = name
	this.Settings = settings
	return &this
}

// NewSnowflakeIntegrationAccountCreateAttributesWithDefaults instantiates a new SnowflakeIntegrationAccountCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeIntegrationAccountCreateAttributesWithDefaults() *SnowflakeIntegrationAccountCreateAttributes {
	this := SnowflakeIntegrationAccountCreateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value.
func (o *SnowflakeIntegrationAccountCreateAttributes) GetAuthentication() SnowflakeIntegrationAccountAuthenticationRequest {
	if o == nil {
		var ret SnowflakeIntegrationAccountAuthenticationRequest
		return ret
	}
	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountCreateAttributes) GetAuthenticationOk() (*SnowflakeIntegrationAccountAuthenticationRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value.
func (o *SnowflakeIntegrationAccountCreateAttributes) SetAuthentication(v SnowflakeIntegrationAccountAuthenticationRequest) {
	o.Authentication = v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *SnowflakeIntegrationAccountCreateAttributes) GetDataflows() SnowflakeIntegrationDataflowsRequest {
	if o == nil || o.Dataflows == nil {
		var ret SnowflakeIntegrationDataflowsRequest
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountCreateAttributes) GetDataflowsOk() (*SnowflakeIntegrationDataflowsRequest, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *SnowflakeIntegrationAccountCreateAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given SnowflakeIntegrationDataflowsRequest and assigns it to the Dataflows field.
func (o *SnowflakeIntegrationAccountCreateAttributes) SetDataflows(v SnowflakeIntegrationDataflowsRequest) {
	o.Dataflows = &v
}

// GetName returns the Name field value.
func (o *SnowflakeIntegrationAccountCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SnowflakeIntegrationAccountCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value.
func (o *SnowflakeIntegrationAccountCreateAttributes) GetSettings() SnowflakeIntegrationAccountSettingsRequest {
	if o == nil {
		var ret SnowflakeIntegrationAccountSettingsRequest
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountCreateAttributes) GetSettingsOk() (*SnowflakeIntegrationAccountSettingsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *SnowflakeIntegrationAccountCreateAttributes) SetSettings(v SnowflakeIntegrationAccountSettingsRequest) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeIntegrationAccountCreateAttributes) MarshalJSON() ([]byte, error) {
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
func (o *SnowflakeIntegrationAccountCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *SnowflakeIntegrationAccountAuthenticationRequest `json:"authentication"`
		Dataflows      *SnowflakeIntegrationDataflowsRequest             `json:"dataflows,omitempty"`
		Name           *string                                           `json:"name"`
		Settings       *SnowflakeIntegrationAccountSettingsRequest       `json:"settings"`
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
	if all.Authentication.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
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
