// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeIntegrationAccountResponseAttributes Attributes of a Snowflake integration account returned in responses.
type SnowflakeIntegrationAccountResponseAttributes struct {
	// Authentication configured on the Snowflake integration account.
	Authentication SnowflakeIntegrationAccountAuthenticationResponse `json:"authentication"`
	// Data Datadog collects from Snowflake, keyed by dataflow id.
	Dataflows SnowflakeIntegrationDataflowsResponse `json:"dataflows"`
	// Human-readable name of the Snowflake integration account.
	Name string `json:"name"`
	// Settings configured on the Snowflake integration account.
	Settings SnowflakeIntegrationAccountSettingsResponse `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSnowflakeIntegrationAccountResponseAttributes instantiates a new SnowflakeIntegrationAccountResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeIntegrationAccountResponseAttributes(authentication SnowflakeIntegrationAccountAuthenticationResponse, dataflows SnowflakeIntegrationDataflowsResponse, name string, settings SnowflakeIntegrationAccountSettingsResponse) *SnowflakeIntegrationAccountResponseAttributes {
	this := SnowflakeIntegrationAccountResponseAttributes{}
	this.Authentication = authentication
	this.Dataflows = dataflows
	this.Name = name
	this.Settings = settings
	return &this
}

// NewSnowflakeIntegrationAccountResponseAttributesWithDefaults instantiates a new SnowflakeIntegrationAccountResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeIntegrationAccountResponseAttributesWithDefaults() *SnowflakeIntegrationAccountResponseAttributes {
	this := SnowflakeIntegrationAccountResponseAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value.
func (o *SnowflakeIntegrationAccountResponseAttributes) GetAuthentication() SnowflakeIntegrationAccountAuthenticationResponse {
	if o == nil {
		var ret SnowflakeIntegrationAccountAuthenticationResponse
		return ret
	}
	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountResponseAttributes) GetAuthenticationOk() (*SnowflakeIntegrationAccountAuthenticationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value.
func (o *SnowflakeIntegrationAccountResponseAttributes) SetAuthentication(v SnowflakeIntegrationAccountAuthenticationResponse) {
	o.Authentication = v
}

// GetDataflows returns the Dataflows field value.
func (o *SnowflakeIntegrationAccountResponseAttributes) GetDataflows() SnowflakeIntegrationDataflowsResponse {
	if o == nil {
		var ret SnowflakeIntegrationDataflowsResponse
		return ret
	}
	return o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountResponseAttributes) GetDataflowsOk() (*SnowflakeIntegrationDataflowsResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dataflows, true
}

// SetDataflows sets field value.
func (o *SnowflakeIntegrationAccountResponseAttributes) SetDataflows(v SnowflakeIntegrationDataflowsResponse) {
	o.Dataflows = v
}

// GetName returns the Name field value.
func (o *SnowflakeIntegrationAccountResponseAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SnowflakeIntegrationAccountResponseAttributes) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value.
func (o *SnowflakeIntegrationAccountResponseAttributes) GetSettings() SnowflakeIntegrationAccountSettingsResponse {
	if o == nil {
		var ret SnowflakeIntegrationAccountSettingsResponse
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountResponseAttributes) GetSettingsOk() (*SnowflakeIntegrationAccountSettingsResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *SnowflakeIntegrationAccountResponseAttributes) SetSettings(v SnowflakeIntegrationAccountSettingsResponse) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeIntegrationAccountResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["authentication"] = o.Authentication
	toSerialize["dataflows"] = o.Dataflows
	toSerialize["name"] = o.Name
	toSerialize["settings"] = o.Settings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeIntegrationAccountResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *SnowflakeIntegrationAccountAuthenticationResponse `json:"authentication"`
		Dataflows      *SnowflakeIntegrationDataflowsResponse             `json:"dataflows"`
		Name           *string                                            `json:"name"`
		Settings       *SnowflakeIntegrationAccountSettingsResponse       `json:"settings"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Authentication == nil {
		return fmt.Errorf("required field authentication missing")
	}
	if all.Dataflows == nil {
		return fmt.Errorf("required field dataflows missing")
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
	if all.Dataflows.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Dataflows = *all.Dataflows
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
