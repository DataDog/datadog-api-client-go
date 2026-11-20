// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionEventQueryDefinition Test-only generated schema for the dashboard versioning proof of concept.
type FormulaAndFunctionEventQueryDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	CalculatedFields []EventsCalculatedField `json:"calculated_fields,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute FormulaAndFunctionEventQueryDefinitionCompute `json:"compute"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CrossOrgUuids []string `json:"cross_org_uuids,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaAndFunctionEventsDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Extractions []EventsDataExtraction `json:"extractions,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy *FormulaAndFunctionEventQueryGroupByConfig `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Indexes []string `json:"indexes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Search *FormulaAndFunctionEventQueryDefinitionSearch `json:"search,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SemanticMode *FormulaSemanticMode `json:"semantic_mode,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Storage *StorageType `json:"storage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaAndFunctionEventQueryDefinition instantiates a new FormulaAndFunctionEventQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaAndFunctionEventQueryDefinition(compute FormulaAndFunctionEventQueryDefinitionCompute, dataSource FormulaAndFunctionEventsDataSource, name string) *FormulaAndFunctionEventQueryDefinition {
	this := FormulaAndFunctionEventQueryDefinition{}
	this.Compute = compute
	this.DataSource = dataSource
	this.Name = name
	return &this
}

// NewFormulaAndFunctionEventQueryDefinitionWithDefaults instantiates a new FormulaAndFunctionEventQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaAndFunctionEventQueryDefinitionWithDefaults() *FormulaAndFunctionEventQueryDefinition {
	this := FormulaAndFunctionEventQueryDefinition{}
	return &this
}

// GetCalculatedFields returns the CalculatedFields field value if set, zero value otherwise.
func (o *FormulaAndFunctionEventQueryDefinition) GetCalculatedFields() []EventsCalculatedField {
	if o == nil || o.CalculatedFields == nil {
		var ret []EventsCalculatedField
		return ret
	}
	return o.CalculatedFields
}

// GetCalculatedFieldsOk returns a tuple with the CalculatedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetCalculatedFieldsOk() (*[]EventsCalculatedField, bool) {
	if o == nil || o.CalculatedFields == nil {
		return nil, false
	}
	return &o.CalculatedFields, true
}

// HasCalculatedFields returns a boolean if a field has been set.
func (o *FormulaAndFunctionEventQueryDefinition) HasCalculatedFields() bool {
	return o != nil && o.CalculatedFields != nil
}

// SetCalculatedFields gets a reference to the given []EventsCalculatedField and assigns it to the CalculatedFields field.
func (o *FormulaAndFunctionEventQueryDefinition) SetCalculatedFields(v []EventsCalculatedField) {
	o.CalculatedFields = v
}

