// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaApmDependencyStatsQuery Test-only generated schema for the dashboard versioning proof of concept.
type FormulaApmDependencyStatsQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaApmDependencyStatsQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Env string `json:"env"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	IsUpstream *bool `json:"is_upstream,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	OperationName string `json:"operation_name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	PrimaryTagName *string `json:"primary_tag_name,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	PrimaryTagValue *string `json:"primary_tag_value,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResourceName string `json:"resource_name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Service string `json:"service"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Stat ApmDependencyStat `json:"stat"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaApmDependencyStatsQuery instantiates a new FormulaApmDependencyStatsQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaApmDependencyStatsQuery(dataSource FormulaApmDependencyStatsQueryDataSource, env string, name string, operationName string, resourceName string, service string, stat ApmDependencyStat) *FormulaApmDependencyStatsQuery {
	this := FormulaApmDependencyStatsQuery{}
	this.DataSource = dataSource
	this.Env = env
	this.Name = name
	this.OperationName = operationName
	this.ResourceName = resourceName
	this.Service = service
	this.Stat = stat
	return &this
}

// NewFormulaApmDependencyStatsQueryWithDefaults instantiates a new FormulaApmDependencyStatsQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaApmDependencyStatsQueryWithDefaults() *FormulaApmDependencyStatsQuery {
	this := FormulaApmDependencyStatsQuery{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *FormulaApmDependencyStatsQuery) GetDataSource() FormulaApmDependencyStatsQueryDataSource {
	if o == nil {
		var ret FormulaApmDependencyStatsQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaApmDependencyStatsQuery) GetDataSourceOk() (*FormulaApmDependencyStatsQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaApmDependencyStatsQuery) SetDataSource(v FormulaApmDependencyStatsQueryDataSource) {
	o.DataSource = v
}

// GetEnv returns the Env field value.
func (o *FormulaApmDependencyStatsQuery) GetEnv() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *FormulaApmDependencyStatsQuery) GetEnvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Env, true
}

// SetEnv sets field value.
func (o *FormulaApmDependencyStatsQuery) SetEnv(v string) {
	o.Env = v
}

// GetIsUpstream returns the IsUpstream field value if set, zero value otherwise.
func (o *FormulaApmDependencyStatsQuery) GetIsUpstream() bool {
	if o == nil || o.IsUpstream == nil {
		var ret bool
		return ret
	}
	return *o.IsUpstream
}

// GetIsUpstreamOk returns a tuple with the IsUpstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaApmDependencyStatsQuery) GetIsUpstreamOk() (*bool, bool) {
	if o == nil || o.IsUpstream == nil {
		return nil, false
	}
	return o.IsUpstream, true
}

// HasIsUpstream returns a boolean if a field has been set.
func (o *FormulaApmDependencyStatsQuery) HasIsUpstream() bool {
	return o != nil && o.IsUpstream != nil
}

// SetIsUpstream gets a reference to the given bool and assigns it to the IsUpstream field.
func (o *FormulaApmDependencyStatsQuery) SetIsUpstream(v bool) {
	o.IsUpstream = &v
}

// GetName returns the Name field value.
func (o *FormulaApmDependencyStatsQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaApmDependencyStatsQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaApmDependencyStatsQuery) SetName(v string) {
	o.Name = v
}

// GetOperationName returns the OperationName field value.
func (o *FormulaApmDependencyStatsQuery) GetOperationName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OperationName
}

// GetOperationNameOk returns a tuple with the OperationName field value
// and a boolean to check if the value has been set.
func (o *FormulaApmDependencyStatsQuery) GetOperationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationName, true
}

// SetOperationName sets field value.
func (o *FormulaApmDependencyStatsQuery) SetOperationName(v string) {
	o.OperationName = v
}

// GetPrimaryTagName returns the PrimaryTagName field value if set, zero value otherwise.
func (o *FormulaApmDependencyStatsQuery) GetPrimaryTagName() string {
	if o == nil || o.PrimaryTagName == nil {
		var ret string
		return ret
	}
	return *o.PrimaryTagName
}

// GetPrimaryTagNameOk returns a tuple with the PrimaryTagName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaApmDependencyStatsQuery) GetPrimaryTagNameOk() (*string, bool) {
	if o == nil || o.PrimaryTagName == nil {
		return nil, false
	}
	return o.PrimaryTagName, true
}

// HasPrimaryTagName returns a boolean if a field has been set.
func (o *FormulaApmDependencyStatsQuery) HasPrimaryTagName() bool {
	return o != nil && o.PrimaryTagName != nil
}

