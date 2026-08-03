// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDeploymentConfigureV2Attributes Attributes for creating a new v2 configuration deployment.
type FleetDeploymentConfigureV2Attributes struct {
	// Ordered list of configuration file operations to perform on the target hosts.
	ConfigOperations []FleetDeploymentOperation `json:"config_operations"`
	// Set to `true` to validate the configuration and resolve target hosts and packages
	// without deploying anything. Returns a 200 with the validation result instead of
	// creating and starting a real deployment.
	DryRun *bool `json:"dry_run,omitempty"`
	// Query used to filter and select target hosts for the deployment. Uses the Datadog query syntax.
	FilterQuery string `json:"filter_query"`
	// List of packages and their target versions to additionally deploy alongside
	// the configuration change.
	TargetPackages []FleetDeploymentConfigureV2Package `json:"target_packages,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetDeploymentConfigureV2Attributes instantiates a new FleetDeploymentConfigureV2Attributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetDeploymentConfigureV2Attributes(configOperations []FleetDeploymentOperation, filterQuery string) *FleetDeploymentConfigureV2Attributes {
	this := FleetDeploymentConfigureV2Attributes{}
	this.ConfigOperations = configOperations
	this.FilterQuery = filterQuery
	return &this
}

// NewFleetDeploymentConfigureV2AttributesWithDefaults instantiates a new FleetDeploymentConfigureV2Attributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetDeploymentConfigureV2AttributesWithDefaults() *FleetDeploymentConfigureV2Attributes {
	this := FleetDeploymentConfigureV2Attributes{}
	return &this
}

// GetConfigOperations returns the ConfigOperations field value.
func (o *FleetDeploymentConfigureV2Attributes) GetConfigOperations() []FleetDeploymentOperation {
	if o == nil {
		var ret []FleetDeploymentOperation
		return ret
	}
	return o.ConfigOperations
}

// GetConfigOperationsOk returns a tuple with the ConfigOperations field value
// and a boolean to check if the value has been set.
func (o *FleetDeploymentConfigureV2Attributes) GetConfigOperationsOk() (*[]FleetDeploymentOperation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigOperations, true
}

// SetConfigOperations sets field value.
func (o *FleetDeploymentConfigureV2Attributes) SetConfigOperations(v []FleetDeploymentOperation) {
	o.ConfigOperations = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *FleetDeploymentConfigureV2Attributes) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentConfigureV2Attributes) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *FleetDeploymentConfigureV2Attributes) HasDryRun() bool {
	return o != nil && o.DryRun != nil
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *FleetDeploymentConfigureV2Attributes) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFilterQuery returns the FilterQuery field value.
func (o *FleetDeploymentConfigureV2Attributes) GetFilterQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FilterQuery
}

// GetFilterQueryOk returns a tuple with the FilterQuery field value
// and a boolean to check if the value has been set.
func (o *FleetDeploymentConfigureV2Attributes) GetFilterQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterQuery, true
}

// SetFilterQuery sets field value.
func (o *FleetDeploymentConfigureV2Attributes) SetFilterQuery(v string) {
	o.FilterQuery = v
}

// GetTargetPackages returns the TargetPackages field value if set, zero value otherwise.
func (o *FleetDeploymentConfigureV2Attributes) GetTargetPackages() []FleetDeploymentConfigureV2Package {
	if o == nil || o.TargetPackages == nil {
		var ret []FleetDeploymentConfigureV2Package
		return ret
	}
	return o.TargetPackages
}

// GetTargetPackagesOk returns a tuple with the TargetPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentConfigureV2Attributes) GetTargetPackagesOk() (*[]FleetDeploymentConfigureV2Package, bool) {
	if o == nil || o.TargetPackages == nil {
		return nil, false
	}
	return &o.TargetPackages, true
}

// HasTargetPackages returns a boolean if a field has been set.
func (o *FleetDeploymentConfigureV2Attributes) HasTargetPackages() bool {
	return o != nil && o.TargetPackages != nil
}

// SetTargetPackages gets a reference to the given []FleetDeploymentConfigureV2Package and assigns it to the TargetPackages field.
func (o *FleetDeploymentConfigureV2Attributes) SetTargetPackages(v []FleetDeploymentConfigureV2Package) {
	o.TargetPackages = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetDeploymentConfigureV2Attributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["config_operations"] = o.ConfigOperations
	if o.DryRun != nil {
		toSerialize["dry_run"] = o.DryRun
	}
	toSerialize["filter_query"] = o.FilterQuery
	if o.TargetPackages != nil {
		toSerialize["target_packages"] = o.TargetPackages
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetDeploymentConfigureV2Attributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigOperations *[]FleetDeploymentOperation         `json:"config_operations"`
		DryRun           *bool                               `json:"dry_run,omitempty"`
		FilterQuery      *string                             `json:"filter_query"`
		TargetPackages   []FleetDeploymentConfigureV2Package `json:"target_packages,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConfigOperations == nil {
		return fmt.Errorf("required field config_operations missing")
	}
	if all.FilterQuery == nil {
		return fmt.Errorf("required field filter_query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"config_operations", "dry_run", "filter_query", "target_packages"})
	} else {
		return err
	}
	o.ConfigOperations = *all.ConfigOperations
	o.DryRun = all.DryRun
	o.FilterQuery = *all.FilterQuery
	o.TargetPackages = all.TargetPackages

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
