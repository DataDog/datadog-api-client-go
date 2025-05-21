// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OCIMetricsConfig The definition of `OCIMetricsConfig` object.
type OCIMetricsConfig struct {
	// The compartment tag filters to apply to metric collection. Each value represents a Datadog tag in the format key:value.
	CompartmentTagFilters []string `json:"compartment_tag_filters,omitempty"`
	// Enable or disable metric collection. Enabled by default for all services.
	Enabled *bool `json:"enabled,omitempty"`
	// The list of services to exclude from metric collection.
	ExcludedServices []string `json:"excluded_services,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOCIMetricsConfig instantiates a new OCIMetricsConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOCIMetricsConfig() *OCIMetricsConfig {
	this := OCIMetricsConfig{}
	return &this
}

// NewOCIMetricsConfigWithDefaults instantiates a new OCIMetricsConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOCIMetricsConfigWithDefaults() *OCIMetricsConfig {
	this := OCIMetricsConfig{}
	return &this
}

// GetCompartmentTagFilters returns the CompartmentTagFilters field value if set, zero value otherwise.
func (o *OCIMetricsConfig) GetCompartmentTagFilters() []string {
	if o == nil || o.CompartmentTagFilters == nil {
		var ret []string
		return ret
	}
	return o.CompartmentTagFilters
}

// GetCompartmentTagFiltersOk returns a tuple with the CompartmentTagFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCIMetricsConfig) GetCompartmentTagFiltersOk() (*[]string, bool) {
	if o == nil || o.CompartmentTagFilters == nil {
		return nil, false
	}
	return &o.CompartmentTagFilters, true
}

// HasCompartmentTagFilters returns a boolean if a field has been set.
func (o *OCIMetricsConfig) HasCompartmentTagFilters() bool {
	return o != nil && o.CompartmentTagFilters != nil
}

// SetCompartmentTagFilters gets a reference to the given []string and assigns it to the CompartmentTagFilters field.
func (o *OCIMetricsConfig) SetCompartmentTagFilters(v []string) {
	o.CompartmentTagFilters = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OCIMetricsConfig) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCIMetricsConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OCIMetricsConfig) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OCIMetricsConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExcludedServices returns the ExcludedServices field value if set, zero value otherwise.
func (o *OCIMetricsConfig) GetExcludedServices() []string {
	if o == nil || o.ExcludedServices == nil {
		var ret []string
		return ret
	}
	return o.ExcludedServices
}

// GetExcludedServicesOk returns a tuple with the ExcludedServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCIMetricsConfig) GetExcludedServicesOk() (*[]string, bool) {
	if o == nil || o.ExcludedServices == nil {
		return nil, false
	}
	return &o.ExcludedServices, true
}

// HasExcludedServices returns a boolean if a field has been set.
func (o *OCIMetricsConfig) HasExcludedServices() bool {
	return o != nil && o.ExcludedServices != nil
}

// SetExcludedServices gets a reference to the given []string and assigns it to the ExcludedServices field.
func (o *OCIMetricsConfig) SetExcludedServices(v []string) {
	o.ExcludedServices = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OCIMetricsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CompartmentTagFilters != nil {
		toSerialize["compartment_tag_filters"] = o.CompartmentTagFilters
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ExcludedServices != nil {
		toSerialize["excluded_services"] = o.ExcludedServices
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OCIMetricsConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompartmentTagFilters []string `json:"compartment_tag_filters,omitempty"`
		Enabled               *bool    `json:"enabled,omitempty"`
		ExcludedServices      []string `json:"excluded_services,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compartment_tag_filters", "enabled", "excluded_services"})
	} else {
		return err
	}
	o.CompartmentTagFilters = all.CompartmentTagFilters
	o.Enabled = all.Enabled
	o.ExcludedServices = all.ExcludedServices

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
