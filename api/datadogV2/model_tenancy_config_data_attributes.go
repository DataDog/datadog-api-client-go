// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TenancyConfigDataAttributes The definition of `TenancyConfigDataAttributes` object.
type TenancyConfigDataAttributes struct {
	// The config version. It is not recommended to add or change this value, as it is determined internally.
	ConfigVersion *int64 `json:"config_version,omitempty"`
	// Enable or disable cost collection.
	CostCollectionEnabled *bool `json:"cost_collection_enabled,omitempty"`
	// The OCID of the compartment containing Datadog managed resources.
	DdCompartmentId *string `json:"dd_compartment_id,omitempty"`
	// The OCID of the resource manager stack for creating Datadog managed resources.
	DdStackId *string `json:"dd_stack_id,omitempty"`
	// The home region of the tenancy to be integrated.
	HomeRegion *string `json:"home_region,omitempty"`
	// The definition of `OCILogsConfig` object.
	LogsConfig *OCILogsConfig `json:"logs_config,omitempty"`
	// The definition of `OCIMetricsConfig` object.
	MetricsConfig *OCIMetricsConfig `json:"metrics_config,omitempty"`
	// The definition of `RegionsConfig` object.
	RegionsConfig *RegionsConfig `json:"regions_config,omitempty"`
	// Enable or disable resource collection.
	ResourceCollectionEnabled *bool `json:"resource_collection_enabled,omitempty"`
	// The attribute's tenancy_name.
	TenancyName *string `json:"tenancy_name,omitempty"`
	// The OCID of the user needed to authenticate and collect data.
	UserOcid *string `json:"user_ocid,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTenancyConfigDataAttributes instantiates a new TenancyConfigDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTenancyConfigDataAttributes() *TenancyConfigDataAttributes {
	this := TenancyConfigDataAttributes{}
	return &this
}

// NewTenancyConfigDataAttributesWithDefaults instantiates a new TenancyConfigDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTenancyConfigDataAttributesWithDefaults() *TenancyConfigDataAttributes {
	this := TenancyConfigDataAttributes{}
	return &this
}

// GetConfigVersion returns the ConfigVersion field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetConfigVersion() int64 {
	if o == nil || o.ConfigVersion == nil {
		var ret int64
		return ret
	}
	return *o.ConfigVersion
}

// GetConfigVersionOk returns a tuple with the ConfigVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetConfigVersionOk() (*int64, bool) {
	if o == nil || o.ConfigVersion == nil {
		return nil, false
	}
	return o.ConfigVersion, true
}

// HasConfigVersion returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasConfigVersion() bool {
	return o != nil && o.ConfigVersion != nil
}

// SetConfigVersion gets a reference to the given int64 and assigns it to the ConfigVersion field.
func (o *TenancyConfigDataAttributes) SetConfigVersion(v int64) {
	o.ConfigVersion = &v
}

// GetCostCollectionEnabled returns the CostCollectionEnabled field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetCostCollectionEnabled() bool {
	if o == nil || o.CostCollectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CostCollectionEnabled
}

// GetCostCollectionEnabledOk returns a tuple with the CostCollectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetCostCollectionEnabledOk() (*bool, bool) {
	if o == nil || o.CostCollectionEnabled == nil {
		return nil, false
	}
	return o.CostCollectionEnabled, true
}

// HasCostCollectionEnabled returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasCostCollectionEnabled() bool {
	return o != nil && o.CostCollectionEnabled != nil
}

// SetCostCollectionEnabled gets a reference to the given bool and assigns it to the CostCollectionEnabled field.
func (o *TenancyConfigDataAttributes) SetCostCollectionEnabled(v bool) {
	o.CostCollectionEnabled = &v
}

// GetDdCompartmentId returns the DdCompartmentId field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetDdCompartmentId() string {
	if o == nil || o.DdCompartmentId == nil {
		var ret string
		return ret
	}
	return *o.DdCompartmentId
}

// GetDdCompartmentIdOk returns a tuple with the DdCompartmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetDdCompartmentIdOk() (*string, bool) {
	if o == nil || o.DdCompartmentId == nil {
		return nil, false
	}
	return o.DdCompartmentId, true
}

// HasDdCompartmentId returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasDdCompartmentId() bool {
	return o != nil && o.DdCompartmentId != nil
}

// SetDdCompartmentId gets a reference to the given string and assigns it to the DdCompartmentId field.
func (o *TenancyConfigDataAttributes) SetDdCompartmentId(v string) {
	o.DdCompartmentId = &v
}

// GetDdStackId returns the DdStackId field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetDdStackId() string {
	if o == nil || o.DdStackId == nil {
		var ret string
		return ret
	}
	return *o.DdStackId
}

// GetDdStackIdOk returns a tuple with the DdStackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetDdStackIdOk() (*string, bool) {
	if o == nil || o.DdStackId == nil {
		return nil, false
	}
	return o.DdStackId, true
}

// HasDdStackId returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasDdStackId() bool {
	return o != nil && o.DdStackId != nil
}

// SetDdStackId gets a reference to the given string and assigns it to the DdStackId field.
func (o *TenancyConfigDataAttributes) SetDdStackId(v string) {
	o.DdStackId = &v
}

// GetHomeRegion returns the HomeRegion field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetHomeRegion() string {
	if o == nil || o.HomeRegion == nil {
		var ret string
		return ret
	}
	return *o.HomeRegion
}

// GetHomeRegionOk returns a tuple with the HomeRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetHomeRegionOk() (*string, bool) {
	if o == nil || o.HomeRegion == nil {
		return nil, false
	}
	return o.HomeRegion, true
}

// HasHomeRegion returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasHomeRegion() bool {
	return o != nil && o.HomeRegion != nil
}

// SetHomeRegion gets a reference to the given string and assigns it to the HomeRegion field.
func (o *TenancyConfigDataAttributes) SetHomeRegion(v string) {
	o.HomeRegion = &v
}

// GetLogsConfig returns the LogsConfig field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetLogsConfig() OCILogsConfig {
	if o == nil || o.LogsConfig == nil {
		var ret OCILogsConfig
		return ret
	}
	return *o.LogsConfig
}

// GetLogsConfigOk returns a tuple with the LogsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetLogsConfigOk() (*OCILogsConfig, bool) {
	if o == nil || o.LogsConfig == nil {
		return nil, false
	}
	return o.LogsConfig, true
}

// HasLogsConfig returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasLogsConfig() bool {
	return o != nil && o.LogsConfig != nil
}

// SetLogsConfig gets a reference to the given OCILogsConfig and assigns it to the LogsConfig field.
func (o *TenancyConfigDataAttributes) SetLogsConfig(v OCILogsConfig) {
	o.LogsConfig = &v
}

// GetMetricsConfig returns the MetricsConfig field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetMetricsConfig() OCIMetricsConfig {
	if o == nil || o.MetricsConfig == nil {
		var ret OCIMetricsConfig
		return ret
	}
	return *o.MetricsConfig
}

// GetMetricsConfigOk returns a tuple with the MetricsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetMetricsConfigOk() (*OCIMetricsConfig, bool) {
	if o == nil || o.MetricsConfig == nil {
		return nil, false
	}
	return o.MetricsConfig, true
}

// HasMetricsConfig returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasMetricsConfig() bool {
	return o != nil && o.MetricsConfig != nil
}

// SetMetricsConfig gets a reference to the given OCIMetricsConfig and assigns it to the MetricsConfig field.
func (o *TenancyConfigDataAttributes) SetMetricsConfig(v OCIMetricsConfig) {
	o.MetricsConfig = &v
}

// GetRegionsConfig returns the RegionsConfig field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetRegionsConfig() RegionsConfig {
	if o == nil || o.RegionsConfig == nil {
		var ret RegionsConfig
		return ret
	}
	return *o.RegionsConfig
}

// GetRegionsConfigOk returns a tuple with the RegionsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetRegionsConfigOk() (*RegionsConfig, bool) {
	if o == nil || o.RegionsConfig == nil {
		return nil, false
	}
	return o.RegionsConfig, true
}

// HasRegionsConfig returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasRegionsConfig() bool {
	return o != nil && o.RegionsConfig != nil
}

// SetRegionsConfig gets a reference to the given RegionsConfig and assigns it to the RegionsConfig field.
func (o *TenancyConfigDataAttributes) SetRegionsConfig(v RegionsConfig) {
	o.RegionsConfig = &v
}

// GetResourceCollectionEnabled returns the ResourceCollectionEnabled field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetResourceCollectionEnabled() bool {
	if o == nil || o.ResourceCollectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ResourceCollectionEnabled
}

// GetResourceCollectionEnabledOk returns a tuple with the ResourceCollectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetResourceCollectionEnabledOk() (*bool, bool) {
	if o == nil || o.ResourceCollectionEnabled == nil {
		return nil, false
	}
	return o.ResourceCollectionEnabled, true
}

// HasResourceCollectionEnabled returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasResourceCollectionEnabled() bool {
	return o != nil && o.ResourceCollectionEnabled != nil
}

// SetResourceCollectionEnabled gets a reference to the given bool and assigns it to the ResourceCollectionEnabled field.
func (o *TenancyConfigDataAttributes) SetResourceCollectionEnabled(v bool) {
	o.ResourceCollectionEnabled = &v
}

// GetTenancyName returns the TenancyName field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetTenancyName() string {
	if o == nil || o.TenancyName == nil {
		var ret string
		return ret
	}
	return *o.TenancyName
}

// GetTenancyNameOk returns a tuple with the TenancyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetTenancyNameOk() (*string, bool) {
	if o == nil || o.TenancyName == nil {
		return nil, false
	}
	return o.TenancyName, true
}

// HasTenancyName returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasTenancyName() bool {
	return o != nil && o.TenancyName != nil
}

// SetTenancyName gets a reference to the given string and assigns it to the TenancyName field.
func (o *TenancyConfigDataAttributes) SetTenancyName(v string) {
	o.TenancyName = &v
}

// GetUserOcid returns the UserOcid field value if set, zero value otherwise.
func (o *TenancyConfigDataAttributes) GetUserOcid() string {
	if o == nil || o.UserOcid == nil {
		var ret string
		return ret
	}
	return *o.UserOcid
}

// GetUserOcidOk returns a tuple with the UserOcid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyConfigDataAttributes) GetUserOcidOk() (*string, bool) {
	if o == nil || o.UserOcid == nil {
		return nil, false
	}
	return o.UserOcid, true
}

// HasUserOcid returns a boolean if a field has been set.
func (o *TenancyConfigDataAttributes) HasUserOcid() bool {
	return o != nil && o.UserOcid != nil
}

// SetUserOcid gets a reference to the given string and assigns it to the UserOcid field.
func (o *TenancyConfigDataAttributes) SetUserOcid(v string) {
	o.UserOcid = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TenancyConfigDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConfigVersion != nil {
		toSerialize["config_version"] = o.ConfigVersion
	}
	if o.CostCollectionEnabled != nil {
		toSerialize["cost_collection_enabled"] = o.CostCollectionEnabled
	}
	if o.DdCompartmentId != nil {
		toSerialize["dd_compartment_id"] = o.DdCompartmentId
	}
	if o.DdStackId != nil {
		toSerialize["dd_stack_id"] = o.DdStackId
	}
	if o.HomeRegion != nil {
		toSerialize["home_region"] = o.HomeRegion
	}
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
	if o.TenancyName != nil {
		toSerialize["tenancy_name"] = o.TenancyName
	}
	if o.UserOcid != nil {
		toSerialize["user_ocid"] = o.UserOcid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TenancyConfigDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigVersion             *int64            `json:"config_version,omitempty"`
		CostCollectionEnabled     *bool             `json:"cost_collection_enabled,omitempty"`
		DdCompartmentId           *string           `json:"dd_compartment_id,omitempty"`
		DdStackId                 *string           `json:"dd_stack_id,omitempty"`
		HomeRegion                *string           `json:"home_region,omitempty"`
		LogsConfig                *OCILogsConfig    `json:"logs_config,omitempty"`
		MetricsConfig             *OCIMetricsConfig `json:"metrics_config,omitempty"`
		RegionsConfig             *RegionsConfig    `json:"regions_config,omitempty"`
		ResourceCollectionEnabled *bool             `json:"resource_collection_enabled,omitempty"`
		TenancyName               *string           `json:"tenancy_name,omitempty"`
		UserOcid                  *string           `json:"user_ocid,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"config_version", "cost_collection_enabled", "dd_compartment_id", "dd_stack_id", "home_region", "logs_config", "metrics_config", "regions_config", "resource_collection_enabled", "tenancy_name", "user_ocid"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConfigVersion = all.ConfigVersion
	o.CostCollectionEnabled = all.CostCollectionEnabled
	o.DdCompartmentId = all.DdCompartmentId
	o.DdStackId = all.DdStackId
	o.HomeRegion = all.HomeRegion
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
	o.TenancyName = all.TenancyName
	o.UserOcid = all.UserOcid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
