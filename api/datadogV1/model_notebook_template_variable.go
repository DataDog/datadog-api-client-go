// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookTemplateVariable Notebook template variable.
type NotebookTemplateVariable struct {
	// The list of values that the template variable drop-down is limited to.
	AvailableValues datadog.NullableList[string] `json:"available_values,omitempty"`
	// Query used to dynamically populate the list of available values for the template variable.
	AvailableValuesQuery *NotebookTemplateVariableAvailableValuesQuery `json:"available_values_query,omitempty"`
	// Mapping of data source names to template variable values.
	DataSourceMappings map[string]string `json:"data_source_mappings,omitempty"`
	// (deprecated) The default value for the template variable on notebook load.
	// Cannot be used in conjunction with `defaults`.
	// Deprecated
	Default datadog.NullableString `json:"default,omitempty"`
	// One or many default values for the template variable. Cannot be used in conjunction with `default`.
	Defaults []string `json:"defaults,omitempty"`
	// The name of the variable.
	Name string `json:"name"`
	// The placement of the template variable in the notebook.
	Placement *string `json:"placement,omitempty"`
	// The tag prefix associated with the variable. Only tags with this prefix appear in the variable drop-down.
	Prefix datadog.NullableString `json:"prefix,omitempty"`
	// The type of the template variable.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewNotebookTemplateVariable instantiates a new NotebookTemplateVariable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebookTemplateVariable(name string) *NotebookTemplateVariable {
	this := NotebookTemplateVariable{}
	this.Name = name
	return &this
}

// NewNotebookTemplateVariableWithDefaults instantiates a new NotebookTemplateVariable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebookTemplateVariableWithDefaults() *NotebookTemplateVariable {
	this := NotebookTemplateVariable{}
	return &this
}

// GetAvailableValues returns the AvailableValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotebookTemplateVariable) GetAvailableValues() []string {
	if o == nil || o.AvailableValues.Get() == nil {
		var ret []string
		return ret
	}
	return *o.AvailableValues.Get()
}

// GetAvailableValuesOk returns a tuple with the AvailableValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *NotebookTemplateVariable) GetAvailableValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableValues.Get(), o.AvailableValues.IsSet()
}

// HasAvailableValues returns a boolean if a field has been set.
func (o *NotebookTemplateVariable) HasAvailableValues() bool {
	return o != nil && o.AvailableValues.IsSet()
}

// SetAvailableValues gets a reference to the given datadog.NullableList[string] and assigns it to the AvailableValues field.
func (o *NotebookTemplateVariable) SetAvailableValues(v []string) {
	o.AvailableValues.Set(&v)
}

// SetAvailableValuesNil sets the value for AvailableValues to be an explicit nil.
func (o *NotebookTemplateVariable) SetAvailableValuesNil() {
	o.AvailableValues.Set(nil)
}

// UnsetAvailableValues ensures that no value is present for AvailableValues, not even an explicit nil.
func (o *NotebookTemplateVariable) UnsetAvailableValues() {
	o.AvailableValues.Unset()
}

// GetAvailableValuesQuery returns the AvailableValuesQuery field value if set, zero value otherwise.
func (o *NotebookTemplateVariable) GetAvailableValuesQuery() NotebookTemplateVariableAvailableValuesQuery {
	if o == nil || o.AvailableValuesQuery == nil {
		var ret NotebookTemplateVariableAvailableValuesQuery
		return ret
	}
	return *o.AvailableValuesQuery
}

// GetAvailableValuesQueryOk returns a tuple with the AvailableValuesQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookTemplateVariable) GetAvailableValuesQueryOk() (*NotebookTemplateVariableAvailableValuesQuery, bool) {
	if o == nil || o.AvailableValuesQuery == nil {
		return nil, false
	}
	return o.AvailableValuesQuery, true
}

// HasAvailableValuesQuery returns a boolean if a field has been set.
func (o *NotebookTemplateVariable) HasAvailableValuesQuery() bool {
	return o != nil && o.AvailableValuesQuery != nil
}

// SetAvailableValuesQuery gets a reference to the given NotebookTemplateVariableAvailableValuesQuery and assigns it to the AvailableValuesQuery field.
func (o *NotebookTemplateVariable) SetAvailableValuesQuery(v NotebookTemplateVariableAvailableValuesQuery) {
	o.AvailableValuesQuery = &v
}

// GetDataSourceMappings returns the DataSourceMappings field value if set, zero value otherwise.
func (o *NotebookTemplateVariable) GetDataSourceMappings() map[string]string {
	if o == nil || o.DataSourceMappings == nil {
		var ret map[string]string
		return ret
	}
	return o.DataSourceMappings
}

// GetDataSourceMappingsOk returns a tuple with the DataSourceMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookTemplateVariable) GetDataSourceMappingsOk() (*map[string]string, bool) {
	if o == nil || o.DataSourceMappings == nil {
		return nil, false
	}
	return &o.DataSourceMappings, true
}

// HasDataSourceMappings returns a boolean if a field has been set.
func (o *NotebookTemplateVariable) HasDataSourceMappings() bool {
	return o != nil && o.DataSourceMappings != nil
}

// SetDataSourceMappings gets a reference to the given map[string]string and assigns it to the DataSourceMappings field.
func (o *NotebookTemplateVariable) SetDataSourceMappings(v map[string]string) {
	o.DataSourceMappings = v
}

