// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetDefinition The definition of a dataset, including its data source, name, indexes, and columns.
type SecurityMonitoringDatasetDefinition struct {
	// The columns in the dataset.
	Columns []SecurityMonitoringDatasetDefinitionColumn `json:"columns"`
	// The data source for the dataset.
	DataSource string `json:"data_source"`
	// The indexes to use for the dataset.
	Indexes []string `json:"indexes,omitempty"`
	// The name of the dataset. Must start with a letter, contain only lowercase letters, numbers, and underscores, and be at most 255 characters long.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringDatasetDefinition instantiates a new SecurityMonitoringDatasetDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringDatasetDefinition(columns []SecurityMonitoringDatasetDefinitionColumn, dataSource string, name string) *SecurityMonitoringDatasetDefinition {
	this := SecurityMonitoringDatasetDefinition{}
	this.Columns = columns
	this.DataSource = dataSource
	this.Name = name
	return &this
}

// NewSecurityMonitoringDatasetDefinitionWithDefaults instantiates a new SecurityMonitoringDatasetDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringDatasetDefinitionWithDefaults() *SecurityMonitoringDatasetDefinition {
	this := SecurityMonitoringDatasetDefinition{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *SecurityMonitoringDatasetDefinition) GetColumns() []SecurityMonitoringDatasetDefinitionColumn {
	if o == nil {
		var ret []SecurityMonitoringDatasetDefinitionColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDefinition) GetColumnsOk() (*[]SecurityMonitoringDatasetDefinitionColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *SecurityMonitoringDatasetDefinition) SetColumns(v []SecurityMonitoringDatasetDefinitionColumn) {
	o.Columns = v
}

// GetDataSource returns the DataSource field value.
func (o *SecurityMonitoringDatasetDefinition) GetDataSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDefinition) GetDataSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *SecurityMonitoringDatasetDefinition) SetDataSource(v string) {
	o.DataSource = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *SecurityMonitoringDatasetDefinition) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDefinition) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *SecurityMonitoringDatasetDefinition) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *SecurityMonitoringDatasetDefinition) SetIndexes(v []string) {
	o.Indexes = v
}

// GetName returns the Name field value.
func (o *SecurityMonitoringDatasetDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SecurityMonitoringDatasetDefinition) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringDatasetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["columns"] = o.Columns
	toSerialize["data_source"] = o.DataSource
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringDatasetDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns    *[]SecurityMonitoringDatasetDefinitionColumn `json:"columns"`
		DataSource *string                                      `json:"data_source"`
		Indexes    []string                                     `json:"indexes,omitempty"`
		Name       *string                                      `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Columns == nil {
		return fmt.Errorf("required field columns missing")
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"columns", "data_source", "indexes", "name"})
	} else {
		return err
	}
	o.Columns = *all.Columns
	o.DataSource = *all.DataSource
	o.Indexes = all.Indexes
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
