// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReferentialQueryFragmentManagedResource Test-only generated schema for the dashboard versioning proof of concept.
type ReferentialQueryFragmentManagedResource struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource ReferentialQueryFragmentManagedResourceDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Filter *string `json:"filter,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TableName string `json:"table_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewReferentialQueryFragmentManagedResource instantiates a new ReferentialQueryFragmentManagedResource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReferentialQueryFragmentManagedResource(dataSource ReferentialQueryFragmentManagedResourceDataSource, name string, tableName string) *ReferentialQueryFragmentManagedResource {
	this := ReferentialQueryFragmentManagedResource{}
	this.DataSource = dataSource
	this.Name = name
	this.TableName = tableName
	return &this
}

// NewReferentialQueryFragmentManagedResourceWithDefaults instantiates a new ReferentialQueryFragmentManagedResource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReferentialQueryFragmentManagedResourceWithDefaults() *ReferentialQueryFragmentManagedResource {
	this := ReferentialQueryFragmentManagedResource{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *ReferentialQueryFragmentManagedResource) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferentialQueryFragmentManagedResource) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *ReferentialQueryFragmentManagedResource) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *ReferentialQueryFragmentManagedResource) SetAlias(v string) {
	o.Alias = &v
}

// GetDataSource returns the DataSource field value.
func (o *ReferentialQueryFragmentManagedResource) GetDataSource() ReferentialQueryFragmentManagedResourceDataSource {
	if o == nil {
		var ret ReferentialQueryFragmentManagedResourceDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *ReferentialQueryFragmentManagedResource) GetDataSourceOk() (*ReferentialQueryFragmentManagedResourceDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *ReferentialQueryFragmentManagedResource) SetDataSource(v ReferentialQueryFragmentManagedResourceDataSource) {
	o.DataSource = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ReferentialQueryFragmentManagedResource) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferentialQueryFragmentManagedResource) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ReferentialQueryFragmentManagedResource) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *ReferentialQueryFragmentManagedResource) SetFilter(v string) {
	o.Filter = &v
}

// GetName returns the Name field value.
func (o *ReferentialQueryFragmentManagedResource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReferentialQueryFragmentManagedResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ReferentialQueryFragmentManagedResource) SetName(v string) {
	o.Name = v
}

// GetTableName returns the TableName field value.
func (o *ReferentialQueryFragmentManagedResource) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *ReferentialQueryFragmentManagedResource) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value.
func (o *ReferentialQueryFragmentManagedResource) SetTableName(v string) {
	o.TableName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReferentialQueryFragmentManagedResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	toSerialize["data_source"] = o.DataSource
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	toSerialize["name"] = o.Name
	toSerialize["table_name"] = o.TableName
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReferentialQueryFragmentManagedResource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias      *string                                            `json:"alias,omitempty"`
		DataSource *ReferentialQueryFragmentManagedResourceDataSource `json:"data_source"`
		Filter     *string                                            `json:"filter,omitempty"`
		Name       *string                                            `json:"name"`
		TableName  *string                                            `json:"table_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.TableName == nil {
		return fmt.Errorf("required field table_name missing")
	}

	hasInvalidField := false
	o.Alias = all.Alias
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Filter = all.Filter
	o.Name = *all.Name
	o.TableName = *all.TableName

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
