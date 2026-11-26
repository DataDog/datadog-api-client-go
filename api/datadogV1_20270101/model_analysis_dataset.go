// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnalysisDataset Test-only generated schema for the dashboard versioning proof of concept.
type AnalysisDataset struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource AnalysisDatasetDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query AnalysisDataset20270101Query `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAnalysisDataset instantiates a new AnalysisDataset object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnalysisDataset(dataSource AnalysisDatasetDataSource, name string, query AnalysisDataset20270101Query) *AnalysisDataset {
	this := AnalysisDataset{}
	this.DataSource = dataSource
	this.Name = name
	this.Query = query
	return &this
}

// NewAnalysisDatasetWithDefaults instantiates a new AnalysisDataset object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnalysisDatasetWithDefaults() *AnalysisDataset {
	this := AnalysisDataset{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *AnalysisDataset) GetDataSource() AnalysisDatasetDataSource {
	if o == nil {
		var ret AnalysisDatasetDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *AnalysisDataset) GetDataSourceOk() (*AnalysisDatasetDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *AnalysisDataset) SetDataSource(v AnalysisDatasetDataSource) {
	o.DataSource = v
}

// GetName returns the Name field value.
func (o *AnalysisDataset) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AnalysisDataset) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AnalysisDataset) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value.
func (o *AnalysisDataset) GetQuery() AnalysisDataset20270101Query {
	if o == nil {
		var ret AnalysisDataset20270101Query
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *AnalysisDataset) GetQueryOk() (*AnalysisDataset20270101Query, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *AnalysisDataset) SetQuery(v AnalysisDataset20270101Query) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnalysisDataset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnalysisDataset) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource *AnalysisDatasetDataSource    `json:"data_source"`
		Name       *string                       `json:"name"`
		Query      *AnalysisDataset20270101Query `json:"query"`
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
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}

	hasInvalidField := false
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Name = *all.Name
	o.Query = *all.Query

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
