// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CompositeAudienceFiltersQuery Test-only generated schema for the dashboard versioning proof of concept.
type CompositeAudienceFiltersQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Accounts []CompositeAudienceFiltersQueryAccountsItems `json:"accounts,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource CompositeAudienceFiltersQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Segments []CompositeAudienceFiltersQuerySegmentsItems `json:"segments,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Users []CompositeAudienceFiltersQueryUsersItems `json:"users,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCompositeAudienceFiltersQuery instantiates a new CompositeAudienceFiltersQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCompositeAudienceFiltersQuery(dataSource CompositeAudienceFiltersQueryDataSource, name string) *CompositeAudienceFiltersQuery {
	this := CompositeAudienceFiltersQuery{}
	this.DataSource = dataSource
	this.Name = name
	return &this
}

// NewCompositeAudienceFiltersQueryWithDefaults instantiates a new CompositeAudienceFiltersQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCompositeAudienceFiltersQueryWithDefaults() *CompositeAudienceFiltersQuery {
	this := CompositeAudienceFiltersQuery{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *CompositeAudienceFiltersQuery) GetAccounts() []CompositeAudienceFiltersQueryAccountsItems {
	if o == nil || o.Accounts == nil {
		var ret []CompositeAudienceFiltersQueryAccountsItems
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeAudienceFiltersQuery) GetAccountsOk() (*[]CompositeAudienceFiltersQueryAccountsItems, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return &o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *CompositeAudienceFiltersQuery) HasAccounts() bool {
	return o != nil && o.Accounts != nil
}

// SetAccounts gets a reference to the given []CompositeAudienceFiltersQueryAccountsItems and assigns it to the Accounts field.
func (o *CompositeAudienceFiltersQuery) SetAccounts(v []CompositeAudienceFiltersQueryAccountsItems) {
	o.Accounts = v
}

// GetDataSource returns the DataSource field value.
func (o *CompositeAudienceFiltersQuery) GetDataSource() CompositeAudienceFiltersQueryDataSource {
	if o == nil {
		var ret CompositeAudienceFiltersQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *CompositeAudienceFiltersQuery) GetDataSourceOk() (*CompositeAudienceFiltersQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *CompositeAudienceFiltersQuery) SetDataSource(v CompositeAudienceFiltersQueryDataSource) {
	o.DataSource = v
}

// GetName returns the Name field value.
func (o *CompositeAudienceFiltersQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CompositeAudienceFiltersQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CompositeAudienceFiltersQuery) SetName(v string) {
	o.Name = v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *CompositeAudienceFiltersQuery) GetSegments() []CompositeAudienceFiltersQuerySegmentsItems {
	if o == nil || o.Segments == nil {
		var ret []CompositeAudienceFiltersQuerySegmentsItems
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeAudienceFiltersQuery) GetSegmentsOk() (*[]CompositeAudienceFiltersQuerySegmentsItems, bool) {
	if o == nil || o.Segments == nil {
		return nil, false
	}
	return &o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *CompositeAudienceFiltersQuery) HasSegments() bool {
	return o != nil && o.Segments != nil
}

// SetSegments gets a reference to the given []CompositeAudienceFiltersQuerySegmentsItems and assigns it to the Segments field.
func (o *CompositeAudienceFiltersQuery) SetSegments(v []CompositeAudienceFiltersQuerySegmentsItems) {
	o.Segments = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *CompositeAudienceFiltersQuery) GetUsers() []CompositeAudienceFiltersQueryUsersItems {
	if o == nil || o.Users == nil {
		var ret []CompositeAudienceFiltersQueryUsersItems
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeAudienceFiltersQuery) GetUsersOk() (*[]CompositeAudienceFiltersQueryUsersItems, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *CompositeAudienceFiltersQuery) HasUsers() bool {
	return o != nil && o.Users != nil
}

// SetUsers gets a reference to the given []CompositeAudienceFiltersQueryUsersItems and assigns it to the Users field.
func (o *CompositeAudienceFiltersQuery) SetUsers(v []CompositeAudienceFiltersQueryUsersItems) {
	o.Users = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CompositeAudienceFiltersQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["name"] = o.Name
	if o.Segments != nil {
		toSerialize["segments"] = o.Segments
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CompositeAudienceFiltersQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Accounts   []CompositeAudienceFiltersQueryAccountsItems `json:"accounts,omitempty"`
		DataSource *CompositeAudienceFiltersQueryDataSource     `json:"data_source"`
		Name       *string                                      `json:"name"`
		Segments   []CompositeAudienceFiltersQuerySegmentsItems `json:"segments,omitempty"`
		Users      []CompositeAudienceFiltersQueryUsersItems    `json:"users,omitempty"`
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

	hasInvalidField := false
	o.Accounts = all.Accounts
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Name = *all.Name
	o.Segments = all.Segments
	o.Users = all.Users

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
