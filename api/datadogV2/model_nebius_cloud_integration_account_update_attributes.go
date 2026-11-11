// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudIntegrationAccountUpdateAttributes Writable attributes used to update a Nebius Cloud integration account. Every field is optional; only the fields provided are changed. When `dataflows` is provided, only the dataflow ids included in the request are modified; dataflows omitted from the map keep their current configuration, as do the settings of an included dataflow that provides only `enabled`.
type NebiusCloudIntegrationAccountUpdateAttributes struct {
	// Authentication for updating the Nebius Cloud integration account. Exactly one method is set.
	Authentication *NebiusCloudIntegrationAccountAuthenticationUpdate `json:"authentication,omitempty"`
	// Data Datadog collects from Nebius Cloud, keyed by dataflow id. Each dataflow turns on a distinct kind of collection: set `enabled` to stop or start it, and use `settings` to configure what it collects. Defaults listed on each dataflow apply when the account is created; on update, omitted fields keep their current values. Every dataflow except `nebius-cloud-cost-collection` runs as the service account in `authentication`, using its authorized key pair and the static API key; enabling one of them without those keys is rejected with a `422` response, and one enabled with keys that later expire is stored but collects no data.
	Dataflows *NebiusCloudIntegrationDataflowsRequest `json:"dataflows,omitempty"`
	// Human-readable name of the Nebius Cloud integration account.
	Name *string `json:"name,omitempty"`
	// Settings for updating the Nebius Cloud integration account. Only the fields provided are changed.
	Settings *NebiusCloudIntegrationAccountSettingsUpdate `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNebiusCloudIntegrationAccountUpdateAttributes instantiates a new NebiusCloudIntegrationAccountUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNebiusCloudIntegrationAccountUpdateAttributes() *NebiusCloudIntegrationAccountUpdateAttributes {
	this := NebiusCloudIntegrationAccountUpdateAttributes{}
	return &this
}

// NewNebiusCloudIntegrationAccountUpdateAttributesWithDefaults instantiates a new NebiusCloudIntegrationAccountUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNebiusCloudIntegrationAccountUpdateAttributesWithDefaults() *NebiusCloudIntegrationAccountUpdateAttributes {
	this := NebiusCloudIntegrationAccountUpdateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountUpdateAttributes) GetAuthentication() NebiusCloudIntegrationAccountAuthenticationUpdate {
	if o == nil || o.Authentication == nil {
		var ret NebiusCloudIntegrationAccountAuthenticationUpdate
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountUpdateAttributes) GetAuthenticationOk() (*NebiusCloudIntegrationAccountAuthenticationUpdate, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountUpdateAttributes) HasAuthentication() bool {
	return o != nil && o.Authentication != nil
}

// SetAuthentication gets a reference to the given NebiusCloudIntegrationAccountAuthenticationUpdate and assigns it to the Authentication field.
func (o *NebiusCloudIntegrationAccountUpdateAttributes) SetAuthentication(v NebiusCloudIntegrationAccountAuthenticationUpdate) {
	o.Authentication = &v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountUpdateAttributes) GetDataflows() NebiusCloudIntegrationDataflowsRequest {
	if o == nil || o.Dataflows == nil {
		var ret NebiusCloudIntegrationDataflowsRequest
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountUpdateAttributes) GetDataflowsOk() (*NebiusCloudIntegrationDataflowsRequest, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountUpdateAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given NebiusCloudIntegrationDataflowsRequest and assigns it to the Dataflows field.
func (o *NebiusCloudIntegrationAccountUpdateAttributes) SetDataflows(v NebiusCloudIntegrationDataflowsRequest) {
	o.Dataflows = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NebiusCloudIntegrationAccountUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationAccountUpdateAttributes) GetSettings() NebiusCloudIntegrationAccountSettingsUpdate {
	if o == nil || o.Settings == nil {
		var ret NebiusCloudIntegrationAccountSettingsUpdate
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationAccountUpdateAttributes) GetSettingsOk() (*NebiusCloudIntegrationAccountSettingsUpdate, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationAccountUpdateAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given NebiusCloudIntegrationAccountSettingsUpdate and assigns it to the Settings field.
func (o *NebiusCloudIntegrationAccountUpdateAttributes) SetSettings(v NebiusCloudIntegrationAccountSettingsUpdate) {
	o.Settings = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NebiusCloudIntegrationAccountUpdateAttributes) MarshalJSON() ([]byte, error) {
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
func (o *NebiusCloudIntegrationAccountUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *NebiusCloudIntegrationAccountAuthenticationUpdate `json:"authentication,omitempty"`
		Dataflows      *NebiusCloudIntegrationDataflowsRequest            `json:"dataflows,omitempty"`
		Name           *string                                            `json:"name,omitempty"`
		Settings       *NebiusCloudIntegrationAccountSettingsUpdate       `json:"settings,omitempty"`
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
