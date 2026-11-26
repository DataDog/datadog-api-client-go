// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AudienceFilters Test-only generated schema for the dashboard versioning proof of concept.
type AudienceFilters struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Accounts []AudienceFiltersAccountsItems `json:"accounts,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Segments []AudienceFiltersSegmentsItems `json:"segments,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Users []AudienceFiltersUsersItems `json:"users,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAudienceFilters instantiates a new AudienceFilters object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAudienceFilters() *AudienceFilters {
	this := AudienceFilters{}
	return &this
}

// NewAudienceFiltersWithDefaults instantiates a new AudienceFilters object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAudienceFiltersWithDefaults() *AudienceFilters {
	this := AudienceFilters{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *AudienceFilters) GetAccounts() []AudienceFiltersAccountsItems {
	if o == nil || o.Accounts == nil {
		var ret []AudienceFiltersAccountsItems
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceFilters) GetAccountsOk() (*[]AudienceFiltersAccountsItems, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return &o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *AudienceFilters) HasAccounts() bool {
	return o != nil && o.Accounts != nil
}

// SetAccounts gets a reference to the given []AudienceFiltersAccountsItems and assigns it to the Accounts field.
func (o *AudienceFilters) SetAccounts(v []AudienceFiltersAccountsItems) {
	o.Accounts = v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *AudienceFilters) GetSegments() []AudienceFiltersSegmentsItems {
	if o == nil || o.Segments == nil {
		var ret []AudienceFiltersSegmentsItems
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceFilters) GetSegmentsOk() (*[]AudienceFiltersSegmentsItems, bool) {
	if o == nil || o.Segments == nil {
		return nil, false
	}
	return &o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *AudienceFilters) HasSegments() bool {
	return o != nil && o.Segments != nil
}

// SetSegments gets a reference to the given []AudienceFiltersSegmentsItems and assigns it to the Segments field.
func (o *AudienceFilters) SetSegments(v []AudienceFiltersSegmentsItems) {
	o.Segments = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *AudienceFilters) GetUsers() []AudienceFiltersUsersItems {
	if o == nil || o.Users == nil {
		var ret []AudienceFiltersUsersItems
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceFilters) GetUsersOk() (*[]AudienceFiltersUsersItems, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *AudienceFilters) HasUsers() bool {
	return o != nil && o.Users != nil
}

// SetUsers gets a reference to the given []AudienceFiltersUsersItems and assigns it to the Users field.
func (o *AudienceFilters) SetUsers(v []AudienceFiltersUsersItems) {
	o.Users = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AudienceFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	if o.Segments != nil {
		toSerialize["segments"] = o.Segments
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AudienceFilters) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Accounts []AudienceFiltersAccountsItems `json:"accounts,omitempty"`
		Segments []AudienceFiltersSegmentsItems `json:"segments,omitempty"`
		Users    []AudienceFiltersUsersItems    `json:"users,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Accounts = all.Accounts
	o.Segments = all.Segments
	o.Users = all.Users

	return nil
}
