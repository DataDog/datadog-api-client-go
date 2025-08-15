// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateAppsDatastoreFromImportRequestDataAttributes The definition of `CreateAppsDatastoreFromImportRequestDataAttributes` object.
type CreateAppsDatastoreFromImportRequestDataAttributes struct {
	// The `attributes` `description`.
	Description *string `json:"description,omitempty"`
	// The `attributes` `name`.
	Name string `json:"name"`
	// The `attributes` `org_access`.
	OrgAccess *string `json:"org_access,omitempty"`
	// The `attributes` `primary_column_name`.
	PrimaryColumnName string `json:"primary_column_name"`
	// The `attributes` `primary_key_generation_strategy`.
	PrimaryKeyGenerationStrategy *CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy `json:"primary_key_generation_strategy,omitempty"`
	// The `attributes` `values`.
	Values []map[string]interface{} `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateAppsDatastoreFromImportRequestDataAttributes instantiates a new CreateAppsDatastoreFromImportRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateAppsDatastoreFromImportRequestDataAttributes(name string, primaryColumnName string, values []map[string]interface{}) *CreateAppsDatastoreFromImportRequestDataAttributes {
	this := CreateAppsDatastoreFromImportRequestDataAttributes{}
	this.Name = name
	this.PrimaryColumnName = primaryColumnName
	this.Values = values
	return &this
}

// NewCreateAppsDatastoreFromImportRequestDataAttributesWithDefaults instantiates a new CreateAppsDatastoreFromImportRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateAppsDatastoreFromImportRequestDataAttributesWithDefaults() *CreateAppsDatastoreFromImportRequestDataAttributes {
	this := CreateAppsDatastoreFromImportRequestDataAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetOrgAccess returns the OrgAccess field value if set, zero value otherwise.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) GetOrgAccess() string {
	if o == nil || o.OrgAccess == nil {
		var ret string
		return ret
	}
	return *o.OrgAccess
}

// GetOrgAccessOk returns a tuple with the OrgAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) GetOrgAccessOk() (*string, bool) {
	if o == nil || o.OrgAccess == nil {
		return nil, false
	}
	return o.OrgAccess, true
}

// HasOrgAccess returns a boolean if a field has been set.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) HasOrgAccess() bool {
	return o != nil && o.OrgAccess != nil
}

// SetOrgAccess gets a reference to the given string and assigns it to the OrgAccess field.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) SetOrgAccess(v string) {
	o.OrgAccess = &v
}

// GetPrimaryColumnName returns the PrimaryColumnName field value.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) GetPrimaryColumnName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PrimaryColumnName
}

// GetPrimaryColumnNameOk returns a tuple with the PrimaryColumnName field value
// and a boolean to check if the value has been set.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) GetPrimaryColumnNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryColumnName, true
}

// SetPrimaryColumnName sets field value.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) SetPrimaryColumnName(v string) {
	o.PrimaryColumnName = v
}

// GetPrimaryKeyGenerationStrategy returns the PrimaryKeyGenerationStrategy field value if set, zero value otherwise.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) GetPrimaryKeyGenerationStrategy() CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy {
	if o == nil || o.PrimaryKeyGenerationStrategy == nil {
		var ret CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy
		return ret
	}
	return *o.PrimaryKeyGenerationStrategy
}

// GetPrimaryKeyGenerationStrategyOk returns a tuple with the PrimaryKeyGenerationStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) GetPrimaryKeyGenerationStrategyOk() (*CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy, bool) {
	if o == nil || o.PrimaryKeyGenerationStrategy == nil {
		return nil, false
	}
	return o.PrimaryKeyGenerationStrategy, true
}

// HasPrimaryKeyGenerationStrategy returns a boolean if a field has been set.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) HasPrimaryKeyGenerationStrategy() bool {
	return o != nil && o.PrimaryKeyGenerationStrategy != nil
}

// SetPrimaryKeyGenerationStrategy gets a reference to the given CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy and assigns it to the PrimaryKeyGenerationStrategy field.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) SetPrimaryKeyGenerationStrategy(v CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy) {
	o.PrimaryKeyGenerationStrategy = &v
}

// GetValues returns the Values field value.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) GetValues() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) GetValuesOk() (*[]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) SetValues(v []map[string]interface{}) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateAppsDatastoreFromImportRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name
	if o.OrgAccess != nil {
		toSerialize["org_access"] = o.OrgAccess
	}
	toSerialize["primary_column_name"] = o.PrimaryColumnName
	if o.PrimaryKeyGenerationStrategy != nil {
		toSerialize["primary_key_generation_strategy"] = o.PrimaryKeyGenerationStrategy
	}
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateAppsDatastoreFromImportRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description                  *string                                                                         `json:"description,omitempty"`
		Name                         *string                                                                         `json:"name"`
		OrgAccess                    *string                                                                         `json:"org_access,omitempty"`
		PrimaryColumnName            *string                                                                         `json:"primary_column_name"`
		PrimaryKeyGenerationStrategy *CreateAppsDatastoreFromImportRequestDataAttributesPrimaryKeyGenerationStrategy `json:"primary_key_generation_strategy,omitempty"`
		Values                       *[]map[string]interface{}                                                       `json:"values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.PrimaryColumnName == nil {
		return fmt.Errorf("required field primary_column_name missing")
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "name", "org_access", "primary_column_name", "primary_key_generation_strategy", "values"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.Name = *all.Name
	o.OrgAccess = all.OrgAccess
	o.PrimaryColumnName = *all.PrimaryColumnName
	if all.PrimaryKeyGenerationStrategy != nil && !all.PrimaryKeyGenerationStrategy.IsValid() {
		hasInvalidField = true
	} else {
		o.PrimaryKeyGenerationStrategy = all.PrimaryKeyGenerationStrategy
	}
	o.Values = *all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
