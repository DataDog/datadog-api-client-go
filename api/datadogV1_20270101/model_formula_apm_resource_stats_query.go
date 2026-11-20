// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaApmResourceStatsQuery Test-only generated schema for the dashboard versioning proof of concept.
type FormulaApmResourceStatsQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaApmResourceStatsQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Env string `json:"env"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []string `json:"group_by"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	OperationName string `json:"operation_name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	PeerService *string `json:"peer.service,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	PrimaryTagName *string `json:"primary_tag_name,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	PrimaryTagValue *string `json:"primary_tag_value,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResourceName *string `json:"resource_name,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Service string `json:"service"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Stat ApmResourceStat `json:"stat"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaApmResourceStatsQuery instantiates a new FormulaApmResourceStatsQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaApmResourceStatsQuery(dataSource FormulaApmResourceStatsQueryDataSource, env string, groupBy []string, name string, operationName string, service string, stat ApmResourceStat) *FormulaApmResourceStatsQuery {
	this := FormulaApmResourceStatsQuery{}
	this.DataSource = dataSource
	this.Env = env
	this.GroupBy = groupBy
	this.Name = name
	this.OperationName = operationName
	this.Service = service
	this.Stat = stat
	return &this
}

// NewFormulaApmResourceStatsQueryWithDefaults instantiates a new FormulaApmResourceStatsQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaApmResourceStatsQueryWithDefaults() *FormulaApmResourceStatsQuery {
	this := FormulaApmResourceStatsQuery{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *FormulaApmResourceStatsQuery) GetDataSource() FormulaApmResourceStatsQueryDataSource {
	if o == nil {
		var ret FormulaApmResourceStatsQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaApmResourceStatsQuery) GetDataSourceOk() (*FormulaApmResourceStatsQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaApmResourceStatsQuery) SetDataSource(v FormulaApmResourceStatsQueryDataSource) {
	o.DataSource = v
}

// GetEnv returns the Env field value.
func (o *FormulaApmResourceStatsQuery) GetEnv() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *FormulaApmResourceStatsQuery) GetEnvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Env, true
}

// SetEnv sets field value.
func (o *FormulaApmResourceStatsQuery) SetEnv(v string) {
	o.Env = v
}

// GetGroupBy returns the GroupBy field value.
func (o *FormulaApmResourceStatsQuery) GetGroupBy() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value
// and a boolean to check if the value has been set.
func (o *FormulaApmResourceStatsQuery) GetGroupByOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// SetGroupBy sets field value.
func (o *FormulaApmResourceStatsQuery) SetGroupBy(v []string) {
	o.GroupBy = v
}

// GetName returns the Name field value.
func (o *FormulaApmResourceStatsQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaApmResourceStatsQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaApmResourceStatsQuery) SetName(v string) {
	o.Name = v
}

// GetOperationName returns the OperationName field value.
func (o *FormulaApmResourceStatsQuery) GetOperationName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OperationName
}

// GetOperationNameOk returns a tuple with the OperationName field value
// and a boolean to check if the value has been set.
func (o *FormulaApmResourceStatsQuery) GetOperationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationName, true
}

// SetOperationName sets field value.
func (o *FormulaApmResourceStatsQuery) SetOperationName(v string) {
	o.OperationName = v
}

// GetPeerService returns the PeerService field value if set, zero value otherwise.
func (o *FormulaApmResourceStatsQuery) GetPeerService() string {
	if o == nil || o.PeerService == nil {
		var ret string
		return ret
	}
	return *o.PeerService
}

// GetPeerServiceOk returns a tuple with the PeerService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaApmResourceStatsQuery) GetPeerServiceOk() (*string, bool) {
	if o == nil || o.PeerService == nil {
		return nil, false
	}
	return o.PeerService, true
}

// HasPeerService returns a boolean if a field has been set.
func (o *FormulaApmResourceStatsQuery) HasPeerService() bool {
	return o != nil && o.PeerService != nil
}

// SetPeerService gets a reference to the given string and assigns it to the PeerService field.
func (o *FormulaApmResourceStatsQuery) SetPeerService(v string) {
	o.PeerService = &v
}

// GetPrimaryTagName returns the PrimaryTagName field value if set, zero value otherwise.
func (o *FormulaApmResourceStatsQuery) GetPrimaryTagName() string {
	if o == nil || o.PrimaryTagName == nil {
		var ret string
		return ret
	}
	return *o.PrimaryTagName
}

// GetPrimaryTagNameOk returns a tuple with the PrimaryTagName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaApmResourceStatsQuery) GetPrimaryTagNameOk() (*string, bool) {
	if o == nil || o.PrimaryTagName == nil {
		return nil, false
	}
	return o.PrimaryTagName, true
}

