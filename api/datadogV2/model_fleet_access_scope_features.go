// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetAccessScopeFeatures Fleet actions authorized on agents matching the access scope query.
type FleetAccessScopeFeatures struct {
	// Whether the scope permits configuration changes.
	Configure *bool `json:"configure,omitempty"`
	// Whether the scope permits Agent installation.
	Install *bool `json:"install,omitempty"`
	// Whether the scope permits package upgrades.
	Upgrade *bool `json:"upgrade,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetAccessScopeFeatures instantiates a new FleetAccessScopeFeatures object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetAccessScopeFeatures() *FleetAccessScopeFeatures {
	this := FleetAccessScopeFeatures{}
	var configure bool = false
	this.Configure = &configure
	var install bool = false
	this.Install = &install
	var upgrade bool = false
	this.Upgrade = &upgrade
	return &this
}

// NewFleetAccessScopeFeaturesWithDefaults instantiates a new FleetAccessScopeFeatures object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetAccessScopeFeaturesWithDefaults() *FleetAccessScopeFeatures {
	this := FleetAccessScopeFeatures{}
	var configure bool = false
	this.Configure = &configure
	var install bool = false
	this.Install = &install
	var upgrade bool = false
	this.Upgrade = &upgrade
	return &this
}

// GetConfigure returns the Configure field value if set, zero value otherwise.
func (o *FleetAccessScopeFeatures) GetConfigure() bool {
	if o == nil || o.Configure == nil {
		var ret bool
		return ret
	}
	return *o.Configure
}

// GetConfigureOk returns a tuple with the Configure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAccessScopeFeatures) GetConfigureOk() (*bool, bool) {
	if o == nil || o.Configure == nil {
		return nil, false
	}
	return o.Configure, true
}

// HasConfigure returns a boolean if a field has been set.
func (o *FleetAccessScopeFeatures) HasConfigure() bool {
	return o != nil && o.Configure != nil
}

// SetConfigure gets a reference to the given bool and assigns it to the Configure field.
func (o *FleetAccessScopeFeatures) SetConfigure(v bool) {
	o.Configure = &v
}

// GetInstall returns the Install field value if set, zero value otherwise.
func (o *FleetAccessScopeFeatures) GetInstall() bool {
	if o == nil || o.Install == nil {
		var ret bool
		return ret
	}
	return *o.Install
}

// GetInstallOk returns a tuple with the Install field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAccessScopeFeatures) GetInstallOk() (*bool, bool) {
	if o == nil || o.Install == nil {
		return nil, false
	}
	return o.Install, true
}

// HasInstall returns a boolean if a field has been set.
func (o *FleetAccessScopeFeatures) HasInstall() bool {
	return o != nil && o.Install != nil
}

// SetInstall gets a reference to the given bool and assigns it to the Install field.
func (o *FleetAccessScopeFeatures) SetInstall(v bool) {
	o.Install = &v
}

// GetUpgrade returns the Upgrade field value if set, zero value otherwise.
func (o *FleetAccessScopeFeatures) GetUpgrade() bool {
	if o == nil || o.Upgrade == nil {
		var ret bool
		return ret
	}
	return *o.Upgrade
}

// GetUpgradeOk returns a tuple with the Upgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAccessScopeFeatures) GetUpgradeOk() (*bool, bool) {
	if o == nil || o.Upgrade == nil {
		return nil, false
	}
	return o.Upgrade, true
}

// HasUpgrade returns a boolean if a field has been set.
func (o *FleetAccessScopeFeatures) HasUpgrade() bool {
	return o != nil && o.Upgrade != nil
}

// SetUpgrade gets a reference to the given bool and assigns it to the Upgrade field.
func (o *FleetAccessScopeFeatures) SetUpgrade(v bool) {
	o.Upgrade = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetAccessScopeFeatures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Configure != nil {
		toSerialize["configure"] = o.Configure
	}
	if o.Install != nil {
		toSerialize["install"] = o.Install
	}
	if o.Upgrade != nil {
		toSerialize["upgrade"] = o.Upgrade
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetAccessScopeFeatures) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Configure *bool `json:"configure,omitempty"`
		Install   *bool `json:"install,omitempty"`
		Upgrade   *bool `json:"upgrade,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"configure", "install", "upgrade"})
	} else {
		return err
	}
	o.Configure = all.Configure
	o.Install = all.Install
	o.Upgrade = all.Upgrade

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
