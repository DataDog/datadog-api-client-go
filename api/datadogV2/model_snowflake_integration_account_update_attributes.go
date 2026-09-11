// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeIntegrationAccountUpdateAttributes Writable attributes used to update a Snowflake integration account. Every field is optional; only the fields provided are changed. When `dataflows` is provided, only the dataflow ids included in the request are modified; dataflows omitted from the map keep their current configuration, as do the settings of an included dataflow that provides only `enabled`. `authentication` is the exception to partial updates: when provided it is replaced as a whole, so it must carry every field that creating an account requires.
type SnowflakeIntegrationAccountUpdateAttributes struct {
	// RSA key pair authentication, the only method Snowflake integration accounts support. Generate an RSA key pair and assign the public key to the Snowflake user named in `settings.username`. Because an update replaces this object as a whole, every required field must be sent again on each update, even when only one of them is changing.
	Authentication *SnowflakeIntegrationAccountAuthenticationRequest `json:"authentication,omitempty"`
	// Data Datadog collects from Snowflake, keyed by dataflow id. Each dataflow turns on a distinct kind of collection: set `enabled` to start or stop it, and use `settings` to configure what it collects. Defaults listed on each dataflow apply when the account is created; on update, omitted fields keep their current values. Every dataflow reads from Snowflake as the user in `settings.username`, so that user's role must be granted access to the underlying views; a dataflow enabled without those grants is stored but collects no data.
	Dataflows *SnowflakeIntegrationDataflowsRequest `json:"dataflows,omitempty"`
	// Human-readable name of the Snowflake integration account. Must be
	// unique within your Datadog organization.
	Name *string `json:"name,omitempty"`
	// Settings for updating the Snowflake integration account. Only the fields provided are changed.
	Settings *SnowflakeIntegrationAccountSettingsUpdate `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSnowflakeIntegrationAccountUpdateAttributes instantiates a new SnowflakeIntegrationAccountUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeIntegrationAccountUpdateAttributes() *SnowflakeIntegrationAccountUpdateAttributes {
	this := SnowflakeIntegrationAccountUpdateAttributes{}
	return &this
}

// NewSnowflakeIntegrationAccountUpdateAttributesWithDefaults instantiates a new SnowflakeIntegrationAccountUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeIntegrationAccountUpdateAttributesWithDefaults() *SnowflakeIntegrationAccountUpdateAttributes {
	this := SnowflakeIntegrationAccountUpdateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *SnowflakeIntegrationAccountUpdateAttributes) GetAuthentication() SnowflakeIntegrationAccountAuthenticationRequest {
	if o == nil || o.Authentication == nil {
		var ret SnowflakeIntegrationAccountAuthenticationRequest
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountUpdateAttributes) GetAuthenticationOk() (*SnowflakeIntegrationAccountAuthenticationRequest, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *SnowflakeIntegrationAccountUpdateAttributes) HasAuthentication() bool {
	return o != nil && o.Authentication != nil
}

// SetAuthentication gets a reference to the given SnowflakeIntegrationAccountAuthenticationRequest and assigns it to the Authentication field.
func (o *SnowflakeIntegrationAccountUpdateAttributes) SetAuthentication(v SnowflakeIntegrationAccountAuthenticationRequest) {
	o.Authentication = &v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *SnowflakeIntegrationAccountUpdateAttributes) GetDataflows() SnowflakeIntegrationDataflowsRequest {
	if o == nil || o.Dataflows == nil {
		var ret SnowflakeIntegrationDataflowsRequest
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountUpdateAttributes) GetDataflowsOk() (*SnowflakeIntegrationDataflowsRequest, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *SnowflakeIntegrationAccountUpdateAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given SnowflakeIntegrationDataflowsRequest and assigns it to the Dataflows field.
func (o *SnowflakeIntegrationAccountUpdateAttributes) SetDataflows(v SnowflakeIntegrationDataflowsRequest) {
	o.Dataflows = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SnowflakeIntegrationAccountUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SnowflakeIntegrationAccountUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SnowflakeIntegrationAccountUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *SnowflakeIntegrationAccountUpdateAttributes) GetSettings() SnowflakeIntegrationAccountSettingsUpdate {
	if o == nil || o.Settings == nil {
		var ret SnowflakeIntegrationAccountSettingsUpdate
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationAccountUpdateAttributes) GetSettingsOk() (*SnowflakeIntegrationAccountSettingsUpdate, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *SnowflakeIntegrationAccountUpdateAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given SnowflakeIntegrationAccountSettingsUpdate and assigns it to the Settings field.
func (o *SnowflakeIntegrationAccountUpdateAttributes) SetSettings(v SnowflakeIntegrationAccountSettingsUpdate) {
	o.Settings = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeIntegrationAccountUpdateAttributes) MarshalJSON() ([]byte, error) {
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
func (o *SnowflakeIntegrationAccountUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *SnowflakeIntegrationAccountAuthenticationRequest `json:"authentication,omitempty"`
		Dataflows      *SnowflakeIntegrationDataflowsRequest             `json:"dataflows,omitempty"`
		Name           *string                                           `json:"name,omitempty"`
		Settings       *SnowflakeIntegrationAccountSettingsUpdate        `json:"settings,omitempty"`
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
	if all.Authentication != nil && all.Authentication.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
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
