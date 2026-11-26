// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PublishedAnalysisQueryFragment Test-only generated schema for the dashboard versioning proof of concept.
type PublishedAnalysisQueryFragment struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource PublishedAnalysisQueryFragmentDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DatasetId string `json:"dataset_id"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Filter *string `json:"filter,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewPublishedAnalysisQueryFragment instantiates a new PublishedAnalysisQueryFragment object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPublishedAnalysisQueryFragment(dataSource PublishedAnalysisQueryFragmentDataSource, datasetId string, name string) *PublishedAnalysisQueryFragment {
	this := PublishedAnalysisQueryFragment{}
	this.DataSource = dataSource
	this.DatasetId = datasetId
	this.Name = name
	return &this
}

// NewPublishedAnalysisQueryFragmentWithDefaults instantiates a new PublishedAnalysisQueryFragment object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPublishedAnalysisQueryFragmentWithDefaults() *PublishedAnalysisQueryFragment {
	this := PublishedAnalysisQueryFragment{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *PublishedAnalysisQueryFragment) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishedAnalysisQueryFragment) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *PublishedAnalysisQueryFragment) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *PublishedAnalysisQueryFragment) SetAlias(v string) {
	o.Alias = &v
}

// GetDataSource returns the DataSource field value.
func (o *PublishedAnalysisQueryFragment) GetDataSource() PublishedAnalysisQueryFragmentDataSource {
	if o == nil {
		var ret PublishedAnalysisQueryFragmentDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *PublishedAnalysisQueryFragment) GetDataSourceOk() (*PublishedAnalysisQueryFragmentDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *PublishedAnalysisQueryFragment) SetDataSource(v PublishedAnalysisQueryFragmentDataSource) {
	o.DataSource = v
}

// GetDatasetId returns the DatasetId field value.
func (o *PublishedAnalysisQueryFragment) GetDatasetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value
// and a boolean to check if the value has been set.
func (o *PublishedAnalysisQueryFragment) GetDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetId, true
}

// SetDatasetId sets field value.
func (o *PublishedAnalysisQueryFragment) SetDatasetId(v string) {
	o.DatasetId = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *PublishedAnalysisQueryFragment) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishedAnalysisQueryFragment) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *PublishedAnalysisQueryFragment) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *PublishedAnalysisQueryFragment) SetFilter(v string) {
	o.Filter = &v
}

// GetName returns the Name field value.
func (o *PublishedAnalysisQueryFragment) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PublishedAnalysisQueryFragment) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *PublishedAnalysisQueryFragment) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PublishedAnalysisQueryFragment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["dataset_id"] = o.DatasetId
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	toSerialize["name"] = o.Name
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PublishedAnalysisQueryFragment) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias      *string                                   `json:"alias,omitempty"`
		DataSource *PublishedAnalysisQueryFragmentDataSource `json:"data_source"`
		DatasetId  *string                                   `json:"dataset_id"`
		Filter     *string                                   `json:"filter,omitempty"`
		Name       *string                                   `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.DatasetId == nil {
		return fmt.Errorf("required field dataset_id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}

	hasInvalidField := false
	o.Alias = all.Alias
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.DatasetId = *all.DatasetId
	o.Filter = all.Filter
	o.Name = *all.Name

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
