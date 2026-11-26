// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsEventOccurrenceQuerySearch Test-only generated schema for the dashboard versioning proof of concept.
type ProductAnalyticsEventOccurrenceQuerySearch struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Occurrences OccurrenceFilter `json:"occurrences"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewProductAnalyticsEventOccurrenceQuerySearch instantiates a new ProductAnalyticsEventOccurrenceQuerySearch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsEventOccurrenceQuerySearch(occurrences OccurrenceFilter, query string) *ProductAnalyticsEventOccurrenceQuerySearch {
	this := ProductAnalyticsEventOccurrenceQuerySearch{}
	this.Occurrences = occurrences
	this.Query = query
	return &this
}

// NewProductAnalyticsEventOccurrenceQuerySearchWithDefaults instantiates a new ProductAnalyticsEventOccurrenceQuerySearch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsEventOccurrenceQuerySearchWithDefaults() *ProductAnalyticsEventOccurrenceQuerySearch {
	this := ProductAnalyticsEventOccurrenceQuerySearch{}
	return &this
}

// GetOccurrences returns the Occurrences field value.
func (o *ProductAnalyticsEventOccurrenceQuerySearch) GetOccurrences() OccurrenceFilter {
	if o == nil {
		var ret OccurrenceFilter
		return ret
	}
	return o.Occurrences
}

// GetOccurrencesOk returns a tuple with the Occurrences field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsEventOccurrenceQuerySearch) GetOccurrencesOk() (*OccurrenceFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Occurrences, true
}

// SetOccurrences sets field value.
func (o *ProductAnalyticsEventOccurrenceQuerySearch) SetOccurrences(v OccurrenceFilter) {
	o.Occurrences = v
}

// GetQuery returns the Query field value.
func (o *ProductAnalyticsEventOccurrenceQuerySearch) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsEventOccurrenceQuerySearch) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ProductAnalyticsEventOccurrenceQuerySearch) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsEventOccurrenceQuerySearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["occurrences"] = o.Occurrences
	toSerialize["query"] = o.Query
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsEventOccurrenceQuerySearch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Occurrences *OccurrenceFilter `json:"occurrences"`
		Query       *string           `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Occurrences == nil {
		return fmt.Errorf("required field occurrences missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}

	hasInvalidField := false
	if all.Occurrences.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Occurrences = *all.Occurrences
	o.Query = *all.Query

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
