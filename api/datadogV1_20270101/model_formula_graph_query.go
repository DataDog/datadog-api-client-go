// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaGraphQuery Test-only generated schema for the dashboard versioning proof of concept.
type FormulaGraphQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute FormulaGraphQueryCompute `json:"compute"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaGraphQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []FormulaGraphQueryGroupByItems `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Search GraphQuerySearch `json:"search"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SemanticMode *FormulaSemanticMode `json:"semantic_mode,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Storage *StorageType `json:"storage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaGraphQuery instantiates a new FormulaGraphQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaGraphQuery(compute FormulaGraphQueryCompute, dataSource FormulaGraphQueryDataSource, name string, search GraphQuerySearch) *FormulaGraphQuery {
	this := FormulaGraphQuery{}
	this.Compute = compute
	this.DataSource = dataSource
	this.Name = name
	this.Search = search
	return &this
}

// NewFormulaGraphQueryWithDefaults instantiates a new FormulaGraphQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaGraphQueryWithDefaults() *FormulaGraphQuery {
	this := FormulaGraphQuery{}
	return &this
}

// GetCompute returns the Compute field value.
func (o *FormulaGraphQuery) GetCompute() FormulaGraphQueryCompute {
	if o == nil {
		var ret FormulaGraphQueryCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *FormulaGraphQuery) GetComputeOk() (*FormulaGraphQueryCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *FormulaGraphQuery) SetCompute(v FormulaGraphQueryCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaGraphQuery) GetDataSource() FormulaGraphQueryDataSource {
	if o == nil {
		var ret FormulaGraphQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaGraphQuery) GetDataSourceOk() (*FormulaGraphQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaGraphQuery) SetDataSource(v FormulaGraphQueryDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *FormulaGraphQuery) GetGroupBy() []FormulaGraphQueryGroupByItems {
	if o == nil || o.GroupBy == nil {
		var ret []FormulaGraphQueryGroupByItems
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaGraphQuery) GetGroupByOk() (*[]FormulaGraphQueryGroupByItems, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *FormulaGraphQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []FormulaGraphQueryGroupByItems and assigns it to the GroupBy field.
func (o *FormulaGraphQuery) SetGroupBy(v []FormulaGraphQueryGroupByItems) {
	o.GroupBy = v
}

// GetName returns the Name field value.
func (o *FormulaGraphQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaGraphQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaGraphQuery) SetName(v string) {
	o.Name = v
}

// GetSearch returns the Search field value.
func (o *FormulaGraphQuery) GetSearch() GraphQuerySearch {
	if o == nil {
		var ret GraphQuerySearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *FormulaGraphQuery) GetSearchOk() (*GraphQuerySearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *FormulaGraphQuery) SetSearch(v GraphQuerySearch) {
	o.Search = v
}

// GetSemanticMode returns the SemanticMode field value if set, zero value otherwise.
func (o *FormulaGraphQuery) GetSemanticMode() FormulaSemanticMode {
	if o == nil || o.SemanticMode == nil {
		var ret FormulaSemanticMode
		return ret
	}
	return *o.SemanticMode
}

// GetSemanticModeOk returns a tuple with the SemanticMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaGraphQuery) GetSemanticModeOk() (*FormulaSemanticMode, bool) {
	if o == nil || o.SemanticMode == nil {
		return nil, false
	}
	return o.SemanticMode, true
}

// HasSemanticMode returns a boolean if a field has been set.
func (o *FormulaGraphQuery) HasSemanticMode() bool {
	return o != nil && o.SemanticMode != nil
}

// SetSemanticMode gets a reference to the given FormulaSemanticMode and assigns it to the SemanticMode field.
func (o *FormulaGraphQuery) SetSemanticMode(v FormulaSemanticMode) {
	o.SemanticMode = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *FormulaGraphQuery) GetStorage() StorageType {
	if o == nil || o.Storage == nil {
		var ret StorageType
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaGraphQuery) GetStorageOk() (*StorageType, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *FormulaGraphQuery) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given StorageType and assigns it to the Storage field.
func (o *FormulaGraphQuery) SetStorage(v StorageType) {
	o.Storage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaGraphQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["compute"] = o.Compute
	toSerialize["data_source"] = o.DataSource
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["name"] = o.Name
	toSerialize["search"] = o.Search
	if o.SemanticMode != nil {
		toSerialize["semantic_mode"] = o.SemanticMode
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaGraphQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute      *FormulaGraphQueryCompute       `json:"compute"`
		DataSource   *FormulaGraphQueryDataSource    `json:"data_source"`
		GroupBy      []FormulaGraphQueryGroupByItems `json:"group_by,omitempty"`
		Name         *string                         `json:"name"`
		Search       *GraphQuerySearch               `json:"search"`
		SemanticMode *FormulaSemanticMode            `json:"semantic_mode,omitempty"`
		Storage      *StorageType                    `json:"storage,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Compute == nil {
		return fmt.Errorf("required field compute missing")
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Search == nil {
		return fmt.Errorf("required field search missing")
	}

	hasInvalidField := false
	if all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = *all.Compute
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.GroupBy = all.GroupBy
	o.Name = *all.Name
	if all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = *all.Search
	if all.SemanticMode != nil && !all.SemanticMode.IsValid() {
		hasInvalidField = true
	} else {
		o.SemanticMode = all.SemanticMode
	}
	if all.Storage != nil && !all.Storage.IsValid() {
		hasInvalidField = true
	} else {
		o.Storage = all.Storage
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