// GetCompute returns the Compute field value.
func (o *FormulaAndFunctionEventQueryDefinition) GetCompute() FormulaAndFunctionEventQueryDefinitionCompute {
	if o == nil {
		var ret FormulaAndFunctionEventQueryDefinitionCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetComputeOk() (*FormulaAndFunctionEventQueryDefinitionCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *FormulaAndFunctionEventQueryDefinition) SetCompute(v FormulaAndFunctionEventQueryDefinitionCompute) {
	o.Compute = v
}

// GetCrossOrgUuids returns the CrossOrgUuids field value if set, zero value otherwise.
func (o *FormulaAndFunctionEventQueryDefinition) GetCrossOrgUuids() []string {
	if o == nil || o.CrossOrgUuids == nil {
		var ret []string
		return ret
	}
	return o.CrossOrgUuids
}

// GetCrossOrgUuidsOk returns a tuple with the CrossOrgUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetCrossOrgUuidsOk() (*[]string, bool) {
	if o == nil || o.CrossOrgUuids == nil {
		return nil, false
	}
	return &o.CrossOrgUuids, true
}

// HasCrossOrgUuids returns a boolean if a field has been set.
func (o *FormulaAndFunctionEventQueryDefinition) HasCrossOrgUuids() bool {
	return o != nil && o.CrossOrgUuids != nil
}

// SetCrossOrgUuids gets a reference to the given []string and assigns it to the CrossOrgUuids field.
func (o *FormulaAndFunctionEventQueryDefinition) SetCrossOrgUuids(v []string) {
	o.CrossOrgUuids = v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaAndFunctionEventQueryDefinition) GetDataSource() FormulaAndFunctionEventsDataSource {
	if o == nil {
		var ret FormulaAndFunctionEventsDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetDataSourceOk() (*FormulaAndFunctionEventsDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaAndFunctionEventQueryDefinition) SetDataSource(v FormulaAndFunctionEventsDataSource) {
	o.DataSource = v
}

// GetExtractions returns the Extractions field value if set, zero value otherwise.
func (o *FormulaAndFunctionEventQueryDefinition) GetExtractions() []EventsDataExtraction {
	if o == nil || o.Extractions == nil {
		var ret []EventsDataExtraction
		return ret
	}
	return o.Extractions
}

// GetExtractionsOk returns a tuple with the Extractions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetExtractionsOk() (*[]EventsDataExtraction, bool) {
	if o == nil || o.Extractions == nil {
		return nil, false
	}
	return &o.Extractions, true
}

// HasExtractions returns a boolean if a field has been set.
func (o *FormulaAndFunctionEventQueryDefinition) HasExtractions() bool {
	return o != nil && o.Extractions != nil
}

// SetExtractions gets a reference to the given []EventsDataExtraction and assigns it to the Extractions field.
func (o *FormulaAndFunctionEventQueryDefinition) SetExtractions(v []EventsDataExtraction) {
	o.Extractions = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *FormulaAndFunctionEventQueryDefinition) GetGroupBy() FormulaAndFunctionEventQueryGroupByConfig {
	if o == nil || o.GroupBy == nil {
		var ret FormulaAndFunctionEventQueryGroupByConfig
		return ret
	}
	return *o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetGroupByOk() (*FormulaAndFunctionEventQueryGroupByConfig, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *FormulaAndFunctionEventQueryDefinition) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given FormulaAndFunctionEventQueryGroupByConfig and assigns it to the GroupBy field.
func (o *FormulaAndFunctionEventQueryDefinition) SetGroupBy(v FormulaAndFunctionEventQueryGroupByConfig) {
	o.GroupBy = &v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *FormulaAndFunctionEventQueryDefinition) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *FormulaAndFunctionEventQueryDefinition) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *FormulaAndFunctionEventQueryDefinition) SetIndexes(v []string) {
	o.Indexes = v
}

// GetName returns the Name field value.
func (o *FormulaAndFunctionEventQueryDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaAndFunctionEventQueryDefinition) SetName(v string) {
	o.Name = v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *FormulaAndFunctionEventQueryDefinition) GetSearch() FormulaAndFunctionEventQueryDefinitionSearch {
	if o == nil || o.Search == nil {
		var ret FormulaAndFunctionEventQueryDefinitionSearch
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetSearchOk() (*FormulaAndFunctionEventQueryDefinitionSearch, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *FormulaAndFunctionEventQueryDefinition) HasSearch() bool {
	return o != nil && o.Search != nil
}

// SetSearch gets a reference to the given FormulaAndFunctionEventQueryDefinitionSearch and assigns it to the Search field.
func (o *FormulaAndFunctionEventQueryDefinition) SetSearch(v FormulaAndFunctionEventQueryDefinitionSearch) {
	o.Search = &v
}

// GetSemanticMode returns the SemanticMode field value if set, zero value otherwise.
func (o *FormulaAndFunctionEventQueryDefinition) GetSemanticMode() FormulaSemanticMode {
	if o == nil || o.SemanticMode == nil {
		var ret FormulaSemanticMode
		return ret
	}
	return *o.SemanticMode
}

// GetSemanticModeOk returns a tuple with the SemanticMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetSemanticModeOk() (*FormulaSemanticMode, bool) {
	if o == nil || o.SemanticMode == nil {
		return nil, false
	}
	return o.SemanticMode, true
}

// HasSemanticMode returns a boolean if a field has been set.
func (o *FormulaAndFunctionEventQueryDefinition) HasSemanticMode() bool {
	return o != nil && o.SemanticMode != nil
}

// SetSemanticMode gets a reference to the given FormulaSemanticMode and assigns it to the SemanticMode field.
func (o *FormulaAndFunctionEventQueryDefinition) SetSemanticMode(v FormulaSemanticMode) {
	o.SemanticMode = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *FormulaAndFunctionEventQueryDefinition) GetStorage() StorageType {
	if o == nil || o.Storage == nil {
		var ret StorageType
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetStorageOk() (*StorageType, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *FormulaAndFunctionEventQueryDefinition) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given StorageType and assigns it to the Storage field.
func (o *FormulaAndFunctionEventQueryDefinition) SetStorage(v StorageType) {
	o.Storage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaAndFunctionEventQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CalculatedFields != nil {
		toSerialize["calculated_fields"] = o.CalculatedFields
	}
	toSerialize["compute"] = o.Compute
	if o.CrossOrgUuids != nil {
		toSerialize["cross_org_uuids"] = o.CrossOrgUuids
	}
	toSerialize["data_source"] = o.DataSource
	if o.Extractions != nil {
		toSerialize["extractions"] = o.Extractions
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	toSerialize["name"] = o.Name
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}
	if o.SemanticMode != nil {
		toSerialize["semantic_mode"] = o.SemanticMode
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaAndFunctionEventQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CalculatedFields []EventsCalculatedField                        `json:"calculated_fields,omitempty"`
		Compute          *FormulaAndFunctionEventQueryDefinitionCompute `json:"compute"`
		CrossOrgUuids    []string                                       `json:"cross_org_uuids,omitempty"`
		DataSource       *FormulaAndFunctionEventsDataSource            `json:"data_source"`
		Extractions      []EventsDataExtraction                         `json:"extractions,omitempty"`
		GroupBy          *FormulaAndFunctionEventQueryGroupByConfig     `json:"group_by,omitempty"`
		Indexes          []string                                       `json:"indexes,omitempty"`
		Name             *string                                        `json:"name"`
		Search           *FormulaAndFunctionEventQueryDefinitionSearch  `json:"search,omitempty"`
		SemanticMode     *FormulaSemanticMode                           `json:"semantic_mode,omitempty"`
		Storage          *StorageType                                   `json:"storage,omitempty"`
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

	hasInvalidField := false
	o.CalculatedFields = all.CalculatedFields
	if all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = *all.Compute
	o.CrossOrgUuids = all.CrossOrgUuids
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Extractions = all.Extractions
	o.GroupBy = all.GroupBy
	o.Indexes = all.Indexes
	o.Name = *all.Name
	if all.Search != nil && all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = all.Search
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
