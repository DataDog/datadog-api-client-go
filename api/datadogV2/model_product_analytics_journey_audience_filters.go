// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyAudienceFilters Restricts the journey to an audience built from named sub-queries.
// Sub-query names must be unique across `users`, `segments`, and `accounts`.
type ProductAnalyticsJourneyAudienceFilters struct {
	// Named account sub-queries.
	Accounts []ProductAnalyticsJourneyAudienceAccountQuery `json:"accounts,omitempty"`
	// Boolean expression combining the sub-query names with `AND`, `OR`, and `NOT`.
	// When empty, all sub-queries are combined with `AND`.
	Formula *string `json:"formula,omitempty"`
	// Named segment sub-queries.
	Segments []ProductAnalyticsJourneyAudienceSegmentQuery `json:"segments,omitempty"`
	// Named user sub-queries.
	Users []ProductAnalyticsJourneyAudienceUserQuery `json:"users,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneyAudienceFilters instantiates a new ProductAnalyticsJourneyAudienceFilters object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneyAudienceFilters() *ProductAnalyticsJourneyAudienceFilters {
	this := ProductAnalyticsJourneyAudienceFilters{}
	return &this
}

// NewProductAnalyticsJourneyAudienceFiltersWithDefaults instantiates a new ProductAnalyticsJourneyAudienceFilters object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneyAudienceFiltersWithDefaults() *ProductAnalyticsJourneyAudienceFilters {
	this := ProductAnalyticsJourneyAudienceFilters{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyAudienceFilters) GetAccounts() []ProductAnalyticsJourneyAudienceAccountQuery {
	if o == nil || o.Accounts == nil {
		var ret []ProductAnalyticsJourneyAudienceAccountQuery
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyAudienceFilters) GetAccountsOk() (*[]ProductAnalyticsJourneyAudienceAccountQuery, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return &o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyAudienceFilters) HasAccounts() bool {
	return o != nil && o.Accounts != nil
}

// SetAccounts gets a reference to the given []ProductAnalyticsJourneyAudienceAccountQuery and assigns it to the Accounts field.
func (o *ProductAnalyticsJourneyAudienceFilters) SetAccounts(v []ProductAnalyticsJourneyAudienceAccountQuery) {
	o.Accounts = v
}

// GetFormula returns the Formula field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyAudienceFilters) GetFormula() string {
	if o == nil || o.Formula == nil {
		var ret string
		return ret
	}
	return *o.Formula
}

// GetFormulaOk returns a tuple with the Formula field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyAudienceFilters) GetFormulaOk() (*string, bool) {
	if o == nil || o.Formula == nil {
		return nil, false
	}
	return o.Formula, true
}

// HasFormula returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyAudienceFilters) HasFormula() bool {
	return o != nil && o.Formula != nil
}

// SetFormula gets a reference to the given string and assigns it to the Formula field.
func (o *ProductAnalyticsJourneyAudienceFilters) SetFormula(v string) {
	o.Formula = &v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyAudienceFilters) GetSegments() []ProductAnalyticsJourneyAudienceSegmentQuery {
	if o == nil || o.Segments == nil {
		var ret []ProductAnalyticsJourneyAudienceSegmentQuery
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyAudienceFilters) GetSegmentsOk() (*[]ProductAnalyticsJourneyAudienceSegmentQuery, bool) {
	if o == nil || o.Segments == nil {
		return nil, false
	}
	return &o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyAudienceFilters) HasSegments() bool {
	return o != nil && o.Segments != nil
}

// SetSegments gets a reference to the given []ProductAnalyticsJourneyAudienceSegmentQuery and assigns it to the Segments field.
func (o *ProductAnalyticsJourneyAudienceFilters) SetSegments(v []ProductAnalyticsJourneyAudienceSegmentQuery) {
	o.Segments = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyAudienceFilters) GetUsers() []ProductAnalyticsJourneyAudienceUserQuery {
	if o == nil || o.Users == nil {
		var ret []ProductAnalyticsJourneyAudienceUserQuery
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyAudienceFilters) GetUsersOk() (*[]ProductAnalyticsJourneyAudienceUserQuery, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyAudienceFilters) HasUsers() bool {
	return o != nil && o.Users != nil
}

// SetUsers gets a reference to the given []ProductAnalyticsJourneyAudienceUserQuery and assigns it to the Users field.
func (o *ProductAnalyticsJourneyAudienceFilters) SetUsers(v []ProductAnalyticsJourneyAudienceUserQuery) {
	o.Users = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneyAudienceFilters) MarshalJSON() ([]byte, error) {
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
func (o *ProductAnalyticsJourneyAudienceFilters) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Accounts []ProductAnalyticsJourneyAudienceAccountQuery `json:"accounts,omitempty"`
		Formula  *string                                       `json:"formula,omitempty"`
		Segments []ProductAnalyticsJourneyAudienceSegmentQuery `json:"segments,omitempty"`
		Users    []ProductAnalyticsJourneyAudienceUserQuery    `json:"users,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
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
