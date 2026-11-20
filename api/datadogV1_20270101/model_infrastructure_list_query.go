// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// InfrastructureListQuery Test-only generated schema for the dashboard versioning proof of concept.
type InfrastructureListQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource InfrastructureListQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryString string `json:"query_string"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResourceType ContainerView `json:"resource_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewInfrastructureListQuery instantiates a new InfrastructureListQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInfrastructureListQuery(dataSource InfrastructureListQueryDataSource, queryString string, resourceType ContainerView) *InfrastructureListQuery {
	this := InfrastructureListQuery{}
	this.DataSource = dataSource
	this.QueryString = queryString
	this.ResourceType = resourceType
	return &this
}

// NewInfrastructureListQueryWithDefaults instantiates a new InfrastructureListQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInfrastructureListQueryWithDefaults() *InfrastructureListQuery {
	this := InfrastructureListQuery{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *InfrastructureListQuery) GetDataSource() InfrastructureListQueryDataSource {
	if o == nil {
		var ret InfrastructureListQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *InfrastructureListQuery) GetDataSourceOk() (*InfrastructureListQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *InfrastructureListQuery) SetDataSource(v InfrastructureListQueryDataSource) {
	o.DataSource = v
}

// GetQueryString returns the QueryString field value.
func (o *InfrastructureListQuery) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *InfrastructureListQuery) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value.
func (o *InfrastructureListQuery) SetQueryString(v string) {
	o.QueryString = v
}

// GetResourceType returns the ResourceType field value.
func (o *InfrastructureListQuery) GetResourceType() ContainerView {
	if o == nil {
		var ret ContainerView
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *InfrastructureListQuery) GetResourceTypeOk() (*ContainerView, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value.
func (o *InfrastructureListQuery) SetResourceType(v ContainerView) {
	o.ResourceType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InfrastructureListQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["query_string"] = o.QueryString
	toSerialize["resource_type"] = o.ResourceType
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InfrastructureListQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource   *InfrastructureListQueryDataSource `json:"data_source"`
		QueryString  *string                            `json:"query_string"`
		ResourceType *ContainerView                     `json:"resource_type"`
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
	if all.ResourceType == nil {
		return fmt.Errorf("required field resource_type missing")
	}

	hasInvalidField := false
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.QueryString = *all.QueryString
	o.ResourceType = *all.ResourceType

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
