// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HistoricalDiscoveredProcessesDataset Test-only generated schema for the dashboard versioning proof of concept.
type HistoricalDiscoveredProcessesDataset struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource HistoricalDiscoveredProcessesDatasetDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Timestamp string `json:"timestamp"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHistoricalDiscoveredProcessesDataset instantiates a new HistoricalDiscoveredProcessesDataset object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHistoricalDiscoveredProcessesDataset(dataSource HistoricalDiscoveredProcessesDatasetDataSource, name string, timestamp string) *HistoricalDiscoveredProcessesDataset {
	this := HistoricalDiscoveredProcessesDataset{}
	this.DataSource = dataSource
	this.Name = name
	this.Timestamp = timestamp
	return &this
}

// NewHistoricalDiscoveredProcessesDatasetWithDefaults instantiates a new HistoricalDiscoveredProcessesDataset object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHistoricalDiscoveredProcessesDatasetWithDefaults() *HistoricalDiscoveredProcessesDataset {
	this := HistoricalDiscoveredProcessesDataset{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *HistoricalDiscoveredProcessesDataset) GetDataSource() HistoricalDiscoveredProcessesDatasetDataSource {
	if o == nil {
		var ret HistoricalDiscoveredProcessesDatasetDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *HistoricalDiscoveredProcessesDataset) GetDataSourceOk() (*HistoricalDiscoveredProcessesDatasetDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *HistoricalDiscoveredProcessesDataset) SetDataSource(v HistoricalDiscoveredProcessesDatasetDataSource) {
	o.DataSource = v
}

// GetName returns the Name field value.
func (o *HistoricalDiscoveredProcessesDataset) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HistoricalDiscoveredProcessesDataset) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *HistoricalDiscoveredProcessesDataset) SetName(v string) {
	o.Name = v
}

// GetTimestamp returns the Timestamp field value.
func (o *HistoricalDiscoveredProcessesDataset) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *HistoricalDiscoveredProcessesDataset) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value.
func (o *HistoricalDiscoveredProcessesDataset) SetTimestamp(v string) {
	o.Timestamp = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HistoricalDiscoveredProcessesDataset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["name"] = o.Name
	toSerialize["timestamp"] = o.Timestamp
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HistoricalDiscoveredProcessesDataset) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource *HistoricalDiscoveredProcessesDatasetDataSource `json:"data_source"`
		Name       *string                                         `json:"name"`
		Timestamp  *string                                         `json:"timestamp"`
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
	if all.Timestamp == nil {
		return fmt.Errorf("required field timestamp missing")
	}

	hasInvalidField := false
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Name = *all.Name
	o.Timestamp = *all.Timestamp

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
