// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AudienceFiltersUsersItems Test-only generated schema for the dashboard versioning proof of concept.
type AudienceFiltersUsersItems struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAudienceFiltersUsersItems instantiates a new AudienceFiltersUsersItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAudienceFiltersUsersItems(name string, query string) *AudienceFiltersUsersItems {
	this := AudienceFiltersUsersItems{}
	this.Name = name
	this.Query = query
	return &this
}

// NewAudienceFiltersUsersItemsWithDefaults instantiates a new AudienceFiltersUsersItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAudienceFiltersUsersItemsWithDefaults() *AudienceFiltersUsersItems {
	this := AudienceFiltersUsersItems{}
	return &this
}

// GetName returns the Name field value.
func (o *AudienceFiltersUsersItems) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AudienceFiltersUsersItems) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AudienceFiltersUsersItems) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value.
func (o *AudienceFiltersUsersItems) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *AudienceFiltersUsersItems) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *AudienceFiltersUsersItems) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AudienceFiltersUsersItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AudienceFiltersUsersItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name  *string `json:"name"`
		Query *string `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	o.Name = *all.Name
	o.Query = *all.Query

	return nil
}
