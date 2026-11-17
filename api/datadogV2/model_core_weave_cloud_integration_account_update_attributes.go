// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CoreWeaveCloudIntegrationAccountUpdateAttributes Writable attributes used to update a CoreWeave integration account. Every field is optional; only the fields provided are changed. When `dataflows` is provided, only the dataflow ids included in the request are modified; dataflows omitted from the map keep their current configuration.
type CoreWeaveCloudIntegrationAccountUpdateAttributes struct {
	// Authentication for updating the CoreWeave integration account. Exactly one method is set.
	Authentication *CoreWeaveCloudIntegrationAccountAuthenticationUpdate `json:"authentication,omitempty"`
	// Data Datadog collects from CoreWeave, keyed by dataflow id. Each dataflow turns on a distinct kind of collection: set `enabled` to start or stop it. Defaults listed on each dataflow apply when the account is created; on update, omitted dataflow ids keep their current values. Collection runs with the API access token in `authentication`, whose user must have the Observability Viewer role; a dataflow enabled without that role is stored but collects no data.
	Dataflows *CoreWeaveCloudIntegrationDataflowsRequest `json:"dataflows,omitempty"`
	// Human-readable name of the CoreWeave integration account.
	Name *string `json:"name,omitempty"`
	// Settings for updating the CoreWeave integration account. Only the fields provided are changed.
	Settings *CoreWeaveCloudIntegrationAccountSettingsUpdate `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCoreWeaveCloudIntegrationAccountUpdateAttributes instantiates a new CoreWeaveCloudIntegrationAccountUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCoreWeaveCloudIntegrationAccountUpdateAttributes() *CoreWeaveCloudIntegrationAccountUpdateAttributes {
	this := CoreWeaveCloudIntegrationAccountUpdateAttributes{}
	return &this
}

// NewCoreWeaveCloudIntegrationAccountUpdateAttributesWithDefaults instantiates a new CoreWeaveCloudIntegrationAccountUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCoreWeaveCloudIntegrationAccountUpdateAttributesWithDefaults() *CoreWeaveCloudIntegrationAccountUpdateAttributes {
	this := CoreWeaveCloudIntegrationAccountUpdateAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *CoreWeaveCloudIntegrationAccountUpdateAttributes) GetAuthentication() CoreWeaveCloudIntegrationAccountAuthenticationUpdate {
	if o == nil || o.Authentication == nil {
		var ret CoreWeaveCloudIntegrationAccountAuthenticationUpdate
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreWeaveCloudIntegrationAccountUpdateAttributes) GetAuthenticationOk() (*CoreWeaveCloudIntegrationAccountAuthenticationUpdate, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *CoreWeaveCloudIntegrationAccountUpdateAttributes) HasAuthentication() bool {
	return o != nil && o.Authentication != nil
}

// SetAuthentication gets a reference to the given CoreWeaveCloudIntegrationAccountAuthenticationUpdate and assigns it to the Authentication field.
func (o *CoreWeaveCloudIntegrationAccountUpdateAttributes) SetAuthentication(v CoreWeaveCloudIntegrationAccountAuthenticationUpdate) {
	o.Authentication = &v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *CoreWeaveCloudIntegrationAccountUpdateAttributes) GetDataflows() CoreWeaveCloudIntegrationDataflowsRequest {
	if o == nil || o.Dataflows == nil {
		var ret CoreWeaveCloudIntegrationDataflowsRequest
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreWeaveCloudIntegrationAccountUpdateAttributes) GetDataflowsOk() (*CoreWeaveCloudIntegrationDataflowsRequest, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *CoreWeaveCloudIntegrationAccountUpdateAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given CoreWeaveCloudIntegrationDataflowsRequest and assigns it to the Dataflows field.
func (o *CoreWeaveCloudIntegrationAccountUpdateAttributes) SetDataflows(v CoreWeaveCloudIntegrationDataflowsRequest) {
	o.Dataflows = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreWeaveCloudIntegrationAccountUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreWeaveCloudIntegrationAccountUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreWeaveCloudIntegrationAccountUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreWeaveCloudIntegrationAccountUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *CoreWeaveCloudIntegrationAccountUpdateAttributes) GetSettings() CoreWeaveCloudIntegrationAccountSettingsUpdate {
	if o == nil || o.Settings == nil {
		var ret CoreWeaveCloudIntegrationAccountSettingsUpdate
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreWeaveCloudIntegrationAccountUpdateAttributes) GetSettingsOk() (*CoreWeaveCloudIntegrationAccountSettingsUpdate, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *CoreWeaveCloudIntegrationAccountUpdateAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given CoreWeaveCloudIntegrationAccountSettingsUpdate and assigns it to the Settings field.
func (o *CoreWeaveCloudIntegrationAccountUpdateAttributes) SetSettings(v CoreWeaveCloudIntegrationAccountSettingsUpdate) {
	o.Settings = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CoreWeaveCloudIntegrationAccountUpdateAttributes) MarshalJSON() ([]byte, error) {
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
func (o *CoreWeaveCloudIntegrationAccountUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *CoreWeaveCloudIntegrationAccountAuthenticationUpdate `json:"authentication,omitempty"`
		Dataflows      *CoreWeaveCloudIntegrationDataflowsRequest            `json:"dataflows,omitempty"`
		Name           *string                                               `json:"name,omitempty"`
		Settings       *CoreWeaveCloudIntegrationAccountSettingsUpdate       `json:"settings,omitempty"`
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