// HasPrimaryTagName returns a boolean if a field has been set.
func (o *FormulaApmResourceStatsQuery) HasPrimaryTagName() bool {
	return o != nil && o.PrimaryTagName != nil
}

// SetPrimaryTagName gets a reference to the given string and assigns it to the PrimaryTagName field.
func (o *FormulaApmResourceStatsQuery) SetPrimaryTagName(v string) {
	o.PrimaryTagName = &v
}

// GetPrimaryTagValue returns the PrimaryTagValue field value if set, zero value otherwise.
func (o *FormulaApmResourceStatsQuery) GetPrimaryTagValue() string {
	if o == nil || o.PrimaryTagValue == nil {
		var ret string
		return ret
	}
	return *o.PrimaryTagValue
}

// GetPrimaryTagValueOk returns a tuple with the PrimaryTagValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaApmResourceStatsQuery) GetPrimaryTagValueOk() (*string, bool) {
	if o == nil || o.PrimaryTagValue == nil {
		return nil, false
	}
	return o.PrimaryTagValue, true
}

// HasPrimaryTagValue returns a boolean if a field has been set.
func (o *FormulaApmResourceStatsQuery) HasPrimaryTagValue() bool {
	return o != nil && o.PrimaryTagValue != nil
}

// SetPrimaryTagValue gets a reference to the given string and assigns it to the PrimaryTagValue field.
func (o *FormulaApmResourceStatsQuery) SetPrimaryTagValue(v string) {
	o.PrimaryTagValue = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *FormulaApmResourceStatsQuery) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaApmResourceStatsQuery) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *FormulaApmResourceStatsQuery) HasResourceName() bool {
	return o != nil && o.ResourceName != nil
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *FormulaApmResourceStatsQuery) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetService returns the Service field value.
func (o *FormulaApmResourceStatsQuery) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *FormulaApmResourceStatsQuery) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *FormulaApmResourceStatsQuery) SetService(v string) {
	o.Service = v
}

// GetStat returns the Stat field value.
func (o *FormulaApmResourceStatsQuery) GetStat() ApmResourceStat {
	if o == nil {
		var ret ApmResourceStat
		return ret
	}
	return o.Stat
}

// GetStatOk returns a tuple with the Stat field value
// and a boolean to check if the value has been set.
func (o *FormulaApmResourceStatsQuery) GetStatOk() (*ApmResourceStat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stat, true
}

// SetStat sets field value.
func (o *FormulaApmResourceStatsQuery) SetStat(v ApmResourceStat) {
	o.Stat = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FormulaApmResourceStatsQuery) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaApmResourceStatsQuery) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FormulaApmResourceStatsQuery) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *FormulaApmResourceStatsQuery) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaApmResourceStatsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["env"] = o.Env
	toSerialize["group_by"] = o.GroupBy
	toSerialize["name"] = o.Name
	toSerialize["operation_name"] = o.OperationName
	if o.PeerService != nil {
		toSerialize["peer.service"] = o.PeerService
	}
	if o.PrimaryTagName != nil {
		toSerialize["primary_tag_name"] = o.PrimaryTagName
	}
	if o.PrimaryTagValue != nil {
		toSerialize["primary_tag_value"] = o.PrimaryTagValue
	}
	if o.ResourceName != nil {
		toSerialize["resource_name"] = o.ResourceName
	}
	toSerialize["service"] = o.Service
	toSerialize["stat"] = o.Stat
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaApmResourceStatsQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource      *FormulaApmResourceStatsQueryDataSource `json:"data_source"`
		Env             *string                                 `json:"env"`
		GroupBy         *[]string                               `json:"group_by"`
		Name            *string                                 `json:"name"`
		OperationName   *string                                 `json:"operation_name"`
		PeerService     *string                                 `json:"peer.service,omitempty"`
		PrimaryTagName  *string                                 `json:"primary_tag_name,omitempty"`
		PrimaryTagValue *string                                 `json:"primary_tag_value,omitempty"`
		ResourceName    *string                                 `json:"resource_name,omitempty"`
		Service         *string                                 `json:"service"`
		Stat            *ApmResourceStat                        `json:"stat"`
		Version         *string                                 `json:"version,omitempty"`
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
	if all.GroupBy == nil {
		return fmt.Errorf("required field group_by missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OperationName == nil {
		return fmt.Errorf("required field operation_name missing")
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
	o.GroupBy = *all.GroupBy
	o.Name = *all.Name
	o.OperationName = *all.OperationName
	o.PeerService = all.PeerService
	o.PrimaryTagName = all.PrimaryTagName
	o.PrimaryTagValue = all.PrimaryTagValue
	o.ResourceName = all.ResourceName
	o.Service = *all.Service
	if !all.Stat.IsValid() {
		hasInvalidField = true
	} else {
		o.Stat = *all.Stat
	}
	o.Version = all.Version

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
