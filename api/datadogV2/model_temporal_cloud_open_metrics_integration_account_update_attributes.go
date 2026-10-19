// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes Writable attributes used to update a Temporal Cloud OpenMetrics integration account. Every field is optional; only the fields provided are changed.
type TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes struct {
	// Authentication for updating the Temporal Cloud OpenMetrics integration account. Exactly one method is set.
	Authentication *TemporalCloudOpenMetricsIntegrationAccountAuthenticationUpdate `json:"authentication,omitempty"`
	// Human-readable name of the Temporal Cloud OpenMetrics integration account.
	Name *string `json:"name,omitempty"`
	// Settings for updating the Temporal Cloud OpenMetrics integration account. Only the fields provided are changed.
	Settings *TemporalCloudOpenMetricsIntegrationAccountSettingsUpdate `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTemporalCloudOpenMetricsIntegrationAccountUpdateAttributes instantiates a new TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTemporalCloudOpenMetricsIntegrationAccountUpdateAttributes() *TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes {
	this := TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes{}
	return &this
}

// NewTemporalCloudOpenMetricsIntegrationAccountUpdateAttributesWithDefaults instantiates a new TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTemporalCloudOpenMetricsIntegrationAccountUpdateAttributesWithDefaults() *TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes {
	this := TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes) GetAuthentication() TemporalCloudOpenMetricsIntegrationAccountAuthenticationUpdate {
	if o == nil || o.Authentication == nil {
		var ret TemporalCloudOpenMetricsIntegrationAccountAuthenticationUpdate
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes) GetAuthenticationOk() (*TemporalCloudOpenMetricsIntegrationAccountAuthenticationUpdate, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes) HasAuthentication() bool {
	return o != nil && o.Authentication != nil
}

// SetAuthentication gets a reference to the given TemporalCloudOpenMetricsIntegrationAccountAuthenticationUpdate and assigns it to the Authentication field.
func (o *TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes) SetAuthentication(v TemporalCloudOpenMetricsIntegrationAccountAuthenticationUpdate) {
	o.Authentication = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes) GetSettings() TemporalCloudOpenMetricsIntegrationAccountSettingsUpdate {
	if o == nil || o.Settings == nil {
		var ret TemporalCloudOpenMetricsIntegrationAccountSettingsUpdate
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes) GetSettingsOk() (*TemporalCloudOpenMetricsIntegrationAccountSettingsUpdate, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given TemporalCloudOpenMetricsIntegrationAccountSettingsUpdate and assigns it to the Settings field.
func (o *TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes) SetSettings(v TemporalCloudOpenMetricsIntegrationAccountSettingsUpdate) {
	o.Settings = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes) MarshalJSON() ([]byte, error) {
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
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TemporalCloudOpenMetricsIntegrationAccountUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *TemporalCloudOpenMetricsIntegrationAccountAuthenticationUpdate `json:"authentication,omitempty"`
		Name           *string                                                         `json:"name,omitempty"`
		Settings       *TemporalCloudOpenMetricsIntegrationAccountSettingsUpdate       `json:"settings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"authentication", "name", "settings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Authentication = all.Authentication
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