// SetPrimaryTagName gets a reference to the given string and assigns it to the PrimaryTagName field.
func (o *FormulaApmDependencyStatsQuery) SetPrimaryTagName(v string) {
	o.PrimaryTagName = &v
}

// GetPrimaryTagValue returns the PrimaryTagValue field value if set, zero value otherwise.
func (o *FormulaApmDependencyStatsQuery) GetPrimaryTagValue() string {
	if o == nil || o.PrimaryTagValue == nil {
		var ret string
		return ret
	}
	return *o.PrimaryTagValue
}

// GetPrimaryTagValueOk returns a tuple with the PrimaryTagValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaApmDependencyStatsQuery) GetPrimaryTagValueOk() (*string, bool) {
	if o == nil || o.PrimaryTagValue == nil {
		return nil, false
	}
	return o.PrimaryTagValue, true
}

// HasPrimaryTagValue returns a boolean if a field has been set.
func (o *FormulaApmDependencyStatsQuery) HasPrimaryTagValue() bool {
	return o != nil && o.PrimaryTagValue != nil
}

// SetPrimaryTagValue gets a reference to the given string and assigns it to the PrimaryTagValue field.
func (o *FormulaApmDependencyStatsQuery) SetPrimaryTagValue(v string) {
	o.PrimaryTagValue = &v
}

// GetResourceName returns the ResourceName field value.
func (o *FormulaApmDependencyStatsQuery) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *FormulaApmDependencyStatsQuery) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value.
func (o *FormulaApmDependencyStatsQuery) SetResourceName(v string) {
	o.ResourceName = v
}

// GetService returns the Service field value.
func (o *FormulaApmDependencyStatsQuery) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *FormulaApmDependencyStatsQuery) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *FormulaApmDependencyStatsQuery) SetService(v string) {
	o.Service = v
}

// GetStat returns the Stat field value.
func (o *FormulaApmDependencyStatsQuery) GetStat() ApmDependencyStat {
	if o == nil {
		var ret ApmDependencyStat
		return ret
	}
	return o.Stat
}

// GetStatOk returns a tuple with the Stat field value
// and a boolean to check if the value has been set.
func (o *FormulaApmDependencyStatsQuery) GetStatOk() (*ApmDependencyStat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stat, true
}

// SetStat sets field value.
func (o *FormulaApmDependencyStatsQuery) SetStat(v ApmDependencyStat) {
	o.Stat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaApmDependencyStatsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["env"] = o.Env
	if o.IsUpstream != nil {
		toSerialize["is_upstream"] = o.IsUpstream
	}
	toSerialize["name"] = o.Name
	toSerialize["operation_name"] = o.OperationName
	if o.PrimaryTagName != nil {
		toSerialize["primary_tag_name"] = o.PrimaryTagName
	}
	if o.PrimaryTagValue != nil {
		toSerialize["primary_tag_value"] = o.PrimaryTagValue
	}
	toSerialize["resource_name"] = o.ResourceName
	toSerialize["service"] = o.Service
	toSerialize["stat"] = o.Stat
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaApmDependencyStatsQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource      *FormulaApmDependencyStatsQueryDataSource `json:"data_source"`
		Env             *string                                   `json:"env"`
		IsUpstream      *bool                                     `json:"is_upstream,omitempty"`
		Name            *string                                   `json:"name"`
		OperationName   *string                                   `json:"operation_name"`
		PrimaryTagName  *string                                   `json:"primary_tag_name,omitempty"`
		PrimaryTagValue *string                                   `json:"primary_tag_value,omitempty"`
		ResourceName    *string                                   `json:"resource_name"`
		Service         *string                                   `json:"service"`
		Stat            *ApmDependencyStat                        `json:"stat"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Env == nil {
		return fmt.Errorf("required field env missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OperationName == nil {
		return fmt.Errorf("required field operation_name missing")
	}
	if all.ResourceName == nil {
		return fmt.Errorf("required field resource_name missing")
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}
	if all.Stat == nil {
		return fmt.Errorf("required field stat missing")
	}

	hasInvalidField := false
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Env = *all.Env
	o.IsUpstream = all.IsUpstream
	o.Name = *all.Name
	o.OperationName = *all.OperationName
	o.PrimaryTagName = all.PrimaryTagName
	o.PrimaryTagValue = all.PrimaryTagValue
	o.ResourceName = *all.ResourceName
	o.Service = *all.Service
	if !all.Stat.IsValid() {
		hasInvalidField = true
	} else {
		o.Stat = *all.Stat
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
