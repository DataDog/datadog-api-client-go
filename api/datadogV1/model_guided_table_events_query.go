// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableEventsQuery An events-platform query fragment used as source data for a guided table. Covers logs, RUM, spans, CI pipelines, events, and product analytics.
type GuidedTableEventsQuery struct {
	// Display alias for the query.
	Alias *string `json:"alias,omitempty"`
	// Events data source.
	DataSource GuidedTableEventsQueryDataSource `json:"data_source"`
	// Indexes to search for events.
	Indexes []string `json:"indexes,omitempty"`
	// Variable name used to reference this query in columns and formulas.
	Name string `json:"name"`
	// Search filter for matching events.
	Search *GuidedTableEventsQuerySearch `json:"search,omitempty"`
	// Storage tier to target for an events-platform query in a guided table.
	Storage *GuidedTableStorageType `json:"storage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGuidedTableEventsQuery instantiates a new GuidedTableEventsQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGuidedTableEventsQuery(dataSource GuidedTableEventsQueryDataSource, name string) *GuidedTableEventsQuery {
	this := GuidedTableEventsQuery{}
	this.DataSource = dataSource
	this.Name = name
	return &this
}

// NewGuidedTableEventsQueryWithDefaults instantiates a new GuidedTableEventsQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGuidedTableEventsQueryWithDefaults() *GuidedTableEventsQuery {
	this := GuidedTableEventsQuery{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *GuidedTableEventsQuery) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableEventsQuery) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *GuidedTableEventsQuery) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *GuidedTableEventsQuery) SetAlias(v string) {
	o.Alias = &v
}

// GetDataSource returns the DataSource field value.
func (o *GuidedTableEventsQuery) GetDataSource() GuidedTableEventsQueryDataSource {
	if o == nil {
		var ret GuidedTableEventsQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *GuidedTableEventsQuery) GetDataSourceOk() (*GuidedTableEventsQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *GuidedTableEventsQuery) SetDataSource(v GuidedTableEventsQueryDataSource) {
	o.DataSource = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *GuidedTableEventsQuery) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableEventsQuery) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *GuidedTableEventsQuery) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *GuidedTableEventsQuery) SetIndexes(v []string) {
	o.Indexes = v
}

// GetName returns the Name field value.
func (o *GuidedTableEventsQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GuidedTableEventsQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *GuidedTableEventsQuery) SetName(v string) {
	o.Name = v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *GuidedTableEventsQuery) GetSearch() GuidedTableEventsQuerySearch {
	if o == nil || o.Search == nil {
		var ret GuidedTableEventsQuerySearch
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableEventsQuery) GetSearchOk() (*GuidedTableEventsQuerySearch, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *GuidedTableEventsQuery) HasSearch() bool {
	return o != nil && o.Search != nil
}

// SetSearch gets a reference to the given GuidedTableEventsQuerySearch and assigns it to the Search field.
func (o *GuidedTableEventsQuery) SetSearch(v GuidedTableEventsQuerySearch) {
	o.Search = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *GuidedTableEventsQuery) GetStorage() GuidedTableStorageType {
	if o == nil || o.Storage == nil {
		var ret GuidedTableStorageType
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableEventsQuery) GetStorageOk() (*GuidedTableStorageType, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *GuidedTableEventsQuery) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given GuidedTableStorageType and assigns it to the Storage field.
func (o *GuidedTableEventsQuery) SetStorage(v GuidedTableStorageType) {
	o.Storage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GuidedTableEventsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	toSerialize["data_source"] = o.DataSource
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	toSerialize["name"] = o.Name
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GuidedTableEventsQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias      *string                           `json:"alias,omitempty"`
		DataSource *GuidedTableEventsQueryDataSource `json:"data_source"`
		Indexes    []string                          `json:"indexes,omitempty"`
		Name       *string                           `json:"name"`
		Search     *GuidedTableEventsQuerySearch     `json:"search,omitempty"`
		Storage    *GuidedTableStorageType           `json:"storage,omitempty"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alias", "data_source", "indexes", "name", "search", "storage"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Alias = all.Alias
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Indexes = all.Indexes
	o.Name = *all.Name
	if all.Search != nil && all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = all.Search
	if all.Storage != nil && !all.Storage.IsValid() {
		hasInvalidField = true
	} else {
		o.Storage = all.Storage
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