// GetDefault returns the Default field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *NotebookTemplateVariable) GetDefault() string {
	if o == nil || o.Default.Get() == nil {
		var ret string
		return ret
	}
	return *o.Default.Get()
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
// Deprecated
func (o *NotebookTemplateVariable) GetDefaultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Default.Get(), o.Default.IsSet()
}

// HasDefault returns a boolean if a field has been set.
func (o *NotebookTemplateVariable) HasDefault() bool {
	return o != nil && o.Default.IsSet()
}

// SetDefault gets a reference to the given datadog.NullableString and assigns it to the Default field.
// Deprecated
func (o *NotebookTemplateVariable) SetDefault(v string) {
	o.Default.Set(&v)
}

// SetDefaultNil sets the value for Default to be an explicit nil.
func (o *NotebookTemplateVariable) SetDefaultNil() {
	o.Default.Set(nil)
}

// UnsetDefault ensures that no value is present for Default, not even an explicit nil.
func (o *NotebookTemplateVariable) UnsetDefault() {
	o.Default.Unset()
}

// GetDefaults returns the Defaults field value if set, zero value otherwise.
func (o *NotebookTemplateVariable) GetDefaults() []string {
	if o == nil || o.Defaults == nil {
		var ret []string
		return ret
	}
	return o.Defaults
}

// GetDefaultsOk returns a tuple with the Defaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookTemplateVariable) GetDefaultsOk() (*[]string, bool) {
	if o == nil || o.Defaults == nil {
		return nil, false
	}
	return &o.Defaults, true
}

// HasDefaults returns a boolean if a field has been set.
func (o *NotebookTemplateVariable) HasDefaults() bool {
	return o != nil && o.Defaults != nil
}

// SetDefaults gets a reference to the given []string and assigns it to the Defaults field.
func (o *NotebookTemplateVariable) SetDefaults(v []string) {
	o.Defaults = v
}

// GetName returns the Name field value.
func (o *NotebookTemplateVariable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NotebookTemplateVariable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *NotebookTemplateVariable) SetName(v string) {
	o.Name = v
}

// GetPlacement returns the Placement field value if set, zero value otherwise.
func (o *NotebookTemplateVariable) GetPlacement() string {
	if o == nil || o.Placement == nil {
		var ret string
		return ret
	}
	return *o.Placement
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookTemplateVariable) GetPlacementOk() (*string, bool) {
	if o == nil || o.Placement == nil {
		return nil, false
	}
	return o.Placement, true
}

// HasPlacement returns a boolean if a field has been set.
func (o *NotebookTemplateVariable) HasPlacement() bool {
	return o != nil && o.Placement != nil
}

// SetPlacement gets a reference to the given string and assigns it to the Placement field.
func (o *NotebookTemplateVariable) SetPlacement(v string) {
	o.Placement = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotebookTemplateVariable) GetPrefix() string {
	if o == nil || o.Prefix.Get() == nil {
		var ret string
		return ret
	}
	return *o.Prefix.Get()
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *NotebookTemplateVariable) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefix.Get(), o.Prefix.IsSet()
}

// HasPrefix returns a boolean if a field has been set.
func (o *NotebookTemplateVariable) HasPrefix() bool {
	return o != nil && o.Prefix.IsSet()
}

// SetPrefix gets a reference to the given datadog.NullableString and assigns it to the Prefix field.
func (o *NotebookTemplateVariable) SetPrefix(v string) {
	o.Prefix.Set(&v)
}

// SetPrefixNil sets the value for Prefix to be an explicit nil.
func (o *NotebookTemplateVariable) SetPrefixNil() {
	o.Prefix.Set(nil)
}

// UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil.
func (o *NotebookTemplateVariable) UnsetPrefix() {
	o.Prefix.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NotebookTemplateVariable) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookTemplateVariable) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NotebookTemplateVariable) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NotebookTemplateVariable) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebookTemplateVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AvailableValues.IsSet() {
		toSerialize["available_values"] = o.AvailableValues.Get()
	}
	if o.AvailableValuesQuery != nil {
		toSerialize["available_values_query"] = o.AvailableValuesQuery
	}
	if o.DataSourceMappings != nil {
		toSerialize["data_source_mappings"] = o.DataSourceMappings
	}
	if o.Default.IsSet() {
		toSerialize["default"] = o.Default.Get()
	}
	if o.Defaults != nil {
		toSerialize["defaults"] = o.Defaults
	}
	toSerialize["name"] = o.Name
	if o.Placement != nil {
		toSerialize["placement"] = o.Placement
	}
	if o.Prefix.IsSet() {
		toSerialize["prefix"] = o.Prefix.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebookTemplateVariable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AvailableValues      datadog.NullableList[string]                  `json:"available_values,omitempty"`
		AvailableValuesQuery *NotebookTemplateVariableAvailableValuesQuery `json:"available_values_query,omitempty"`
		DataSourceMappings   map[string]string                             `json:"data_source_mappings,omitempty"`
		Default              datadog.NullableString                        `json:"default,omitempty"`
		Defaults             []string                                      `json:"defaults,omitempty"`
		Name                 *string                                       `json:"name"`
		Placement            *string                                       `json:"placement,omitempty"`
		Prefix               datadog.NullableString                        `json:"prefix,omitempty"`
		Type                 *string                                       `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	o.AvailableValues = all.AvailableValues
	o.AvailableValuesQuery = all.AvailableValuesQuery
	o.DataSourceMappings = all.DataSourceMappings
	o.Default = all.Default
	o.Defaults = all.Defaults
	o.Name = *all.Name
	o.Placement = all.Placement
	o.Prefix = all.Prefix
	o.Type = all.Type

	return nil
}
