// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendPointsRequestAnyOfAnyOfItemQuery Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendPointsRequestAnyOfAnyOfItemQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource NetworkDevicesDataSourceName `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryString string `json:"query_string"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGeomapBackendPointsRequestAnyOfAnyOfItemQuery instantiates a new GeomapBackendPointsRequestAnyOfAnyOfItemQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeomapBackendPointsRequestAnyOfAnyOfItemQuery(dataSource NetworkDevicesDataSourceName, queryString string) *GeomapBackendPointsRequestAnyOfAnyOfItemQuery {
	this := GeomapBackendPointsRequestAnyOfAnyOfItemQuery{}
	this.DataSource = dataSource
	this.QueryString = queryString
	return &this
}

// NewGeomapBackendPointsRequestAnyOfAnyOfItemQueryWithDefaults instantiates a new GeomapBackendPointsRequestAnyOfAnyOfItemQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeomapBackendPointsRequestAnyOfAnyOfItemQueryWithDefaults() *GeomapBackendPointsRequestAnyOfAnyOfItemQuery {
	this := GeomapBackendPointsRequestAnyOfAnyOfItemQuery{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *GeomapBackendPointsRequestAnyOfAnyOfItemQuery) GetDataSource() NetworkDevicesDataSourceName {
	if o == nil {
		var ret NetworkDevicesDataSourceName
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *GeomapBackendPointsRequestAnyOfAnyOfItemQuery) GetDataSourceOk() (*NetworkDevicesDataSourceName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *GeomapBackendPointsRequestAnyOfAnyOfItemQuery) SetDataSource(v NetworkDevicesDataSourceName) {
	o.DataSource = v
}

// GetQueryString returns the QueryString field value.
func (o *GeomapBackendPointsRequestAnyOfAnyOfItemQuery) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *GeomapBackendPointsRequestAnyOfAnyOfItemQuery) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value.
func (o *GeomapBackendPointsRequestAnyOfAnyOfItemQuery) SetQueryString(v string) {
	o.QueryString = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeomapBackendPointsRequestAnyOfAnyOfItemQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["query_string"] = o.QueryString
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeomapBackendPointsRequestAnyOfAnyOfItemQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource  *NetworkDevicesDataSourceName `json:"data_source"`
		QueryString *string                       `json:"query_string"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.QueryString == nil {
		return fmt.Errorf("required field query_string missing")
	}

	hasInvalidField := false
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.QueryString = *all.QueryString

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
