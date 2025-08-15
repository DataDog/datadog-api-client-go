// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateAppsDatastoreRequestDataAttributes The definition of `CreateAppsDatastoreRequestDataAttributes` object.
type CreateAppsDatastoreRequestDataAttributes struct {
	// The `attributes` `description`.
	Description *string `json:"description,omitempty"`
	// The `attributes` `name`.
	Name string `json:"name"`
	// The `attributes` `org_access`.
	OrgAccess *string `json:"org_access,omitempty"`
	// The `attributes` `primary_column_name`.
	PrimaryColumnName string `json:"primary_column_name"`
	// The `attributes` `primary_key_generation_strategy`.
	PrimaryKeyGenerationStrategy *string `json:"primary_key_generation_strategy,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateAppsDatastoreRequestDataAttributes instantiates a new CreateAppsDatastoreRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateAppsDatastoreRequestDataAttributes(name string, primaryColumnName string) *CreateAppsDatastoreRequestDataAttributes {
	this := CreateAppsDatastoreRequestDataAttributes{}
	this.Name = name
	this.PrimaryColumnName = primaryColumnName
	return &this
}

// NewCreateAppsDatastoreRequestDataAttributesWithDefaults instantiates a new CreateAppsDatastoreRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateAppsDatastoreRequestDataAttributesWithDefaults() *CreateAppsDatastoreRequestDataAttributes {
	this := CreateAppsDatastoreRequestDataAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateAppsDatastoreRequestDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppsDatastoreRequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateAppsDatastoreRequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateAppsDatastoreRequestDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value.
func (o *CreateAppsDatastoreRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAppsDatastoreRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CreateAppsDatastoreRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetOrgAccess returns the OrgAccess field value if set, zero value otherwise.
func (o *CreateAppsDatastoreRequestDataAttributes) GetOrgAccess() string {
	if o == nil || o.OrgAccess == nil {
		var ret string
		return ret
	}
	return *o.OrgAccess
}

// GetOrgAccessOk returns a tuple with the OrgAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppsDatastoreRequestDataAttributes) GetOrgAccessOk() (*string, bool) {
	if o == nil || o.OrgAccess == nil {
		return nil, false
	}
	return o.OrgAccess, true
}

// HasOrgAccess returns a boolean if a field has been set.
func (o *CreateAppsDatastoreRequestDataAttributes) HasOrgAccess() bool {
	return o != nil && o.OrgAccess != nil
}

// SetOrgAccess gets a reference to the given string and assigns it to the OrgAccess field.
func (o *CreateAppsDatastoreRequestDataAttributes) SetOrgAccess(v string) {
	o.OrgAccess = &v
}

// GetPrimaryColumnName returns the PrimaryColumnName field value.
func (o *CreateAppsDatastoreRequestDataAttributes) GetPrimaryColumnName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PrimaryColumnName
}

// GetPrimaryColumnNameOk returns a tuple with the PrimaryColumnName field value
// and a boolean to check if the value has been set.
func (o *CreateAppsDatastoreRequestDataAttributes) GetPrimaryColumnNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryColumnName, true
}

// SetPrimaryColumnName sets field value.
func (o *CreateAppsDatastoreRequestDataAttributes) SetPrimaryColumnName(v string) {
	o.PrimaryColumnName = v
}

// GetPrimaryKeyGenerationStrategy returns the PrimaryKeyGenerationStrategy field value if set, zero value otherwise.
func (o *CreateAppsDatastoreRequestDataAttributes) GetPrimaryKeyGenerationStrategy() string {
	if o == nil || o.PrimaryKeyGenerationStrategy == nil {
		var ret string
		return ret
	}
	return *o.PrimaryKeyGenerationStrategy
}

// GetPrimaryKeyGenerationStrategyOk returns a tuple with the PrimaryKeyGenerationStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppsDatastoreRequestDataAttributes) GetPrimaryKeyGenerationStrategyOk() (*string, bool) {
	if o == nil || o.PrimaryKeyGenerationStrategy == nil {
		return nil, false
	}
	return o.PrimaryKeyGenerationStrategy, true
}

// HasPrimaryKeyGenerationStrategy returns a boolean if a field has been set.
func (o *CreateAppsDatastoreRequestDataAttributes) HasPrimaryKeyGenerationStrategy() bool {
	return o != nil && o.PrimaryKeyGenerationStrategy != nil
}

// SetPrimaryKeyGenerationStrategy gets a reference to the given string and assigns it to the PrimaryKeyGenerationStrategy field.
func (o *CreateAppsDatastoreRequestDataAttributes) SetPrimaryKeyGenerationStrategy(v string) {
	o.PrimaryKeyGenerationStrategy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateAppsDatastoreRequestDataAttributes) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateAppsDatastoreRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description                  *string `json:"description,omitempty"`
		Name                         *string `json:"name"`
		OrgAccess                    *string `json:"org_access,omitempty"`
		PrimaryColumnName            *string `json:"primary_column_name"`
		PrimaryKeyGenerationStrategy *string `json:"primary_key_generation_strategy,omitempty"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "name", "org_access", "primary_column_name", "primary_key_generation_strategy"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Name = *all.Name
	o.OrgAccess = all.OrgAccess
	o.PrimaryColumnName = *all.PrimaryColumnName
	o.PrimaryKeyGenerationStrategy = all.PrimaryKeyGenerationStrategy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
