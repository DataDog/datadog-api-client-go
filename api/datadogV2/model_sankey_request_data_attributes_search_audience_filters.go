// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyRequestDataAttributesSearchAudienceFilters
type SankeyRequestDataAttributesSearchAudienceFilters struct {
	//
	Accounts []SankeyRequestDataAttributesSearchAudienceFiltersAccountsItems `json:"accounts,omitempty"`
	//
	Formula *string `json:"formula,omitempty"`
	//
	Segments []SankeyRequestDataAttributesSearchAudienceFiltersSegmentsItems `json:"segments,omitempty"`
	//
	Users []SankeyRequestDataAttributesSearchAudienceFiltersUsersItems `json:"users,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSankeyRequestDataAttributesSearchAudienceFilters instantiates a new SankeyRequestDataAttributesSearchAudienceFilters object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSankeyRequestDataAttributesSearchAudienceFilters() *SankeyRequestDataAttributesSearchAudienceFilters {
	this := SankeyRequestDataAttributesSearchAudienceFilters{}
	return &this
}

// NewSankeyRequestDataAttributesSearchAudienceFiltersWithDefaults instantiates a new SankeyRequestDataAttributesSearchAudienceFilters object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSankeyRequestDataAttributesSearchAudienceFiltersWithDefaults() *SankeyRequestDataAttributesSearchAudienceFilters {
	this := SankeyRequestDataAttributesSearchAudienceFilters{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *SankeyRequestDataAttributesSearchAudienceFilters) GetAccounts() []SankeyRequestDataAttributesSearchAudienceFiltersAccountsItems {
	if o == nil || o.Accounts == nil {
		var ret []SankeyRequestDataAttributesSearchAudienceFiltersAccountsItems
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRequestDataAttributesSearchAudienceFilters) GetAccountsOk() (*[]SankeyRequestDataAttributesSearchAudienceFiltersAccountsItems, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return &o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *SankeyRequestDataAttributesSearchAudienceFilters) HasAccounts() bool {
	return o != nil && o.Accounts != nil
}

// SetAccounts gets a reference to the given []SankeyRequestDataAttributesSearchAudienceFiltersAccountsItems and assigns it to the Accounts field.
func (o *SankeyRequestDataAttributesSearchAudienceFilters) SetAccounts(v []SankeyRequestDataAttributesSearchAudienceFiltersAccountsItems) {
	o.Accounts = v
}

// GetFormula returns the Formula field value if set, zero value otherwise.
func (o *SankeyRequestDataAttributesSearchAudienceFilters) GetFormula() string {
	if o == nil || o.Formula == nil {
		var ret string
		return ret
	}
	return *o.Formula
}

// GetFormulaOk returns a tuple with the Formula field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRequestDataAttributesSearchAudienceFilters) GetFormulaOk() (*string, bool) {
	if o == nil || o.Formula == nil {
		return nil, false
	}
	return o.Formula, true
}

// HasFormula returns a boolean if a field has been set.
func (o *SankeyRequestDataAttributesSearchAudienceFilters) HasFormula() bool {
	return o != nil && o.Formula != nil
}

// SetFormula gets a reference to the given string and assigns it to the Formula field.
func (o *SankeyRequestDataAttributesSearchAudienceFilters) SetFormula(v string) {
	o.Formula = &v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *SankeyRequestDataAttributesSearchAudienceFilters) GetSegments() []SankeyRequestDataAttributesSearchAudienceFiltersSegmentsItems {
	if o == nil || o.Segments == nil {
		var ret []SankeyRequestDataAttributesSearchAudienceFiltersSegmentsItems
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRequestDataAttributesSearchAudienceFilters) GetSegmentsOk() (*[]SankeyRequestDataAttributesSearchAudienceFiltersSegmentsItems, bool) {
	if o == nil || o.Segments == nil {
		return nil, false
	}
	return &o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *SankeyRequestDataAttributesSearchAudienceFilters) HasSegments() bool {
	return o != nil && o.Segments != nil
}

// SetSegments gets a reference to the given []SankeyRequestDataAttributesSearchAudienceFiltersSegmentsItems and assigns it to the Segments field.
func (o *SankeyRequestDataAttributesSearchAudienceFilters) SetSegments(v []SankeyRequestDataAttributesSearchAudienceFiltersSegmentsItems) {
	o.Segments = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *SankeyRequestDataAttributesSearchAudienceFilters) GetUsers() []SankeyRequestDataAttributesSearchAudienceFiltersUsersItems {
	if o == nil || o.Users == nil {
		var ret []SankeyRequestDataAttributesSearchAudienceFiltersUsersItems
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRequestDataAttributesSearchAudienceFilters) GetUsersOk() (*[]SankeyRequestDataAttributesSearchAudienceFiltersUsersItems, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *SankeyRequestDataAttributesSearchAudienceFilters) HasUsers() bool {
	return o != nil && o.Users != nil
}

// SetUsers gets a reference to the given []SankeyRequestDataAttributesSearchAudienceFiltersUsersItems and assigns it to the Users field.
func (o *SankeyRequestDataAttributesSearchAudienceFilters) SetUsers(v []SankeyRequestDataAttributesSearchAudienceFiltersUsersItems) {
	o.Users = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SankeyRequestDataAttributesSearchAudienceFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	if o.Formula != nil {
		toSerialize["formula"] = o.Formula
	}
	if o.Segments != nil {
		toSerialize["segments"] = o.Segments
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SankeyRequestDataAttributesSearchAudienceFilters) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Accounts []SankeyRequestDataAttributesSearchAudienceFiltersAccountsItems `json:"accounts,omitempty"`
		Formula  *string                                                         `json:"formula,omitempty"`
		Segments []SankeyRequestDataAttributesSearchAudienceFiltersSegmentsItems `json:"segments,omitempty"`
		Users    []SankeyRequestDataAttributesSearchAudienceFiltersUsersItems    `json:"users,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"accounts", "formula", "segments", "users"})
	} else {
		return err
	}
	o.Accounts = all.Accounts
	o.Formula = all.Formula
	o.Segments = all.Segments
	o.Users = all.Users

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
