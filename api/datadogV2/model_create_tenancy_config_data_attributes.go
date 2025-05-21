// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateTenancyConfigDataAttributes The definition of `CreateTenancyConfigDataAttributes` object.
type CreateTenancyConfigDataAttributes struct {
	// The auth credentials of the user. Consists of a public key fingerprint and private key.
	AuthCredentials AuthCredentials `json:"auth_credentials"`
	// The config version. It is not recommended to add or change this value, as it is determined internally.
	ConfigVersion *int64 `json:"config_version,omitempty"`
	// The OCID of the compartment containing Datadog managed resources.
	DdCompartmentId *string `json:"dd_compartment_id,omitempty"`
	// The OCID of the resource manager stack for creating Datadog managed resources.
	DdStackId *string `json:"dd_stack_id,omitempty"`
	// The home region of the tenancy to be integrated.
	HomeRegion string `json:"home_region"`
	// The definition of `OCILogsConfig` object.
	LogsConfig *OCILogsConfig `json:"logs_config,omitempty"`
	// The definition of `OCIMetricsConfig` object.
	MetricsConfig *OCIMetricsConfig `json:"metrics_config,omitempty"`
	// The definition of `RegionsConfig` object.
	RegionsConfig *RegionsConfig `json:"regions_config,omitempty"`
	// Enable or disable resource collection.
	ResourceCollectionEnabled *bool `json:"resource_collection_enabled,omitempty"`
	// The OCID of the user needed to authenticate and collect data.
	UserOcid string `json:"user_ocid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateTenancyConfigDataAttributes instantiates a new CreateTenancyConfigDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateTenancyConfigDataAttributes(authCredentials AuthCredentials, homeRegion string, userOcid string) *CreateTenancyConfigDataAttributes {
	this := CreateTenancyConfigDataAttributes{}
	this.AuthCredentials = authCredentials
	this.HomeRegion = homeRegion
	this.UserOcid = userOcid
	return &this
}

// NewCreateTenancyConfigDataAttributesWithDefaults instantiates a new CreateTenancyConfigDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateTenancyConfigDataAttributesWithDefaults() *CreateTenancyConfigDataAttributes {
	this := CreateTenancyConfigDataAttributes{}
	return &this
}

// GetAuthCredentials returns the AuthCredentials field value.
func (o *CreateTenancyConfigDataAttributes) GetAuthCredentials() AuthCredentials {
	if o == nil {
		var ret AuthCredentials
		return ret
	}
	return o.AuthCredentials
}

// GetAuthCredentialsOk returns a tuple with the AuthCredentials field value
// and a boolean to check if the value has been set.
func (o *CreateTenancyConfigDataAttributes) GetAuthCredentialsOk() (*AuthCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthCredentials, true
}

// SetAuthCredentials sets field value.
func (o *CreateTenancyConfigDataAttributes) SetAuthCredentials(v AuthCredentials) {
	o.AuthCredentials = v
}

// GetConfigVersion returns the ConfigVersion field value if set, zero value otherwise.
func (o *CreateTenancyConfigDataAttributes) GetConfigVersion() int64 {
	if o == nil || o.ConfigVersion == nil {
		var ret int64
		return ret
	}
	return *o.ConfigVersion
}

// GetConfigVersionOk returns a tuple with the ConfigVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenancyConfigDataAttributes) GetConfigVersionOk() (*int64, bool) {
	if o == nil || o.ConfigVersion == nil {
		return nil, false
	}
	return o.ConfigVersion, true
}

// HasConfigVersion returns a boolean if a field has been set.
func (o *CreateTenancyConfigDataAttributes) HasConfigVersion() bool {
	return o != nil && o.ConfigVersion != nil
}

// SetConfigVersion gets a reference to the given int64 and assigns it to the ConfigVersion field.
func (o *CreateTenancyConfigDataAttributes) SetConfigVersion(v int64) {
	o.ConfigVersion = &v
}

// GetDdCompartmentId returns the DdCompartmentId field value if set, zero value otherwise.
func (o *CreateTenancyConfigDataAttributes) GetDdCompartmentId() string {
	if o == nil || o.DdCompartmentId == nil {
		var ret string
		return ret
	}
	return *o.DdCompartmentId
}

// GetDdCompartmentIdOk returns a tuple with the DdCompartmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenancyConfigDataAttributes) GetDdCompartmentIdOk() (*string, bool) {
	if o == nil || o.DdCompartmentId == nil {
		return nil, false
	}
	return o.DdCompartmentId, true
}

// HasDdCompartmentId returns a boolean if a field has been set.
func (o *CreateTenancyConfigDataAttributes) HasDdCompartmentId() bool {
	return o != nil && o.DdCompartmentId != nil
}

// SetDdCompartmentId gets a reference to the given string and assigns it to the DdCompartmentId field.
func (o *CreateTenancyConfigDataAttributes) SetDdCompartmentId(v string) {
	o.DdCompartmentId = &v
}

// GetDdStackId returns the DdStackId field value if set, zero value otherwise.
func (o *CreateTenancyConfigDataAttributes) GetDdStackId() string {
	if o == nil || o.DdStackId == nil {
		var ret string
		return ret
	}
	return *o.DdStackId
}

// GetDdStackIdOk returns a tuple with the DdStackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenancyConfigDataAttributes) GetDdStackIdOk() (*string, bool) {
	if o == nil || o.DdStackId == nil {
		return nil, false
	}
	return o.DdStackId, true
}

// HasDdStackId returns a boolean if a field has been set.
func (o *CreateTenancyConfigDataAttributes) HasDdStackId() bool {
	return o != nil && o.DdStackId != nil
}

// SetDdStackId gets a reference to the given string and assigns it to the DdStackId field.
func (o *CreateTenancyConfigDataAttributes) SetDdStackId(v string) {
	o.DdStackId = &v
}

// GetHomeRegion returns the HomeRegion field value.
func (o *CreateTenancyConfigDataAttributes) GetHomeRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.HomeRegion
}

// GetHomeRegionOk returns a tuple with the HomeRegion field value
// and a boolean to check if the value has been set.
func (o *CreateTenancyConfigDataAttributes) GetHomeRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HomeRegion, true
}

// SetHomeRegion sets field value.
func (o *CreateTenancyConfigDataAttributes) SetHomeRegion(v string) {
	o.HomeRegion = v
}

// GetLogsConfig returns the LogsConfig field value if set, zero value otherwise.
func (o *CreateTenancyConfigDataAttributes) GetLogsConfig() OCILogsConfig {
	if o == nil || o.LogsConfig == nil {
		var ret OCILogsConfig
		return ret
	}
	return *o.LogsConfig
}

// GetLogsConfigOk returns a tuple with the LogsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenancyConfigDataAttributes) GetLogsConfigOk() (*OCILogsConfig, bool) {
	if o == nil || o.LogsConfig == nil {
		return nil, false
	}
	return o.LogsConfig, true
}

// HasLogsConfig returns a boolean if a field has been set.
func (o *CreateTenancyConfigDataAttributes) HasLogsConfig() bool {
	return o != nil && o.LogsConfig != nil
}

// SetLogsConfig gets a reference to the given OCILogsConfig and assigns it to the LogsConfig field.
func (o *CreateTenancyConfigDataAttributes) SetLogsConfig(v OCILogsConfig) {
	o.LogsConfig = &v
}

// GetMetricsConfig returns the MetricsConfig field value if set, zero value otherwise.
func (o *CreateTenancyConfigDataAttributes) GetMetricsConfig() OCIMetricsConfig {
	if o == nil || o.MetricsConfig == nil {
		var ret OCIMetricsConfig
		return ret
	}
	return *o.MetricsConfig
}

// GetMetricsConfigOk returns a tuple with the MetricsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenancyConfigDataAttributes) GetMetricsConfigOk() (*OCIMetricsConfig, bool) {
	if o == nil || o.MetricsConfig == nil {
		return nil, false
	}
	return o.MetricsConfig, true
}

// HasMetricsConfig returns a boolean if a field has been set.
func (o *CreateTenancyConfigDataAttributes) HasMetricsConfig() bool {
	return o != nil && o.MetricsConfig != nil
}

// SetMetricsConfig gets a reference to the given OCIMetricsConfig and assigns it to the MetricsConfig field.
func (o *CreateTenancyConfigDataAttributes) SetMetricsConfig(v OCIMetricsConfig) {
	o.MetricsConfig = &v
}

// GetRegionsConfig returns the RegionsConfig field value if set, zero value otherwise.
func (o *CreateTenancyConfigDataAttributes) GetRegionsConfig() RegionsConfig {
	if o == nil || o.RegionsConfig == nil {
		var ret RegionsConfig
		return ret
	}
	return *o.RegionsConfig
}

// GetRegionsConfigOk returns a tuple with the RegionsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenancyConfigDataAttributes) GetRegionsConfigOk() (*RegionsConfig, bool) {
	if o == nil || o.RegionsConfig == nil {
		return nil, false
	}
	return o.RegionsConfig, true
}

// HasRegionsConfig returns a boolean if a field has been set.
func (o *CreateTenancyConfigDataAttributes) HasRegionsConfig() bool {
	return o != nil && o.RegionsConfig != nil
}

// SetRegionsConfig gets a reference to the given RegionsConfig and assigns it to the RegionsConfig field.
func (o *CreateTenancyConfigDataAttributes) SetRegionsConfig(v RegionsConfig) {
	o.RegionsConfig = &v
}

// GetResourceCollectionEnabled returns the ResourceCollectionEnabled field value if set, zero value otherwise.
func (o *CreateTenancyConfigDataAttributes) GetResourceCollectionEnabled() bool {
	if o == nil || o.ResourceCollectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ResourceCollectionEnabled
}

// GetResourceCollectionEnabledOk returns a tuple with the ResourceCollectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenancyConfigDataAttributes) GetResourceCollectionEnabledOk() (*bool, bool) {
	if o == nil || o.ResourceCollectionEnabled == nil {
		return nil, false
	}
	return o.ResourceCollectionEnabled, true
}

// HasResourceCollectionEnabled returns a boolean if a field has been set.
func (o *CreateTenancyConfigDataAttributes) HasResourceCollectionEnabled() bool {
	return o != nil && o.ResourceCollectionEnabled != nil
}

// SetResourceCollectionEnabled gets a reference to the given bool and assigns it to the ResourceCollectionEnabled field.
func (o *CreateTenancyConfigDataAttributes) SetResourceCollectionEnabled(v bool) {
	o.ResourceCollectionEnabled = &v
}

// GetUserOcid returns the UserOcid field value.
func (o *CreateTenancyConfigDataAttributes) GetUserOcid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserOcid
}

// GetUserOcidOk returns a tuple with the UserOcid field value
// and a boolean to check if the value has been set.
func (o *CreateTenancyConfigDataAttributes) GetUserOcidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserOcid, true
}

// SetUserOcid sets field value.
func (o *CreateTenancyConfigDataAttributes) SetUserOcid(v string) {
	o.UserOcid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateTenancyConfigDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_credentials"] = o.AuthCredentials
	if o.ConfigVersion != nil {
		toSerialize["config_version"] = o.ConfigVersion
	}
	if o.DdCompartmentId != nil {
		toSerialize["dd_compartment_id"] = o.DdCompartmentId
	}
	if o.DdStackId != nil {
		toSerialize["dd_stack_id"] = o.DdStackId
	}
	toSerialize["home_region"] = o.HomeRegion
	if o.LogsConfig != nil {
		toSerialize["logs_config"] = o.LogsConfig
	}
	if o.MetricsConfig != nil {
		toSerialize["metrics_config"] = o.MetricsConfig
	}
	if o.RegionsConfig != nil {
		toSerialize["regions_config"] = o.RegionsConfig
	}
	if o.ResourceCollectionEnabled != nil {
		toSerialize["resource_collection_enabled"] = o.ResourceCollectionEnabled
	}
	toSerialize["user_ocid"] = o.UserOcid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateTenancyConfigDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthCredentials           *AuthCredentials  `json:"auth_credentials"`
		ConfigVersion             *int64            `json:"config_version,omitempty"`
		DdCompartmentId           *string           `json:"dd_compartment_id,omitempty"`
		DdStackId                 *string           `json:"dd_stack_id,omitempty"`
		HomeRegion                *string           `json:"home_region"`
		LogsConfig                *OCILogsConfig    `json:"logs_config,omitempty"`
		MetricsConfig             *OCIMetricsConfig `json:"metrics_config,omitempty"`
		RegionsConfig             *RegionsConfig    `json:"regions_config,omitempty"`
		ResourceCollectionEnabled *bool             `json:"resource_collection_enabled,omitempty"`
		UserOcid                  *string           `json:"user_ocid"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthCredentials == nil {
		return fmt.Errorf("required field auth_credentials missing")
	}
	if all.HomeRegion == nil {
		return fmt.Errorf("required field home_region missing")
	}
	if all.UserOcid == nil {
		return fmt.Errorf("required field user_ocid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth_credentials", "config_version", "dd_compartment_id", "dd_stack_id", "home_region", "logs_config", "metrics_config", "regions_config", "resource_collection_enabled", "user_ocid"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AuthCredentials.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AuthCredentials = *all.AuthCredentials
	o.ConfigVersion = all.ConfigVersion
	o.DdCompartmentId = all.DdCompartmentId
	o.DdStackId = all.DdStackId
	o.HomeRegion = *all.HomeRegion
	if all.LogsConfig != nil && all.LogsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LogsConfig = all.LogsConfig
	if all.MetricsConfig != nil && all.MetricsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MetricsConfig = all.MetricsConfig
	if all.RegionsConfig != nil && all.RegionsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RegionsConfig = all.RegionsConfig
	o.ResourceCollectionEnabled = all.ResourceCollectionEnabled
	o.UserOcid = *all.UserOcid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
