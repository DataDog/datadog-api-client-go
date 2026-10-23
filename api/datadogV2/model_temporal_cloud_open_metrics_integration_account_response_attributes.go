// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TemporalCloudOpenMetricsIntegrationAccountResponseAttributes Attributes of a Temporal Cloud OpenMetrics integration account returned in responses.
type TemporalCloudOpenMetricsIntegrationAccountResponseAttributes struct {
	// Authentication configured on the Temporal Cloud OpenMetrics integration account.
	Authentication *TemporalCloudOpenMetricsIntegrationAccountAuthenticationResponse `json:"authentication,omitempty"`
	// Data Datadog collects from Temporal Cloud OpenMetrics, keyed by dataflow id.
	Dataflows *TemporalCloudOpenMetricsIntegrationDataflowsResponse `json:"dataflows,omitempty"`
	// Human-readable name of the Temporal Cloud OpenMetrics integration account.
	Name string `json:"name"`
	// Settings configured on the Temporal Cloud OpenMetrics integration account.
	Settings TemporalCloudOpenMetricsIntegrationAccountSettingsResponse `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTemporalCloudOpenMetricsIntegrationAccountResponseAttributes instantiates a new TemporalCloudOpenMetricsIntegrationAccountResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTemporalCloudOpenMetricsIntegrationAccountResponseAttributes(name string, settings TemporalCloudOpenMetricsIntegrationAccountSettingsResponse) *TemporalCloudOpenMetricsIntegrationAccountResponseAttributes {
	this := TemporalCloudOpenMetricsIntegrationAccountResponseAttributes{}
	this.Name = name
	this.Settings = settings
	return &this
}

// NewTemporalCloudOpenMetricsIntegrationAccountResponseAttributesWithDefaults instantiates a new TemporalCloudOpenMetricsIntegrationAccountResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTemporalCloudOpenMetricsIntegrationAccountResponseAttributesWithDefaults() *TemporalCloudOpenMetricsIntegrationAccountResponseAttributes {
	this := TemporalCloudOpenMetricsIntegrationAccountResponseAttributes{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *TemporalCloudOpenMetricsIntegrationAccountResponseAttributes) GetAuthentication() TemporalCloudOpenMetricsIntegrationAccountAuthenticationResponse {
	if o == nil || o.Authentication == nil {
		var ret TemporalCloudOpenMetricsIntegrationAccountAuthenticationResponse
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountResponseAttributes) GetAuthenticationOk() (*TemporalCloudOpenMetricsIntegrationAccountAuthenticationResponse, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountResponseAttributes) HasAuthentication() bool {
	return o != nil && o.Authentication != nil
}

// SetAuthentication gets a reference to the given TemporalCloudOpenMetricsIntegrationAccountAuthenticationResponse and assigns it to the Authentication field.
func (o *TemporalCloudOpenMetricsIntegrationAccountResponseAttributes) SetAuthentication(v TemporalCloudOpenMetricsIntegrationAccountAuthenticationResponse) {
	o.Authentication = &v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *TemporalCloudOpenMetricsIntegrationAccountResponseAttributes) GetDataflows() TemporalCloudOpenMetricsIntegrationDataflowsResponse {
	if o == nil || o.Dataflows == nil {
		var ret TemporalCloudOpenMetricsIntegrationDataflowsResponse
		return ret
	}
	return *o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountResponseAttributes) GetDataflowsOk() (*TemporalCloudOpenMetricsIntegrationDataflowsResponse, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountResponseAttributes) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given TemporalCloudOpenMetricsIntegrationDataflowsResponse and assigns it to the Dataflows field.
func (o *TemporalCloudOpenMetricsIntegrationAccountResponseAttributes) SetDataflows(v TemporalCloudOpenMetricsIntegrationDataflowsResponse) {
	o.Dataflows = &v
}

// GetName returns the Name field value.
func (o *TemporalCloudOpenMetricsIntegrationAccountResponseAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *TemporalCloudOpenMetricsIntegrationAccountResponseAttributes) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value.
func (o *TemporalCloudOpenMetricsIntegrationAccountResponseAttributes) GetSettings() TemporalCloudOpenMetricsIntegrationAccountSettingsResponse {
	if o == nil {
		var ret TemporalCloudOpenMetricsIntegrationAccountSettingsResponse
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *TemporalCloudOpenMetricsIntegrationAccountResponseAttributes) GetSettingsOk() (*TemporalCloudOpenMetricsIntegrationAccountSettingsResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *TemporalCloudOpenMetricsIntegrationAccountResponseAttributes) SetSettings(v TemporalCloudOpenMetricsIntegrationAccountSettingsResponse) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TemporalCloudOpenMetricsIntegrationAccountResponseAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["name"] = o.Name
	toSerialize["settings"] = o.Settings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TemporalCloudOpenMetricsIntegrationAccountResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *TemporalCloudOpenMetricsIntegrationAccountAuthenticationResponse `json:"authentication,omitempty"`
		Dataflows      *TemporalCloudOpenMetricsIntegrationDataflowsResponse             `json:"dataflows,omitempty"`
		Name           *string                                                           `json:"name"`
		Settings       *TemporalCloudOpenMetricsIntegrationAccountSettingsResponse       `json:"settings"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	o.Authentication = all.Authentication
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
