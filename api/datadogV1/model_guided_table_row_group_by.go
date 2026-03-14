// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableRowGroupBy Defines one group-by dimension used to produce the rows of a guided table.
type GuidedTableRowGroupBy struct {
	// Display name used as the column header for this group-by.
	Alias *string `json:"alias,omitempty"`
	// Text formatting configuration for this group-by column.
	Format *GuidedTableRowGroupByFormat `json:"format,omitempty"`
	// Mapping from each source query to the field name used as the grouping key. Different data sources may use different field names for the same concept.
	GroupKeys []GuidedTableGroupKey `json:"group_keys"`
	// Whether this group-by column is hidden in the rendered table.
	IsHidden *bool `json:"is_hidden,omitempty"`
	// Auto-generated name for this group-by parameter.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGuidedTableRowGroupBy instantiates a new GuidedTableRowGroupBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGuidedTableRowGroupBy(groupKeys []GuidedTableGroupKey, name string) *GuidedTableRowGroupBy {
	this := GuidedTableRowGroupBy{}
	this.GroupKeys = groupKeys
	this.Name = name
	return &this
}

// NewGuidedTableRowGroupByWithDefaults instantiates a new GuidedTableRowGroupBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGuidedTableRowGroupByWithDefaults() *GuidedTableRowGroupBy {
	this := GuidedTableRowGroupBy{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *GuidedTableRowGroupBy) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableRowGroupBy) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *GuidedTableRowGroupBy) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *GuidedTableRowGroupBy) SetAlias(v string) {
	o.Alias = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *GuidedTableRowGroupBy) GetFormat() GuidedTableRowGroupByFormat {
	if o == nil || o.Format == nil {
		var ret GuidedTableRowGroupByFormat
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableRowGroupBy) GetFormatOk() (*GuidedTableRowGroupByFormat, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *GuidedTableRowGroupBy) HasFormat() bool {
	return o != nil && o.Format != nil
}

// SetFormat gets a reference to the given GuidedTableRowGroupByFormat and assigns it to the Format field.
func (o *GuidedTableRowGroupBy) SetFormat(v GuidedTableRowGroupByFormat) {
	o.Format = &v
}

// GetGroupKeys returns the GroupKeys field value.
func (o *GuidedTableRowGroupBy) GetGroupKeys() []GuidedTableGroupKey {
	if o == nil {
		var ret []GuidedTableGroupKey
		return ret
	}
	return o.GroupKeys
}

// GetGroupKeysOk returns a tuple with the GroupKeys field value
// and a boolean to check if the value has been set.
func (o *GuidedTableRowGroupBy) GetGroupKeysOk() (*[]GuidedTableGroupKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupKeys, true
}

// SetGroupKeys sets field value.
func (o *GuidedTableRowGroupBy) SetGroupKeys(v []GuidedTableGroupKey) {
	o.GroupKeys = v
}

// GetIsHidden returns the IsHidden field value if set, zero value otherwise.
func (o *GuidedTableRowGroupBy) GetIsHidden() bool {
	if o == nil || o.IsHidden == nil {
		var ret bool
		return ret
	}
	return *o.IsHidden
}

// GetIsHiddenOk returns a tuple with the IsHidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableRowGroupBy) GetIsHiddenOk() (*bool, bool) {
	if o == nil || o.IsHidden == nil {
		return nil, false
	}
	return o.IsHidden, true
}

// HasIsHidden returns a boolean if a field has been set.
func (o *GuidedTableRowGroupBy) HasIsHidden() bool {
	return o != nil && o.IsHidden != nil
}

// SetIsHidden gets a reference to the given bool and assigns it to the IsHidden field.
func (o *GuidedTableRowGroupBy) SetIsHidden(v bool) {
	o.IsHidden = &v
}

// GetName returns the Name field value.
func (o *GuidedTableRowGroupBy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GuidedTableRowGroupBy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *GuidedTableRowGroupBy) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GuidedTableRowGroupBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	toSerialize["group_keys"] = o.GroupKeys
	if o.IsHidden != nil {
		toSerialize["is_hidden"] = o.IsHidden
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GuidedTableRowGroupBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias     *string                      `json:"alias,omitempty"`
		Format    *GuidedTableRowGroupByFormat `json:"format,omitempty"`
		GroupKeys *[]GuidedTableGroupKey       `json:"group_keys"`
		IsHidden  *bool                        `json:"is_hidden,omitempty"`
		Name      *string                      `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.GroupKeys == nil {
		return fmt.Errorf("required field group_keys missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alias", "format", "group_keys", "is_hidden", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Alias = all.Alias
	if all.Format != nil && all.Format.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Format = all.Format
	o.GroupKeys = *all.GroupKeys
	o.IsHidden = all.IsHidden
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
