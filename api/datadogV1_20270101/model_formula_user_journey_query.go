// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaUserJourneyQuery Test-only generated schema for the dashboard versioning proof of concept.
type FormulaUserJourneyQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute FormulaUserJourneyQueryCompute `json:"compute"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaUserJourneyQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ExcludeAnonymousTraffic *bool `json:"exclude_anonymous_traffic,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []UserJourneyGroupBy `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Search FormulaUserJourneyQuery20270101Search `json:"search"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Storage *StorageType `json:"storage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaUserJourneyQuery instantiates a new FormulaUserJourneyQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaUserJourneyQuery(compute FormulaUserJourneyQueryCompute, dataSource FormulaUserJourneyQueryDataSource, name string, search FormulaUserJourneyQuery20270101Search) *FormulaUserJourneyQuery {
	this := FormulaUserJourneyQuery{}
	this.Compute = compute
	this.DataSource = dataSource
	this.Name = name
	this.Search = search
	return &this
}

// NewFormulaUserJourneyQueryWithDefaults instantiates a new FormulaUserJourneyQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaUserJourneyQueryWithDefaults() *FormulaUserJourneyQuery {
	this := FormulaUserJourneyQuery{}
	return &this
}

// GetCompute returns the Compute field value.
func (o *FormulaUserJourneyQuery) GetCompute() FormulaUserJourneyQueryCompute {
	if o == nil {
		var ret FormulaUserJourneyQueryCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *FormulaUserJourneyQuery) GetComputeOk() (*FormulaUserJourneyQueryCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *FormulaUserJourneyQuery) SetCompute(v FormulaUserJourneyQueryCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaUserJourneyQuery) GetDataSource() FormulaUserJourneyQueryDataSource {
	if o == nil {
		var ret FormulaUserJourneyQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaUserJourneyQuery) GetDataSourceOk() (*FormulaUserJourneyQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaUserJourneyQuery) SetDataSource(v FormulaUserJourneyQueryDataSource) {
	o.DataSource = v
}

// GetExcludeAnonymousTraffic returns the ExcludeAnonymousTraffic field value if set, zero value otherwise.
func (o *FormulaUserJourneyQuery) GetExcludeAnonymousTraffic() bool {
	if o == nil || o.ExcludeAnonymousTraffic == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeAnonymousTraffic
}

// GetExcludeAnonymousTrafficOk returns a tuple with the ExcludeAnonymousTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaUserJourneyQuery) GetExcludeAnonymousTrafficOk() (*bool, bool) {
	if o == nil || o.ExcludeAnonymousTraffic == nil {
		return nil, false
	}
	return o.ExcludeAnonymousTraffic, true
}

// HasExcludeAnonymousTraffic returns a boolean if a field has been set.
func (o *FormulaUserJourneyQuery) HasExcludeAnonymousTraffic() bool {
	return o != nil && o.ExcludeAnonymousTraffic != nil
}

// SetExcludeAnonymousTraffic gets a reference to the given bool and assigns it to the ExcludeAnonymousTraffic field.
func (o *FormulaUserJourneyQuery) SetExcludeAnonymousTraffic(v bool) {
	o.ExcludeAnonymousTraffic = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *FormulaUserJourneyQuery) GetGroupBy() []UserJourneyGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []UserJourneyGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaUserJourneyQuery) GetGroupByOk() (*[]UserJourneyGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *FormulaUserJourneyQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []UserJourneyGroupBy and assigns it to the GroupBy field.
func (o *FormulaUserJourneyQuery) SetGroupBy(v []UserJourneyGroupBy) {
	o.GroupBy = v
}

// GetName returns the Name field value.
func (o *FormulaUserJourneyQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaUserJourneyQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaUserJourneyQuery) SetName(v string) {
	o.Name = v
}

// GetSearch returns the Search field value.
func (o *FormulaUserJourneyQuery) GetSearch() FormulaUserJourneyQuery20270101Search {
	if o == nil {
		var ret FormulaUserJourneyQuery20270101Search
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *FormulaUserJourneyQuery) GetSearchOk() (*FormulaUserJourneyQuery20270101Search, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *FormulaUserJourneyQuery) SetSearch(v FormulaUserJourneyQuery20270101Search) {
	o.Search = v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *FormulaUserJourneyQuery) GetStorage() StorageType {
	if o == nil || o.Storage == nil {
		var ret StorageType
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaUserJourneyQuery) GetStorageOk() (*StorageType, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *FormulaUserJourneyQuery) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given StorageType and assigns it to the Storage field.
func (o *FormulaUserJourneyQuery) SetStorage(v StorageType) {
	o.Storage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaUserJourneyQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["compute"] = o.Compute
	toSerialize["data_source"] = o.DataSource
	if o.ExcludeAnonymousTraffic != nil {
		toSerialize["exclude_anonymous_traffic"] = o.ExcludeAnonymousTraffic
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["name"] = o.Name
	toSerialize["search"] = o.Search
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaUserJourneyQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute                 *FormulaUserJourneyQueryCompute        `json:"compute"`
		DataSource              *FormulaUserJourneyQueryDataSource     `json:"data_source"`
		ExcludeAnonymousTraffic *bool                                  `json:"exclude_anonymous_traffic,omitempty"`
		GroupBy                 []UserJourneyGroupBy                   `json:"group_by,omitempty"`
		Name                    *string                                `json:"name"`
		Search                  *FormulaUserJourneyQuery20270101Search `json:"search"`
		Storage                 *StorageType                           `json:"storage,omitempty"`
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
	o.ExcludeAnonymousTraffic = all.ExcludeAnonymousTraffic
	o.GroupBy = all.GroupBy
	o.Name = *all.Name
	o.Search = *all.Search
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
