// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TemporalCloudOpenMetricsIntegrationAccountCreateAttributes Writable attributes used to create a Temporal Cloud OpenMetrics integration account.
type TemporalCloudOpenMetricsIntegrationAccountCreateAttributes struct {
	// Authentication for creating the Temporal Cloud OpenMetrics integration account. Exactly one method is set.
	Authentication TemporalCloudOpenMetricsIntegrationAccountAuthenticationRequest `json:"authentication"`
	// Human-readable name of the Temporal Cloud OpenMetrics integration account.
	Name string `json:"name"`
	// Settings for creating the Temporal Cloud OpenMetrics integration account.
	Settings TemporalCloudOpenMetricsIntegrationAccountSettingsRequest `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTemporalCloudOpenMetricsIntegrationAccountCreateAttributes instantiates a new TemporalCloudOpenMetricsIntegrationAccountCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTemporalCloudOpenMetricsIntegrationAccountCreateAttributes(authentication TemporalCloudOpenMetricsIntegrationAccountAuthenticationRequest, name string, settings TemporalCloudOpenMetricsIntegrationAccountSettingsRequest) *TemporalCloudOpenMetricsIntegrationAccountCreateAttributes {
	this := TemporalCloudOpenMetricsIntegrationAccountCreateAttributes{}
	this.Authentication = authentication
	this.Name = name
	this.Settings = settings
	return &this
}

// NewTemporalCloudOpenMetricsIntegrationAccountCreateAttributesWithDefaults instantiates a new TemporalCloudOpenMetricsIntegrationAccountCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTemporalCloudOpenMetricsIntegrationAccountCreateAttributesWithDefaults() *TemporalCloudOpenMetricsIntegrationAccountCreateAttributes {
	this := TemporalCloudOpenMetricsIntegrationAccountCreateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value.
func (o *TemporalCloudOpenMetricsIntegrationAccountCreateAttributes) GetAuthentication() TemporalCloudOpenMetricsIntegrationAccountAuthenticationRequest {
	if o == nil {
		var ret TemporalCloudOpenMetricsIntegrationAccountAuthenticationRequest
		return ret
	}
	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountCreateAttributes) GetAuthenticationOk() (*TemporalCloudOpenMetricsIntegrationAccountAuthenticationRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value.
func (o *TemporalCloudOpenMetricsIntegrationAccountCreateAttributes) SetAuthentication(v TemporalCloudOpenMetricsIntegrationAccountAuthenticationRequest) {
	o.Authentication = v
}

// GetName returns the Name field value.
func (o *TemporalCloudOpenMetricsIntegrationAccountCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *TemporalCloudOpenMetricsIntegrationAccountCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value.
func (o *TemporalCloudOpenMetricsIntegrationAccountCreateAttributes) GetSettings() TemporalCloudOpenMetricsIntegrationAccountSettingsRequest {
	if o == nil {
		var ret TemporalCloudOpenMetricsIntegrationAccountSettingsRequest
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountCreateAttributes) GetSettingsOk() (*TemporalCloudOpenMetricsIntegrationAccountSettingsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *TemporalCloudOpenMetricsIntegrationAccountCreateAttributes) SetSettings(v TemporalCloudOpenMetricsIntegrationAccountSettingsRequest) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TemporalCloudOpenMetricsIntegrationAccountCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["authentication"] = o.Authentication
	toSerialize["name"] = o.Name
	toSerialize["settings"] = o.Settings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TemporalCloudOpenMetricsIntegrationAccountCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *TemporalCloudOpenMetricsIntegrationAccountAuthenticationRequest `json:"authentication"`
		Name           *string                                                          `json:"name"`
		Settings       *TemporalCloudOpenMetricsIntegrationAccountSettingsRequest       `json:"settings"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"authentication", "name", "settings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Authentication = *all.Authentication
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
