// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetIntegrationsByStatusV2 Integrations organized by their status.
type FleetIntegrationsByStatusV2 struct {
	// Configuration files for integrations.
	ConfigurationFiles []FleetConfigurationFileV2 `json:"configuration_files,omitempty"`
	// Integrations with errors.
	ErrorIntegrations []FleetIntegrationDetailsV2 `json:"error_integrations,omitempty"`
	// Detected but not configured integrations.
	MissingIntegrations []FleetDetectedIntegration `json:"missing_integrations,omitempty"`
	// Integrations with warnings.
	WarningIntegrations []FleetIntegrationDetailsV2 `json:"warning_integrations,omitempty"`
	// Integrations that are working correctly.
	WorkingIntegrations []FleetIntegrationDetailsV2 `json:"working_integrations,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetIntegrationsByStatusV2 instantiates a new FleetIntegrationsByStatusV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetIntegrationsByStatusV2() *FleetIntegrationsByStatusV2 {
	this := FleetIntegrationsByStatusV2{}
	return &this
}

// NewFleetIntegrationsByStatusV2WithDefaults instantiates a new FleetIntegrationsByStatusV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetIntegrationsByStatusV2WithDefaults() *FleetIntegrationsByStatusV2 {
	this := FleetIntegrationsByStatusV2{}
	return &this
}

// GetConfigurationFiles returns the ConfigurationFiles field value if set, zero value otherwise.
func (o *FleetIntegrationsByStatusV2) GetConfigurationFiles() []FleetConfigurationFileV2 {
	if o == nil || o.ConfigurationFiles == nil {
		var ret []FleetConfigurationFileV2
		return ret
	}
	return o.ConfigurationFiles
}

// GetConfigurationFilesOk returns a tuple with the ConfigurationFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationsByStatusV2) GetConfigurationFilesOk() (*[]FleetConfigurationFileV2, bool) {
	if o == nil || o.ConfigurationFiles == nil {
		return nil, false
	}
	return &o.ConfigurationFiles, true
}

// HasConfigurationFiles returns a boolean if a field has been set.
func (o *FleetIntegrationsByStatusV2) HasConfigurationFiles() bool {
	return o != nil && o.ConfigurationFiles != nil
}

// SetConfigurationFiles gets a reference to the given []FleetConfigurationFileV2 and assigns it to the ConfigurationFiles field.
func (o *FleetIntegrationsByStatusV2) SetConfigurationFiles(v []FleetConfigurationFileV2) {
	o.ConfigurationFiles = v
}

// GetErrorIntegrations returns the ErrorIntegrations field value if set, zero value otherwise.
func (o *FleetIntegrationsByStatusV2) GetErrorIntegrations() []FleetIntegrationDetailsV2 {
	if o == nil || o.ErrorIntegrations == nil {
		var ret []FleetIntegrationDetailsV2
		return ret
	}
	return o.ErrorIntegrations
}

// GetErrorIntegrationsOk returns a tuple with the ErrorIntegrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationsByStatusV2) GetErrorIntegrationsOk() (*[]FleetIntegrationDetailsV2, bool) {
	if o == nil || o.ErrorIntegrations == nil {
		return nil, false
	}
	return &o.ErrorIntegrations, true
}

// HasErrorIntegrations returns a boolean if a field has been set.
func (o *FleetIntegrationsByStatusV2) HasErrorIntegrations() bool {
	return o != nil && o.ErrorIntegrations != nil
}

// SetErrorIntegrations gets a reference to the given []FleetIntegrationDetailsV2 and assigns it to the ErrorIntegrations field.
func (o *FleetIntegrationsByStatusV2) SetErrorIntegrations(v []FleetIntegrationDetailsV2) {
	o.ErrorIntegrations = v
}

// GetMissingIntegrations returns the MissingIntegrations field value if set, zero value otherwise.
func (o *FleetIntegrationsByStatusV2) GetMissingIntegrations() []FleetDetectedIntegration {
	if o == nil || o.MissingIntegrations == nil {
		var ret []FleetDetectedIntegration
		return ret
	}
	return o.MissingIntegrations
}

// GetMissingIntegrationsOk returns a tuple with the MissingIntegrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationsByStatusV2) GetMissingIntegrationsOk() (*[]FleetDetectedIntegration, bool) {
	if o == nil || o.MissingIntegrations == nil {
		return nil, false
	}
	return &o.MissingIntegrations, true
}

// HasMissingIntegrations returns a boolean if a field has been set.
func (o *FleetIntegrationsByStatusV2) HasMissingIntegrations() bool {
	return o != nil && o.MissingIntegrations != nil
}

// SetMissingIntegrations gets a reference to the given []FleetDetectedIntegration and assigns it to the MissingIntegrations field.
func (o *FleetIntegrationsByStatusV2) SetMissingIntegrations(v []FleetDetectedIntegration) {
	o.MissingIntegrations = v
}

// GetWarningIntegrations returns the WarningIntegrations field value if set, zero value otherwise.
func (o *FleetIntegrationsByStatusV2) GetWarningIntegrations() []FleetIntegrationDetailsV2 {
	if o == nil || o.WarningIntegrations == nil {
		var ret []FleetIntegrationDetailsV2
		return ret
	}
	return o.WarningIntegrations
}

// GetWarningIntegrationsOk returns a tuple with the WarningIntegrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationsByStatusV2) GetWarningIntegrationsOk() (*[]FleetIntegrationDetailsV2, bool) {
	if o == nil || o.WarningIntegrations == nil {
		return nil, false
	}
	return &o.WarningIntegrations, true
}

// HasWarningIntegrations returns a boolean if a field has been set.
func (o *FleetIntegrationsByStatusV2) HasWarningIntegrations() bool {
	return o != nil && o.WarningIntegrations != nil
}

// SetWarningIntegrations gets a reference to the given []FleetIntegrationDetailsV2 and assigns it to the WarningIntegrations field.
func (o *FleetIntegrationsByStatusV2) SetWarningIntegrations(v []FleetIntegrationDetailsV2) {
	o.WarningIntegrations = v
}

// GetWorkingIntegrations returns the WorkingIntegrations field value if set, zero value otherwise.
func (o *FleetIntegrationsByStatusV2) GetWorkingIntegrations() []FleetIntegrationDetailsV2 {
	if o == nil || o.WorkingIntegrations == nil {
		var ret []FleetIntegrationDetailsV2
		return ret
	}
	return o.WorkingIntegrations
}

// GetWorkingIntegrationsOk returns a tuple with the WorkingIntegrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationsByStatusV2) GetWorkingIntegrationsOk() (*[]FleetIntegrationDetailsV2, bool) {
	if o == nil || o.WorkingIntegrations == nil {
		return nil, false
	}
	return &o.WorkingIntegrations, true
}

// HasWorkingIntegrations returns a boolean if a field has been set.
func (o *FleetIntegrationsByStatusV2) HasWorkingIntegrations() bool {
	return o != nil && o.WorkingIntegrations != nil
}

// SetWorkingIntegrations gets a reference to the given []FleetIntegrationDetailsV2 and assigns it to the WorkingIntegrations field.
func (o *FleetIntegrationsByStatusV2) SetWorkingIntegrations(v []FleetIntegrationDetailsV2) {
	o.WorkingIntegrations = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetIntegrationsByStatusV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConfigurationFiles != nil {
		toSerialize["configuration_files"] = o.ConfigurationFiles
	}
	if o.ErrorIntegrations != nil {
		toSerialize["error_integrations"] = o.ErrorIntegrations
	}
	if o.MissingIntegrations != nil {
		toSerialize["missing_integrations"] = o.MissingIntegrations
	}
	if o.WarningIntegrations != nil {
		toSerialize["warning_integrations"] = o.WarningIntegrations
	}
	if o.WorkingIntegrations != nil {
		toSerialize["working_integrations"] = o.WorkingIntegrations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetIntegrationsByStatusV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigurationFiles  []FleetConfigurationFileV2  `json:"configuration_files,omitempty"`
		ErrorIntegrations   []FleetIntegrationDetailsV2 `json:"error_integrations,omitempty"`
		MissingIntegrations []FleetDetectedIntegration  `json:"missing_integrations,omitempty"`
		WarningIntegrations []FleetIntegrationDetailsV2 `json:"warning_integrations,omitempty"`
		WorkingIntegrations []FleetIntegrationDetailsV2 `json:"working_integrations,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"configuration_files", "error_integrations", "missing_integrations", "warning_integrations", "working_integrations"})
	} else {
		return err
	}
	o.ConfigurationFiles = all.ConfigurationFiles
	o.ErrorIntegrations = all.ErrorIntegrations
	o.MissingIntegrations = all.MissingIntegrations
	o.WarningIntegrations = all.WarningIntegrations
	o.WorkingIntegrations = all.WorkingIntegrations

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
